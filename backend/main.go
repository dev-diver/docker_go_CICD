package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// 모든 서버에서의 요청을 허용하는 CORS 설정
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Static("/", "../dist")

	app.Get("/test", func(c *fiber.Ctx) error {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file")
			panic("Error loading .env file") // 패닉 발생
		}

		envVars := make(map[string]string)
		for _, env := range os.Environ() {
			pair := string(env)
			envVars[pair[0:4]] = pair[5:]
		}

		return c.JSON(envVars)
	})

	// 5001 포트에서 애플리케이션 실행
	app.Listen(":5001")
}
