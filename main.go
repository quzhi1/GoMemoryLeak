package main

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/pprof"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(pprof.New())

	// Health
	app.Get("/", health)

	// Define endpoints
	app.Get("/append-slice", appendSlice)
	app.Get("/hanging-go-routine", hangingGoRoutine)

	// Start
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

func health(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"status": "healthy",
	})
}

var globalSlice = make([]int64, 0)

func appendSlice(c *fiber.Ctx) error {
	globalSlice = append(globalSlice, time.Now().Unix())
	return c.JSON(map[string]int{
		"sliceSize": len(globalSlice),
	})
}

func hangingGoRoutine(c *fiber.Ctx) error {
	go time.Sleep(time.Hour * 24)
	return c.SendString("Created a sleeper")
}
