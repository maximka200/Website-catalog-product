package productgprc

import (
	"context"
	"fmt"

	productv1 "github.com/maximka200/protobuff_product/gen"
	"google.golang.org/grpc"
)

type Products interface {
	NewProduct(ctx context.Context, imageURL string, title string, description string, discount uint8, price int64, currency int32, productURL string) (int64, error)
	DeleteProduct(ctx context.Context, id int64) (bool, error)
	GetProduct(ctx context.Context, id int64) (*productv1.GetProductResponse, error)
	GetProducts(ctx context.Context, count int64) (*productv1.GetProductsResponse, error)
}

type serverAPI struct {
	productv1.UnimplementedProductServer
	product Products
}

func RegisterServ(gRPC *grpc.Server, product Products) {
	productv1.RegisterProductServer(gRPC, &serverAPI{product: product})
}

func (s *serverAPI) NewProduct(ctx context.Context, req *productv1.NewProductRequest) (*productv1.NewProductResponse, error) {
	const op = "productgprc.NewProduct"

	rq, err := s.product.NewProduct(ctx, req.GetImageURL(), req.GetTitle(), req.GetDescription(), uint8(req.GetDiscount()), req.GetPrice(), req.GetCurrency(), req.GetProductURL())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &productv1.NewProductResponse{Id: rq}, nil
}

func (s *serverAPI) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	const op = "productgprc.GetProduct"

	resp, err := s.product.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	//todo: marshall model.Product in productv1.GetProductResponse
	return &productv1.GetProductResponse{Id: resp.Id, ImageURL: resp.ImageURL, Title: resp.Title,
		Description: resp.Description, Price: resp.Price, Currency: resp.Currency, ProductURL: resp.ProductURL}, nil
}

func (s *serverAPI) DeleteProduct(ctx context.Context, req *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	const op = "productgprc.DeleteProduct"

	resp, err := s.product.DeleteProduct(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &productv1.DeleteProductResponse{IsDelete: resp}, nil
}

func (s *serverAPI) GetProducts(ctx context.Context, req *productv1.GetProductsRequest) (*productv1.GetProductsResponse, error) {
	const op = "productgprc.GetProducts"

	resp, err := s.product.GetProducts(ctx, req.GetCount())
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return resp, nil
}
