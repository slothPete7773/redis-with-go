package handler

import (
	"go-redis-k6/service"

	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogService service.CatalogService
}

func NewCatalogHandler(catalogService service.CatalogService) CatalogHandler {
	return catalogHandler{
		catalogService: catalogService,
	}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {
	// producst, err := h.GetProducts()
	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}
	return c.JSON(response)

}
