package domain

import "time"

type Image struct {
	ID            int64
	OriginalURL   string
	LocalName     string
	FileExtension string
	FileSize      int64
	DownloadDate  time.Time
}
