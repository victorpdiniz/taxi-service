package models

import (
    "gorm.io/gorm"
)

type Motorista struct {
    gorm.Model
    Nome                    string         `json:"nome" gorm:"not null"`
    Email                   string         `json:"email" gorm:"uniqueIndex;not null"`
    Status                  StatusMotorista `json:"status" gorm:"not null;default:'disponivel'"`
    NotificacoesHabilitadas bool           `json:"notificacoes_habilitadas" gorm:"default:true"`
}

type StatusMotorista string

const (
    MotoristaDisponivel StatusMotorista = "disponivel"
    MotoristaOcupado    StatusMotorista = "ocupado"
)