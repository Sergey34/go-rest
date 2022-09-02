package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-rest/model"
	"go-rest/service"
	"net/http"
)

type BarController struct {
	barService service.BarService
}

func NewBarController(barService service.BarService) *BarController {
	return &BarController{barService: barService}
}

func (c BarController) bar(fiberCtx *fiber.Ctx) error {
	var bar model.Bar
	if err := json.Unmarshal(fiberCtx.Body(), &bar); err != nil {
		return respondWithError(fiberCtx, http.StatusConflict, "message %v", err)
	}
	result, err := c.barService.Bar(bar)
	if err != nil {
		return respondWithError(fiberCtx, http.StatusConflict, "message %v", "topic.Topic")
	}
	return respondWithJson(fiberCtx, http.StatusCreated, result)
}
