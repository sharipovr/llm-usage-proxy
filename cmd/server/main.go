package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sharipovr/llm-usage-proxy/internal/config"
	"github.com/sharipovr/llm-usage-proxy/internal/middleware"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(log)

	cfg := config.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger(log))

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Info("listening on", slog.String("addr", cfg.ListenAddr))
	if err := srv.ListenAndServe(); err != nil {
		log.Error("server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
