package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-rest/controller"
	"go-rest/service"
	"os"
	"os/signal"
	"syscall"
)

var (
	ctx = context.WithValue(context.Background(), "requestId", "main")
)

func main() {
	barService := service.NewBarService()
	fooService := service.NewFooService()

	controllers := controller.ApiControllers{
		FooController:     controller.NewFooController(fooService),
		BarController:     controller.NewBarController(barService),
		GeneralController: controller.NewGeneralController(),
	}

	app := controller.CreateApi(ctx, controllers)

	registerShutdownHook(app)

	serverBind := ":8080"
	if err := app.Listen(serverBind); err != nil {
		println(err.Error())
	}

	println("MaaS server gracefully finished")
}

func registerShutdownHook(app *fiber.App) {
	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		// We received an interrupt signal, shut down.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			println("MaaS error during server shutdown: %v", err)
		}
	}()
}
