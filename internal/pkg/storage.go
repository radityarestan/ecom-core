package pkg

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
	"mime/multipart"
	"time"
)

type Storage struct {
	Cl          *storage.Client
	ProjectID   string
	BucketName  string
	ProductPath string
}

func (s *Storage) UploadFile(file multipart.File, fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	wc := s.Cl.Bucket(s.BucketName).Object(s.ProductPath + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}
