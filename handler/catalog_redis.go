package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis-k6/service"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv  service.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogService service.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{
		catalogSrv:  catalogService,
		redisClient: redisClient,
	}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {

	key := "handler::GetProducts"

	if jsonData, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("From Redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(jsonData)
	}

	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	if data, err := json.Marshal(response); err == nil {

		h.redisClient.Set(context.Background(), key, string(data), time.Second*15)
	}
	fmt.Println("from Database")
	return c.JSON(response)
}
