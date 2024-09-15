package productgprc

import (
	"context"
	"fmt"
	"productservice/internal/models"

	productv1 "github.com/maximka200/protobuff_product/gen"
	"google.golang.org/grpc"
)

type Products interface {
	NewProduct(ctx context.Context, id int64, imageURL string, title string, Description string, Price int64, Currency int32) (int64, error)
	DeleteProduct(ctx context.Context, id int64) (bool, error)
	GetProduct(ctx context.Context, id int64) (models.Product, error)
}

type serverAPI struct {
	productv1.UnimplementedProductServer
	product Products
}

func RegisterServ(gRPC *grpc.Server, protuct Products) {
	productv1.RegisterProductServer(gRPC, &serverAPI{product: protuct})
}

func (s *serverAPI) NewProduct(ctx context.Context, req *productv1.NewProductRequest) (*productv1.NewProductResponse, error) {
	const op = "serverAPI.NewProduct"

	rq, err := s.product.NewProduct(context.Background(), 1, "", "", "", 1, 1)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &productv1.NewProductResponse{Id: rq}, nil
}

func (s *serverAPI) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	const op = "serverAPI.GetProduct"

	_, err := s.product.GetProduct(context.Background(), req.GetId())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &productv1.GetProductResponse{}, nil
}

func (s *serverAPI) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	const op = "serverAPI.DeleteProduct"

	resp, err := s.product.DeleteProduct(context.Background(), req.GetId())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &productv1.DeleteProductResponse{IsDelete: resp}, nil
}
