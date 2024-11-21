package repository

import (
	"context"

	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/domain/repository"
)

type ContainerRepository interface {
	InsertFunction(metadata model.FunctionMetadata) error
}

type containerRepository struct {
	functionRepo repository.FunctionRepository
	fileRepo     repository.FileRepository
	packageRepo  repository.PackageRepository
}

func NewContainerRepository(functionRepo repository.FunctionRepository, fileRepo repository.FileRepository, packageRepo repository.PackageRepository) ContainerRepository {
	return &containerRepository{
		functionRepo: functionRepo,
		fileRepo:     fileRepo,
		packageRepo:  packageRepo,
	}
}

func (c *containerRepository) InsertFunction(metadata model.FunctionMetadata) error {
	ctx := context.Background()

	tx, err := c.functionRepo.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	packageID, err := c.packageRepo.GetPackageID(metadata.File.Package.Name, metadata.File.Package.Link, tx)
	if err != nil {
		return err
	}

	fileID, err := c.fileRepo.GetFileID(metadata.File.Name, packageID, tx)
	if err != nil {
		return err
	}

	return c.functionRepo.InsertFunction(metadata.Name, metadata.Signature, fileID, metadata.Comment, tx)
}
