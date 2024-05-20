package checkHealthRepository

import (
	"database/sql"
	"fp_pinjaman_online/src/checkHealth"
)

type checkHealthRepository struct {
	db *sql.DB
}

func NewCheckHealthRepository(db *sql.DB) checkHealth.CheckHealthRepository {
	return &checkHealthRepository{db}
}

func (repo *checkHealthRepository) RetrieveVersion() (string, error) {
	version, query := "", "SELECT version FROM version_apps"
	if err := repo.db.QueryRow(query).Scan(&version); err != nil {
		return "", err
	}

	return version, nil
}
