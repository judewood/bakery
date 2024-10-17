package s3client

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type S3Communicator interface {
	GetDataFromS3(url string) ([]byte, error)
}

type S3Client struct {
}

func New() *S3Client {
	return &S3Client{}
}

func (s *S3Client) GetDataFromS3(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Get(url)
	if err != nil || response == nil {
		slog.Warn("failed to get response from S3", "error", err)
	}
	slog.Debug("Status code from S3 request", "code", response.StatusCode)
	switch response.StatusCode {
	case http.StatusOK:
		slog.Debug("request returned OK", "url", url)
	case http.StatusBadRequest:
		slog.Warn("request returned bad request", "url", url)
		return []byte{}, fmt.Errorf("client error %+v", response.StatusCode)
	case http.StatusNotFound:
		slog.Warn("request returned not found", "url", url)
		return []byte{}, fmt.Errorf("client error %+v", response.StatusCode)
	default:
		return nil, errors.New("request returned internal service error")
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize response body for url: %s err: %w", url, err)
	}
	return payload, nil
}
