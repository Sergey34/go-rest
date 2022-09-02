package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

type ApiControllers struct {
	FooController     *FooController
	BarController     *BarController
	GeneralController *GeneralController
}

func CreateApi(ctx context.Context, controllers ApiControllers) *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: 30 * time.Second,
	})
	app.Get("/api-version", controllers.GeneralController.ApiVersion)

	apiV1 := app.Group("/api/v1")

	apiV1.Get("/foo", controllers.FooController.foo)

	apiV1.Post("/bar", controllers.BarController.bar)

	return app
}
