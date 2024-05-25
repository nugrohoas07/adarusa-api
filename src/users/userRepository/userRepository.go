package userRepository

import (
	"database/sql"
	"fmt"
	"fp_pinjaman_online/model/dcFormDto"
	"fp_pinjaman_online/model/debiturFormDto"
	"fp_pinjaman_online/model/userDto"
	"fp_pinjaman_online/src/users"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) users.UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(req userDto.CreateRequest, roleId int) error {
    query := "INSERT INTO users(email, password, role_id) VALUES ($1,$2, $3)"
    _, err := repo.db.Exec(query, req.Email, req.Password, roleId)
	return err
}

func (repo *userRepository) Login(req userDto.LoginRequest) (string, error) {
	var hashedPassword string
	
	query := "SELECT password FROM users WHERE email=$1"
	err := repo.db.QueryRow(query, req.Email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", nil
	}

	return hashedPassword, nil
}

func (repo *userRepository) UserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)"
	err := repo.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (repo *userRepository) GetUserByEmail(email string) (userDto.User, error) {
	var user userDto.User
	query := `
        SELECT u.id, u.email, u.password, r.roles_name 
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE u.email=$1`
	err := repo.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password, &user.Roles)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with email %s not found", email)
		}
		return user, err
	}
	return user, nil
}

