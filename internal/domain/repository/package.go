package repository

import (
	"errors"

	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v4"
)

const (
	selectPackage = `SELECT id 
		FROM public.packages 
		WHERE name = $1 AND link = $2;`

	insertPackage = `INSERT INTO public.packages (name, link) 
		VALUES ($1, $2) 
		RETURNING id;`
)

type PackageRepository interface {
	GetPackageID(name, link string, tx postgres.Transaction) (int64, error)
}

type packageRepository struct {
	db *postgres.Database
}

func NewPackageRepository(db *postgres.Database) PackageRepository {
	return &packageRepository{
		db: db,
	}
}

func (p *packageRepository) GetPackageID(name, link string, tx postgres.Transaction) (int64, error) {
	var packageID int64
	err := tx.QueryRow(selectPackage, name, link).Scan(&packageID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return packageID, err
		}

		err = tx.QueryRow(insertPackage, name, link).Scan(&packageID)
	}

	return packageID, err
}
