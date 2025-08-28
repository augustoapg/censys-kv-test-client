package routes

import (
	"github.com/augustoapg/censysKvTestClient/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/test_deletion", app.VerificationHandler.VerifyDeletion)
	r.Post("/test_overwrite", app.VerificationHandler.VerifyOverwrite)

	return r
}
