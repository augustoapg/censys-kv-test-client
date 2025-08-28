package main

import (
	"net/http"
	"os"
	"time"

	"github.com/augustoapg/censysKvTestClient/internal/app"
	"github.com/augustoapg/censysKvTestClient/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	app, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Server running on port: %s", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatalf("Error starting server: %v", err)
	}
}
