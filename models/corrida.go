package models

import "time"

// CorridaStatus representa o status de uma corrida
const (
	StatusProcurandoMotorista      = "procurando_motorista"
	StatusMotoristaEncontrado      = "motorista_encontrado"
	StatusCorridaIniciada          = "corrida_iniciada"
	StatusEmAndamento              = "em_andamento" // Mantido para compatibilidade
	StatusAtrasado                 = "atrasado"
	StatusConcluidaAntecedencia    = "concluída com antecedência"
	StatusConcluidaNoTempo         = "concluída no tempo previsto"
	StatusCanceladaPorExcessoTempo = "cancelada por excesso de tempo"
	StatusCanceladaPeloUsuario     = "cancelada pelo usuário"
)

type Corrida struct {
	ID             int
	MotoristaID    int
	PassageiroID   int
	Origem         string
	Destino        string
	TempoEstimado  int // minutos
	TempoDecorrido int // minutos
	Preco          float64 // valor da corrida
	Status         string
	BonusAplicado  bool
	DataInicio     time.Time
	DataFim        *time.Time
	Avaliacao	   *int
	Destino		   string
	MotoristaLat   float64
	MotoristaLng   float64
}
