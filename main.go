package main

import (
	"fmt"
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"shellrean.id/back-end/internal/api"
	"shellrean.id/back-end/internal/config"
	"shellrean.id/back-end/internal/connection"
	"shellrean.id/back-end/internal/repository"
	"shellrean.id/back-end/internal/service"
)

func main() {
	cnf := config.Get()
	db := connection.GetDatabase(cnf.Database)
	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{
			Key: []byte(cnf.JWT.Key),
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// LOG KE CONSOLE
			fmt.Println("JWT ERROR =>", err)

			// KIRIM KE RESPONSE (sementara, untuk debug)
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	bookingRepository := repository.NewBooking(db)
	bookingService := service.NewBookingService(bookingRepository)
	userRepository := repository.NewUser(db)
	userService := service.NewAuthService(cnf, userRepository)
	roomRepository := repository.NewRoom(db)
	roomService := service.NewRoomService(roomRepository)
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
	_ = bookingService

	api.NewBookingApi(app, bookingService, jwtMidd)
	api.NewAuthApi(app, userService)
	api.NewRoomApi(app, roomService, jwtMidd)
	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		panic(err)
	}
}
