package app

import (
	"log"
	"os"

	"github.com/augustoapg/censysKvTestClient/internal/api"
)

type App struct {
	Logger              *log.Logger
	VerificationHandler *api.VerificationHandler
}

func NewApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	verificationHandler := api.NewVerificationHandler(logger)

	return &App{
		Logger:              logger,
		VerificationHandler: verificationHandler,
	}, nil
}
