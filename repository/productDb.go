package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})

	err := mockData(db)
	if err != nil {
		log.Fatal(err)
	}

	return productRepositoryDB{
		db: db,
	}
}

func mockData(db *gorm.DB) error {
	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []product{}
	for i := 0; i < 100; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("product-%v", i+1),
			Quantity: i * 10,
		})
	}

	return db.Create(&products).Error
}

func (p productRepositoryDB) GetProducts() ([]product, error) {

	return nil, nil
}
