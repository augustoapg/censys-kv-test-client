package app

import (
	"log"
	"os"

	"github.com/augustoapg/censysKvTestClient/internal/api"
	"github.com/augustoapg/censysKvTestClient/internal/services"
)

type App struct {
	Logger              *log.Logger
	VerificationHandler *api.VerificationHandler
	KVStoreService      *services.KVStoreService
}

func NewApp(kvStoreUrl string) (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	kvService, err := services.NewKVStoreService(logger, kvStoreUrl)
	if err != nil {
		return nil, err
	}

	verificationHandler := api.NewVerificationHandler(logger, kvService)

	app := &App{
		Logger:              logger,
		VerificationHandler: verificationHandler,
		KVStoreService:      kvService,
	}

	return app, nil
}
