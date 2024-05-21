package debtCollectorRepository

import (
	"database/sql"
	"fmt"
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
	"fp_pinjaman_online/src/debtCollector"
	"strings"
	"time"
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
			return debtCollectorEntity.Tugas{}, fmt.Errorf("tugas with id: %v not found", tugasId)
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

func (repo *debtCollectorRepository) UpdateLogTugasById(storedLog debtCollectorEntity.LogTugas, updateLogPayload debtCollectorDto.UpdateLogTugasPayload) error {
	if strings.TrimSpace(updateLogPayload.Description) != "" {
		storedLog.Description = updateLogPayload.Description
	}
	query := "UPDATE log_tugas SET description = $1,updated_at = $2 WHERE id = $3"
	_, err := repo.db.Exec(query, storedLog.Description, time.Now(), storedLog.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *debtCollectorRepository) SoftDeleteLogTugasById(logTugasId string) error {
	query := "UPDATE log_tugas SET deleted_at = $1 WHERE id = $2"
	_, err := repo.db.Exec(query, time.Now(), logTugasId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *debtCollectorRepository) SelectLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error) {
	var logTugas debtCollectorEntity.LogTugas
	query := "SELECT id,tugas_id,description,created_at,updated_at FROM log_tugas WHERE id = $1 AND deleted_at IS NULL"
	err := repo.db.QueryRow(query, logTugasId).Scan(&logTugas.ID, &logTugas.TugasId, &logTugas.Description, &logTugas.CreatedAt, &logTugas.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return debtCollectorEntity.LogTugas{}, fmt.Errorf("log tugas with id: %v not found", logTugasId)
		}
		return debtCollectorEntity.LogTugas{}, err
	}
	return logTugas, nil
}
