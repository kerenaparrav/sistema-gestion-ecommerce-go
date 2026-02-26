package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"sistema-gestion-ecommerce-go/internal/domain"
	"sistema-gestion-ecommerce-go/internal/repository"
	"sistema-gestion-ecommerce-go/internal/service"
)

func main() {
	mux := http.NewServeMux()

	repo := repository.NewInMemoryProductRepository()
	svc := service.NewProductService(repo)

	// Prueba rápida
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")
	})

	// CRUD base
	mux.HandleFunc("/productos", productosHandler(svc))
	mux.HandleFunc("/productos/", productoByIDHandler(svc))

	// Endpoint para demostrar CONCURRENCIA (crea productos en paralelo)
	mux.HandleFunc("/productos/cargar-demo", cargarDemoConcurrenteHandler(svc))

	fmt.Println("✅ API corriendo en: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("❌ Error:", err)
	}
}

// /productos  (GET lista, POST crea)
func productosHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(svc.ListProducts())

		case http.MethodPost:
			var p domain.Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}

			svc.RegisterProduct(p)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"mensaje":  "Producto registrado correctamente",
				"producto": p,
			})

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}

// /productos/{id} (GET uno, PUT actualiza, DELETE elimina)
func productoByIDHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extraer id del path: /productos/123
		idStr := strings.TrimPrefix(r.URL.Path, "/productos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {

		case http.MethodGet:
			p, ok := svc.GetProductByID(id)
			if !ok {
				http.Error(w, "Producto no encontrado", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)

		case http.MethodPut:
			var p domain.Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}

			updated, ok := svc.UpdateProduct(id, p)
			if !ok {
				http.Error(w, "Producto no encontrado", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"mensaje":  "Producto actualizado correctamente",
				"producto": updated,
			})

		case http.MethodDelete:
			ok := svc.DeleteProduct(id)
			if !ok {
				http.Error(w, "Producto no encontrado", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"mensaje": "Producto eliminado correctamente",
				"id":      id,
			})

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}

// /productos/cargar-demo  (GET) -> Crea varios productos en goroutines (CONCURRENCIA)
func cargarDemoConcurrenteHandler(svc *service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		demo := []domain.Product{
			{ID: 101, Name: "Mouse", Price: 15.99},
			{ID: 102, Name: "Teclado", Price: 25.50},
			{ID: 103, Name: "Monitor", Price: 199.99},
			{ID: 104, Name: "Audífonos", Price: 45.00},
		}

		var wg sync.WaitGroup
		wg.Add(len(demo))

		for _, p := range demo {
			p := p // evitar problema de variable en goroutine
			go func() {
				defer wg.Done()
				svc.RegisterProduct(p)
			}()
		}

		wg.Wait()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"mensaje": "Carga concurrente completada",
			"total":   len(demo),
		})
	}
}
