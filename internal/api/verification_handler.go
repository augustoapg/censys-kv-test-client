package api

import (
	"log"
	"net/http"
)

type VerificationHandler struct {
	Logger *log.Logger
}

func NewVerificationHandler(logger *log.Logger) *VerificationHandler {
	return &VerificationHandler{
		Logger: logger,
	}
}

func (v *VerificationHandler) VerifyDeletion(w http.ResponseWriter, r *http.Request) {
	v.Logger.Println("verify deletion")
}

func (v *VerificationHandler) VerifyOverwrite(w http.ResponseWriter, r *http.Request) {
	v.Logger.Println("verify overwrite")
}
