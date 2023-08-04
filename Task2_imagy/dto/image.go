package dto

import "github.com/MasoudHeydari/Exercise_1/Task2_imagy/domain"

type DownloadImageFromURLRequest struct {
	URLPath   string
	LocalName string
	DstPath   string
}

type DownloadImageFromURLResponse struct {
	ImageName string
}

type ListImageRequest struct{}

type ListImageResponse struct {
	Images []domain.Image
}

type DownloadImageRequest struct {
	ImageName string
}

type DownloadImageResponse struct{}
