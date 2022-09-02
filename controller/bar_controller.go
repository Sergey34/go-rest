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

// HealthCheck godoc
// @Summary BarController controller.
// @Description BarController will update your bar request.
// @Tags BarController
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/bar [post]
// @Param request body model.Bar true "query params"
func (c BarController) bar(fiberCtx *fiber.Ctx) error {
	var bar model.Bar
	if err := json.Unmarshal(fiberCtx.Body(), &bar); err != nil {
		return respondWithError(fiberCtx, http.StatusConflict, "message %v", err)
	}
	result, err := c.barService.Bar(bar)
	if err != nil {
		return respondWithError(fiberCtx, http.StatusConflict, "message %v", "topic.Topic")
	}
	return respondWithJson(fiberCtx, http.StatusOK, result)
}
