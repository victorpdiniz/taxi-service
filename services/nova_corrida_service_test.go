package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"taxi-service/models"
)

// Helper para criar uma notificação válida - CORRIGIDO para usar campos corretos
func createValidNotificacao() *models.NotificacaoCorrida {
	return &models.NotificacaoCorrida{
		CorridaID:       1,
		MotoristaID:     1,
		PassageiroNome:  "Maria Silva",
		Origem:  "Rua A, 123", // Campo correto
		Destino: "Rua B, 456", // Campo correto
		Valor:   12.5,   // Campo correto
		DistanciaKm:     1.5,
		TempoEstimado:   23,
	}
}

// Helper para limpar dados entre testes
func limparDadosTeste() {
	// Escrever array vazio para limpar dados
	writeNotificacoesCorrida([]models.NotificacaoCorrida{})
}

// Helper para criar motoristas mock para testes
func setupMotoristasParaTeste() error {
	_, err := criarDadosMockMotoristas()
	return err
}

func TestCreateNotificacaoCorrida(t *testing.T) {
	limparDadosTeste()
	setupMotoristasParaTeste()

	t.Run("Successful Creation", func(t *testing.T) {
		notificacao := createValidNotificacao()

		err := CreateNotificacaoCorrida(notificacao)

		require.NoError(t, err)
		assert.NotZero(t, notificacao.ID)
		assert.Equal(t, models.NotificacaoPendente, notificacao.Status)
		assert.False(t, notificacao.CreatedAt.IsZero())
		assert.False(t, notificacao.UpdatedAt.IsZero())
		assert.False(t, notificacao.ExpiraEm.IsZero())
		assert.True(t, notificacao.ExpiraEm.After(time.Now()))

		// Verificar se expira em aproximadamente 20 segundos
		expectedExpiry := time.Now().Add(20 * time.Second)
		assert.WithinDuration(t, expectedExpiry, notificacao.ExpiraEm, 1*time.Second)
	})

	t.Run("Multiple Notifications Have Different IDs", func(t *testing.T) {
		limparDadosTeste()

		notificacao1 := createValidNotificacao()
		notificacao2 := createValidNotificacao()
		notificacao2.MotoristaID = 2

		err1 := CreateNotificacaoCorrida(notificacao1)
		err2 := CreateNotificacaoCorrida(notificacao2)

		require.NoError(t, err1)
		require.NoError(t, err2)
		assert.NotEqual(t, notificacao1.ID, notificacao2.ID)
		assert.Equal(t, notificacao1.ID+1, notificacao2.ID)
	})

	t.Run("Notification Auto Expiration", func(t *testing.T) {
		limparDadosTeste()

		notificacao := createValidNotificacao()
		err := CreateNotificacaoCorrida(notificacao)
		require.NoError(t, err)

		// Verificar que inicialmente está pendente
		notificacaoInicial, err := GetNotificacaoCorrida(notificacao.ID)
		require.NoError(t, err)
		assert.Equal(t, models.NotificacaoPendente, notificacaoInicial.Status)

		// Simular expiração alterando ExpiraEm para o passado
		notificacoes, err := readNotificacoesCorrida()
		require.NoError(t, err)

		for i := range notificacoes {
			if notificacoes[i].ID == notificacao.ID {
				notificacoes[i].ExpiraEm = time.Now().Add(-1 * time.Second)
				break
			}
		}
		err = writeNotificacoesCorrida(notificacoes)
		require.NoError(t, err)

		// Tentar aceitar uma notificação expirada deve retornar erro
		err = AceitarNotificacaoCorrida(notificacao.ID, notificacao.MotoristaID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "notificacao expired")
	})
}

