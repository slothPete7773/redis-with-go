package service

import "go-redis-k6/repository"

type catalogService struct {
	productRepo repository.ProductRepository
}

func NewCatalogService(productRepo repository.ProductRepository) CatalogService {
	return catalogService{
		productRepo: productRepo,
	}
}

func (s catalogService) GetProducts() (products []Product, err error) {
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productsDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}
	return products, nil
}
