package routes

import (
	"os"
	"taxi_service/controllers"
	"taxi_service/repositories"
	"taxi_service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/", logger.New())

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
	})

	SetupDummyRoutes(api)
	SetupMotoristaRoutes(api)
}

func SetupMotoristaRoutes(api fiber.Router) {
	// Inicializar dependências
	motoristaRepo := repositories.NewJSONMotoristaRepository()

	// Escolher serviço de email baseado no ambiente
	var emailService services.EmailService
	if os.Getenv("APP_ENV") == "production" {
		emailService = services.NewSMTPEmailServiceFromEnv()
	} else {
		emailService = services.NewMockEmailService()
	}

	motoristaService := services.NewMotoristaService(motoristaRepo, emailService)
	motoristaController := controllers.NewMotoristaController(motoristaService)

	// Grupo de rotas da API
	apiGroup := api.Group("/api")

	// Rotas de motoristas
	motoristas := apiGroup.Group("/motoristas")

	// Cadastro de motorista
	motoristas.Post("/", motoristaController.CadastrarMotorista)

	// Buscar motorista
	motoristas.Get("/:id", motoristaController.BuscarMotorista)

	// Upload de documentos
	motoristas.Post("/:id/documentos", motoristaController.UploadDocumento)

	// Validar documentos
	motoristas.Post("/:id/validar-documentos", motoristaController.ValidarDocumentos)

	// Aprovar motorista
	motoristas.Put("/:id/aprovar", motoristaController.AprovarMotorista)

	// Rejeitar motorista
	motoristas.Put("/:id/rejeitar", motoristaController.RejeitarMotorista)

	// Utilitários
	motoristas.Post("/verificar-senha", motoristaController.VerificarForcaSenha)
	motoristas.Post("/validar-documento", motoristaController.ValidarDocumentoUpload)
}
