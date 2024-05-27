package adminInterface

import (
	"fp_pinjaman_online/model/dto/adminDto"
	adminEntity "fp_pinjaman_online/model/entity/admin"
	"time"
)

type AdminRepository interface {
	RetrieveUserStatusById(id int) (*adminEntity.UserCompleteInfo, error)
	UpdateUserStatus(id int, status string) error
	RetrieveUserLimitByAdmin(userID int) (*adminEntity.UserCompleteInfoLoanLimit, error)
	RetrievePinjamanById(loanID int) (*adminEntity.Pinjaman, error)
	UpdateLoanStatus(loanID int, status string) error
	InsertCicilan(loanID int, dueDate time.Time, amount float64, status string) error
}

type AdminUsecase interface {
	VerifyAndUpdateUser(req adminDto.RequestUpdateStatusUser) (adminDto.AdminResponse, error)
	VerifyAndCreateCicilan(req adminDto.RequestVerifyLoan) (adminDto.LoanResponse, error)
}
