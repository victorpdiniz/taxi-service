package models

import "time"

// CorridaStatus representa o status de uma corrida
const (
	StatusEmAndamento              = "em_andamento"
	StatusAtrasado                 = "atrasado"
	StatusConcluidaAntecedencia    = "concluída com antecedência"
	StatusConcluidaNoTempo         = "concluída no tempo previsto"
	StatusCanceladaPorExcessoTempo = "cancelada por excesso de tempo"
)

type Corrida struct {
	ID             int
	MotoristaID    int
	PassageiroID   int
	TempoEstimado  int // minutos
	TempoDecorrido int // minutos
	Preco          float64 // valor da corrida
	Status         string
	BonusAplicado  bool
	DataInicio     time.Time
	DataFim        *time.Time
}