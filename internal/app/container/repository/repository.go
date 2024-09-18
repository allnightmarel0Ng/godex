package repository

import (
	"context"
	"errors"
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v4"
)

type ContainerRepository interface {
	InsertFunction(metadata model.FunctionMetadata) error
	GetFunctionBySignature(signature string) ([]model.FunctionMetadata, error)
}

type containerRepository struct {
	db *postgres.Database
}

func NewContainerRepository(db *postgres.Database) ContainerRepository {
	return &containerRepository{
		db: db,
	}
}

func (c *containerRepository) InsertFunction(metadata model.FunctionMetadata) error {
	ctx := context.Background()

	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	var packageID int64
	err = tx.QueryRow(ctx, "SELECT id FROM public.packages WHERE name = $1 AND link = $2;",
		metadata.File.Package.Name,
		metadata.File.Package.Link).Scan(&packageID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		err = tx.QueryRow(ctx, "INSERT INTO public.packages (name, link) VALUES ($1, $2) RETURNING id;",
			metadata.File.Package.Name,
			metadata.File.Package.Link).Scan(&packageID)

		if err != nil {
			return err
		}
	}

	var fileID int64
	err = tx.QueryRow(ctx, "SELECT id FROM public.files WHERE name = $1 AND package_id = $2;",
		metadata.File.Name,
		packageID).Scan(&fileID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		err = tx.QueryRow(ctx, "INSERT INTO public.files (name, package_id) VALUES ($1, $2) RETURNING id;",
			metadata.File.Name,
			packageID).Scan(&fileID)

		if err != nil {
			return err
		}
	}

	var functionID int64
	err = tx.QueryRow(ctx, "SELECT id FROM public.functions WHERE name = $1 AND signature = $2 AND file_id = $3 AND comment = $4;",
		metadata.Name, metadata.Signature, fileID, metadata.Comment).Scan(&functionID)
	if err != nil {
		_, err = tx.Exec(ctx, "INSERT INTO public.functions (name, signature, file_id, comment) VALUES ($1, $2, $3, $4);",
			metadata.Name, metadata.Signature, fileID, metadata.Comment)
	}

	return err
}

func (c *containerRepository) GetFunctionBySignature(signature string) ([]model.FunctionMetadata, error) {
	rows, err := c.db.Query("SELECT * FROM public.functions WHERE signature = $1;", signature)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("function for such signature wasn't found")
	}
	defer rows.Close()

	var result []model.FunctionMetadata

	for rows.Next() {
		var functionName, sign, comment string
		var functionID, fileID int
		err = rows.Scan(&functionID, &functionName, &sign, &fileID, &comment)
		if err != nil {
			return nil, err
		}

		var fileName string
		var packageID int
		err = c.db.QueryRow("SELECT * FROM public.files WHERE id = $1;", fileID).Scan(&fileID, &fileName, &packageID)
		if err != nil {
			return nil, err
		}

		var packageName, link string
		err = c.db.QueryRow("SELECT * FROM public.packages WHERE id = $1;", packageID).Scan(&packageID, &packageName, &link)

		result = append(result, model.FunctionMetadata{
			Name:      functionName,
			Signature: sign,
			Comment:   comment,
			File: model.FileMetadata{
				Name: fileName,
				Package: model.PackageMetadata{
					Name: packageName,
					Link: link,
				},
			},
		})
	}

	return result, err
}
