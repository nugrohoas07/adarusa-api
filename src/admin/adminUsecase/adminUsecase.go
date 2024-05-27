package adminUsecase

import (
	"fmt"
	"fp_pinjaman_online/model/dto/adminDto"
	adminEntity "fp_pinjaman_online/model/entity/admin"
	"fp_pinjaman_online/pkg/validation"
	adminInterface "fp_pinjaman_online/src/admin"
	"log"
	"math"
	"time"
)

type adminUsecase struct {
	repo adminInterface.AdminRepository
}

func NewAdminUsecase(repo adminInterface.AdminRepository) adminInterface.AdminUsecase {
	return &adminUsecase{repo: repo}
}

func (uc *adminUsecase) VerifyAndUpdateUser(req adminDto.RequestUpdateStatusUser) (adminDto.AdminResponse, error) {

	user, err := uc.repo.RetrieveUserStatusById(req.ID)
	if err != nil {
		log.Printf("Error retrieving user with ID %d: %v", req.ID, err)
		return adminDto.AdminResponse{}, err
	}

	if user == nil {
		log.Printf("No user found with ID %d", req.ID)
		return adminDto.AdminResponse{}, fmt.Errorf("no user found with ID %d", req.ID)
	}

	if req.Status == "verified" && !validation.ValidateUserComplete(*user) {
		log.Printf("Verification failed for user ID %d: incomplete user information", req.ID)
		return adminDto.AdminResponse{}, fmt.Errorf("verification failed: missing bank account information for user ID %d", req.ID)
	}

	if user.Status != req.Status {
		err := uc.repo.UpdateUserStatus(req.ID, req.Status)
		if err != nil {
			log.Printf("Failed to update status for user ID %d: %v", req.ID, err)
			return adminDto.AdminResponse{}, err
		}

		if req.Status == "verified" {
			now := time.Now()
			user.VerifiedAt = &now
		}
	}
	return adminDto.AdminResponse{
		ID:     user.UserID,
		Email:  user.Email,
		Status: user.Status,
	}, nil
}

func (uc *adminUsecase) VerifyAndCreateCicilan(req adminDto.RequestVerifyLoan) (adminDto.LoanResponse, error) {
	loan, err := uc.repo.RetrievePinjamanById(req.LoanID)
	if err != nil {
		log.Printf("Error retrieving loan: %v", err)
		return adminDto.LoanResponse{}, err
	}

	if req.StatusPengajuan == "rejected" {
		return adminDto.LoanResponse{}, fmt.Errorf("loan application is rejected")
	}

	if req.StatusPengajuan == "completed" {
		return adminDto.LoanResponse{}, fmt.Errorf("loan application is complete")
	}

	user, err := uc.repo.RetrieveUserLimitByAdmin(req.UserID)
	if err != nil {
		log.Printf("Error retrieving user : %v", err)
		return adminDto.LoanResponse{}, err
	}

	if user.MaxPinjaman.Float64 >= loan.JumlahPinjaman {
		err = uc.repo.UpdateLoanStatus(loan.ID, req.StatusPengajuan)
		if err != nil {
			return adminDto.LoanResponse{}, err
		}

		if req.StatusPengajuan == "approved" {
			return uc.CreatePaymentSchedule(loan)
		}
		return adminDto.LoanResponse{}, nil
	}

	return adminDto.LoanResponse{}, fmt.Errorf("loan amount exceeds the user's loan limit")
}

func (uc *adminUsecase) CreatePaymentSchedule(loan *adminEntity.Pinjaman) (adminDto.LoanResponse, error) {
	monthlyPayment := CalculateMonthlyPayment(loan.JumlahPinjaman, loan.BungaPerBulan*12, loan.Tenor)
	dueDate := time.Now()

	for i := 1; i <= loan.Tenor; i++ {
		dueDate = dueDate.AddDate(0, 1, 0)
		if err := uc.repo.InsertCicilan(loan.ID, dueDate, monthlyPayment, "unpaid"); err != nil {
			return adminDto.LoanResponse{}, err
		}
	}

	return adminDto.LoanResponse{}, nil
}

func CalculateMonthlyPayment(principal float64, annualInterestRate float64, tenor int) float64 {
	monthlyInterestRate := annualInterestRate / 12
	numberOfPayments := float64(tenor)

	// Menghitung pembayaran bulanan menggunakan rumus anuitas
	monthlyPayment := principal * (monthlyInterestRate * math.Pow(1+monthlyInterestRate, numberOfPayments)) / (math.Pow(1+monthlyInterestRate, numberOfPayments) - 1)

	return monthlyPayment
}
