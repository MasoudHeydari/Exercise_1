package dto

import "github.com/MasoudHeydari/Exercise_1/Task2_imagy/domain"

type DownloadImageRequest struct {
	URLPath   string
	LocalName string
	DstPath   string
}

type DownloadImageResponse struct {
	ImageName string
}

type ListImageRequest struct{}

type ListImageResponse struct {
	Images []domain.Image
}
