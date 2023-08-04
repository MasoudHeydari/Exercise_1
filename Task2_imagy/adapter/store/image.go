package store

import (
	"context"
	"fmt"

	entSql "entgo.io/ent/dialect/sql"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store/model"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/contract"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/ent"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/ent/image"
	"github.com/labstack/gommon/log"
)

type Interactor struct {
	db *Database
}

func NewImageStoreInteractor(db *Database) contract.ImageStoreInteractor {
	return &Interactor{
		db: db,
	}
}

func (i *Interactor) Create(ctx context.Context, img model.Image) (domainImage model.Image, err error) {
	// we have to DB queries, execute them in single db transaction
	tx, err := i.db.Client.BeginTx(ctx, &entSql.TxOptions{})
	if err != nil {
		return model.Image{}, nil
	}
	defer func() {
		if err == nil {
			return
		}

		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Errorf("Create rollback failed - error: %w\n", rollBackErr)
		}
	}()
	entImg, err := tx.Image.Query().Where(image.LocalNameEQ(img.LocalName)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		// server faced with an internal error - 500 internal server error
		return model.Image{}, fmt.Errorf("500 - %w", err)
	}

	if entImg != nil {
		domainImage = entImageToModelImage(*entImg)
		return domainImage, fmt.Errorf("409 - confilict, image already exists in database")
	}

	// we can guarantee that there is no image in db with the name of image.LocalName
	entImg, err = tx.Image.Create().
		SetOriginalURL(img.OriginalURL).
		SetLocalName(img.LocalName).
		SetFileExtension(img.FileExtension).
		SetFileSize(img.FileSize).
		Save(ctx)
	if err != nil {
		return model.Image{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.Image{}, err
	}

	domainImage = entImageToModelImage(*entImg)
	return
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

func (i *Interactor) DoesExit(ctx context.Context, imageName string) (model.Image, error) {
	entImg, err := i.db.Client.Image.Query().Where(image.LocalNameEQ(imageName)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// image not exits in DB - 404 Not Found error code
			return model.Image{}, fmt.Errorf("404 - %s not exits in database %w", imageName, err)
		}
		// server faced with an internal error - 500 internal server error
		return model.Image{}, fmt.Errorf("500 - %w", err)
	}
	return model.Image{
		ID:            entImg.ID,
		OriginalURL:   entImg.OriginalURL,
		LocalName:     entImg.LocalName,
		FileExtension: entImg.FileExtension,
		FileSize:      entImg.FileSize,
		DownloadDate:  entImg.DownloadDate,
	}, nil
}

func entImageToModelImage(entImg ent.Image) model.Image {
	return model.Image{
		ID:            entImg.ID,
		OriginalURL:   entImg.OriginalURL,
		LocalName:     entImg.LocalName,
		FileExtension: entImg.FileExtension,
		FileSize:      entImg.FileSize,
		DownloadDate:  entImg.DownloadDate,
	}
}
