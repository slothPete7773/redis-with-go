package main

import (
	"fmt"
	"go-redis-k6/repository"
	"go-redis-k6/service"

	// "github.com/go-sql-driver/mysql"
	// _ "gorm.io/driver/mysql"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisDb := initRedis()
	_ = redisDb

	// productRepo := repository.NewProductDB(db)
	productRepo := repository.NewProductDB(db)
	productService := service.NewCatalogServiceRedis(productRepo, redisDb)
	// _ = productRepo
	products, err := productService.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Products: %v", products)

	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	time.Sleep(time.Millisecond * 10)
	// 	return c.SendString("Hello WOrld")
	// })

	// app.Listen(":8000")
}

func init() {

}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
func initDatabase() *gorm.DB {
	dialector := mysql.Open("root:mariadb@tcp(localhost:3306)/mariadb")
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}

	return db
}