func TestGetNotificacaoCorrida(t *testing.T) {
	limparDadosTeste()

	t.Run("Successful Get", func(t *testing.T) {
		notificacao := createValidNotificacao()
		err := CreateNotificacaoCorrida(notificacao)
		require.NoError(t, err)

		resultado, err := GetNotificacaoCorrida(notificacao.ID)

		require.NoError(t, err)
		assert.Equal(t, notificacao.ID, resultado.ID)
		assert.Equal(t, notificacao.PassageiroNome, resultado.PassageiroNome)
		assert.Equal(t, notificacao.MotoristaID, resultado.MotoristaID)
		assert.Equal(t, notificacao.CorridaID, resultado.CorridaID)
		assert.Equal(t, notificacao.Origem, resultado.Origem)
		assert.Equal(t, notificacao.Destino, resultado.Destino)
	})

	t.Run("Notification Not Found", func(t *testing.T) {
		_, err := GetNotificacaoCorrida(999)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "notificacao not found")
	})
}

func TestListNotificacoesCorrida(t *testing.T) {
	limparDadosTeste()

	t.Run("List Empty Notifications", func(t *testing.T) {
		notificacoes, err := ListNotificacoesCorrida()

		require.NoError(t, err)
		assert.Empty(t, notificacoes)
	})

	t.Run("List Multiple Notifications", func(t *testing.T) {
		notificacao1 := createValidNotificacao()
		notificacao2 := createValidNotificacao()
		notificacao2.MotoristaID = 2
		notificacao2.PassageiroNome = "João Santos"

		err1 := CreateNotificacaoCorrida(notificacao1)
		err2 := CreateNotificacaoCorrida(notificacao2)
		require.NoError(t, err1)
		require.NoError(t, err2)

		notificacoes, err := ListNotificacoesCorrida()

		require.NoError(t, err)
		assert.Len(t, notificacoes, 2)

		// Verificar se as notificações estão na lista
		found1 := false
		found2 := false
		for _, n := range notificacoes {
			if n.ID == notificacao1.ID {
				found1 = true
				assert.Equal(t, "Maria Silva", n.PassageiroNome)
			}
			if n.ID == notificacao2.ID {
				found2 = true
				assert.Equal(t, "João Santos", n.PassageiroNome)
			}
		}
		assert.True(t, found1, "Notificação 1 não encontrada na lista")
		assert.True(t, found2, "Notificação 2 não encontrada na lista")
	})
}

func TestRecusarNotificacaoENotificarProximo(t *testing.T) {
	limparDadosTeste()
	setupMotoristasParaTeste()

	t.Run("Reject Notification And Notify Next Driver", func(t *testing.T) {
		// Criar primeira notificação para motorista ID 1
		notificacao := createValidNotificacao()
		notificacao.CorridaID = 100 // ID único para este teste
		err := CreateNotificacaoCorrida(notificacao)
		require.NoError(t, err)

		t.Logf("Criada notificação ID %d para corrida %d, motorista %d", notificacao.ID, notificacao.CorridaID, notificacao.MotoristaID)

		// Contar notificações antes da recusa
		notificacoesAntes, err := ListNotificacoesCorrida()
		require.NoError(t, err)
		countAntes := len(notificacoesAntes)

		// Recusar a notificação
		err = RecusarNotificacaoCorrida(notificacao.ID, notificacao.MotoristaID)
		require.NoError(t, err)

		// Verificar que foi marcada como recusada
		notificacaoRecusada, err := GetNotificacaoCorrida(notificacao.ID)
		require.NoError(t, err)
		assert.Equal(t, models.NotificacaoRecusada, notificacaoRecusada.Status)

		// Aguardar um pouco para a goroutine criar nova notificação
		time.Sleep(3 * time.Second)

		// Verificar se nova notificação foi criada
		notificacoesDepois, err := ListNotificacoesCorrida()
		require.NoError(t, err)
		countDepois := len(notificacoesDepois)

		assert.Greater(t, countDepois, countAntes, "Nova notificação deveria ter sido criada")

		// LOG todas as notificações para debug
		t.Logf("Notificações após recusa:")
		for _, n := range notificacoesDepois {
			t.Logf("- ID: %d, CorridaID: %d, MotoristaID: %d, Status: %s", n.ID, n.CorridaID, n.MotoristaID, n.Status)
		}

		// Encontrar a nova notificação
		var novaNotificacao *models.NotificacaoCorrida
		for _, n := range notificacoesDepois {
			if n.CorridaID == 100 && n.ID != notificacao.ID && n.Status == models.NotificacaoPendente {
				novaNotificacao = &n
				break
			}
		}

		require.NotNil(t, novaNotificacao, "Nova notificação para próximo motorista não foi encontrada")
		assert.NotEqual(t, notificacao.MotoristaID, novaNotificacao.MotoristaID, "Nova notificação deve ser para motorista diferente")
		assert.Equal(t, uint(100), novaNotificacao.CorridaID, "Nova notificação deve ser para mesma corrida")
		assert.Equal(t, models.NotificacaoPendente, novaNotificacao.Status, "Nova notificação deve estar pendente")
	})
}

