package adminUsecase

import (
	"fp_pinjaman_online/model/dto/adminDto"
	adminInterface "fp_pinjaman_online/src/admin"
	"log"
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
		return adminDto.AdminResponse{}, err
	}

	if user.Email == "" || user.AccountNumber == "" || user.BankName == "" ||
		user.EmergencyContact == "" || user.EmergencyPhone == "" || user.JobName == "" ||
		user.NIK == "" || user.FullName == "" || user.PersonalPhoneNumber == "" ||
		user.PersonalAddress == "" || user.City == "" || user.FotoKTP == "" ||
		user.FotoSelfie == "" {
		log.Printf("Incomplete data for user ID %d. Setting status to 'rejected'", req.ID)
		if err := uc.repo.UpdateUserStatus(req.ID, "rejected"); err != nil {
			log.Printf("Failed to update status for user ID %d: %v", req.ID, err)
			return adminDto.AdminResponse{}, err
		}
		return adminDto.AdminResponse{}, nil
	}

	if user.Status != req.Status {
		if err := uc.repo.UpdateUserStatus(req.ID, req.Status); err != nil {
			log.Printf("Failed to update status for user ID %d: %v", req.ID, err)
			return adminDto.AdminResponse{}, err
		}

		if req.Status == "verified" {
			now := time.Now()
			user.VerifiedAt = &now
		}
	}

	response := adminDto.AdminResponse{
		ID:         user.UserID,
		Email:      user.Email,
		Status:     req.Status,
		VerifiedAt: "",
	}
	if user.VerifiedAt != nil {
		response.VerifiedAt = user.VerifiedAt.Format(time.RFC3339)
	}

	return response, nil
}
