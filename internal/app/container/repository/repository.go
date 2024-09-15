package repository

import (
	"errors"

	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
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
	var packageID int64
	row := c.db.QueryRow(
		"SELECT id FROM public.packages WHERE name == $1 AND link == $2;",
		metadata.File.Package.Name,
		metadata.File.Package.Link)
	if row == nil {
		result, err := c.db.Exec(
			"INSERT INTO public.packages (name, link) VALUES ($1, $2);",
			metadata.File.Package.Name,
			metadata.File.Package.Link)

		if err != nil {
			return err
		}

		if result == nil {
			return errors.New("unable to INSERT new package in to database")
		}

		packageID, err = result.LastInsertId()
		if err != nil {
			return err
		}
	} else {
		err := row.Scan(&packageID)
		if err != nil {
			return err
		}
	}

	var fileID int64
	row = c.db.QueryRow(
		"SELECT id FROM public.files WHERE name == $1 AND package_id == $2;",
		metadata.File.Name,
		packageID)
	if row == nil {
		result, err := c.db.Exec(
			"INSERT INTO public.files (name, link) VALUES ($1, $2);",
			metadata.File.Name,
			fileID)

		if err != nil {
			return err
		}

		if result == nil {
			return errors.New("unable to INSERT new file in to database")
		}

		fileID, err = result.LastInsertId()
		if err != nil {
			return err
		}
	} else {
		err := row.Scan(&fileID)
		if err != nil {
			return err
		}
	}

	row = c.db.QueryRow(
		"SELECT id FROM public.functions WHERE name == $1 AND signature == $2 AND file_id == $3 AND comment == $4;",
		metadata.Name, metadata.Signature, fileID, metadata.Comment)
	if row == nil {
		result, err := c.db.Exec(
			"INSERT INTO public.functions (name, signature, file_id, comment) VALUES ($1, $2, $3, $4);",
			metadata.Name, metadata.Signature, fileID, metadata.Comment)

		if err != nil {
			return err
		}

		if result == nil {
			return errors.New("unable to INSERT new function in to database")
		}
	}

	return nil
}

func (c *containerRepository) GetFunctionBySignature(signature string) ([]model.FunctionMetadata, error) {
	rows, err := c.db.Query("SELECT * FROM public.functions WHERE signature == $1;", signature)
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

		row := c.db.QueryRow("SELECT * FROM public.files WHERE id == $1;", fileID)
		if row == nil {
			return nil, errors.New("file for such signature wasn't found")
		}

		var fileName string
		var packageID int
		err = row.Scan(&fileID, &fileName, &packageID)
		if err != nil {
			return nil, err
		}

		row = c.db.QueryRow("SELECT * FROM public.packages WHERE id == $1;", packageID)
		if row == nil {
			return nil, errors.New("package for such signature wasn't found")
		}

		var packageName, link string
		err = row.Scan(&packageID, &packageName, &link)
		if err != nil {
			return nil, err
		}

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

	return result, nil
}
