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
	imageStoreInteractor contract.ImageStoreInteractor
}

func New(imageStoreInteractor contract.ImageStoreInteractor) contract.ImageInteractor {
	return &Interactor{
		imageStoreInteractor: imageStoreInteractor,
	}
}

func (i *Interactor) Download(ctx context.Context, req dto.DownloadImageRequest) (dto.DownloadImageResponse, error) {
	ImageAbsPath := fmt.Sprintf("%s/%s", req.RootStoragePath, req.ImageName)
	fInfo, err := os.Stat(ImageAbsPath)
	if err != nil {
		return dto.DownloadImageResponse{}, fmt.Errorf("404 - %w", err)
	}
	if fInfo.Name() != req.ImageName {
		return dto.DownloadImageResponse{}, fmt.Errorf("404 - failed to locate %s file", req.ImageName)
	}
	err = i.imageStoreInteractor.DoesExit(ctx, req.ImageName)
	if err != nil {
		return dto.DownloadImageResponse{}, err
	}
	return dto.DownloadImageResponse{ImageAbsPath: ImageAbsPath}, nil
}

func (i *Interactor) List(ctx context.Context, req dto.ListImageRequest) (dto.ListImageResponse, error) {
	modelImages, err := i.imageStoreInteractor.List(ctx)
	if err != nil {
		return dto.ListImageResponse{}, err
	}

	return dto.ListImageResponse{
		Images: convertModelImagesToDomainImages(modelImages),
	}, nil
}

func (i *Interactor) DownloadFromURL(ctx context.Context, req dto.DownloadImageFromURLRequest) (dto.DownloadImageFromURLResponse, error) {
	res, err := http.Get(req.URLPath)
	if err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	defer res.Body.Close()
	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		return dto.DownloadImageFromURLResponse{}, fmt.Errorf("failed to downlaod image from %s url, status code: %d", req.URLPath, statusCode)
	}
	contentType := res.Header.Get("Content-Type")
	if err = checkResponseContentType(contentType); err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	fileExt, err := extractFileExtension(contentType)
	if err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	fileName := fmt.Sprintf("%s.%s", req.LocalName, fileExt)
	f, err := os.Create(filepath.Join(req.DstPath, fileName))
	if err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	fileSize, err := io.Copy(f, res.Body)
	if err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	image := model.Image{
		OriginalURL:   req.URLPath,
		LocalName:     fileName,
		FileExtension: fileExt,
		FileSize:      fileSize,
	}
	err = i.imageStoreInteractor.Create(ctx, image)
	if err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	return dto.DownloadImageFromURLResponse{ImageName: fileName}, nil
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
