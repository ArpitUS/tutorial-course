package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "time"
)

func main() {
    app := fiber.New()

    app.Get("/metrics", func(c *fiber.Ctx) error {
        now := time.Now().Format(time.RFC3339)
        fmt.Println("Monitoring check at:", now)
        return c.JSON(fiber.Map{
            "service": "mlops-monitor",
            "status":  "healthy",
            "timestamp": now,
        })
    })

    app.Listen(":4000")
}
