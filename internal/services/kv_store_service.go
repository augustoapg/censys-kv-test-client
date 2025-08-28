package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type KV struct {
	Key       string     `json:"key"`
	Value     string     `json:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// KVResponse represents the wrapped response from the API
type KVResponse struct {
	KV KV `json:"kv"`
}

type KVStoreService struct {
	Logger     *log.Logger
	Client     *http.Client
	KvStoreUrl string
}

func NewKVStoreService(logger *log.Logger, kvStoreUrl string) (*KVStoreService, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// adds default value if not provided
	if kvStoreUrl == "" {
		kvStoreUrl = "http://localhost:8080"
	}

	return &KVStoreService{
		Logger:     logger,
		Client:     client,
		KvStoreUrl: kvStoreUrl,
	}, nil
}

// GetKV retrieves a key-value pair from the kv-store.
// Returns the kv pair or error if the request fails.
func (s *KVStoreService) GetKV(key string) (KV, error) {
	if key == "" {
		return KV{}, fmt.Errorf("key is required")
	}

	url := fmt.Sprintf("%s/kv/%s", s.KvStoreUrl, key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.Logger.Printf("[GetKV] error creating request: %v", err)
		return KV{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		s.Logger.Printf("[GetKV] error sending request: %v", err)
		return KV{}, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var body map[string]any

		if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
			s.Logger.Printf("[GetKV] error decoding error body: %v", err)
			return KV{}, fmt.Errorf("error decoding error body: %v", err)
		}

		message := "unexpected status code"
		if msg, ok := body["message"]; ok {
			message = fmt.Sprintf("%s", msg)
		}

		s.Logger.Printf("[GetKV] unexpected status code: %d. Error message: %v", resp.StatusCode, message)
		return KV{}, fmt.Errorf("unexpected status code: %d. Error message: %v", resp.StatusCode, message)
	}

	var kvResponse KVResponse
	if err := json.NewDecoder(resp.Body).Decode(&kvResponse); err != nil {
		s.Logger.Printf("[GetKV] error decoding body: %v", err)
		return KV{}, fmt.Errorf("error decoding body: %v", err)
	}

	return kvResponse.KV, nil
}

// UpsertKV creates or updates a key-value pair in the kv-store.
// Returns the created or updated kv pair or error if the request fails.
func (s *KVStoreService) UpsertKV(key string, value string) (KV, error) {
	if key == "" {
		return KV{}, fmt.Errorf("key is required")
	}

	url := fmt.Sprintf("%s/kv", s.KvStoreUrl)

	jsonData, err := json.Marshal(KV{
		Key:   key,
		Value: value,
	})

	if err != nil {
		s.Logger.Printf("[UpsertKV] Error marshalling kv: %v", err)
		return KV{}, fmt.Errorf("error marshalling kv: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		s.Logger.Printf("[UpsertKV] Error creating request: %v", err)
		return KV{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		s.Logger.Printf("[UpsertKV] Error sending request: %v", err)
		return KV{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.Logger.Printf("[UpsertKV] Unexpected status code: %d", resp.StatusCode)

		var body map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			s.Logger.Printf("[UpsertKV] Error decoding error response body: %v", err)
		}

		message := "unexpected status code"
		if msg, ok := body["message"]; ok {
			message = fmt.Sprintf("%v", msg)
		}
		s.Logger.Printf("[UpsertKV] unexpected status code: %d. Error message: %v", resp.StatusCode, message)
		return KV{}, fmt.Errorf("%s", message)
	}

	var kvResponse KVResponse
	if err = json.NewDecoder(resp.Body).Decode(&kvResponse); err != nil {
		s.Logger.Printf("[UpsertKV] Error decoding response: %v", err)
		return KV{}, fmt.Errorf("error decoding response: %v", err)
	}

	return kvResponse.KV, nil
}

// DeleteKV deletes a key-value pair from the kv-store.
// Returns error if the request fails.
func (s *KVStoreService) DeleteKV(key string) error {
	if key == "" {
		return fmt.Errorf("key is required")
	}

	url := fmt.Sprintf("%s/kv/%s", s.KvStoreUrl, key)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		s.Logger.Printf("[DeleteKV] Error creating request: %v", err)
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		s.Logger.Printf("[DeleteKV] Error sending request: %v", err)
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		var body map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			s.Logger.Printf("[DeleteKV] Error decoding error response body: %v", err)
		}

		message := "unexpected status code"
		if msg, ok := body["message"]; ok {
			message = fmt.Sprintf("%v", msg)
		}
		s.Logger.Printf("[DeleteKV] unexpected status code: %d. Error message: %v", resp.StatusCode, message)
		return fmt.Errorf("%s", message)
	}

	return nil
}
