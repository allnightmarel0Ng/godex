package repository

import (
	"context"
	"errors"

	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
)

const (
	selectFunction = 
		`SELECT id 
		FROM public.functions 
		WHERE name = $1 AND signature = $2 AND file_id = $3 AND comment = $4;`

	insertFunction = 
		`INSERT INTO public.functions (name, signature, file_id, comment) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;`

	selectBySignature = 
		`SELECT
			f.name AS function_name,
			f.signature AS signature,
			f.comment AS comment,
			file.name AS file_name,
			p.name AS package_name,
			p.link AS package_link
		FROM public.functions AS f
		JOIN public.files AS file ON file.id = f.file_id
		JOIN public.packages AS p ON p.id = file.package_id
		WHERE f.signature = $1`
)

type FunctionRepository interface {
	Begin(ctx context.Context) (postgres.Transaction, error)
	InsertFunction(name, signature string, fileID int64, comment string, tx postgres.Transaction) error
	GetFunctionsBySignature(signature string) ([]model.FunctionMetadata, error)
}

type functionRepository struct {
	db *postgres.Database
}

func NewFunctionRepository(db *postgres.Database) FunctionRepository {
	return &functionRepository{
		db: db,
	}
}

func (f *functionRepository) Begin(ctx context.Context) (postgres.Transaction, error) {
	return f.db.Begin(ctx)
}

func (f *functionRepository) InsertFunction(name, signature string, fileID int64, comment string, tx postgres.Transaction) error {
	var functionID int64
	err := tx.QueryRow(selectFunction, name, signature, fileID, comment).Scan(&functionID)
	if err != nil {
		_, err = tx.Exec(insertFunction, name, signature, fileID, comment)
	}
	return err
}

func (f *functionRepository) GetFunctionsBySignature(signature string) ([]model.FunctionMetadata, error) {
	rows, err := f.db.Query(selectBySignature, signature)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("function for such signature wasn't found")
	}
	defer rows.Close()

	var result []model.FunctionMetadata
	for rows.Next() {
		var functionName, signature, comment, fileName, packageName, packageLink string
		err = rows.Scan(&functionName, &signature, &comment, &fileName, &packageName, &packageLink)
		if err != nil {
			return nil, err
		}

		result = append(result, model.FunctionMetadata{
			Name: functionName,
			Signature: signature,
			Comment: comment,
			File: model.FileMetadata{
				Name: fileName,
				Package: model.PackageMetadata{
					Name: packageName,
					Link: packageLink,
				},
			},
		})
	}

	return result, nil
}