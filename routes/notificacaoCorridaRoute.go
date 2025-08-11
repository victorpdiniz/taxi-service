package routes

import (
    "github.com/gofiber/fiber/v2"
    "taxi-service/controllers"
)

func NotificacaoCorridaRoutes(api fiber.Router) {
    notificacoes := api.Group("/notificacoes")

    // ============= ROTAS CRUD BÁSICAS =============
    // GET /notificacoes - Lista todas as notificações
    notificacoes.Get("/", controllers.ListNotificacoesCorrida)
    
    // GET /notificacoes/:id - Busca notificação por ID
    notificacoes.Get("/:id", controllers.GetNotificacaoCorrida)
    
    // POST /notificacoes - Cria nova notificação
    notificacoes.Post("/", controllers.CreateNotificacaoCorrida)

    // ============= ROTAS DE AÇÕES =============
    // POST /notificacoes/:id/motorista/:motoristaID/accept - Aceitar notificação
    notificacoes.Post("/:id/motorista/:motoristaID/accept", controllers.AceitarNotificacaoCorrida)
    
    // POST /notificacoes/:id/motorista/:motoristaID/refuse - Recusar notificação
    notificacoes.Post("/:id/motorista/:motoristaID/refuse", controllers.RecusarNotificacaoCorrida)

    // ============= ROTAS ADMINISTRATIVAS =============
    // PUT /notificacoes/:id/status - Atualizar status (auxiliar)
    notificacoes.Put("/:id/status", controllers.UpdateNotificacaoStatus)
}