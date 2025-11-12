package main

import (
	"first_l/internal/config"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "Local"
	envProd  = "Production"
	envDev   = "Development"
)

func main(){
	// init config : clean env
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// init logger
	log := setupLogger(cfg.Env)
	// adding a Env to log
	// log.With(slog.String("env", cfg.Env))
	log.Info("starting url-shortner", slog.String("env", cfg.Env))

	

	// storage SQLite

	// init router chi, "chi render"

	// run server
}

func setupLogger(env string) *slog.Logger{
		var log *slog.Logger

		switch env {
		case envLocal:
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case envProd:
			log = slog.New(slog.NewJSONHandler(os.Stdout,  &slog.HandlerOptions{Level: slog.LevelInfo}))
		case envDev:
			log = slog.New(slog.NewJSONHandler(os.Stdout,  &slog.HandlerOptions{Level: slog.LevelDebug}))
		default:
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}
		return log
	}
