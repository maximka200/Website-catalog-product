package service

import (
	"context"
	"productservice/internal/models"
)

type ProductStruct struct {
	storageStruct string // stub
}

func NewProductStruct() *ProductStruct {
	return &ProductStruct{"stub"}
}

func (ps *ProductStruct) NewProduct(ctx context.Context, id int64, imageURL string, title string, Description string, Price int64, Currency int32) (int64, error) {
	return 1, nil
}

func (ps *ProductStruct) DeleteProduct(ctx context.Context, id int64) (bool, error) {
	return true, nil
}

func (ps *ProductStruct) GetProduct(ctx context.Context, id int64) (models.Product, error) {
	return models.Product{Id: 1}, nil
}
