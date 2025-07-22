package models

import (
	"time"
)

type Corrida struct {
	ID               int       `json:"id"`
	Data             string    `json:"data"`         //dia da corrida
	Horario          time.Time `json:"horario"`      // horário de inicio
	Tempo            int       `json:"tempo"`        // tempo para chegar ao destino
	Valor            int       `json:"valor"`        // valor da corrida
	Avaliacao        *int      `json:"avaliacao"`    // avaliacao 1, 2, 3, 4, 5 ou nil, * permite nil
	Status           string    `json:"status"`       // andamento, finalizada ou cancelada
	CPFMotorista     *int      `json:"cpfMotorista"` // chave estrangeira pro motorista responsavel
	LocalDesembarque string    `json:"localDesembarque"`
	MotoristaID      int
	PassageiroID     int
	Preco            float64 // valor da corrida
	BonusAplicado    bool
	DataInicio       time.Time
	DataFim          *time.Time
}

const (
	StatusEmAndamento              = "em_andamento"
	StatusAtrasado                 = "atrasado"
	StatusConcluidaAntecedencia    = "concluída com antecedência"
	StatusConcluidaNoTempo         = "concluída no tempo previsto"
	StatusCanceladaPorExcessoTempo = "cancelada por excesso de tempo"
)