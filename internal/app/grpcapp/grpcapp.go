package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	productgprc "productservice/internal/server/productgrpc"

	"google.golang.org/grpc"
)

type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func NewGRPCApp(log *slog.Logger, port string, productService productgprc.Products) *GRPCApp {
	gRPCServer := grpc.NewServer()
	productgprc.RegisterServ(gRPCServer, productService)
	return &GRPCApp{log, gRPCServer, port}
}

func (app *GRPCApp) Run() error {
	const op = "grpcapp.Run"

	log := app.log.With(slog.String("op", op),
		slog.String("port", app.port),
	)

	log.Info("starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := app.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *GRPCApp) Stop() {
	const op = "grpcapp.Stop"

	app.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.String("port", app.port))

	app.gRPCServer.GracefulStop()
}
