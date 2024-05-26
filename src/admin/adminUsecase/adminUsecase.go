package adminUsecase

import (
	"fmt"
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
		return adminDto.AdminResponse{}, fmt.Errorf("no user found with ID %d", req.ID)
	}

	// Check if bank account information is missing when attempting to verify the user
	if req.Status == "verified" && (!user.AccountNumber.Valid || user.AccountNumber.String == "") {
		log.Printf("Verification failed for user ID %d: missing bank account information", req.ID)
		return adminDto.AdminResponse{}, fmt.Errorf("verification failed: missing bank account information for user ID %d", req.ID)
	}

	// Proceed with updating the user status if all necessary information is present or status is "rejected"
	if user.Status != req.Status {
		err := uc.repo.UpdateUserStatus(req.ID, req.Status)
		if err != nil {
			log.Printf("Failed to update status for user ID %d: %v", req.ID, err)
			return adminDto.AdminResponse{}, err
		}

		// Update the VerifiedAt field if the status is updated to "verified"
		if req.Status == "verified" {
			now := time.Now()
			user.VerifiedAt = &now
		}
	}

	// Prepare the response DTO
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