func TestExpiracaoENotificarProximo(t *testing.T) {
	limparDadosTeste()
	setupMotoristasParaTeste()

	t.Run("Notification Expires And Notifies Next Driver", func(t *testing.T) {
		// Usar corridaID único e diferente de outros testes
		corridaID := uint(200)

		// Criar notificação que expirará
		notificacao := createValidNotificacao()
		notificacao.CorridaID = corridaID
		err := CreateNotificacaoCorrida(notificacao)
		require.NoError(t, err)

		t.Logf("Criada notificação ID %d para corrida %d, motorista %d", notificacao.ID, notificacao.CorridaID, notificacao.MotoristaID)

		// Contar notificações antes
		notificacoesAntes, err := ListNotificacoesCorrida()
		require.NoError(t, err)
		countAntes := len(notificacoesAntes)

		t.Logf("Aguardando 22 segundos para expiração...")

		// Aguardar expiração (20s + buffer)
		time.Sleep(22 * time.Second)

		// Verificar se notificação original expirou
		notificacaoExpirada, err := GetNotificacaoCorrida(notificacao.ID)
		require.NoError(t, err)
		assert.Equal(t, models.NotificacaoExpirada, notificacaoExpirada.Status)

		// Aguardar mais tempo para criação da nova notificação
		time.Sleep(5 * time.Second)

		// Verificar se nova notificação foi criada
		notificacoesDepois, err := ListNotificacoesCorrida()
		require.NoError(t, err)
		countDepois := len(notificacoesDepois)

		// LOG todas as notificações para debug
		t.Logf("Notificações após expiração (total: %d):", countDepois)
		for _, n := range notificacoesDepois {
			t.Logf("- ID: %d, CorridaID: %d, MotoristaID: %d, Status: %s", n.ID, n.CorridaID, n.MotoristaID, n.Status)
		}

		assert.Greater(t, countDepois, countAntes, "Nova notificação deveria ter sido criada após expiração")

		// Encontrar a nova notificação (buscar especificamente pela corridaID 200)
		var novaNotificacao *models.NotificacaoCorrida
		for _, n := range notificacoesDepois {
			if n.CorridaID == corridaID && n.ID != notificacao.ID && n.Status == models.NotificacaoPendente {
				novaNotificacao = &n
				break
			}
		}

		require.NotNil(t, novaNotificacao, "Nova notificação para próximo motorista não foi encontrada após expiração. CorridaID esperado: %d", corridaID)
		assert.NotEqual(t, notificacao.MotoristaID, novaNotificacao.MotoristaID, "Nova notificação deve ser para motorista diferente")
		assert.Equal(t, corridaID, novaNotificacao.CorridaID, "Nova notificação deve ser para mesma corrida")

		t.Logf("Nova notificação criada: ID %d para motorista %d", novaNotificacao.ID, novaNotificacao.MotoristaID)
	})
}

