package image

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store/model"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/domain"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
)

var _ contract.ImageInteractor = &Interactor{}

type Interactor struct {
	db contract.ImageStoreInteractor
}

func New(db contract.ImageStoreInteractor) contract.ImageInteractor {
	return &Interactor{
		db: db,
	}
}

func (i *Interactor) List(ctx context.Context, req dto.ListImageRequest) (dto.ListImageResponse, error) {
	modelImages, err := i.db.List(ctx)
	if err != nil {
		return dto.ListImageResponse{}, err
	}

	return dto.ListImageResponse{
		Images: convertModelImagesToDomainImages(modelImages),
	}, nil
}

func (i *Interactor) Download(ctx context.Context, req dto.DownloadImageRequest) (dto.DownloadImageResponse, error) {
	res, err := http.Get(req.URLPath)
	if err != nil {
		return dto.DownloadImageResponse{}, err
	}
	defer res.Body.Close()
	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		return dto.DownloadImageResponse{}, fmt.Errorf("failed to downlaod image from %s url, status code: %d", req.URLPath, statusCode)
	}
	contentType := res.Header.Get("Content-Type")
	if err = checkResponseContentType(contentType); err != nil {
		return dto.DownloadImageResponse{}, err
	}
	fileExt, err := extractFileExtension(contentType)
	if err != nil {
		return dto.DownloadImageResponse{}, err
	}
	fileName := fmt.Sprintf("%s.%s", req.LocalName, fileExt)
	f, err := os.Create(filepath.Join(req.DstPath, fileName))
	if err != nil {
		return dto.DownloadImageResponse{}, err
	}
	fileSize, err := io.Copy(f, res.Body)
	if err != nil {
		return dto.DownloadImageResponse{}, err
	}
	image := model.Image{
		OriginalURL:   req.URLPath,
		LocalName:     f.Name(),
		FileExtension: fileExt,
		FileSize:      fileSize,
	}
	err = i.db.Create(ctx, image)
	if err != nil {
		return dto.DownloadImageResponse{}, err
	}
	return dto.DownloadImageResponse{ImageName: fileName}, nil
}

func checkResponseContentType(contentType string) error {
	if len(contentType) == 0 {
		return fmt.Errorf("request's content-type cannot be ampty")
	}
	if !strings.Contains(contentType, "image") {
		return fmt.Errorf("unsupported content-type for image downloader - request's content-type is: %s", contentType)
	}
	return nil
}

func extractFileExtension(s string) (string, error) {
	i := strings.Split(s, "/")
	if len(i) <= 1 {
		return "", fmt.Errorf("failed to extraxt file extension")
	}
	return i[1], nil
}

func convertModelImagesToDomainImages(modelImages []model.Image) []domain.Image {
	domainImages := make([]domain.Image, 0)
	for _, modelImage := range modelImages {
		domainImages = append(domainImages, modelImage.ToDomainImage())
	}
	return domainImages
}
