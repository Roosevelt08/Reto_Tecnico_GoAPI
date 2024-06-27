package main

import (
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Obtener solicitudes")
	})

	app.Post("/rotate", func(c *fiber.Ctx) error {
		var data map[string][][]int
		if err := c.BodyParser(&data); err != nil {
			log.Println("Error en JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se puede analizar la matriz"})
		}

		matrix, exists := data["matrix"]
		if !exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se encuentra la clave matrix"})
		}

		rotatedMatrix := rotateMatrix(matrix)
		stats := calculateStatistics(matrix)

		response := map[string]interface{}{
			"rotatedMatrix": rotatedMatrix,
			"statistics":    stats,
		}

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":8080"))
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("rotar matrices")
	})
}

func rotateMatrix(original [][]int) [][]int {
	if len(original) == 0 {
		return original
	}

	n := len(original)
	m := len(original[0])
	rotated := make([][]int, m)

	for i := 0; i < m; i++ {
		rotated[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = original[n-j-1][i]
		}
	}

	return rotated
}

func calculateStatistics(matrix [][]int) map[string]interface{} {
	max, min, sum := matrix[0][0], matrix[0][0], 0
	isDiagonal := true
	totalElements := 0

	for i := range matrix {
		for j := range matrix[i] {
			value := matrix[i][j]
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
			sum += value
			totalElements++
			if i != j && value != 0 {
				isDiagonal = false
			}
		}
	}

	average := float64(sum) / float64(totalElements)

	return map[string]interface{}{
		"max":      max,
		"min":      min,
		"average":  average,
		"sum":      sum,
		"diagonal": isDiagonal,
	}
}

func TestRotateMatrix(t *testing.T) {
	original := [][]int{
		{1, 2},
		{3, 4},
	}
	expected := [][]int{
		{3, 1},
		{4, 2},
	}
	rotated := rotateMatrix(original)
	if !reflect.DeepEqual(rotated, expected) {
		t.Errorf("Expected %v, got %v", expected, rotated)
	}
}

func TestAPI(t *testing.T) {
	app := fiber.New()
	setUpRoutes(app)

	req, _ := http.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
