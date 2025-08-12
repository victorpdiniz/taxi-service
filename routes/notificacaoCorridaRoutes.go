package routes

import (
    "github.com/gofiber/fiber/v2"
    "taxi-service/controllers"
)

func NotificacaoCorridaRoutes(api fiber.Router) {
    notificacoes := api.Group("/notificacoes")

    // ============= ROTAS CRUD BÁSICAS =============
    notificacoes.Get("/", controllers.ListNotificacoesCorrida)
    notificacoes.Get("/:id", controllers.GetNotificacaoCorrida)
    notificacoes.Post("/", controllers.CreateNotificacaoCorrida)
    
    // ============= ROTAS ESPECÍFICAS DE MOTORISTA =============
    notificacoes.Get("/motorista/:motoristaID/pending", controllers.GetNotificacoesPendentesParaMotorista)

    // ============= ROTAS DE AÇÕES =============
    notificacoes.Post("/:id/motorista/:motoristaID/accept", controllers.AceitarNotificacaoCorrida)
    notificacoes.Post("/:id/motorista/:motoristaID/refuse", controllers.RecusarNotificacaoCorrida)

    // ============= ROTAS ADMINISTRATIVAS =============
    notificacoes.Post("/expire", controllers.ExpirarNotificacoesVencidas)
    notificacoes.Put("/:id/status", controllers.UpdateNotificacaoStatus)
}