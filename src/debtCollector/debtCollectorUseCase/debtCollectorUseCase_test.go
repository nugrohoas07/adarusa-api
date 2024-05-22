package debtCollectorUseCase

import (
	"errors"
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
	"fp_pinjaman_online/src/debtCollector"
	"fp_pinjaman_online/src/debtCollector/debtCollectorRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DebtCollectorUseCaseSuite struct {
	suite.Suite
	dcRepoMock *debtCollectorRepository.DebtCollectorRepositoryMock
	usecase    debtCollector.DebtCollectorUseCase
}

func (s *DebtCollectorUseCaseSuite) SetupTest() {
	s.dcRepoMock = &debtCollectorRepository.DebtCollectorRepositoryMock{Mock: mock.Mock{}}
	s.usecase = NewDebtCollectorUseCase(s.dcRepoMock)
}

func TestUsersUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DebtCollectorUseCaseSuite))
}

func (s *DebtCollectorUseCaseSuite) TestCreateLogTugas_Success() {
	mockPayload := debtCollectorDto.NewLogTugasPayload{
		TugasId:     "1",
		Description: "tes isi log",
	}
	s.dcRepoMock.Mock.On("SelectTugasById", mockPayload.TugasId).Return(debtCollectorEntity.Tugas{}, nil)
	s.dcRepoMock.Mock.On("InsertLogTugas", mockPayload).Return(nil)

	err := s.usecase.CreateLogTugas(mockPayload)

	s.NoError(err)
}

func (s *DebtCollectorUseCaseSuite) TestCreateLogTugas_Fail() {
	s.Run("Fail tugas not found", func() {
		mockPayloadNotFound := debtCollectorDto.NewLogTugasPayload{
			TugasId:     "1",
			Description: "tes isi log",
		}

		expectedError := errors.New("error not found")
		s.dcRepoMock.Mock.On("SelectTugasById", mockPayloadNotFound.TugasId).Return(debtCollectorEntity.Tugas{}, expectedError)

		err := s.usecase.CreateLogTugas(mockPayloadNotFound)
		s.Error(err)
		s.Equal(expectedError, err)
	})

	s.Run("Fail internal server error", func() {
		mockPayloadInternalServerError := debtCollectorDto.NewLogTugasPayload{
			TugasId:     "2",
			Description: "tes isi log",
		}

		expectedErrorServer := errors.New("error internal server")
		s.dcRepoMock.Mock.On("SelectTugasById", mockPayloadInternalServerError.TugasId).Return(debtCollectorEntity.Tugas{}, nil)
		s.dcRepoMock.Mock.On("InsertLogTugas", mockPayloadInternalServerError).Return(expectedErrorServer)

		err := s.usecase.CreateLogTugas(mockPayloadInternalServerError)
		s.Error(err)
		s.Equal(expectedErrorServer, err)
	})
}
