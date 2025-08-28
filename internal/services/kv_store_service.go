package services

import (
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

func (s *KVStoreService) GetKV(key string) (KV, error) {
	if key == "" {
		return KV{}, fmt.Errorf("key is required")
	}

	return KV{}, nil
}

func (s *KVStoreService) UpsertKV(key string, value string) (KV, error) {
	if key == "" {
		return KV{}, fmt.Errorf("key is required")
	}

	return KV{}, nil
}

func (s *KVStoreService) DeleteKV(key string) error {
	if key == "" {
		return fmt.Errorf("key is required")
	}

	return nil
}
