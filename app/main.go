package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// 모든 서버에서의 요청을 허용하는 CORS 설정
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	type MessageRequest struct {
		Message string `json:"message" validate:"required"`
	}

	type MessageResponse struct {
		Answer string `json:"answer"`
	}

	var validate = validator.New()

	// /echo 엔드포인트 설정
	app.Post("/echo", func(c *fiber.Ctx) error {
		var request MessageRequest
		if err := c.BodyParser(&request); err != nil {
			log.Println("Error parsing JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		if err := validate.Struct(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "message field is required",
			})
		}

		response := MessageResponse{
			Answer: request.Message,
		}

		return c.JSON(response)
	})

	// 5001 포트에서 애플리케이션 실행
	app.Listen(":5001")
}
