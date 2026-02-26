package repository

import (
	"errors"
	"sync"

	"sistema-gestion-ecommerce-go/internal/domain"
)

// Interfaz (contrato)
type ProductRepository interface {
	AddProduct(product domain.Product)
	GetAllProducts() []domain.Product

	GetByID(id int) (domain.Product, bool)
	Update(id int, product domain.Product) (domain.Product, bool)
	Delete(id int) bool
}

// Implementación en memoria (thread-safe)
type InMemoryProductRepository struct {
	mu       sync.RWMutex
	products []domain.Product
}

// Constructor
func NewInMemoryProductRepository() ProductRepository {
	return &InMemoryProductRepository{
		products: []domain.Product{},
	}
}

func (r *InMemoryProductRepository) AddProduct(product domain.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products = append(r.products, product)
}

func (r *InMemoryProductRepository) GetAllProducts() []domain.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// devolvemos copia para no exponer el slice interno
	out := make([]domain.Product, len(r.products))
	copy(out, r.products)
	return out
}

func (r *InMemoryProductRepository) GetByID(id int) (domain.Product, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, p := range r.products {
		if p.ID == id {
			return p, true
		}
	}
	return domain.Product{}, false
}

func (r *InMemoryProductRepository) Update(id int, product domain.Product) (domain.Product, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.products {
		if p.ID == id {
			// mantenemos el ID del path como fuente de verdad
			product.ID = id
			r.products[i] = product
			return product, true
		}
	}
	return domain.Product{}, false
}

func (r *InMemoryProductRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.products {
		if p.ID == id {
			// eliminar por “swap with last” (rápido)
			r.products[i] = r.products[len(r.products)-1]
			r.products = r.products[:len(r.products)-1]
			return true
		}
	}
	return false
}

// (Opcional) helper por si luego quieres validar cosas
var ErrNotFound = errors.New("producto no encontrado")
