package adminRepository

import (
	"database/sql"
	"fp_pinjaman_online/model/entity"
	adminInterface "fp_pinjaman_online/src/admin"
	"log"
	"time"
)

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) adminInterface.AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) RetrieveUserStatusById(userID int) (*entity.UserCompleteInfo, error) {
	query := `
    SELECT 
        u.id AS user_id,
        u.email,
        u.status,
        r.account_number,
        r.bank_name,
        k.name AS emergency_contact_name,
        k.phone_number AS emergency_contact_phone,
        j.job_name,
        j.gaji AS salary,
        j.office_name,
        d.NIK,
        d.fullname,
        d.phone_number AS personal_phone_number,
        d.address AS personal_address,
        d.city,
        d.foto_ktp,
        d.foto_selfie,
        u.created_at,
        u.updated_at,
        u.verified_at,
        u.deleted_at
    FROM 
        users u
    LEFT JOIN rekening r ON u.id = r.user_id
    LEFT JOIN kontak_darurat k ON u.id = k.user_id
    LEFT JOIN users_job_detail j ON u.id = j.user_id
    LEFT JOIN detail_users d ON u.id = d.user_id
    WHERE u.id = $1;
    `

	uci := entity.UserCompleteInfo{}
	err := r.db.QueryRow(query, userID).Scan(
		&uci.UserID,
		&uci.Email,
		&uci.Status,
		&uci.AccountNumber,
		&uci.BankName,
		&uci.EmergencyContact,
		&uci.EmergencyPhone,
		&uci.JobName,
		&uci.Gaji,
		&uci.OfficeName,
		&uci.NIK,
		&uci.FullName,
		&uci.PersonalPhoneNumber,
		&uci.PersonalAddress,
		&uci.City,
		&uci.FotoKTP,
		&uci.FotoSelfie,
		&uci.CreatedAt,
		&uci.UpdatedAt,
		&uci.VerifiedAt,
		&uci.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with ID: %d", userID)
			return nil, nil // or return a custom error for not found
		}
		log.Printf("Error fetching user details: %v", err)
		return nil, err
	}

	return &uci, nil
}

func (r *adminRepository) UpdateUserStatus(id int, status string) error {
	query := `
    UPDATE users 
    SET status = $2, updated_at = $3
    WHERE id = $1;
    `
	_, err := r.db.Exec(query, id, status, time.Now())
	if err != nil {
		log.Printf("Error updating user status: %v", err)
		return err
	}

	return nil
}
