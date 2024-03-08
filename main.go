package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	app := fiber.New()

	app.Use(logger.New())

	loginLimiter := limiter.New(limiter.Config{
		Max:        3,
		Expiration: 30 * time.Second,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("selamat datang!")
	})

	app.Post("/login", loginLimiter, func(c *fiber.Ctx) error {
		var req Login

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid request",
			})
		}

		if req.Username == "ferizco" && req.Password == "12345" {
			return c.JSON(fiber.Map{
				"message": "login successfull",
			})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid username or password",
			})
		}
	})

	app.Listen(":3000")
}
