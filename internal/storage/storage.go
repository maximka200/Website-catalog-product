package storage

import (
	"context"
	"fmt"
	"productservice/internal/config"
	"productservice/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	productTable = "Products"
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

func (s *StorageStruct) NewProduct(ctx context.Context, imageURL string, title string, description string, discount uint8, price int64, currency int32) (int64, error) {
	const op = "storage.NewProduct"

	stmt, err := s.db.Prepare(fmt.Sprintf(`INSERT INTO %s (image_url, title, description, price, currency,
	discount, product_url) VALUES (?, ?, ?, ?, ?, ?, ?)`, productTable))
	if err != nil {
		return 0, fmt.Errorf("%s: %s", op, err)
	}

	res, err := stmt.QueryContext(ctx, imageURL, title, description, price, currency)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", op, err)
	}

	var id int64

	if err := res.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %s", op, err)
	}

	return id, nil
}

func (s *StorageStruct) DeleteProduct(ctx context.Context, id int64) (bool, error) {
	const op = "storage.DeleteProduct"

	stmt, err := s.db.Prepare("DELETE FROM %s WHERE id=?")
	if err != nil {
		return false, fmt.Errorf("%s: %s", op, err)
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return false, fmt.Errorf("%s: %s", op, err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("%s: %s", op, err)
	} else if count == 0 {
		return false, nil
	}

	return true, nil
}

func (s *StorageStruct) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	const op = "storage.GetProduct"

	var model *models.Product

	stmt, err := s.db.Prepare("SELECT * FROM $s WHERE id=?")
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	res := stmt.QueryRowContext(ctx, id)

	if err := res.Scan(model); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return model, nil
}
