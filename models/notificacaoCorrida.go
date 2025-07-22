package models

import (
    "time"
)

type NotificacaoCorrida struct {
    ID             uint      `json:"id"`
    MotoristaID    uint      `json:"motorista_id"`
    CorridaID      uint      `json:"corrida_id"`
    Origem         string    `json:"origem"`
    Destino        string    `json:"destino"`
    NomePassageiro string    `json:"nome_passageiro"`
    Valor          float64   `json:"valor"`
    ExpiraEm       time.Time `json:"expira_em"`
    Status         string    `json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

const (
    NotificacaoPendente  = "pendente"
    NotificacaoAceita    = "aceita"
    NotificacaoRecusada  = "recusada"
    NotificacaoExpirada  = "expirada"
)