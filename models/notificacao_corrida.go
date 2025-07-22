package models

import (
    "time"
    "gorm.io/gorm"
)

type NotificacaoCorrida struct {
    gorm.Model
    MotoristaID    uint                     `json:"motorista_id" gorm:"not null;index"`
    Motorista      Motorista                `json:"motorista" gorm:"foreignKey:MotoristaID"`
    CorridaID      uint                     `json:"corrida_id" gorm:"not null;index"`
    Corrida        Corrida                  `json:"corrida" gorm:"foreignKey:CorridaID"`
    Origem         string                   `json:"origem" gorm:"not null"`
    Destino        string                   `json:"destino" gorm:"not null"`
    NomePassageiro string                   `json:"nome_passageiro" gorm:"not null"`
    Valor          float64                  `json:"valor" gorm:"not null"`
    ExpiraEm       time.Time                `json:"expira_em" gorm:"not null;index"`
    Status         StatusNotificacaoCorrida `json:"status" gorm:"not null;default:'pendente';index"`
}

type StatusNotificacaoCorrida string

const (
    NotificacaoPendente  StatusNotificacaoCorrida = "pendente"
    NotificacaoAceita    StatusNotificacaoCorrida = "aceita"
    NotificacaoRecusada  StatusNotificacaoCorrida = "recusada"
    NotificacaoExpirada  StatusNotificacaoCorrida = "expirada"
)