package services

import (
	"fmt"
	"sync"
	"time"
	"your-app/models"
)

// CorridaService gerencia a lógica de negócio das corridas.
type CorridaService struct {
	corridas map[int]*models.Corrida
	mutex    sync.RWMutex
	nextID   int
}

// NewCorridaService cria uma nova instância de CorridaService.
func NewCorridaService() *CorridaService {
	return &CorridaService{
		corridas: make(map[int]*models.Corrida),
		nextID:   1,
	}
}

// CriarNovaCorrida cria uma nova corrida e a prepara para ser aceita.
func (s *CorridaService) CriarNovaCorrida(corridaInput models.Corrida) (*models.Corrida, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	corrida := &corridaInput
	corrida.ID = s.nextID
	s.nextID++
	corrida.Status = models.StatusProcurandoMotorista
	corrida.DataInicio = time.Now()

	s.corridas[corrida.ID] = corrida

	return corrida, nil
}

// GetCorridaPorID busca uma corrida pelo seu ID.
func (s *CorridaService) GetCorridaPorID(id int) (*models.Corrida, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	corrida, exists := s.corridas[id]
	if !exists {
		return nil, fmt.Errorf("corrida com ID %d não encontrada", id)
	}
	return corrida, nil
}

// AceitarCorrida permite que um motorista aceite uma corrida.
func (s *CorridaService) AceitarCorrida(corridaID int, motoristaID int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	corrida, exists := s.corridas[corridaID]
	if !exists {
		return fmt.Errorf("corrida com ID %d não encontrada", corridaID)
	}

	if corrida.Status != models.StatusProcurandoMotorista {
		return fmt.Errorf("corrida %d não está mais procurando por motorista", corridaID)
	}

	corrida.Status = models.StatusMotoristaEncontrado
	corrida.MotoristaID = motoristaID
	fmt.Printf("Corrida %d: Motorista %d aceitou a corrida.\n", corrida.ID, corrida.MotoristaID)

	return nil
}

// AtualizarPosicao atualiza a localização do motorista para uma corrida específica.
func (s *CorridaService) AtualizarPosicao(corridaID int, lat, lng float64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	corrida, exists := s.corridas[corridaID]
	if !exists {
		return fmt.Errorf("corrida com ID %d não encontrada", corridaID)
	}

	corrida.MotoristaLat = lat
	corrida.MotoristaLng = lng
	return nil
}

// CancelarCorrida cancela uma corrida que está em andamento.
func (s *CorridaService) CancelarCorrida(corridaID int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	corrida, exists := s.corridas[corridaID]
	if !exists {
		return fmt.Errorf("corrida com ID %d não encontrada", corridaID)
	}

	corrida.Status = models.StatusCanceladaPeloUsuario
	now := time.Now()
    corrida.DataFim = &now
	fmt.Printf("Corrida %d: Cancelada pelo usuário.\n", corrida.ID)

	return nil
}

// FinalizarCorrida finaliza uma corrida com sucesso.
func (s *CorridaService) FinalizarCorrida(corridaID int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	corrida, exists := s.corridas[corridaID]
	if !exists {
		return fmt.Errorf("corrida com ID %d não encontrada", corridaID)
	}

	corrida.Status = models.StatusConcluidaNoTempo // Ou outra lógica para determinar o status final
	now := time.Now()
    corrida.DataFim = &now
	fmt.Printf("Corrida %d: Finalizada com sucesso.\n", corrida.ID)

	return nil
}
