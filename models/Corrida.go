package models

import (
    "time"
    "gorm.io/gorm"
)

type Corrida struct {
    gorm.Model
    ID                   int        `json:"id"`
    Data                 string     `json:"data"`                 // dia da corrida
    Horario              time.Time  `json:"horario"`              // horário de inicio 
    Tempo                int        `json:"tempo"`                // tempo para chegar ao destino
    TempoEstimado        int        `json:"tempoEstimado"`        // tempo estimado em minutos
    TempoDecorrido       int        `json:"tempoDecorrido"`       // tempo decorrido em minutos
    Valor                int        `json:"valor"`                // valor da corrida (original)
    Preco                float64    `json:"preco"`                // valor da corrida (float64)
    Avaliacao            *int       `json:"avaliacao"`            // avaliacao 1, 2, 3, 4, 5 ou nil
    Status               string     `json:"status"`               // andamento, finalizada ou cancelada
    CPFMotorista         *int       `json:"cpfMotorista"`         // chave estrangeira pro motorista responsavel (legacy)
    MotoristaID          int        `json:"motoristaID"`          // ID do motorista
    PassageiroID         int        `json:"passageiroID"`         // ID do passageiro
    LocalDesembarque     string     `json:"localDesembarque"`     // local de desembarque
    BonusAplicado        bool       `json:"bonusAplicado"`        // se bonus foi aplicado
    DataInicio           time.Time  `json:"dataInicio"`           // data/hora de início
    DataFim              *time.Time `json:"dataFim"`              // data/hora de fim (pode ser nil)
}

const (
    // Status originais (mantidos para compatibilidade)
    StatusAndamento  = "andamento"
    StatusCancelada  = "cancelada"
    StatusFinalizada = "finalizada"
    
    // Status estendidos (incorporados do segundo modelo)
    StatusEmAndamento              = "em_andamento"
    StatusAtrasado                 = "atrasado"
    StatusConcluidaAntecedencia    = "concluída com antecedência"
    StatusConcluidaNoTempo         = "concluída no tempo previsto"
    StatusCanceladaPorExcessoTempo = "cancelada por excesso de tempo"
)