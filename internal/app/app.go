package app

import (
	"fmt"
	"log/slog"
	"productservice/internal/app/grpcapp"
	"productservice/internal/config"
	"productservice/internal/service"
	"productservice/internal/storage"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

func NewApp(log *slog.Logger, cfg *config.Config) *App {
	db, err := storage.NewDB(cfg)
	if err != nil {
		panic(fmt.Sprintf("cannot connect with db: %s", err))
	}
	log.Info("db success created")
	storage := storage.NewStorageStruct(db)
	prod := service.NewProductStruct(storage)
	grpcApp := grpcapp.NewGRPCApp(log, cfg.Port, prod)

	return &App{
		GRPCSrv: grpcApp,
	}
}

func (app *App) MustRun() {
	if err := app.GRPCSrv.Run(); err != nil {
		err = fmt.Errorf("error run server: %w", err)
		panic(err)
	}
}
