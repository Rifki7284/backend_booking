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

type reviewsApi struct {
	reviewService domain.ReviewsService
}

func NewReviewApi(app *fiber.App, reviewsService domain.ReviewsService, middleware fiber.Handler) {
	ra := reviewsApi{
		reviewService: reviewsService,
	}
	Protected := app.Group(
		"/review",
		middleware,
		middlewares.RoleMiddleware("Client"),
	)
	Protected.Post("/create", ra.Create)
}
func (ra reviewsApi) Create(ctx *fiber.Ctx) error {
	b, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.CreateReviewRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation error", fails))
	}

	err := ra.reviewService.Create(b, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