func TestBuscarMotoristasDisponiveis(t *testing.T) {
	setupMotoristasParaTeste()

	t.Run("Find Available Drivers Within Radius", func(t *testing.T) {
		// Coordenadas próximas aos motoristas mock (São Paulo)
		lat, lng := -23.5505, -46.6333
		excluirIDs := []uint{}
		raio := 5.0

		motoristas, err := buscarMotoristasDisponiveis(lat, lng, excluirIDs, raio)

		require.NoError(t, err)
		assert.NotEmpty(t, motoristas, "Deveria encontrar motoristas disponíveis")

		// Verificar que todos têm status disponível
		for _, m := range motoristas {
			assert.Equal(t, "disponivel", m.Status)
		}
	})

	t.Run("Exclude Already Notified Drivers", func(t *testing.T) {
		lat, lng := -23.5505, -46.6333
		excluirIDs := []uint{1, 2} // Excluir João Silva e Maria Santos
		raio := 5.0

		motoristas, err := buscarMotoristasDisponiveis(lat, lng, excluirIDs, raio)

		require.NoError(t, err)

		// Verificar que motoristas excluídos não estão na lista
		for _, m := range motoristas {
			assert.NotContains(t, excluirIDs, m.ID, "Motorista %d deveria ter sido excluído", m.ID)
		}
	})

	t.Run("No Available Drivers", func(t *testing.T) {
		lat, lng := -23.5505, -46.6333
		excluirIDs := []uint{1, 2, 3, 4} // Excluir todos os disponíveis
		raio := 5.0

		motoristas, err := buscarMotoristasDisponiveis(lat, lng, excluirIDs, raio)

		require.NoError(t, err)
		assert.Empty(t, motoristas, "Não deveria encontrar motoristas quando todos foram excluídos")
	})
}

func TestObterMotoristasJaNotificados(t *testing.T) {
	limparDadosTeste()

	t.Run("Get Notified Drivers For Ride", func(t *testing.T) {
		corridaID := uint(300)

		// Criar notificações para diferentes motoristas da mesma corrida
		notif1 := createValidNotificacao()
		notif1.CorridaID = corridaID
		notif1.MotoristaID = 1
		err := CreateNotificacaoCorrida(notif1)
		require.NoError(t, err)

		notif2 := createValidNotificacao()
		notif2.CorridaID = corridaID
		notif2.MotoristaID = 2
		err = CreateNotificacaoCorrida(notif2)
		require.NoError(t, err)

		// Criar notificação para corrida diferente
		notif3 := createValidNotificacao()
		notif3.CorridaID = 400
		notif3.MotoristaID = 3
		err = CreateNotificacaoCorrida(notif3)
		require.NoError(t, err)

		// Obter motoristas notificados para corrida 300
		motoristasNotificados, err := obterMotoristasJaNotificados(corridaID)

		require.NoError(t, err)
		assert.Len(t, motoristasNotificados, 2, "Deveria ter 2 motoristas notificados para corrida 300")
		assert.Contains(t, motoristasNotificados, uint(1))
		assert.Contains(t, motoristasNotificados, uint(2))
		assert.NotContains(t, motoristasNotificados, uint(3), "Motorista 3 é de corrida diferente")
	})
}

func TestCalcularDistanciaKm(t *testing.T) {
	t.Run("Calculate Distance Between Two Points", func(t *testing.T) {
		// Coordenadas de São Paulo (aproximadas)
		lat1, lng1 := -23.5505, -46.6333 // Centro de SP
		lat2, lng2 := -23.5515, -46.6343 // Próximo ao centro

		distancia := calcularDistanciaKm(lat1, lng1, lat2, lng2)

		assert.Greater(t, distancia, 0.0, "Distância deve ser maior que zero")
		assert.Less(t, distancia, 2.0, "Distância entre pontos próximos deve ser pequena")
	})

	t.Run("Distance Between Same Point Should Be Zero", func(t *testing.T) {
		lat, lng := -23.5505, -46.6333

		distancia := calcularDistanciaKm(lat, lng, lat, lng)

		assert.Equal(t, 0.0, distancia, "Distância entre mesmo ponto deve ser zero")
	})
}

