package store

import (
	"context"
	"fmt"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store/model"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/ent"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/ent/image"
)

type Interactor struct {
	db *Database
}

func NewImageStoreInteractor(db *Database) contract.ImageStoreInteractor {
	return &Interactor{
		db: db,
	}
}

func (i *Interactor) Create(ctx context.Context, image model.Image) error {
	_, err := i.db.Client.Image.Create().
		SetOriginalURL(image.OriginalURL).
		SetLocalName(image.LocalName).
		SetFileExtension(image.FileExtension).
		SetFileSize(image.FileSize).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i *Interactor) List(ctx context.Context) ([]model.Image, error) {
	entImages, err := i.db.Client.Image.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	modelImages := make([]model.Image, 0)
	for _, entImg := range entImages {
		modelImg := model.Image{
			ID:            entImg.ID,
			OriginalURL:   entImg.OriginalURL,
			LocalName:     entImg.LocalName,
			FileExtension: entImg.FileExtension,
			FileSize:      entImg.FileSize,
			DownloadDate:  entImg.DownloadDate,
		}
		modelImages = append(modelImages, modelImg)
	}
	return modelImages, nil
}

func (i *Interactor) DoesExit(ctx context.Context, imageName string) error {
	doesExits, err := i.db.Client.Image.Query().Where(image.LocalNameEQ(imageName)).Exist(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// image not exits in DB - 404 Not Found error code
			return fmt.Errorf("404 - %w", err)
		}
		// server faced with an internal error - 500 internal server error
		return fmt.Errorf("500 - %w", err)
	}
	if !doesExits {
		return fmt.Errorf("404 - %s not exits in database", imageName)
	}
	return nil
}
