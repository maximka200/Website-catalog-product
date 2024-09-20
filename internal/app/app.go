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
	storage := storage.NewStorageStruct(db, log)
	log = log.With(slog.String("portdb", cfg.DB.Port), slog.String("hostdb", cfg.DBname),
		slog.String("username and password", cfg.DB.Username+" "+cfg.DB.Password))
	log.Info("success connect to db")
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
