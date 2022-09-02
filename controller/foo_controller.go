package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-rest/service"
	"net/http"
)

type FooController struct {
	fooService service.FooService
}

func NewFooController(fooService service.FooService) *FooController {
	return &FooController{fooService: fooService}
}

// HealthCheck godoc
// @Summary FooController controller.
// @Description FooController will return a new foo object.
// @Tags FooController
// @Accept */*
// @Produce json
// @Success 200 {object} model.Foo
// @Router /api/v1/foo [get]
func (c FooController) foo(fiberCtx *fiber.Ctx) error {
	foo, err := c.fooService.Foo(fiberCtx.Body())
	if err != nil {
		return respondWithError(fiberCtx, http.StatusConflict, "message %v", "topic.Topic")
	}
	return respondWithJson(fiberCtx, http.StatusOK, foo)
}
