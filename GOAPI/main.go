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
		AllowOrigins: "http://localhost:8081", // Permitir solicitudes desde el frontend
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New()) // Middleware para registrar todas las solicitudes

	// Configurar rutas
	setUpRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

// Configuración de rutas
func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("rotar matrices")
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
		response := fiber.Map{
			"rotatedMatrix": rotatedMatrix,
			"statistics":    stats,
		}

		return c.JSON(response)
	})
}

// Función para rotar la matriz
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
			rotated[i][j] = original[n-j-1][i] // Transpone y luego invierte las filas
		}
	}

	return rotated
}

// Función para calcular estadísticas
func calculateStatistics(matrix [][]int) map[string]interface{} {
	maxVal := matrix[0][0]
	minVal := matrix[0][0]
	sum := 0
	count := 0
	Diagonal := true

	for i := range matrix {
		for j := range matrix[i] {
			val := matrix[i][j]
			if val > maxVal {
				maxVal = val
			}
			if val < minVal {
				minVal = val
			}
			sum += val
			count++
			if i != j && val != 0 {
				Diagonal = false
			}
		}
	}

	promedio := float64(sum) / float64(count)
	stats := map[string]interface{}{
		"max":      maxVal,
		"min":      minVal,
		"sum":      sum,
		"promedio": promedio,
		"diagonal": Diagonal,
	}

	return stats
}

// Prueba de la función de rotación de matrices
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

// Prueba de la API
func TestAPI(t *testing.T) {
	// Setup
	app := fiber.New()
	setUpRoutes(app) // Configurar rutas

	// Realizar solicitud
	req, _ := http.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	// Verificar la respuesta
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
