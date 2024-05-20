package debtCollectorRepository

import (
	"database/sql"
	"fmt"
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
	"fp_pinjaman_online/src/debtCollector"
)

type debtCollectorRepository struct {
	db *sql.DB
}

func NewDebtCollectorRepository(db *sql.DB) debtCollector.DebtCollectorRepository {
	return &debtCollectorRepository{db}
}

func (repo *debtCollectorRepository) SelectTugasById(tugasId string) (debtCollectorEntity.Tugas, error) {
	var tugas debtCollectorEntity.Tugas
	query := "SELECT id,user_id,collector_id,status FROM claim_tugas WHERE id = $1"
	err := repo.db.QueryRow(query, tugasId).Scan(&tugas.ID, &tugas.UserId, &tugas.CollectorId, &tugas.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return debtCollectorEntity.Tugas{}, fmt.Errorf("tugas not found")
		}
		return debtCollectorEntity.Tugas{}, err
	}
	return tugas, nil
}

func (repo *debtCollectorRepository) InsertLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error {
	query := "INSERT INTO log_tugas(tugas_id,description) VALUES ($1, $2)"
	_, err := repo.db.Exec(query, newLogPayload.TugasId, newLogPayload.Description)
	if err != nil {
		return err
	}
	return nil
}
