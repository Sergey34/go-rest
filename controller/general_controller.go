package controller

import (
	"github.com/gofiber/fiber/v2"
	api_version "go-rest/api-version"
	"net/http"
)

type GeneralController struct {
}

func NewGeneralController() *GeneralController {
	return &GeneralController{}
}

func (g *GeneralController) ApiVersion(ctx *fiber.Ctx) error {
	return respondWithJson(ctx, http.StatusOK, api_version.ApiVersion)
}
