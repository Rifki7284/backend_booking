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

type roomApi struct {
	roomService domain.RoomService
}

func NewRoomApi(app *fiber.App, roomService domain.RoomService, middleware fiber.Handler) {
	ra := roomApi{
		roomService: roomService,
	}
	client := app.Group(
		"/rooms",
		middleware,
		middlewares.RoleMiddleware("Client"),
	)

	client.Get("/", ra.Index)
	client.Post("/create", ra.Create)
	client.Put("/:id", ra.Update)
}
func (ra roomApi) Index(ctx *fiber.Ctx) error {
	r, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, error := ra.roomService.Index(r)
	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(error.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}
func (ra roomApi) Create(ctx *fiber.Ctx) error {
	r, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.CreateRoomRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}
	err := ra.roomService.Create(r, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
func (ra roomApi) Update(ctx *fiber.Ctx) error {
	r, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.UpdateRoomRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}
	req.ID = ctx.Params("id")
	err := ra.roomService.Update(r, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
