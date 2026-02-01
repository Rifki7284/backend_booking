package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
	"shellrean.id/back-end/internal/middlewares"
	"shellrean.id/back-end/internal/util"
)

type bookingApi struct {
	bookingService domain.BookingService
}

func NewBookingApi(app *fiber.App, bookingService domain.BookingService, middleware fiber.Handler) {
	ba := bookingApi{
		bookingService: bookingService,
	}
	Admin := app.Group(
		"/admin/booking",
		middleware,
		middlewares.RoleMiddleware("Admin"),
	)
	Admin.Get("/", ba.Index)
	Protected := app.Group(
		"/booking",
		middleware,
		middlewares.RoleMiddleware("Client"),
	)
	Protected.Get("/detail/:id", ba.Show)
	Protected.Post("/create", ba.Create)
	Protected.Put("/:id", ba.Update)
}
func (ba bookingApi) Index(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	res, error := ba.bookingService.Index(b)
	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(error.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}
func (ba bookingApi) Create(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	claim := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	id := claim["id"].(string)
	var req dto.CreateBookingRequest
	req.UserID = id
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}
	err := ba.bookingService.Create(b, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
func (ba bookingApi) Update(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	claim := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	id := claim["id"].(string)
	defer cancel()
	var req dto.UpdateBookingRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}
	req.ID = ctx.Params("id")
	err := ba.bookingService.Update(b, req, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
func (ba bookingApi) Show(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	id := ctx.Params("id")
	data, err := ba.bookingService.Show(b, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(data))
}
