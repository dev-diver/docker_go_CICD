package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

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

	app.Post("/remove", func(c *fiber.Ctx) error {
		if err := removeClientWithSocket(); err != nil {
			return err
		}
		return c.SendString("Client container restarted successfully")
	})
	// 5001 포트에서 애플리케이션 실행
	app.Listen(":5001")
}

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	log.Printf("Running command: %v %v", command, strings.Join(args, " "))

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	return shell
}

func removeClientWithSocket() error {
	containerName := "frontend"

	if err := execCommand("docker", "compose", "rm", "-f", containerName); err != nil {
		log.Printf("Failed to remove client container: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	return nil
}
