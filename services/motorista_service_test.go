package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"taxi_service/models"
)

// MockMotoristaRepository mock do repositório de motorista
type MockMotoristaRepository struct {
	motoristas map[string]*models.Motorista
}

func NewMockMotoristaRepository() *MockMotoristaRepository {
	return &MockMotoristaRepository{
		motoristas: make(map[string]*models.Motorista),
	}
}

func (m *MockMotoristaRepository) Criar(motorista *models.Motorista) error {
	if _, exists := m.motoristas[motorista.ID]; exists {
		return errors.New("motorista com este ID já existe")
	}
	m.motoristas[motorista.ID] = motorista
	return nil
}

func (m *MockMotoristaRepository) BuscarPorID(id string) (*models.Motorista, error) {
	if motorista, exists := m.motoristas[id]; exists {
		return motorista, nil
	}
	return nil, errors.New("motorista não encontrado")
}

func (m *MockMotoristaRepository) BuscarPorEmail(email string) (*models.Motorista, error) {
	for _, motorista := range m.motoristas {
		if motorista.Email == email {
			return motorista, nil
		}
	}
	return nil, errors.New("motorista não encontrado")
}

func (m *MockMotoristaRepository) BuscarPorCPF(cpf string) (*models.Motorista, error) {
	for _, motorista := range m.motoristas {
		if motorista.CPF == cpf {
			return motorista, nil
		}
	}
	return nil, errors.New("motorista não encontrado")
}

func (m *MockMotoristaRepository) BuscarPorCNH(cnh string) (*models.Motorista, error) {
	for _, motorista := range m.motoristas {
		if motorista.CNH == cnh {
			return motorista, nil
		}
	}
	return nil, errors.New("motorista não encontrado")
}

func (m *MockMotoristaRepository) Atualizar(motorista *models.Motorista) error {
	if _, exists := m.motoristas[motorista.ID]; !exists {
		return errors.New("motorista não encontrado")
	}
	m.motoristas[motorista.ID] = motorista
	return nil
}

func (m *MockMotoristaRepository) Deletar(id string) error {
	if _, exists := m.motoristas[id]; !exists {
		return errors.New("motorista não encontrado")
	}
	delete(m.motoristas, id)
	return nil
}

func (m *MockMotoristaRepository) ListarTodos() ([]*models.Motorista, error) {
	var motoristas []*models.Motorista
	for _, motorista := range m.motoristas {
		motoristas = append(motoristas, motorista)
	}
	return motoristas, nil
}

