package routes

import (
	"github.com/gofiber/fiber/v2"

	"taxi-service/controllers"
)

func DummyRoutes(api fiber.Router) {
	dummy := api.Group("/dummy")

	dummy.Get("/", controllers.ListDummyInfo)
	dummy.Get("/:id", controllers.GetDummyInfo)
	dummy.Post("/", controllers.CreateDummyInfo)
	dummy.Put("/:id", controllers.UpdateDummyInfo)
	dummy.Delete("/:id", controllers.DeleteDummyInfo)
}
