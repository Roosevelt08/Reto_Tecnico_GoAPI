package main

import (
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New()) // Middleware para registrar todas las solicitudes

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:8081", // Ajusta esto a tu configuración de frontend
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// 	AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
	// }))

	// GET básica
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Obtener solicitudes")
	})
	// POST para rotar matrices
	app.Post("/rotate", func(c *fiber.Ctx) error {
		var data map[string][][]int
		if err := c.BodyParser(&data); err != nil {
			log.Println("Error en JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se puede analizar la matriz"})
		}
		matrix := data["matrix"]
		rotatedMatrix := rotateMatrix(matrix)
		return c.JSON(rotatedMatrix)
	})

	log.Fatal(app.Listen(":8080"))
}

// POST para rotar matrices
func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("rotar matrices")
	})

	app.Post("/rotate", func(c *fiber.Ctx) error {
		var matrix [][]int
		if err := c.BodyParser(&matrix); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se puede analizar la matriz"})
		}

		rotatedMatrix := rotateMatrix(matrix)
		return c.JSON(rotatedMatrix)
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
	// Setup
	app := fiber.New()
	setUpRoutes(app) // Asume que tienes una función que configura las rutas

	// Realizar solicitud
	req, _ := http.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	// Verificar la respuesta
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
