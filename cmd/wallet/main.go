package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"os"
	"wallet/internal/config"
	"wallet/internal/http-server/middleware/logger"
	"wallet/internal/lib/logger/handlers/slogpretty"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting wallet", slog.String("env", cfg.Env))

	ctx := context.Background()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.ClientIPFromHeader(middleware.GetClientIP(ctx)))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(logger.New(log))
}

// Инициализация
func setupLogger(env string) *slog.Logger {
	var h slog.Handler

	switch env {
	case config.EnvLocal:
		opts := slogpretty.PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
		h = opts.NewPrettyHandler(os.Stdout)
	case config.EnvDev:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case config.EnvProd:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return slog.New(h)
}
