package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/avran02/decoplan/files/internal/config"
	"github.com/google/uuid"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	getObjectOptions  = minio.GetObjectOptions{}
	makeBucketOptions = minio.MakeBucketOptions{
		Region: config.DefaultLocation,
	}
)

type FilesService interface {
	UploadFile(ctx context.Context, data io.Reader, fileName string) (string, error)
	DownloadFile(ctx context.Context, fileID string) (io.ReadCloser, error)
	DeleteFile(ctx context.Context, fileID string) error
}

type filesService struct {
	minio *minio.Client
}

func (s *filesService) UploadFile(ctx context.Context, data io.Reader, fileName string) (string, error) {
	slog.Info("filesService.UploadFile")
	if err := s.createBucketIfNotExists(ctx); err != nil {
		return "", fmt.Errorf("failed to create bucket: %w", err)
	}

	fileID := uuid.NewString() + getExtatension(fileName)
	if _, err := s.minio.PutObject(ctx, config.UserDataBucket, fileID, data, -1, minio.PutObjectOptions{}); err != nil {
		if errors.Is(err, io.EOF) {
			return fileID, nil
		}
		slog.Error("failed to upload file", "fileID", fileID, "err", err)
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return fileID, nil
}

func (s *filesService) DownloadFile(ctx context.Context, fileID string) (io.ReadCloser, error) {
	slog.Info("filesService.DownloadFile")
	if err := s.createBucketIfNotExists(ctx); err != nil {
		return nil, err
	}

	o, err := s.minio.GetObject(ctx, config.UserDataBucket, fileID, getObjectOptions)
	if err != nil {
		err = fmt.Errorf("failed to get object: %w", err)
		slog.Error(err.Error())
		return nil, err
	}

	return o, nil
}

func (s *filesService) DeleteFile(ctx context.Context, fileID string) error {
	slog.Info("filesService.DeleteFile")
	if err := s.createBucketIfNotExists(ctx); err != nil {
		return err
	}

	slog.Debug("deleting file", "fileID", fileID, "bucket", config.UserDataBucket)
	return s.minio.RemoveObject(ctx, config.UserDataBucket, fileID, minio.RemoveObjectOptions{})
}

func (s *filesService) createBucketIfNotExists(ctx context.Context) error {
	exists, err := s.minio.BucketExists(ctx, config.UserDataBucket)
	if err != nil {
		err = fmt.Errorf("failed to check if bucket exists: %w", err)
		slog.Error(err.Error())
		return err
	}

	if !exists {
		err = s.minio.MakeBucket(ctx, config.UserDataBucket, makeBucketOptions)
		if err != nil {
			err = fmt.Errorf("failed to create bucket: %w", err)
			slog.Error(err.Error())
			return err
		}
	}

	return nil
}

func getExtatension(fileName string) string {
	return strings.ToLower(filepath.Ext(fileName))
}

func New(conf config.Minio) FilesService {
	slog.Info("initializing service")
	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Region: config.DefaultLocation,
		Secure: false,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &filesService{
		minio: minioClient,
	}
}
