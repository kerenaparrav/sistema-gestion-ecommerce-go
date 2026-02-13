package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"sistema-gestion-ecommerce-go/internal/domain"
	"sistema-gestion-ecommerce-go/internal/repository"
	"sistema-gestion-ecommerce-go/internal/service"
)

// Este archivo actúa como capa de presentación (CLI).
// No contiene lógica de negocio ni acceso directo a datos,
// delegando estas responsabilidades a las capas service y repository.

var reader = bufio.NewReader(os.Stdin)

// readLine lee una línea completa desde la consola.
// Se utiliza para permitir entradas con espacios, como nombres de productos.
func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// readFloat solicita un número decimal al usuario.
// Incluye validación para evitar valores no numéricos o negativos.
func readFloat(prompt string) float64 {
	for {
		valueStr := readLine(prompt)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("❌ Error: ingresa un número válido (ej: 2.50).")
			continue
		}
		if value < 0 {
			fmt.Println("❌ Error: el precio no puede ser negativo.")
			continue
		}
		return value
	}
}

// readInt solicita un número entero al usuario.
// Se utiliza para validar entradas como ID u opciones del menú.
func readInt(prompt string) int {
	for {
		valueStr := readLine(prompt)
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("❌ Error: ingresa un número entero válido (ej: 1).")
			continue
		}
		return value
	}
}

func main() {
	// ✅ Encapsulación + Interfaces:
	// main NO guarda productos directamente, solo usa el servicio.
	repo := repository.NewInMemoryProductRepository()
	productService := service.NewProductService(repo)

	// Menú principal del sistema.
	// Controla la interacción con el usuario mediante consola
	// y delega la lógica de negocio a la capa de servicios.
	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("1. Registrar producto")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Salir")

		opcion := readInt("Seleccione una opción: ")

		switch opcion {
		// Se utiliza un switch para controlar las opciones del menú
		// y mantener el flujo del programa organizado.

		case 1:
			fmt.Println("\n--- Registrar producto ---")

			id := readInt("Ingrese ID del producto: ")
			name := readLine("Ingrese nombre del producto: ")
			price := readFloat("Ingrese precio del producto: ")

			product := domain.Product{
				ID:    id,
				Name:  name,
				Price: price,
			}

			productService.RegisterProduct(product)
			fmt.Println("✅ Producto registrado correctamente.")

		case 2:
			fmt.Println("\n--- Listado de productos ---")

			products := productService.ListProducts()
			if len(products) == 0 {
				fmt.Println("📭 No hay productos registrados.")
				break
			}

			for _, p := range products {
				fmt.Printf("ID: %d | Nombre: %s | Precio: %.2f\n", p.ID, p.Name, p.Price)
			}

		case 3:
			fmt.Println("Saliendo del sistema. Hasta luego 👋")
			return

		default:
			fmt.Println("❌ Opción inválida. Intente de nuevo.")
		}
	}
}
