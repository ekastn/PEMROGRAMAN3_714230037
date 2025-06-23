package router

import (
	"w4backend/config/middleware"
	"w4backend/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	api := app.Group("/api", middleware.Middlewares("admin"))

	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa", handler.GetAllMahasiswa)

	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)
	api.Post("/mahasiswa", handler.InsertMahasiswa)
	api.Put("/mahasiswa/:npm", handler.UpdateMahasiswa)
	api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa)
}
