package debiturRepository_test

import (
	"database/sql"
	"fp_pinjaman_online/src/debitur/debiturRepository"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddPengajuanPinjaman(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := debiturRepository.NewDebiturRepository(db)

	query := regexp.QuoteMeta("INSERT INTO pinjaman (user_id, jumlah_pinjaman, tenor, description, bunga_per_bulan, status_pengajuan) VALUES ($1, $2, $3, $4, $5, 'pending')")
	mock.ExpectExec(query).WithArgs(1, 1000.0, 6, "Test description", 0.09).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddPengajuanPinjaman(1, 1000.0, 6, "Test description")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetPengajuanPinjaman(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := debiturRepository.NewDebiturRepository(db)

	query := regexp.QuoteMeta("SELECT id, jumlah_pinjaman, tenor, description, bunga_per_bulan, status_pengajuan, created_at, updated_at FROM pinjaman WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC")
	rows := sqlmock.NewRows([]string{"id", "jumlah_pinjaman", "tenor", "description", "bunga_per_bulan", "status_pengajuan", "created_at", "updated_at"}).
		AddRow(1, 1000.0, 6, "Test description", 0.09, "pending", time.Now(), time.Now())

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	result, err := repo.GetPengajuanPinjaman(1)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCicilan(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := debiturRepository.NewDebiturRepository(db)

	queryCount := regexp.QuoteMeta("SELECT COUNT(*) FROM cicilan WHERE pinjaman_id = $1")
	querySelect := regexp.QuoteMeta("SELECT id, pinjaman_id, tanggal_jatuh_tempo, tanggal_selesai_bayar, jumlah_bayar, status FROM cicilan WHERE pinjaman_id = $1")

	mock.ExpectQuery(queryCount).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	rows := sqlmock.NewRows([]string{"id", "pinjaman_id", "tanggal_jatuh_tempo", "tanggal_selesai_bayar", "jumlah_bayar", "status"}).
		AddRow(1, 1, time.Now(), time.Now(), 1000.0, "unpaid")
	mock.ExpectQuery(querySelect).WithArgs("1").WillReturnRows(rows)

	result, paging, err := repo.GetCicilan(1, 10, 0, "1", "")
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, 1, paging.TotalData)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCicilanPayment(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := debiturRepository.NewDebiturRepository(db)

	querySelect := regexp.QuoteMeta("SELECT id FROM cicilan WHERE pinjaman_id = $1 AND status = 'unpaid'")
	queryUpdate := regexp.QuoteMeta("UPDATE cicilan SET status = 'paid' WHERE id = $1")

	mock.ExpectQuery(querySelect).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec(queryUpdate).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.CicilanPayment(1, 1000.0)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCicilanPayment_NoCicilan(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := debiturRepository.NewDebiturRepository(db)

	querySelect := regexp.QuoteMeta("SELECT id FROM cicilan WHERE pinjaman_id = $1 AND status = 'unpaid'")

	mock.ExpectQuery(querySelect).WithArgs(1).WillReturnError(sql.ErrNoRows)

	_, err = repo.CicilanPayment(1, 1000.0)
	assert.Error(t, err)
	assert.Equal(t, "anda tidak mempunyai cicilan", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}
