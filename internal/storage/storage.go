package storage

import (
	"context"
	"fmt"
	"productservice/internal/config"
	"productservice/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type StorageStruct struct {
	db *sqlx.DB
}

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	op := "storage.NewSqlxDB"

	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLmode))
	if err != nil {
		return nil, fmt.Errorf("%s:%s", err, op)
	}
	return db, nil
}

func NewStorageStruct(db *sqlx.DB) *StorageStruct {
	return &StorageStruct{db: db}
}

func (s *StorageStruct) NewProduct(ctx context.Context, imageURL string, title string, description string, price int64, currency int32) (int64, error) {
	const op = "storage.NewProduct"

	return 0, nil
}

func (s *StorageStruct) DeleteProduct(ctx context.Context, id int64) (bool, error) {
	const op = "storage.DeleteProduct"

	return false, nil
}

func (s *StorageStruct) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	const op = "storage.GetProduct"

	return &models.Product{Title: "pedic"}, nil
}
