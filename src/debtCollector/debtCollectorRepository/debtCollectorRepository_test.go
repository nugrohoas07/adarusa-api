package debtCollectorRepository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
	"fp_pinjaman_online/src/debtCollector"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type DebtCollectorRepoTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo debtCollector.DebtCollectorRepository
}

func (s *DebtCollectorRepoTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	s.repo = NewDebtCollectorRepository(db)
}

func TestDebtCollectorRepoTestSuite(t *testing.T) {
	suite.Run(t, new(DebtCollectorRepoTestSuite))
}

func (s *DebtCollectorRepoTestSuite) TestSelectAllLateDebitur_Success() {
	dcCityMock := "Malang"
	pageMock := 1
	sizeMock := 1
	rows := sqlmock.NewRows([]string{"id", "fullname", "address", "unpaid"}).
		AddRow("1", "user satu", "address 1", 1000000)

	query := regexp.QuoteMeta(`SELECT u.id,du.fullname,du.address,SUM(c.jumlah_bayar) AS unpaid
		FROM cicilan c
		INNER JOIN pinjaman p ON c.pinjaman_id = p.id
		INNER JOIN users u ON p.user_id = u.id
		INNER JOIN detail_users du ON u.id = du.user_id
		LEFT JOIN claim_tugas ct ON u.id = ct.user_id AND ct.status = 'ongoing'
		WHERE c.tanggal_jatuh_tempo < $1 AND c.status = 'unpaid'
		AND du.city ILIKE '%' || $2 || '%'
		AND ct.user_id IS NULL
		GROUP BY u.id, du.fullname, du.address
		LIMIT $3 OFFSET $4`)
	countQuery := regexp.QuoteMeta(`SELECT COUNT(*) FROM (SELECT DISTINCT ON (u.id) u.id
		FROM cicilan c
		INNER JOIN pinjaman p ON c.pinjaman_id = p.id
		INNER JOIN users u ON p.user_id = u.id
		INNER JOIN detail_users du ON u.id = du.user_id
		LEFT JOIN claim_tugas ct ON u.id = ct.user_id AND ct.status = 'ongoing'
		WHERE c.tanggal_jatuh_tempo < $1 AND c.status = 'unpaid'
		AND du.city ILIKE '%' || $2 || '%'
		AND ct.user_id IS NULL);`)

	offset := (pageMock - 1) * sizeMock
	s.mock.ExpectQuery(query).WithArgs(AnyTime{}, dcCityMock, sizeMock, offset).WillReturnRows(rows)
	s.mock.ExpectQuery(countQuery).WithArgs(AnyTime{}, dcCityMock).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))
	expData := []debtCollectorEntity.LateDebtor{
		{
			ID:           "1",
			FullName:     "user satu",
			Address:      "address 1",
			UnpaidAmount: 1000000,
		},
	}
	expPaging := json.Paging{
		Page:      1,
		TotalData: 3,
	}

	data, paging, err := s.repo.SelectAllLateDebitur(dcCityMock, pageMock, sizeMock)
	s.NoError(err)
	s.Equal(expPaging, paging)
	s.Equal(expData, data)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DebtCollectorRepoTestSuite) TestInsertLogTugas_Success() {
	mockPayload := debtCollectorDto.NewLogTugasPayload{
		TugasId:     "1",
		Description: "test log",
	}

	query := regexp.QuoteMeta(`INSERT INTO log_tugas(tugas_id,description) VALUES ($1, $2)`)
	s.mock.ExpectExec(query).WithArgs(mockPayload.TugasId, mockPayload.Description).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.InsertLogTugas(mockPayload)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DebtCollectorRepoTestSuite) TestInsertLogTugas_Fail() {
	mockPayload := debtCollectorDto.NewLogTugasPayload{
		TugasId:     "1",
		Description: "test log",
	}

	query := regexp.QuoteMeta(`INSERT INTO log_tugas(tugas_id,description) VALUES ($1, $2)`)
	expError := fmt.Errorf("internal server error")
	s.mock.ExpectExec(query).WithArgs(mockPayload.TugasId, mockPayload.Description).WillReturnError(expError)

	err := s.repo.InsertLogTugas(mockPayload)
	s.Error(err)
	s.Equal(expError, err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DebtCollectorRepoTestSuite) TestUpdateLogTugasById_Success() {
	mockStoredLog := debtCollectorEntity.LogTugas{
		ID:          "1",
		TugasId:     "1",
		Description: "test log",
	}
	mockPayload := debtCollectorDto.UpdateLogTugasPayload{
		Description: "test log update",
	}

	query := regexp.QuoteMeta(`UPDATE log_tugas SET description = $1,updated_at = $2 WHERE id = $3`)
	s.mock.ExpectExec(query).WithArgs(mockPayload.Description, AnyTime{}, mockStoredLog.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.UpdateLogTugasById(mockStoredLog, mockPayload)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DebtCollectorRepoTestSuite) TestUpdateLogTugasById_Fail() {
	mockStoredLog := debtCollectorEntity.LogTugas{
		ID:          "1",
		TugasId:     "1",
		Description: "test log",
	}
	mockPayload := debtCollectorDto.UpdateLogTugasPayload{
		Description: "test log update",
	}

	query := regexp.QuoteMeta(`UPDATE log_tugas SET description = $1,updated_at = $2 WHERE id = $3`)
	expError := fmt.Errorf("internal server error")
	s.mock.ExpectExec(query).WithArgs(mockPayload.Description, AnyTime{}, mockStoredLog.ID).WillReturnError(expError)

	err := s.repo.UpdateLogTugasById(mockStoredLog, mockPayload)
	s.Error(err)
	s.Equal(expError, err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DebtCollectorRepoTestSuite) TestSoftDeleteLogTugasById_Success() {
	mockLogTugasId := "1"

	query := regexp.QuoteMeta(`UPDATE log_tugas SET deleted_at = $1 WHERE id = $2`)
	s.mock.ExpectExec(query).WithArgs(AnyTime{}, mockLogTugasId).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.SoftDeleteLogTugasById(mockLogTugasId)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
}
