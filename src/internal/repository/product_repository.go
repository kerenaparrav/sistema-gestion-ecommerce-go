package repository

import "sistema-gestion-ecommerce-go/internal/domain"

// Interfaz (contrato)
type ProductRepository interface {
	AddProduct(product domain.Product)
	GetAllProducts() []domain.Product
}

// Implementación en memoria
type InMemoryProductRepository struct {
	products []domain.Product
}

// Constructor ✅ (ESTO FALTABA)
func NewInMemoryProductRepository() ProductRepository {
	return &InMemoryProductRepository{
		products: []domain.Product{},
	}
}

func (r *InMemoryProductRepository) AddProduct(product domain.Product) {
	r.products = append(r.products, product)
}

func (r *InMemoryProductRepository) GetAllProducts() []domain.Product {
	return r.products
}
