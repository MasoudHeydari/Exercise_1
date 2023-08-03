package domain

import "time"

type Image struct {
	ID            int64     `json:"id"`
	OriginalURL   string    `json:"original_url"`
	LocalName     string    `json:"local_name"`
	FileExtension string    `json:"file_extension"`
	FileSize      int64     `json:"file_size"`
	DownloadDate  time.Time `json:"download_date"`
}
