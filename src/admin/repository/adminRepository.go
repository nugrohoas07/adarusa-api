package adminRepository

import (
	"database/sql"
	adminEntity "fp_pinjaman_online/model/entity/admin"
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

func (r *adminRepository) RetrieveUserStatusById(userID int) (*adminEntity.UserCompleteInfo, error) {
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
        j.gaji AS gaji,
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

	uci := adminEntity.UserCompleteInfo{}
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
			return nil, nil
		}
		log.Printf("Error fetching user details: %v", err)
		return nil, err
	}

	return &uci, nil
}

func (r *adminRepository) UpdateUserStatus(userID int, status string) error {
	query := `
	UPDATE users
		SET status = $2,
    		verified_at = NOW(),
    		updated_at = NOW()
		WHERE id = $1;
    `
	_, err := r.db.Exec(query, userID, status)
	if err != nil {
		log.Printf("Error updating user status: %v", err)
		return err
	}

	return nil
}

func (r *adminRepository) RetrieveUserLimitByAdmin(userID int) (*adminEntity.UserCompleteInfoLoanLimit, error) {
	query := `
    SELECT 
        u.id AS user_id,
        u.email,
        u.status,
        uj.job_name,
        uj.gaji AS gaji,
        uj.office_name,
        du.NIK,
        du.fullname,
        du.phone_number AS personal_phone_number,
        du.address AS personal_address,
        du.city,
        du.foto_ktp,
        du.foto_selfie,
        lp.max_pinjaman
    FROM 
        users u
    LEFT JOIN users_job_detail uj ON u.id = uj.user_id
    LEFT JOIN detail_users du ON u.id = du.user_id
    LEFT JOIN limit_pinjaman lp ON du.limit_id = lp.id
    WHERE u.id = $1;
    `

	uci := adminEntity.UserCompleteInfoLoanLimit{}
	err := r.db.QueryRow(query, userID).Scan(
		&uci.UserID,
		&uci.Email,
		&uci.Status,
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
		&uci.MaxPinjaman,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with ID: %d", userID)
			return nil, nil
		}
		log.Printf("Error fetching user details: %v", err)
		return nil, err
	}

	return &uci, nil
}

func (r *adminRepository) RetrievePinjamanById(loanID int) (*adminEntity.Pinjaman, error) {
	query := `
        SELECT id, user_id, admin_id, jumlah_pinjaman, tenor, bunga_per_bulan, description, 
               status_pengajuan, created_at, updated_at, deleted_at
        FROM pinjaman
        WHERE id = $1;
    `

	var pinjaman adminEntity.Pinjaman

	row := r.db.QueryRow(query, loanID)
	err := row.Scan(
		&pinjaman.ID, &pinjaman.UserID, &pinjaman.AdminID, &pinjaman.JumlahPinjaman,
		&pinjaman.Tenor, &pinjaman.BungaPerBulan, &pinjaman.Description,
		&pinjaman.StatusPengajuan, &pinjaman.CreatedAt, &pinjaman.UpdatedAt, &pinjaman.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No Pinjaman found with ID %d", loanID)
			return nil, err
		}
		log.Printf("Error querying for Pinjaman with ID %d: %v", loanID, err)
		return nil, err
	}

	return &pinjaman, nil
}

func (r *adminRepository) UpdateLoanStatus(loanID int, status string) error {
	query := `
	UPDATE pinjaman
		SET status_pengajuan = $2,
    		updated_at = NOW()
		WHERE id = $1;
    `
	_, err := r.db.Exec(query, loanID, status)
	if err != nil {
		log.Printf("Error updating pinjaman status: %v", err)
		return err
	}

	return nil
}

func (r *adminRepository) InsertCicilan(loanID int, dueDate time.Time, amount float64, status string) error {
	query := `
        INSERT INTO cicilan (pinjaman_id, tanggal_jatuh_tempo, jumlah_bayar, status)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(query, loanID, dueDate, amount, status)
	if err != nil {
		return err
	}
	return nil
}
