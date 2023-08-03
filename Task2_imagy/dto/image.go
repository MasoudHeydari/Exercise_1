package dto

type DownloadImageRequest struct {
	URLPath   string
	LocalName string
	DstPath   string
}

type DownloadImageResponse struct {
	ImageName string
}
