package services

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"os"
	"path/filepath"
	"sync"
	"taxi-service/models"
	"time"
)

const notificacaoCorridaFile = "../data/notificacao_corrida.json"
const motoristasDisponiveis = "../data/motoristas_disponiveis.json"

// Mutex global para controle de concorrência
var notificacaoMutex sync.RWMutex

// ============= FUNÇÕES AUXILIARES DE ARQUIVO =============

func readNotificacoesCorrida() ([]models.NotificacaoCorrida, error) {
	// Criar diretório se não existir
	dataDir := filepath.Dir(notificacaoCorridaFile)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	file, err := os.Open(notificacaoCorridaFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.NotificacaoCorrida{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var notificacoes []models.NotificacaoCorrida
	err = json.NewDecoder(file).Decode(&notificacoes)
	if err != nil {
		return nil, err
	}

	return notificacoes, nil
}

func writeNotificacoesCorrida(notificacoes []models.NotificacaoCorrida) error {
	// Garantir que o diretório existe antes de escrever
	dataDir := filepath.Dir(notificacaoCorridaFile)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(notificacoes, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(notificacaoCorridaFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ============= FUNÇÕES AUXILIARES PARA MOTORISTAS =============

// MotoristaDisponivel representa um motorista disponível
type MotoristaDisponivel struct {
	ID        uint    `json:"id"`
	Nome      string  `json:"nome"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Status    string  `json:"status"` // "disponivel", "ocupado"
	Avaliacao float64 `json:"avaliacao"`
}

// readMotoristasDisponiveis - Lê motoristas disponíveis do arquivo
func readMotoristasDisponiveis() ([]MotoristaDisponivel, error) {
	dataDir := filepath.Dir(motoristasDisponiveis)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	file, err := os.Open(motoristasDisponiveis)
	if err != nil {
		if os.IsNotExist(err) {
			// Criar arquivo com dados mock se não existir
			return criarDadosMockMotoristas()
		}
		return nil, err
	}
	defer file.Close()

	var motoristas []MotoristaDisponivel
	err = json.NewDecoder(file).Decode(&motoristas)
	if err != nil {
		return nil, err
	}

	return motoristas, nil
}

// criarDadosMockMotoristas - Cria dados mock para testes
func criarDadosMockMotoristas() ([]MotoristaDisponivel, error) {
	motoristas := []MotoristaDisponivel{
		{ID: 1, Nome: "João Silva", Lat: -23.5505, Lng: -46.6333, Status: "disponivel", Avaliacao: 4.8},
		{ID: 2, Nome: "Maria Santos", Lat: -23.5515, Lng: -46.6343, Status: "disponivel", Avaliacao: 4.9},
		{ID: 3, Nome: "Pedro Costa", Lat: -23.5525, Lng: -46.6353, Status: "disponivel", Avaliacao: 4.7},
		{ID: 4, Nome: "Ana Oliveira", Lat: -23.5535, Lng: -46.6363, Status: "disponivel", Avaliacao: 4.6},
		{ID: 5, Nome: "Carlos Pereira", Lat: -23.5545, Lng: -46.6373, Status: "ocupado", Avaliacao: 4.9},
	}

	// Salvar dados mock
	dataDir := filepath.Dir(motoristasDisponiveis)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(motoristas, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(motoristasDisponiveis, data, 0644)
	if err != nil {
		return nil, err
	}

	return motoristas, nil
}

// buscarMotoristasDisponiveis - Busca motoristas disponíveis próximos, excluindo os já notificados
func buscarMotoristasDisponiveis(lat, lng float64, excluirMotoristasIDs []uint, raioKm float64) ([]MotoristaDisponivel, error) {
	motoristas, err := readMotoristasDisponiveis()
	if err != nil {
		return nil, err
	}

	var motoristasElegiveis []MotoristaDisponivel

	for _, motorista := range motoristas {
		// Pular se não está disponível
		if motorista.Status != "disponivel" {
			continue
		}

		// Pular se está na lista de exclusão
		jaNotificado := false
		for _, excluirID := range excluirMotoristasIDs {
			if motorista.ID == excluirID {
				jaNotificado = true
				break
			}
		}
		if jaNotificado {
			continue
		}

		// Calcular distância
		distancia := calcularDistanciaKm(lat, lng, motorista.Lat, motorista.Lng)
		if distancia <= raioKm {
			motoristasElegiveis = append(motoristasElegiveis, motorista)
		}
	}

	// Ordenar por avaliação (melhor primeiro) e depois por distância
	for i := 0; i < len(motoristasElegiveis); i++ {
		for j := i + 1; j < len(motoristasElegiveis); j++ {
			// Priorizar por avaliação
			if motoristasElegiveis[j].Avaliacao > motoristasElegiveis[i].Avaliacao {
				motoristasElegiveis[i], motoristasElegiveis[j] = motoristasElegiveis[j], motoristasElegiveis[i]
			}
		}
	}

	return motoristasElegiveis, nil
}

// calcularDistanciaKm - Calcula distância entre dois pontos usando fórmula de Haversine
func calcularDistanciaKm(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadiusKm = 6371.0

	lat1Rad := lat1 * math.Pi / 180
	lng1Rad := lng1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lng2Rad := lng2 * math.Pi / 180

	dlat := lat2Rad - lat1Rad
	dlng := lng2Rad - lng1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

// obterMotoristasJaNotificados - Obtém lista de motoristas já notificados para uma corrida
func obterMotoristasJaNotificados(corridaID uint) ([]uint, error) {
	notificacoes, err := readNotificacoesCorrida()
	if err != nil {
		return nil, err
	}

	var motoristasNotificados []uint
	for _, notificacao := range notificacoes {
		if notificacao.CorridaID == corridaID {
			motoristasNotificados = append(motoristasNotificados, notificacao.MotoristaID)
		}
	}

	return motoristasNotificados, nil
}

// notificarProximoMotorista - Encontra e notifica o próximo motorista disponível
func notificarProximoMotorista(corridaID uint, origemLat, origemLng float64, notificacaoOriginal models.NotificacaoCorrida) error {
	log.Printf("Iniciando busca por próximo motorista para corrida %d", corridaID)

	// Obter lista de motoristas já notificados
	motoristasJaNotificados, err := obterMotoristasJaNotificados(corridaID)
	if err != nil {
		log.Printf("Erro ao obter motoristas já notificados: %v", err)
		return err
	}

	log.Printf("Motoristas já notificados para corrida %d: %v", corridaID, motoristasJaNotificados)

	// Buscar motoristas disponíveis num raio de 5km
	motoristasDisponiveis, err := buscarMotoristasDisponiveis(origemLat, origemLng, motoristasJaNotificados, 5.0)
	if err != nil {
		log.Printf("Erro ao buscar motoristas disponíveis: %v", err)
		return err
	}

	// Se não há mais motoristas disponíveis
	if len(motoristasDisponiveis) == 0 {
		log.Printf("Nenhum motorista disponível encontrado para corrida %d", corridaID)
		return nil // Não é erro, apenas não há mais motoristas
	}

	// Selecionar o primeiro motorista (melhor avaliado)
	proximoMotorista := motoristasDisponiveis[0]
	log.Printf("Próximo motorista selecionado: %s (ID: %d, Avaliação: %.1f)",
		proximoMotorista.Nome, proximoMotorista.ID, proximoMotorista.Avaliacao)

	// Calcular nova distância e tempo estimado
	distancia := calcularDistanciaKm(proximoMotorista.Lat, proximoMotorista.Lng, origemLat, origemLng)
	TempoEstimado := float64(distancia * 3) // Estimativa: 3 minutos por km

	// Criar nova notificação baseada na original
	novaNotificacao := &models.NotificacaoCorrida{
		CorridaID:       corridaID,
		MotoristaID:     proximoMotorista.ID,
		PassageiroNome:  notificacaoOriginal.PassageiroNome,
		Origem:  notificacaoOriginal.Origem,
		Destino: notificacaoOriginal.Destino,
		Valor:   notificacaoOriginal.Valor,
		DistanciaKm:     distancia,
		TempoEstimado:   TempoEstimado,
	}

	// Criar a notificação (que iniciará novo timer de 20 segundos)
	if err := CreateNotificacaoCorrida(novaNotificacao); err != nil {
		log.Printf("Erro ao criar nova notificação para motorista %d: %v", proximoMotorista.ID, err)
		return err
	}

	log.Printf("Nova notificação criada (ID: %d) para motorista %s (ID: %d) para corrida %d",
		novaNotificacao.ID, proximoMotorista.Nome, proximoMotorista.ID, corridaID)

	return nil
}

// ============= FUNÇÕES PRINCIPAIS DE SERVIÇO =============

// ListNotificacoesCorrida - Lista todas as notificações
func ListNotificacoesCorrida() ([]models.NotificacaoCorrida, error) {
	return readNotificacoesCorrida()
}

// GetNotificacaoCorrida - Busca notificação por ID
func GetNotificacaoCorrida(id uint) (models.NotificacaoCorrida, error) {

	notificacoes, err := readNotificacoesCorrida()
	if err != nil {
		return models.NotificacaoCorrida{}, err
	}

	for _, notificacao := range notificacoes {
		if notificacao.ID == id {
			return notificacao, nil
		}
	}

	return models.NotificacaoCorrida{}, errors.New("notificacao not found")
}

// AceitarNotificacaoCorrida - Aceita uma notificação de corrida
func AceitarNotificacaoCorrida(notificacaoID uint, motoristaID uint) error {
	notificacaoMutex.Lock()
	defer notificacaoMutex.Unlock()

	// Usar GetNotificacaoCorrida para buscar a notificação
	notificacao, err := GetNotificacaoCorrida(notificacaoID)
	if err != nil {
		return err
	}

	// Verificar se pertence ao motorista
	if notificacao.MotoristaID != motoristaID {
		return errors.New("notificacao does not belong to this motorista")
	}

	// Verificar se ainda está pendente
	if notificacao.Status != models.NotificacaoPendente {
		return errors.New("notificacao already processed")
	}

	// Verificar se não expirou
	if time.Now().After(notificacao.ExpiraEm) {
		// Marcar como expirada
		notificacao.Status = models.NotificacaoExpirada
		notificacao.UpdatedAt = time.Now()

		// Atualizar no array
		notificacoes, err := readNotificacoesCorrida()
		if err != nil {
			return err
		}

		for i := range notificacoes {
			if notificacoes[i].ID == notificacaoID {
				notificacoes[i] = notificacao
				break
			}
		}

		writeNotificacoesCorrida(notificacoes)
		return errors.New("notificacao expired")
	}

	// Aceitar a notificação
	notificacao.Status = models.NotificacaoAceita
	notificacao.UpdatedAt = time.Now()

	// Atualizar no array
	notificacoes, err := readNotificacoesCorrida()
	if err != nil {
		return err
	}

	for i := range notificacoes {
		if notificacoes[i].ID == notificacaoID {
			notificacoes[i] = notificacao
			break
		}
	}

	if err := writeNotificacoesCorrida(notificacoes); err != nil {
		return err
	}

	log.Printf("Notificação %d aceita pelo motorista %d", notificacaoID, motoristaID)
	return nil
}

// RecusarNotificacaoCorrida - Recusa uma notificação de corrida
func RecusarNotificacaoCorrida(notificacaoID uint, motoristaID uint) error {
	notificacaoMutex.Lock()
	defer notificacaoMutex.Unlock()

	// Usar GetNotificacaoCorrida para buscar a notificação
	notificacao, err := GetNotificacaoCorrida(notificacaoID)
	if err != nil {
		return err
	}

	// Verificar se pertence ao motorista
	if notificacao.MotoristaID != motoristaID {
		return errors.New("notificacao does not belong to this motorista")
	}

	// Verificar se ainda está pendente
	if notificacao.Status != models.NotificacaoPendente {
		return errors.New("notificacao already processed")
	}

	// Verificar se não expirou
	if time.Now().After(notificacao.ExpiraEm) {
		// Marcar como expirada
		notificacao.Status = models.NotificacaoExpirada
		notificacao.UpdatedAt = time.Now()

		// Atualizar no array
		notificacoes, err := readNotificacoesCorrida()
		if err != nil {
			return err
		}

		for i := range notificacoes {
			if notificacoes[i].ID == notificacaoID {
				notificacoes[i] = notificacao
				break
			}
		}

		writeNotificacoesCorrida(notificacoes)
		return errors.New("notificacao expired")
	}

	// Recusar a notificação
	notificacao.Status = models.NotificacaoRecusada
	notificacao.UpdatedAt = time.Now()

	// Atualizar no array
	notificacoes, err := readNotificacoesCorrida()
	if err != nil {
		return err
	}

	for i := range notificacoes {
		if notificacoes[i].ID == notificacaoID {
			notificacoes[i] = notificacao
			break
		}
	}

	if err := writeNotificacoesCorrida(notificacoes); err != nil {
		return err
	}

	log.Printf("Notificação %d recusada pelo motorista %d", notificacaoID, motoristaID)

	// Notificar próximo motorista após recusa
	go func() {
		// Aguardar um pouco para evitar concorrência
		time.Sleep(1 * time.Second)

		if err := notificarProximoMotorista(notificacao.CorridaID, -23.5505, -46.6333, notificacao); err != nil {
			log.Printf("Erro ao notificar próximo motorista após recusa: %v", err)
		}
	}()

	return nil
}

// CreateNotificacaoCorrida - Cria nova notificação para motorista
func CreateNotificacaoCorrida(notificacao *models.NotificacaoCorrida) error {
	notificacaoMutex.Lock()
	notificacoes, err := readNotificacoesCorrida()
	if err != nil {
		notificacaoMutex.Unlock()
		return err
	}

	// Atribuir novo ID (máximo + 1)
	var maxID uint = 0
	for _, n := range notificacoes {
		if n.ID > maxID {
			maxID = n.ID
		}
	}
	notificacao.ID = maxID + 1

	// Definir valores padrão
	now := time.Now()
	notificacao.Status = models.NotificacaoPendente
	notificacao.CreatedAt = now
	notificacao.UpdatedAt = now
	notificacao.ExpiraEm = now.Add(20 * time.Second) // Expira em 20 segundos

	// Adicionar nova notificação à lista
	notificacoes = append(notificacoes, *notificacao)

	if err := writeNotificacoesCorrida(notificacoes); err != nil {
		notificacaoMutex.Unlock()
		return err
	}

	notificacaoID := notificacao.ID
	corridaID := notificacao.CorridaID
	notificacaoParaProximo := *notificacao // Cópia para usar na goroutine
	notificacaoMutex.Unlock()

	log.Printf("Notificação %d criada para motorista %d (corrida %d)", notificacaoID, notificacao.MotoristaID, corridaID)

	// Iniciar rotina para expiração automática após 20 segundos
	go func(id uint, corrida uint, notifOriginal models.NotificacaoCorrida) {
		// Aguardar 20 segundos
		time.Sleep(20 * time.Second)

		// Usar mutex para operação thread-safe
		notificacaoMutex.Lock()
		defer notificacaoMutex.Unlock()

		// Usar GetNotificacaoCorrida para buscar a notificação
		notificacaoAtual, err := GetNotificacaoCorrida(id)
		if err != nil {
			// Notificação não encontrada, já foi removida
			return
		}

		// Verificar se ainda está pendente (não foi aceita nem recusada)
		if notificacaoAtual.Status == models.NotificacaoPendente {
			// Marcar como expirada
			notificacaoAtual.Status = models.NotificacaoExpirada
			notificacaoAtual.UpdatedAt = time.Now()

			// Atualizar no array
			notificacoes, err := readNotificacoesCorrida()
			if err != nil {
				return
			}

			for i := range notificacoes {
				if notificacoes[i].ID == id {
					notificacoes[i] = notificacaoAtual
					break
				}
			}

			// Salvar alterações
			if err := writeNotificacoesCorrida(notificacoes); err != nil {
				log.Printf("Erro ao salvar notificação expirada: %v", err)
				return
			}

			log.Printf("Notificação %d expirou automaticamente", id)

			// LÓGICA ADICIONAL: Notificar próximo motorista após expiração
			go func() {
				if err := notificarProximoMotorista(corrida, -23.5505, -46.6333, notifOriginal); err != nil {
					log.Printf("Erro ao notificar próximo motorista após expiração: %v", err)
				}
			}()
		}

		// Se status não é pendente, significa que foi aceita ou recusada
		// Nenhuma ação necessária neste caso
	}(notificacaoID, corridaID, notificacaoParaProximo)

	return nil
}