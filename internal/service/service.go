package service

// this layer is for business logic
import (
	"context"
	"fmt"
	"productservice/internal/models"
)

type ProductStruct struct {
	storageStruct StorageMethod // stub
}

type StorageMethod interface {
	NewProduct(ctx context.Context, imageURL string, title string, description string, price int64, currency int32) (int64, error)
	DeleteProduct(ctx context.Context, id int64) (bool, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
}

func NewProductStruct(storage StorageMethod) *ProductStruct {
	return &ProductStruct{storageStruct: storage}
}

func (ps *ProductStruct) NewProduct(ctx context.Context, imageURL string, title string, description string, price int64, currency int32) (int64, error) {
	const op = "service.NewProduct"

	id, err := ps.storageStruct.NewProduct(ctx, imageURL, title, description, price, currency)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", op, err)
	}

	return id, nil
}

func (ps *ProductStruct) DeleteProduct(ctx context.Context, id int64) (bool, error) {
	const op = "service.DeleteProduct"

	isDelete, err := ps.storageStruct.DeleteProduct(ctx, id)
	if err != nil {
		return false, fmt.Errorf("%s: %s", op, err)
	}

	return isDelete, nil
}

func (ps *ProductStruct) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	const op = "service.GetProduct"

	model, err := ps.storageStruct.GetProduct(ctx, id)
	if err != nil {
		return &models.Product{}, fmt.Errorf("%s: %s", op, err)
	}

	return model, nil
}
