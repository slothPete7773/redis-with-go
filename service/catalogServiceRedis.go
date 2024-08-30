package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis-k6/repository"
	"time"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repository.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repository.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{
		productRepo: productRepo,
		redisClient: redisClient,
	}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {

	key := "service::GetProducts"
	// GET Redis
	if productsJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if err = json.Unmarshal([]byte(productsJson), &products); err == nil {
			fmt.Println("From Service Redis")
			return products, nil
		}
	}

	// Database
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

	// SET Redis
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, data, time.Second*15)
	}

	fmt.Println("From Database.")
	return products, nil
}
