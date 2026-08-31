package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/database"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func generatePresignedURL(s3Client *s3.Client, bucket, key string, expireTime time.Duration) (string, error) {
	presigner := s3.NewPresignClient(s3Client)

	presignedGetURL, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expireTime))
	if err != nil {
		return "", fmt.Errorf("failed to presign get object: %w", err)
	}
	return presignedGetURL.URL, nil
}

func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
	if video.VideoURL == nil {
		return video, nil
	}

	urlParts := strings.Split(*video.VideoURL, ",")
	if len(urlParts) != 2 {
		return video, fmt.Errorf("invalid video URL format: %s", *video.VideoURL)
	}
	bucket := urlParts[0]
	videoKey := urlParts[1]

	signedURL, err := generatePresignedURL(cfg.s3Client, bucket, videoKey, 15*time.Minute)
	if err != nil {
		return video, fmt.Errorf("failed to generate presigned URL: %w", err)
	}
	video.VideoURL = &signedURL
	return video, nil
}
