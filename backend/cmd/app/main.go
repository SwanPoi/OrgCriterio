package main

import (
	"os"

	"github.com/SwanPoi/OrgCriterio/internal/config"
	"github.com/SwanPoi/OrgCriterio/internal/handler"
	"github.com/SwanPoi/OrgCriterio/internal/logger"
	"github.com/SwanPoi/OrgCriterio/internal/storage"
	"github.com/SwanPoi/OrgCriterio/internal/usecase"
)

func main() {
	cfg := config.Must()

	log := logger.Setup(cfg.Env, cfg.Log.Level)
	log.Info("starting application", "env", cfg.Env, "log_level", cfg.Log.Level)

	store := storage.New()
	uc := usecase.New(store)
	srv := handler.NewServer(log, uc)

	log.Info("starting server", "address", cfg.HTTPServer.Address)

	if err := srv.Run(cfg.HTTPServer.Address); err != nil {
		log.Error("failed to run server", "error", err.Error())
		os.Exit(1)
	}
}