package gcs

import (
	"bytes"
	"context"
	"fmt"
	"ioutil"

	"cloud.google.com/go/storage"
	storage "github.com/minamijoyo/tfmigrate-storage"
)

type Storage struct {
	// config is a storage config for s3.
	config *Config
	// client is an instance of S3Client interface to call API.
	// It is intended to be replaced with a mock for testing.
	client Client
}

var _ storage.Storage = (*Storage)(nil)

// NewStorage returns a new instance of Storage.
func NewStorage(config *Config) (*Storage, error) {
	s := &Storage{
		config: config,
		client = storage.NewClient(ctx)
	}
	return s, nil
}

func (s *Storage) Write(ctx context.Context, path string, b []byte) error {
	if err := s.init(ctx); err != nil {
		return err
	}
	w := s.client.Bucket(s.config.Bucket).Object(path).NewWriter(ctx)

	_, err := w.Write(b)
	if err != nil {
		return fmt.Errorf("failed writing to gcs://%s/%s: %w", s.config.Bucket, path, err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed closing writer for gcs://%s/%s: %w", s.config.Bucket, path, err)
	}
	return nil
}

func (s *Storage) Read(ctx context.Context, path string) ([]byte, error) {
	if err := s.init(ctx); err != nil {
		return nil, err
	}
	r, err := s.client.Bucket(s.config.Bucket).Object(path).NewReader(ctx)
	
	_, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("failed reading from gcs://%s/%s: %w", s.config.Bucket, path, err)
	}

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed reading from gcs://%s/%s: %w", s.config.Bucket, path, err)
	}

	err = r.Close()
	if err != nil {
		return nil, fmt.Errorf("failed closing reader for gcs://%s/%s: %w", s.config.Bucket, path, err)
	}
	return body, nil
}

func (s *CSSStorage) Init(ctx context.Context) error {
	if (s.client != nil) {
		return nil
	}
	
}
