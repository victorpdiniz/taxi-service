package models

import (
    "time"
    "gorm.io/gorm"
)

type Corrida struct {
    gorm.Model
    PassageiroID string         `json:"passageiro_id" gorm:"not null"`
    MotoristaID  *string        `json:"motorista_id"`
    Origem       string         `json:"origem" gorm:"not null"`
    Destino      string         `json:"destino" gorm:"not null"`
    Valor        float64        `json:"valor" gorm:"not null"`
    Status       StatusCorrida  `json:"status" gorm:"not null;default:'solicitada'"`
    SolicitadaEm time.Time      `json:"solicitada_em" gorm:"not null"`
    AceitaEm     *time.Time     `json:"aceita_em"`
    IniciadaEm   *time.Time     `json:"iniciada_em"`
    FinalizadaEm *time.Time     `json:"finalizada_em"`
    CanceladaEm  *time.Time     `json:"cancelada_em"`
    DuracaoMinutos *int         `json:"duracao_minutos"` // Duração em minutos
}

type StatusCorrida string

const (
    CorridaSolicitada   StatusCorrida = "solicitada"
    CorridaAceita       StatusCorrida = "aceita"
    CorridaEmAndamento  StatusCorrida = "em_andamento"
    CorridaFinalizada   StatusCorrida = "finalizada"
    CorridaCancelada    StatusCorrida = "cancelada"
)