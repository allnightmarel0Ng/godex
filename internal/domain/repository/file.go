package repository

import (
	"errors"

	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v4"
)

const (
	selectFile = 
		`SELECT id 
		FROM public.files 
		WHERE name = $1 AND package_id = $2;`

	insertFile = 
		`INSERT INTO public.files (name, package_id) 
		VALUES ($1, $2) 
		RETURNING id;`
)

type FileRepository interface {
	GetFileID(name string, packageID int64, tx postgres.Transaction) (int64, error)
}

type fileRepository struct {
	db *postgres.Database
}

func NewFileRepository(db *postgres.Database) FileRepository {
	return &fileRepository{
		db: db,
	}
}

func (f *fileRepository) GetFileID(name string, packageID int64, tx postgres.Transaction) (int64, error) {
	var fileID int64
	err := tx.QueryRow(selectFile, name, packageID).Scan(&fileID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return fileID, err
		}

		err = tx.QueryRow(insertFile, name, packageID).Scan(&fileID)
		if err != nil {
			return fileID, err
		}
	}

	return fileID, nil
}