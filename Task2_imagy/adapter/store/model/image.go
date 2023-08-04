package model

import (
	"time"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/domain"
)

type Image struct {
	ID            int64
	OriginalURL   string
	LocalName     string
	FileExtension string
	FileSize      int64
	DownloadDate  time.Time
}

func (i *Image) ToDomainImage() domain.Image {
	return domain.Image{
		ID:            i.ID,
		OriginalURL:   i.OriginalURL,
		LocalName:     i.LocalName,
		FileExtension: i.FileExtension,
		FileSize:      i.FileSize,
		DownloadDate:  i.DownloadDate,
	}
}
