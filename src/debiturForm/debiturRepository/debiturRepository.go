package debiturRepository

import (
	"database/sql"
	"fmt"
	"fp_pinjaman_online/model/debiturFormDto"
	"fp_pinjaman_online/src/debiturForm"
)

type debiturRepository struct {
    db *sql.DB
}

func NewDebiturDetailRepository(db *sql.DB) debiturForm.DebiturRepository {
    return &debiturRepository{db}
}

func (repo *debiturRepository) CreateDetailDebitur(req debiturFormDto.Debitur) error {
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

func (repo *debiturRepository) UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error {
    _, err := repo.db.Exec(`
    UPDATE detail_users SET foto_ktp=$1, foto_selfie=$2 WHERE user_id=$3
    `, fotoKTP, fotoSelfie, userId)
    return err
}

func (repo *debiturRepository) GetDataByRole(role, status string, limit, offset int) ([]debiturFormDto.DetailDebitur, int, error) {
    var debitur []debiturFormDto.DetailDebitur
    var totalData int

    var query string
    var countQuery string
    var args []interface{}
    args = append(args, role)

    if status != "" {
        query = `SELECT u.id, du.nik, du.fullname, du.phone_number, du.address, du.city, du.foto_ktp, du.foto_selfie
                 FROM users u
                 JOIN detail_users du ON u.id = du.user_id
                 JOIN roles r ON u.role_id = r.id
                 WHERE r.roles_name = $1 AND u.status = $2
                 LIMIT $3 OFFSET $4`
        countQuery = `SELECT count(*)
                      FROM users u
                      JOIN detail_users du ON u.id = du.user_id
                      JOIN roles r ON u.role_id = r.id
                      WHERE r.roles_name = $1 AND u.status = $2`
        args = append(args, status, limit, offset)
    } else {
        query = `SELECT u.id, du.nik, du.fullname, du.phone_number, du.address, du.city, du.foto_ktp, du.foto_selfie
                 FROM users u
                 JOIN detail_users du ON u.id = du.user_id
                 JOIN roles r ON u.role_id = r.id
                 WHERE r.roles_name = $1
                 LIMIT $2 OFFSET $3`
        countQuery = `SELECT count(*)
                      FROM users u
                      JOIN detail_users du ON u.id = du.user_id
                      JOIN roles r ON u.role_id = r.id
                      WHERE r.roles_name = $1`
        args = append(args, limit, offset)
    }

    // Log the query and arguments for debugging
    fmt.Printf("Executing query: %s with args: %v\n", query, args)

    rows, err := repo.db.Query(query, args...)
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()

    for rows.Next() {
        var dbt debiturFormDto.DetailDebitur
        err := rows.Scan(&dbt.UserID, &dbt.Nik, &dbt.Fullname, &dbt.PhoneNumber, &dbt.Address, &dbt.City, &dbt.FotoKtp, &dbt.FotoSelfie)
        if err != nil {
            return nil, 0, err
        }
        debitur = append(debitur, dbt)
    }

    countQueryArgs := []interface{}{role}
    if status != "" {
        countQueryArgs = append(countQueryArgs, status)
    }

    // Log the count query and arguments for debugging
    fmt.Printf("Executing count query: %s with args: %v\n", countQuery, countQueryArgs)

    err = repo.db.QueryRow(countQuery, countQueryArgs...).Scan(&totalData)
    if err != nil {
        return nil, 0, err
    }

    return debitur, totalData, nil
}

