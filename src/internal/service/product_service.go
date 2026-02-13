package service

import (
	"sistema-gestion-ecommerce-go/internal/domain"
	"sistema-gestion-ecommerce-go/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

// Constructor ✅ (ESTO TAMBIÉN FALTABA)
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) RegisterProduct(product domain.Product) {
	s.repo.AddProduct(product)
}

func (s *ProductService) ListProducts() []domain.Product {
	return s.repo.GetAllProducts()
}