func TestNotificarProximoMotorista(t *testing.T) {
	limparDadosTeste()
	setupMotoristasParaTeste()

	t.Run("Notify Next Available Driver", func(t *testing.T) {
		corridaID := uint(500)
		origemLat, origemLng := -23.5505, -46.6333

		// Criar notificação original - CORRIGIDO para usar campos corretos
		notificacaoOriginal := models.NotificacaoCorrida{
			CorridaID:       corridaID,
			MotoristaID:     1,
			PassageiroNome:  "Teste Passageiro",
			Origem:  "Origem Teste",
			Destino: "Destino Teste",
			Valor:   20.00,
		}

		// Contar notificações antes
		notificacoesAntes, err := ListNotificacoesCorrida()
		require.NoError(t, err)
		countAntes := len(notificacoesAntes)

		// Notificar próximo motorista
		err = notificarProximoMotorista(corridaID, origemLat, origemLng, notificacaoOriginal)
		require.NoError(t, err)

		// Verificar se nova notificação foi criada
		notificacoesDepois, err := ListNotificacoesCorrida()
		require.NoError(t, err)
		countDepois := len(notificacoesDepois)

		assert.Greater(t, countDepois, countAntes, "Nova notificação deveria ter sido criada")

		// Encontrar a nova notificação
		var novaNotificacao *models.NotificacaoCorrida
		for _, n := range notificacoesDepois {
			if n.CorridaID == corridaID {
				novaNotificacao = &n
				break
			}
		}

		require.NotNil(t, novaNotificacao, "Nova notificação não foi encontrada")
		assert.Equal(t, corridaID, novaNotificacao.CorridaID)
		assert.Equal(t, notificacaoOriginal.PassageiroNome, novaNotificacao.PassageiroNome)
		assert.Equal(t, models.NotificacaoPendente, novaNotificacao.Status)
		assert.NotZero(t, novaNotificacao.MotoristaID)
	})
}

func TestCicloCompletoRecusaEProximoMotorista(t *testing.T) {
	limparDadosTeste()
	setupMotoristasParaTeste()

	t.Run("Complete Cycle: Multiple Rejections And Next Drivers", func(t *testing.T) {
		corridaID := uint(600)

		// Criar primeira notificação
		notificacao1 := createValidNotificacao()
		notificacao1.CorridaID = corridaID
		notificacao1.MotoristaID = 1
		err := CreateNotificacaoCorrida(notificacao1)
		require.NoError(t, err)

		t.Logf("Primeira notificação criada para motorista %d", notificacao1.MotoristaID)

		// Recusar primeira notificação
		err = RecusarNotificacaoCorrida(notificacao1.ID, notificacao1.MotoristaID)
		require.NoError(t, err)

		// Aguardar criação da segunda notificação
		time.Sleep(3 * time.Second)

		// Buscar segunda notificação
		notificacoes, err := ListNotificacoesCorrida()
		require.NoError(t, err)

		var notificacao2 *models.NotificacaoCorrida
		for _, n := range notificacoes {
			if n.CorridaID == corridaID && n.Status == models.NotificacaoPendente {
				notificacao2 = &n
				break
			}
		}

		require.NotNil(t, notificacao2, "Segunda notificação não foi criada")
		assert.NotEqual(t, notificacao1.MotoristaID, notificacao2.MotoristaID, "Segunda notificação deve ser para motorista diferente")

		t.Logf("Segunda notificação criada para motorista %d", notificacao2.MotoristaID)

		// Recusar segunda notificação
		err = RecusarNotificacaoCorrida(notificacao2.ID, notificacao2.MotoristaID)
		require.NoError(t, err)

		// Aguardar criação da terceira notificação
		time.Sleep(3 * time.Second)

		// Verificar se terceira notificação foi criada
		notificacoesFinal, err := ListNotificacoesCorrida()
		require.NoError(t, err)

		var notificacao3 *models.NotificacaoCorrida
		for _, n := range notificacoesFinal {
			if n.CorridaID == corridaID && n.Status == models.NotificacaoPendente && n.ID != notificacao2.ID {
				notificacao3 = &n
				break
			}
		}

		require.NotNil(t, notificacao3, "Terceira notificação não foi criada")
		assert.NotEqual(t, notificacao1.MotoristaID, notificacao3.MotoristaID)
		assert.NotEqual(t, notificacao2.MotoristaID, notificacao3.MotoristaID)

		t.Logf("Terceira notificação criada para motorista %d", notificacao3.MotoristaID)

		// Verificar que temos 3 notificações para a mesma corrida
		contadorCorrida := 0
		for _, n := range notificacoesFinal {
			if n.CorridaID == corridaID {
				contadorCorrida++
			}
		}

		assert.Equal(t, 3, contadorCorrida, "Deveria ter 3 notificações para a corrida %d", corridaID)
	})
}

