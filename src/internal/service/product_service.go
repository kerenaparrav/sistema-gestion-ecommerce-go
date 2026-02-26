package service

import (
	"sistema-gestion-ecommerce-go/internal/domain"
	"sistema-gestion-ecommerce-go/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

// Constructor
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) RegisterProduct(product domain.Product) {
	s.repo.AddProduct(product)
}

func (s *ProductService) ListProducts() []domain.Product {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id int) (domain.Product, bool) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(id int, product domain.Product) (domain.Product, bool) {
	return s.repo.Update(id, product)
}

func (s *ProductService) DeleteProduct(id int) bool {
	return s.repo.Delete(id)
}
