package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/augustoapg/censysKvTestClient/internal/services"
	"github.com/augustoapg/censysKvTestClient/internal/utils"
)

type VerificationHandler struct {
	Logger         *log.Logger
	KvStoreService *services.KVStoreService
}

func NewVerificationHandler(logger *log.Logger, kvStoreService *services.KVStoreService) *VerificationHandler {
	return &VerificationHandler{
		Logger:         logger,
		KvStoreService: kvStoreService,
	}
}

// VerifyDeletion is a handler that verifies if deleting key-value pairs from kv-store is working correctly.
// It creates a key-value pair, retrieves it, deletes it, and then attempts to retrieve it again.
// Returns 200 with success message if successful, or 500 with error message if not.
func (v *VerificationHandler) VerifyDeletion(w http.ResponseWriter, r *http.Request) {
	// 1. Create a key-value pair
	key := "test_key"
	value := "test_value"

	kv, err := v.KvStoreService.UpsertKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyDeletion] error creating kv: %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("error creating kv: %v", err))
		return
	}

	if kv.Key != key || kv.Value != value {
		v.Logger.Printf("[VerifyDeletion] key or value mismatch when creating kv. Expected: %+v, Got: %+v", services.KV{Key: key, Value: value}, kv)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("key or value mismatch when creating kv. Expected: %+v, Got: %+v", services.KV{Key: key, Value: value}, kv))
		return
	}

	// 2. Retrieve kv from store to make sure upsert worked
	kv, err = v.KvStoreService.GetKV(key)
	if err != nil {
		v.Logger.Printf("[VerifyDeletion] error retrieving kv: %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("error retrieving kv: %v", err))
	}

	if kv.Key != key || kv.Value != value {
		v.Logger.Printf("[VerifyDeletion] key or value mismatch when retrieving kv. Expected: %+v, Got: %+v", services.KV{Key: key, Value: value}, kv)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("key or value mismatch when retrieving kv. Expected: %+v, Got: %+v", services.KV{Key: key, Value: value}, kv))
		return
	}

	// 3. Delete kv from store
	err = v.KvStoreService.DeleteKV(key)
	if err != nil {
		v.Logger.Printf("[VerifyDeletion] error deleting kv: %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("error deleting kv: %v", err))
		return
	}

	// 4. Attempt to retrieve kv from store to make sure delete worked
	kv, err = v.KvStoreService.GetKV(key)
	if err == nil || kv.Key != "" || kv.Value != "" {
		v.Logger.Printf("[VerifyDeletion] kv not deleted. Kv retrieved: %+v", kv)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("kv not deleted. Kv retrieved: %+v", kv))
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Deletion test successful")
}

func (v *VerificationHandler) VerifyOverwrite(w http.ResponseWriter, r *http.Request) {
	// 1. Create a key-value pair

	// 2. Retrieve kv from store to make sure upsert worked

	// 3. Overwrite the kv

	// 4. Retrieve kv from store to make sure overwrite worked

	v.Logger.Println("verify overwrite")
}