func TestMotoristaService(t *testing.T) {
	// Setup
	repo := NewMockMotoristaRepository()
	emailService := NewMockEmailService()
	service := NewMotoristaService(repo, emailService)

	// Request base válido para reutilizar em todos os testes
	createValidRequest := func() CadastroMotoristaRequest {
		return CadastroMotoristaRequest{
			Nome:             "João Silva",
			DataNascimento:   "15/03/1990",
			CPF:              "11144477735", // CPF válido
			CNH:              "12345678901",
			CategoriaCNH:     "B",
			ValidadeCNH:      "15/03/2030",
			PlacaVeiculo:     "ABC1234",
			ModeloVeiculo:    "Honda Civic 2020",
			Telefone:         "11999999999",
			Email:            "joao.silva@email.com",
			Senha:            "MinhaSenh@123",
			ConfirmacaoSenha: "MinhaSenh@123",
		}
	}

	var motoristaID string

	t.Run("Cadastro bem-sucedido", func(t *testing.T) {
		request := createValidRequest()

		motorista, err := service.CadastrarMotorista(request)
		require.NoError(t, err)

		assert.NotEmpty(t, motorista.ID)
		assert.Equal(t, "João Silva", motorista.Nome)
		assert.Equal(t, "joao.silva@email.com", motorista.Email)
		assert.Equal(t, models.StatusAguardandoAprovacao, motorista.Status)

		// Salvar ID para testes posteriores
		motoristaID = motorista.ID

		// Verificar se email foi enviado
		emails := emailService.ObterEmailsEnviados()
		assert.Len(t, emails, 1)
		assert.Equal(t, "joao.silva@email.com", emails[0].Para)
		assert.Contains(t, emails[0].Assunto, "Cadastro realizado com sucesso")
	})

	t.Run("Erro - senhas não conferem", func(t *testing.T) {
		request := createValidRequest()
		request.Email = "outro@email.com"          // Email diferente para não conflitar
		request.CPF = "52998224725"                // CPF diferente válido
		request.ConfirmacaoSenha = "MinhaSenh@456" // Senha de confirmação diferente

		_, err := service.CadastrarMotorista(request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Senhas não conferem")
	})

	t.Run("Erro - CPF já cadastrado", func(t *testing.T) {
		request := createValidRequest()
		request.Email = "maria@email.com" // Email diferente
		request.CNH = "98765432109"       // CNH diferente
		// CPF mantém o mesmo do primeiro teste para testar duplicação

		_, err := service.CadastrarMotorista(request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CPF já cadastrado")
	})

	t.Run("Erro - idade mínima", func(t *testing.T) {
		request := createValidRequest()
		request.DataNascimento = "15/03/2010" // 15 anos
		request.CPF = "60968336086"           // CPF válido diferente
		request.Email = "jovem@email.com"     // Email diferente
		request.CNH = "11111111111"           // CNH diferente

		_, err := service.CadastrarMotorista(request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Motorista deve ter pelo menos 18 anos")
	})

	t.Run("Erro - CNH vencida", func(t *testing.T) {
		request := createValidRequest()
		request.ValidadeCNH = "15/03/2020"  // CNH vencida
		request.CPF = "07727709049"         // CPF válido diferente
		request.Email = "vencido@email.com" // Email diferente
		request.CNH = "22222222222"         // CNH diferente

		_, err := service.CadastrarMotorista(request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CNH vencida. Renove sua CNH para prosseguir")
	})

	t.Run("Upload de documento", func(t *testing.T) {
		uploadRequest := UploadDocumentoRequest{
			TipoDocumento:  "CNH",
			CaminhoArquivo: "/uploads/cnh_joao.jpg",
			Formato:        "JPG",
			Tamanho:        2 * 1024 * 1024, // 2MB
		}

		err := service.UploadDocumento(motoristaID, uploadRequest)
		assert.NoError(t, err)

		// Verificar se documento foi adicionado
		motorista, err := service.BuscarMotorista(motoristaID)
		require.NoError(t, err)
		assert.Len(t, motorista.Documentos, 1)
		assert.Equal(t, "CNH", motorista.Documentos[0].TipoDocumento)
		assert.Equal(t, "JPG", motorista.Documentos[0].Formato)
	})

	t.Run("Upload de todos os documentos obrigatórios", func(t *testing.T) {
		// Upload CRLV
		uploadRequest := UploadDocumentoRequest{
			TipoDocumento:  "CRLV",
			CaminhoArquivo: "/uploads/crlv_joao.png",
			Formato:        "PNG",
			Tamanho:        int64(1.5 * 1024 * 1024), // 1.5MB
		}
		err := service.UploadDocumento(motoristaID, uploadRequest)
		assert.NoError(t, err)

		// Upload selfie com CNH
		uploadRequest = UploadDocumentoRequest{
			TipoDocumento:  "selfie_cnh",
			CaminhoArquivo: "/uploads/selfie_joao.jpg",
			Formato:        "JPG",
			Tamanho:        1 * 1024 * 1024, // 1MB
		}
		err = service.UploadDocumento(motoristaID, uploadRequest)
		assert.NoError(t, err)

		// Verificar mudança de status
		motorista, err := service.BuscarMotorista(motoristaID)
		require.NoError(t, err)
		assert.Equal(t, models.StatusDocumentosAnalise, motorista.Status)
		assert.Len(t, motorista.Documentos, 3)

		// Verificar se email de recebimento foi enviado
		emails := emailService.ObterEmailsEnviados()
		assert.GreaterOrEqual(t, len(emails), 2) // Pelo menos confirmação + recebimento
	})

	t.Run("Erro - upload arquivo muito grande", func(t *testing.T) {
		uploadRequest := UploadDocumentoRequest{
			TipoDocumento:  "CNH",
			CaminhoArquivo: "/uploads/cnh_grande.jpg",
			Formato:        "JPG",
			Tamanho:        6 * 1024 * 1024, // 6MB - muito grande
		}

		err := service.UploadDocumento(motoristaID, uploadRequest)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Arquivo muito grande. Tamanho máximo: 5MB")
	})

	t.Run("Erro - formato inválido", func(t *testing.T) {
		uploadRequest := UploadDocumentoRequest{
			TipoDocumento:  "CNH",
			CaminhoArquivo: "/uploads/cnh.txt",
			Formato:        "TXT",
			Tamanho:        1 * 1024 * 1024,
		}

		err := service.UploadDocumento(motoristaID, uploadRequest)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Formato não suportado. Use JPG, PNG ou PDF")
	})

	t.Run("Validação automática de documentos", func(t *testing.T) {
		emailService.LimparEmails() // Limpar emails anteriores

		err := service.ValidarDocumentos(motoristaID)
		assert.NoError(t, err)

		// Verificar mudança de status
		motorista, err := service.BuscarMotorista(motoristaID)
		require.NoError(t, err)
		assert.Equal(t, models.StatusAprovado, motorista.Status)

		// Verificar todos os documentos aprovados
		for _, doc := range motorista.Documentos {
			assert.Equal(t, "aprovado", doc.Status)
		}

		// Verificar se email de aprovação foi enviado
		emails := emailService.ObterEmailsEnviados()
		assert.Len(t, emails, 1)
		assert.Contains(t, emails[0].Assunto, "aprovado")
	})
}

func TestValidarDadosCadastro(t *testing.T) {
	repo := NewMockMotoristaRepository()
	emailService := NewMockEmailService()
	service := NewMotoristaService(repo, emailService)

	// Request base válido - todos os campos preenchidos corretamente
	createValidRequest := func() CadastroMotoristaRequest {
		return CadastroMotoristaRequest{
			Nome:             "João Silva",
			DataNascimento:   "15/03/1990",
			CPF:              "60968336086", // CPF válido
			CNH:              "12345678901",
			CategoriaCNH:     "B",
			ValidadeCNH:      "15/03/2030",
			PlacaVeiculo:     "ABC1234",
			ModeloVeiculo:    "Honda Civic 2020",
			Telefone:         "11999999999",
			Email:            "joao@test.com",
			Senha:            "MinhaSenh@123",
			ConfirmacaoSenha: "MinhaSenh@123",
		}
	}

	tests := []struct {
		name          string
		modifyRequest func(*CadastroMotoristaRequest)
		expectedError string
	}{
		{
			"Nome obrigatório",
			func(r *CadastroMotoristaRequest) { r.Nome = "" },
			"Nome é obrigatório",
		},
		{
			"CPF obrigatório",
			func(r *CadastroMotoristaRequest) { r.CPF = "" },
			"CPF é obrigatório",
		},
		{
			"CNH obrigatória",
			func(r *CadastroMotoristaRequest) { r.CNH = "" },
			"CNH é obrigatória",
		},
		{
			"Email obrigatório",
			func(r *CadastroMotoristaRequest) { r.Email = "" },
			"Email é obrigatório",
		},
		{
			"Senha obrigatória",
			func(r *CadastroMotoristaRequest) { r.Senha = "" },
			"Senha é obrigatória",
		},
		{
			"Telefone obrigatório",
			func(r *CadastroMotoristaRequest) { r.Telefone = "" },
			"Telefone é obrigatório",
		},
		{
			"Placa obrigatória",
			func(r *CadastroMotoristaRequest) { r.PlacaVeiculo = "" },
			"Placa do veículo é obrigatória",
		},
		{
			"CPF inválido",
			func(r *CadastroMotoristaRequest) { r.CPF = "12345678900" },
			"CPF inválido",
		},
		{
			"CNH inválida",
			func(r *CadastroMotoristaRequest) { r.CNH = "123456789" },
			"CNH deve ter 11 dígitos",
		},
		{
			"Email inválido",
			func(r *CadastroMotoristaRequest) { r.Email = "email_invalido" },
			"Formato de email inválido",
		},
		{
			"Telefone inválido",
			func(r *CadastroMotoristaRequest) { r.Telefone = "123456789" },
			"Formato de telefone inválido",
		},
		{
			"Placa inválida",
			func(r *CadastroMotoristaRequest) { r.PlacaVeiculo = "ABC12345" },
			"Formato de placa inválido",
		},
		{
			"Senha fraca",
			func(r *CadastroMotoristaRequest) { r.Senha = "123456" },
			"Senha deve ter pelo menos 8 caracteres",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar uma nova cópia do request válido para cada teste
			request := createValidRequest()
			// Modificar apenas o campo que queremos testar
			tt.modifyRequest(&request)

			err := service.ValidarDadosCadastro(request)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}

func TestVerificarForcaSenha(t *testing.T) {
	repo := NewMockMotoristaRepository()
	emailService := NewMockEmailService()
	service := NewMotoristaService(repo, emailService)

	tests := []struct {
		name          string
		senha         string
		expectedForce string
		expectError   bool
	}{
		{"Senha forte", "MinhaSenh@123", "Forte", false},
		{"Senha média", "MinhaS1!", "Média", false},
		{"Senha fraca", "123456", "Fraca", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			force, err := service.VerificarForcaSenha(tt.senha)
			assert.Equal(t, tt.expectedForce, force)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
