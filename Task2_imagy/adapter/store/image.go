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
