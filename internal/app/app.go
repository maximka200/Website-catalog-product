package app

import (
	"fmt"
	"log/slog"
	"productservice/internal/app/grpcapp"
	"productservice/internal/service"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

func NewApp(log *slog.Logger, grpcPort string) *App { // TTL - time to live
	prod := service.NewProductStruct()
	grpcApp := grpcapp.NewGRPCApp(log, grpcPort, prod)

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
