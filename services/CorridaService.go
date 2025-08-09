package services

import (
	"fmt"
	"math/rand"
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

// CriarNovaCorrida cria uma nova corrida e inicia a simulação.
func (s *CorridaService) CriarNovaCorrida(corridaInput models.Corrida) (*models.Corrida, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	corrida := &corridaInput
	corrida.ID = s.nextID
	s.nextID++
	corrida.Status = models.StatusProcurandoMotorista
	corrida.DataInicio = time.Now()

	s.corridas[corrida.ID] = corrida

	// Inicia a simulação em uma nova goroutine
	go s.simularCorrida(corrida)

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

// simularCorrida executa a lógica de simulação de uma corrida.
func (s *CorridaService) simularCorrida(corrida *models.Corrida) {
	// 1. Procurando motorista
	fmt.Printf("Corrida %d: Procurando motorista...\n", corrida.ID)
	time.Sleep(5 * time.Second) // Simula o tempo de busca

	s.mutex.Lock()
	corrida.Status = models.StatusMotoristaEncontrado
	corrida.MotoristaID = rand.Intn(100) + 1 // Atribui um ID de motorista aleatório
	fmt.Printf("Corrida %d: Motorista %d encontrado. A caminho!\n", corrida.ID, corrida.MotoristaID)
	s.mutex.Unlock()

	// 2. Motorista a caminho
	time.Sleep(5 * time.Second) // Simula o tempo de chegada do motorista

	s.mutex.Lock()
	corrida.Status = models.StatusCorridaIniciada
	fmt.Printf("Corrida %d: Corrida iniciada.\n", corrida.ID)
	s.mutex.Unlock()

	// 3. Corrida em andamento
	// Usaremos o tempo estimado em segundos para a simulação
	tempoEstimadoSegundos := time.Duration(corrida.TempoEstimado) * time.Second
	tempoInicioCorrida := time.Now()
	
	for time.Since(tempoInicioCorrida) < tempoEstimadoSegundos {
		time.Sleep(1 * time.Second)
		s.mutex.Lock()
		corrida.TempoDecorrido = int(time.Since(tempoInicioCorrida).Seconds())
		s.mutex.Unlock()
	}


	// 4. Finalizar a corrida
	s.mutex.Lock()
	now := time.Now()
	corrida.DataFim = &now
	corrida.TempoDecorrido = int(tempoEstimadoSegundos.Seconds())
	// Lógica de bônus/status final pode ser adicionada aqui se necessário
	corrida.Status = models.StatusConcluidaNoTempo
	fmt.Printf("Corrida %d: Corrida finalizada.\n", corrida.ID)
	s.mutex.Unlock()
}