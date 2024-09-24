package service

// this layer is for business logic
import (
	"context"
	"fmt"
	"math/rand"

	productv1 "github.com/maximka200/protobuff_product/gen"
)

type ProductStruct struct {
	storageStruct StorageMethod // stub
}

type StorageMethod interface {
	NewProduct(ctx context.Context, imageURL string, title string, description string, discount uint8, price int64, currency int32, productURL string) (int64, error)
	DeleteProduct(ctx context.Context, id int64) (bool, error)
	GetProduct(ctx context.Context, id int64) (*productv1.GetProductResponse, error)
	GetAvailableId(ctx context.Context) (*[]int, error)
}

func NewProductStruct(storage StorageMethod) *ProductStruct {
	return &ProductStruct{storageStruct: storage}
}

func (ps *ProductStruct) NewProduct(ctx context.Context, imageURL string, title string, description string, discount uint8, price int64, currency int32, productURL string) (int64, error) {
	const op = "service.NewProduct"

	id, err := ps.storageStruct.NewProduct(ctx, imageURL, title, description, discount, price, currency, productURL)
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

func (ps *ProductStruct) GetProduct(ctx context.Context, id int64) (*productv1.GetProductResponse, error) {
	const op = "service.GetProduct"

	model, err := ps.storageStruct.GetProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return model, nil
}

func (ps *ProductStruct) GetProducts(ctx context.Context, count int64) (*productv1.GetProductsResponse, error) {
	const op = "service.GetProducts"

	// todo: caching idList
	idList, err := ps.storageStruct.GetAvailableId(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	idRandomList := getRandomIDs(*idList, int(count))

	idResultList, err := getProductList(ps, ctx, *idRandomList)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &productv1.GetProductsResponse{ProductList: idResultList}, nil
}

func getRandomIDs(ids []int, n int) *[]int {
	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	if n > len(ids) {
		n = len(ids)
	}

	ids = ids[:n]

	return &ids
}

func getProductList(ps *ProductStruct, ctx context.Context, idList []int) ([]*productv1.GetProductResponse, error) {
	const op = "storage.getProductList"
	var result []*productv1.GetProductResponse

	for i := 0; i < len(idList); i++ {
		resp, err := ps.GetProduct(ctx, int64(idList[i]))
		if err != nil {
			return nil, fmt.Errorf("%s: cannot get product: %s", op, err)
		}
		result = append(result, resp)
	}

	return result, nil
}