func TestConcurrentNotificationCreation(t *testing.T) {
	limparDadosTeste()

	t.Run("Multiple Goroutines Creating Notifications", func(t *testing.T) {
		const numGoroutines = 5
		resultChan := make(chan *models.NotificacaoCorrida, numGoroutines)
		errorChan := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(motoristaID uint) {
				notificacao := createValidNotificacao()
				notificacao.MotoristaID = motoristaID
				err := CreateNotificacaoCorrida(notificacao)
				if err != nil {
					errorChan <- err
				} else {
					resultChan <- notificacao
				}
			}(uint(i + 1))
		}

		var notificacoes []*models.NotificacaoCorrida
		var errors []error

		for i := 0; i < numGoroutines; i++ {
			select {
			case notif := <-resultChan:
				notificacoes = append(notificacoes, notif)
			case err := <-errorChan:
				errors = append(errors, err)
			case <-time.After(10 * time.Second):
				t.Fatal("Timeout waiting for goroutines")
			}
		}

		assert.Empty(t, errors, "Não deveria haver erros na criação concorrente")
		assert.Len(t, notificacoes, numGoroutines)

		// Verificar IDs únicos
		ids := make(map[uint]bool)
		for _, notif := range notificacoes {
			assert.False(t, ids[notif.ID], "ID duplicado encontrado: %d", notif.ID)
			ids[notif.ID] = true
		}
	})
}

func TestRealTimeExpirationWithNextDriver(t *testing.T) {
	limparDadosTeste()
	setupMotoristasParaTeste()

	t.Run("Real Time Expiration Creates Next Notification", func(t *testing.T) {
		// Usar corridaID único para evitar conflitos
		corridaID := uint(700)

		// Criar notificação que expirará
		notificacao := createValidNotificacao()
		notificacao.CorridaID = corridaID
		err := CreateNotificacaoCorrida(notificacao)
		require.NoError(t, err)

		t.Logf("Criada notificação ID %d para corrida %d, motorista %d", notificacao.ID, notificacao.CorridaID, notificacao.MotoristaID)
		t.Logf("Aguardando 22 segundos para expiração...")

		// Aguardar expiração real (22 segundos)
		time.Sleep(22 * time.Second)

		// Verificar se notificação original expirou
		notificacaoExpirada, err := GetNotificacaoCorrida(notificacao.ID)
		require.NoError(t, err)
		assert.Equal(t, models.NotificacaoExpirada, notificacaoExpirada.Status)

		// Aguardar mais tempo para criação da próxima notificação
		time.Sleep(5 * time.Second)

		// Verificar se nova notificação foi criada
		notificacoes, err := ListNotificacoesCorrida()
		require.NoError(t, err)

		// LOG todas as notificações para debug
		t.Logf("Notificações encontradas:")
		for _, n := range notificacoes {
			t.Logf("- ID: %d, CorridaID: %d, MotoristaID: %d, Status: %s", n.ID, n.CorridaID, n.MotoristaID, n.Status)
		}

		var novaNotificacao *models.NotificacaoCorrida
		for _, n := range notificacoes {
			// Buscar especificamente pela corridaID 700
			if n.CorridaID == corridaID && n.Status == models.NotificacaoPendente && n.ID != notificacao.ID {
				novaNotificacao = &n
				break
			}
		}

		require.NotNil(t, novaNotificacao, "Nova notificação deveria ter sido criada após expiração para corrida %d", corridaID)
		assert.NotEqual(t, notificacao.MotoristaID, novaNotificacao.MotoristaID, "Nova notificação deve ser para motorista diferente")
		assert.Equal(t, corridaID, novaNotificacao.CorridaID, "Nova notificação deve ser para mesma corrida")

		t.Logf("Nova notificação criada para motorista %d após expiração", novaNotificacao.MotoristaID)
	})
}
