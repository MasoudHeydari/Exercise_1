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
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/config"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/domain"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
)

var _ contract.ImageInteractor = &Interactor{}

type Interactor struct {
	imageStoreInteractor contract.ImageStoreInteractor
}

// New creates new contract.ImageInteractor.
func New(imageStoreInteractor contract.ImageStoreInteractor) contract.ImageInteractor {
	return &Interactor{
		imageStoreInteractor: imageStoreInteractor,
	}
}

// Upload inserts a new image to the Imagy storage and DB.
func (i *Interactor) Upload(ctx context.Context, req dto.UploadImageRequest) (dto.UploadImageResponse, error) {
	if req.ImageFile == nil {
		return dto.UploadImageResponse{}, fmt.Errorf("file cannot be nil")
	}
	imageAbsPath := filepath.Join(config.GetUserContentUploadPath(), req.ImageFile.Filename)
	_, err := os.Stat(imageAbsPath)
	if err == nil {
		return dto.UploadImageResponse{}, fmt.Errorf("409 - confilict, image already exists in filesystem")
	}
	contentType := req.ImageFile.Header.Get("Content-Type")
	if err = checkContentType(contentType); err != nil {
		return dto.UploadImageResponse{}, err
	}
	maxImageSizeInBytes := config.GetMaxImageSizeInBytes()
	fileSize := req.ImageFile.Size
	if fileSize > maxImageSizeInBytes {
		return dto.UploadImageResponse{}, fmt.Errorf("image is too big, max image size is %d MB", maxImageSizeInBytes>>20)
	}
	fileExt, err := extractFileExtension(contentType)
	if err != nil {
		return dto.UploadImageResponse{}, err
	}
	f, err := os.Create(imageAbsPath)
	if err != nil {
		return dto.UploadImageResponse{}, err
	}
	tmpFile, err := req.ImageFile.Open()
	if err != nil {
		return dto.UploadImageResponse{}, err
	}
	defer tmpFile.Close()
	_, err = io.Copy(f, tmpFile)
	if err != nil {
		return dto.UploadImageResponse{}, err
	}
	imageName := req.ImageFile.Filename
	imageURL := fmt.Sprintf("http://%s/api/v1/images/%s", config.GetServerAddressAndPort(), imageName)
	image := model.Image{
		OriginalURL:   imageURL,
		LocalName:     imageName,
		FileExtension: fileExt,
		FileSize:      fileSize,
	}
	modelImage, err := i.imageStoreInteractor.Create(ctx, image)
	if err != nil {
		_ = os.Remove(imageAbsPath)
		return dto.UploadImageResponse{}, err
	}
	return dto.UploadImageResponse{Image: modelImage.ToDomainImage()}, nil
}

// Download returns the image content to the client if it exits in Imagy.
func (i *Interactor) Download(ctx context.Context, req dto.DownloadImageRequest) (dto.DownloadImageResponse, error) {
	str := req.RootStoragePath
	img, err := i.imageStoreInteractor.DoesExit(ctx, req.ImageName)
	if err != nil {
		// not found
		return dto.DownloadImageResponse{}, err
	}
	if strings.Contains(img.OriginalURL, config.GetServerAddressAndPort()) {
		str = config.GetUserContentUploadPath()
	}
	ImageAbsPath := fmt.Sprintf("%s/%s", str, req.ImageName)
	fInfo, err := os.Stat(ImageAbsPath)
	if err != nil {
		return dto.DownloadImageResponse{}, fmt.Errorf("404 - %w", err)
	}
	if fInfo.Name() != req.ImageName {
		return dto.DownloadImageResponse{}, fmt.Errorf("404 - failed to locate %s file", req.ImageName)
	}

	return dto.DownloadImageResponse{ImageAbsPath: ImageAbsPath}, nil
}

// List returns the list of stored images in Imagy.
func (i *Interactor) List(ctx context.Context, req dto.ListImageRequest) (dto.ListImageResponse, error) {
	modelImages, err := i.imageStoreInteractor.List(ctx)
	if err != nil {
		return dto.ListImageResponse{}, err
	}

	return dto.ListImageResponse{
		Images: convertModelImagesToDomainImages(modelImages),
	}, nil
}

// DownloadFromURL makes a GET request to the input URL and if the response code is 200 OK,
// stores the image in storage and insert a new row in DB.
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
	if err = checkContentType(contentType); err != nil {
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
	_, err = i.imageStoreInteractor.Create(ctx, image)
	if err != nil {
		return dto.DownloadImageFromURLResponse{}, err
	}
	return dto.DownloadImageFromURLResponse{ImageName: fileName}, nil
}

// checkContentType checks request content-type.
func checkContentType(contentType string) error {
	if len(contentType) == 0 {
		return fmt.Errorf("request's content-type cannot be ampty")
	}
	if !strings.Contains(contentType, "image") {
		return fmt.Errorf("unsupported content-type for image downloader - request's content-type is: %s", contentType)
	}
	return nil
}

// extractFileExtension extract the file's extension from request's header.
func extractFileExtension(s string) (string, error) {
	i := strings.Split(s, "/")
	if len(i) <= 1 {
		return "", fmt.Errorf("failed to extraxt file extension")
	}
	return i[1], nil
}

// convertModelImagesToDomainImages converts the model.Image to domain.Image.
func convertModelImagesToDomainImages(modelImages []model.Image) []domain.Image {
	domainImages := make([]domain.Image, 0)
	for _, modelImage := range modelImages {
		domainImages = append(domainImages, modelImage.ToDomainImage())
	}
	return domainImages
}
