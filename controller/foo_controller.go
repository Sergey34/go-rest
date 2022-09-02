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

func (c FooController) foo(fiberCtx *fiber.Ctx) error {
	foo, err := c.fooService.Foo(fiberCtx.Body())
	if err != nil {
		return respondWithError(fiberCtx, http.StatusConflict, "message %v", "topic.Topic")
	}
	return respondWithJson(fiberCtx, http.StatusCreated, foo)
}
