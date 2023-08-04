package store

import (
	"context"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store/model"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
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
