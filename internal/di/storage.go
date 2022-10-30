package di

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/radityarestan/ecom-core/internal/pkg"
	"github.com/radityarestan/ecom-core/internal/shared/config"
	"os"
)

func NewStorage(cfg *config.Config) (*pkg.Storage, error) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cfg.GCP.Credential)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &pkg.Storage{
		Cl:          client,
		ProjectID:   cfg.GCP.ProjectID,
		BucketName:  cfg.GCP.BucketName,
		ProductPath: cfg.GCP.ProductStoragePath,
	}, nil
}
