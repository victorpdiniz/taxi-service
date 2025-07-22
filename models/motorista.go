package models

import (
    "time"
)

type Motorista struct {
    ID                      uint      `json:"id"`
    Nome                    string    `json:"nome"`
    Email                   string    `json:"email"`
    Telefone                string    `json:"telefone"`
    Status                  string    `json:"status"`
    NotificacoesHabilitadas bool      `json:"notificacoes_habilitadas"`
    CreatedAt               time.Time `json:"created_at"`
    UpdatedAt               time.Time `json:"updated_at"`
}

const (
    MotoristaDisponivel = "disponivel"
    MotoristaOcupado    = "ocupado"
)