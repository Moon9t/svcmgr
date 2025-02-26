package storage

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type CosmicVault struct {
	s3Client *s3.Client
	bucket   string
}

func NewCosmicVault(bucket string) (*CosmicVault, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &CosmicVault{
		s3Client: s3.NewFromConfig(cfg),
		bucket:   bucket,
	}, nil
}

func (v *CosmicVault) BackupConfig(ctx context.Context, data []byte) error {
	key := fmt.Sprintf("backups/%s/config.enc", time.Now().Format("2006-01-02"))
	_, err := v.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(v.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}
