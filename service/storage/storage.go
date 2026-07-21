package storage

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, key string, body io.Reader) (string, error)
	Delete(ctx context.Context, key string) error
	GetPresignedURL(ctx context.Context, key string) (string, error)
	Download(ctx context.Context, key string) (io.ReadCloser, error)
}
