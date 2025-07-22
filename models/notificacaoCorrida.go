package models

import (
    "time"
)

type NotificacaoStatus string

const (
    NotificacaoPendente  NotificacaoStatus = "pendente"
    NotificacaoAceita    NotificacaoStatus = "aceita"
    NotificacaoRecusada  NotificacaoStatus = "recusada"
    NotificacaoExpirada  NotificacaoStatus = "expirada"
)

type NotificacaoCorrida struct {
    ID              uint              `json:"id"`
    MotoristaID     uint              `json:"motorista_id"`
    CorridaID       uint              `json:"corrida_id"`
    PassageiroNome  string            `json:"passageiro_nome"`  // Corrigido de NomePassageiro
    Valor           float64           `json:"valor"`
    DistanciaKm     float64           `json:"distancia_km"`     // Adicionado
    TempoEstimado   string            `json:"tempo_estimado"`   // Adicionado
    Origem          string            `json:"origem"`
    Destino         string            `json:"destino"`
    Status          NotificacaoStatus `json:"status"`           // Tipado corretamente
    CreatedAt       time.Time         `json:"created_at"`
    UpdatedAt       time.Time         `json:"updated_at"`
    ExpiraEm        time.Time         `json:"expira_em"`
}