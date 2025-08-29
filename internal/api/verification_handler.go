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

// upsertAndValidateKV creates a key-value pair in the KV store and validates it by verifying that
// the key-value pair returned is the same that was sent to the service.
func (v *VerificationHandler) upsertAndValidateKV(key string, value string) error {
	kv, err := v.KvStoreService.UpsertKV(key, value)
	if err != nil {
		return fmt.Errorf("error creating kv: %v", err)
	}

	if kv.Key != key || kv.Value != value {
		return fmt.Errorf("key or value mismatch when creating kv. Expected: %+v, Got: %+v", services.KV{Key: key, Value: value}, kv)
	}

	return nil
}

// getAndValidateKV retrieves a key-value pair from the KV store and validates it by verifying that
// the value returned is the expected one.
func (v *VerificationHandler) getAndValidateKV(key string, expectedValue string) error {
	kv, err := v.KvStoreService.GetKV(key)
	if err != nil {
		return fmt.Errorf("error retrieving kv: %v", err)
	}

	if kv.Key != key || kv.Value != expectedValue {
		return fmt.Errorf("key or value mismatch when retrieving kv. Expected: %+v, Got: %+v", services.KV{Key: key, Value: expectedValue}, kv)
	}

	return nil
}

// VerifyDeletion is a handler that verifies if deleting key-value pairs from kv-store is working correctly.
// It creates a key-value pair, retrieves it, deletes it, and then attempts to retrieve it again.
// Returns 200 with success message if successful, or 500 with error message if not.
func (v *VerificationHandler) VerifyDeletion(w http.ResponseWriter, r *http.Request) {
	key := "test_key"
	value := "test_value"

	// 1. Create a key-value pair
	err := v.upsertAndValidateKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyDeletion] %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// 2. Retrieve kv from store to make sure upsert worked
	err = v.getAndValidateKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyDeletion] %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
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
	kv, err := v.KvStoreService.GetKV(key)
	if err == nil || kv.Key != "" || kv.Value != "" {
		v.Logger.Printf("[VerifyDeletion] kv not deleted. Kv retrieved: %+v", kv)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("kv not deleted. Kv retrieved: %+v", kv))
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Deletion test successful")
}

// VerifyOverwrite is a handler that verifies if overwriting key-value pairs from kv-store is working correctly.
// It creates a key-value pair, retrieves it, overwrites it, and then retrieves it again to make sure the overwrite worked.
// Returns 200 with success message if successful, or 500 with error message if not.
func (v *VerificationHandler) VerifyOverwrite(w http.ResponseWriter, r *http.Request) {
	key := "test_key"
	value := "test_value"

	// 1. Create a key-value pair
	err := v.upsertAndValidateKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyOverwrite] %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// 2. Retrieve kv from store to make sure upsert worked
	err = v.getAndValidateKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyOverwrite] %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// 3. Overwrite the kv
	value = "test_value_2"
	err = v.upsertAndValidateKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyOverwrite] %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// 4. Retrieve kv from store to make sure overwrite worked
	err = v.getAndValidateKV(key, value)
	if err != nil {
		v.Logger.Printf("[VerifyOverwrite] %v", err)
		utils.WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Overwrite test successful")
}
