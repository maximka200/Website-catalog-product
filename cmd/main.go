package main

import (
	"log/slog"
	"os"
	"os/signal"
	"productservice/internal/app"
	"productservice/internal/config"
	"syscall"
)

func main() {
	cfg := config.MustReadConfig()

	log := initLogger(cfg.Env)
	log.Info("logger and config success start")
	application := app.NewApp(log, &cfg)
	log.Info("app success start")
	go application.MustRun()
	// run server

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	application.GRPCSrv.Stop()

	log.Info("application stopped")
}

func initLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
