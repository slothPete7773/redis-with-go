package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRedis(db *gorm.DB, redis *redis.Client) ProductRepository {
	return productRepositoryRedis{
		db:          db,
		redisClient: redis,
	}
}

func (p productRepositoryRedis) GetProducts() (products []product, err error) {

	key := "repository::GetProducts"
	// 1. Cache Redis GET
	productsJson, err := p.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		// Found data in cache
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("From redis")
			return products, nil
		}

	}

	// 2. In-between Get database data
	err = p.db.Order("quantity desc").Limit(10).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// 3. Cache to Redis SET
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	err = p.redisClient.Set(context.Background(), key, string(data), time.Second*20).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database")
	return products, nil
}