func (repo *userRepository) CreateDetailDebitur(req debiturFormDto.Debitur) error {
    tx, err := repo.db.Begin()
    if err != nil {
        return err
    }

    var exists bool
    err = tx.QueryRow(`SELECT EXISTS(SELECT 1 FROM detail_users WHERE user_id=$1)`, req.DetailUser.UserID).Scan(&exists)
    if err != nil {
        tx.Rollback()
        return err
    }
    if exists {
        _, err = tx.Exec(`
        UPDATE detail_users SET limit_id=$1, nik=$2, fullname=$3, phone_number=$4, address=$5, city=$6, foto_ktp=$7, foto_selfie=$8 WHERE user_id=$9`, req.DetailUser.LimitID, req.DetailUser.Nik, req.DetailUser.Fullname, req.DetailUser.PhoneNumber, req.DetailUser.Address, req.DetailUser.City, req.DetailUser.FotoKtp, req.DetailUser.FotoSelfie, req.DetailUser.UserID)
        if err != nil {
            tx.Rollback()
            return err
        }
    } else {
        _, err = tx.Exec(`
        INSERT INTO detail_users (user_id, limit_id, nik, fullname, phone_number, address, city, foto_ktp, foto_selfie)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        `, req.DetailUser.UserID, req.DetailUser.LimitID, req.DetailUser.Nik, req.DetailUser.Fullname, req.DetailUser.PhoneNumber, req.DetailUser.Address, req.DetailUser.City, req.DetailUser.FotoKtp, req.DetailUser.FotoSelfie)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    err = upsertUserJobDetail(tx, req.UserJobs)
    if err != nil {
        tx.Rollback()
        return err
    }

    err = upsertEmergencyContact(tx, req.EmergencyContact)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

func (repo *userRepository) CreateDetailDc(req dcFormDto.DetailDC) error {
    tx, err := repo.db.Begin()
    if err != nil {
        return err
    }

    var exists bool
    err = tx.QueryRow(`SELECT EXISTS(SELECT 1 FROM detail_users WHERE user_id=$1)`, req.UserID).Scan(&exists)
    if err != nil {
        tx.Rollback()
        return err
    }
    if exists {
        _, err = tx.Exec(`
        UPDATE detail_users SET limit_id=$1, nik=$2, fullname=$3, phone_number=$4, address=$5, city=$6, foto_ktp=$7, foto_selfie=$8 WHERE user_id=$9`, req.LimitID, req.Nik, req.Fullname, req.PhoneNumber, req.Address, req.City, req.FotoKtp, req.FotoSelfie, req.UserID)
        if err != nil {
            tx.Rollback()
            return err
        }
    } else {
        _, err = tx.Exec(`
        INSERT INTO detail_users (user_id, limit_id, nik, fullname, phone_number, address, city, foto_ktp, foto_selfie)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        `, req.UserID, req.LimitID, req.Nik, req.Fullname, req.PhoneNumber, req.Address, req.City, req.FotoKtp, req.FotoSelfie)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}

func upsertUserJobDetail(tx *sql.Tx, jobDetail debiturFormDto.UserJobs) error {
    var exists bool
    err := tx.QueryRow(`SELECT EXISTS(SELECT 1 FROM users_job_detail WHERE user_id=$1)`, jobDetail.UserID).Scan(&exists)
    if err != nil {
        return err
    }

    if exists {
        _, err = tx.Exec(`
        UPDATE users_job_detail SET job_name=$1, gaji=$2, office_name=$3, office_contact=$4, address=$5 WHERE user_id=$6
        `, jobDetail.JobName, jobDetail.Salary, jobDetail.OfficeName, jobDetail.OfficeContact, jobDetail.OfficeAddress, jobDetail.UserID)
    } else {
        _, err = tx.Exec(`
        INSERT INTO users_job_detail (user_id, job_name, gaji, office_name, office_contact, address)
        VALUES ($1, $2, $3, $4, $5, $6)
        `, jobDetail.UserID, jobDetail.JobName, jobDetail.Salary, jobDetail.OfficeName, jobDetail.OfficeContact, jobDetail.OfficeAddress)
    }
    return err
}

func upsertEmergencyContact(tx *sql.Tx, contact debiturFormDto.EmergencyContact) error {
    var exists bool
    err := tx.QueryRow(`SELECT EXISTS(SELECT 1 FROM kontak_darurat WHERE user_id=$1)`, contact.UserID).Scan(&exists)
    if err != nil {
        return err
    }

    if exists {
        _, err = tx.Exec(`
        UPDATE kontak_darurat SET name=$1, phone_number=$2 WHERE user_id=$3
        `, contact.Name, contact.PhoneNumber, contact.UserID)
    } else {
        _, err = tx.Exec(`
        INSERT INTO kontak_darurat (user_id, name, phone_number)
        VALUES ($1, $2, $3)
        `, contact.UserID, contact.Name, contact.PhoneNumber)
    }
    return err
}

func (repo *userRepository) UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error {
    _, err := repo.db.Exec(`
    UPDATE detail_users SET foto_ktp=$1, foto_selfie=$2 WHERE user_id=$3
    `, fotoKTP, fotoSelfie, userId)
    return err
}

func (repo *userRepository) GetFullname(userId int) (string, error) {
	var fullname string
	err := repo.db.QueryRow("SELECT fullname FROM detail_users WHERE user_id=$1", userId).Scan(&fullname)
	fmt.Println("fullname:", fullname)
	if err != nil {
		return "", nil
	}

	return fullname, nil
}

func (repo *userRepository) GetDataByRole(role, status string, limit, offset int) ([]debiturFormDto.DetailDebitur, int, error) {
    var debitur []debiturFormDto.DetailDebitur
    var totalData int

    // Base queries with conditional status filter
    query := `
        SELECT u.id, du.nik, du.fullname, du.phone_number, du.address, du.city, du.foto_ktp, du.foto_selfie, du.limit_id
        FROM users u
        JOIN detail_users du ON u.id = du.user_id
        JOIN roles r ON u.role_id = r.id
        JOIN limit_pinjaman lp ON du.limit_id = lp.id
        WHERE r.roles_name = $1
        AND ($2 = '' OR u.status = $2::user_status)
        LIMIT $3 OFFSET $4`
    
    countQuery := `
        SELECT count(*)
        FROM users u
        JOIN detail_users du ON u.id = du.user_id
        JOIN roles r ON u.role_id = r.id
        JOIN limit_pinjaman lp ON du.limit_id = lp.id
        WHERE r.roles_name = $1
        AND ($2 = '' OR u.status = $2::user_status)`

    // Arguments for the queries
    args := []interface{}{role, status, limit, offset}

    // Execute the main query
    rows, err := repo.db.Query(query, args...)
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()

    // Process the result set
    for rows.Next() {
        var dbt debiturFormDto.DetailDebitur
        err := rows.Scan(&dbt.UserID, &dbt.Nik, &dbt.Fullname, &dbt.PhoneNumber, &dbt.Address, &dbt.City, &dbt.FotoKtp, &dbt.FotoSelfie, &dbt.LimitID)
        if err != nil {
            return nil, 0, err
        }
        debitur = append(debitur, dbt)
    }

    // Execute the count query
    countArgs := args[:2] // Use only role and status for the count query
    err = repo.db.QueryRow(countQuery, countArgs...).Scan(&totalData)
    if err != nil {
        return nil, 0, err
    }

    return debitur, totalData, nil
}

