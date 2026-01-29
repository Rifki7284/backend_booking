package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
	"shellrean.id/back-end/internal/middlewares"
	"shellrean.id/back-end/internal/util"
)

type propertiesApi struct {
	propertiesService domain.PropertiesService
}

func NewPropertiesApi(app *fiber.App, propertiesService domain.PropertiesService, middleware fiber.Handler) {
	pa := propertiesApi{
		propertiesService: propertiesService,
	}
	Owner := app.Group(
		"/properties",
		middleware,
		middlewares.RoleMiddleware("Owner"),
	)
	Owner.Get("/", pa.Index)
	Owner.Get("/detail/:id", pa.Show)
	Owner.Post("/create", pa.Create)
	Owner.Put("/:id", pa.Update)
	Owner.Delete("/:id", pa.Delete)
}
func (pa propertiesApi) Index(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, error := pa.propertiesService.Index(b)
	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(error.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}
func (pa propertiesApi) Create(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.CreatePropertiesRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}
	err := pa.propertiesService.Create(b, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
func (pa propertiesApi) Update(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.UpdatePropertiesRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}
	req.ID = ctx.Params("id")
	err := pa.propertiesService.Update(b, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
func (pa propertiesApi) Delete(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	id := ctx.Params("id")
	err := pa.propertiesService.Delete(b, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.SendStatus(http.StatusNoContent)
}
func (pa propertiesApi) Show(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	id := ctx.Params("id")
	data, err := pa.propertiesService.Show(b, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(data))
}
