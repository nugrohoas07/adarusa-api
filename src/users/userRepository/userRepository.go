package userRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"fp_pinjaman_online/model/userDto"
	"fp_pinjaman_online/src/users"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) users.UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(req userDto.CreateRequest) error {
    defaultRoleId := 2  // Default role_id value you want to assign
    query := "INSERT INTO users(email, password, role_id) VALUES ($1,$2,$3)"
    _, err := repo.db.Exec(query, req.Email, req.Password, defaultRoleId)
	if err != nil {
		return err
	}
	return nil
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

func (repo *userRepository) GetUserById(id string) (userDto.User, error) {
	var user userDto.User
	query := "SELECT id, email, password FROM users WHERE id=$1 AND deleted_at IS NULL"
	err := repo.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user id %s not found", id)
		}
		return user, err
	}
	return user, err
}

func (repo *userRepository) GetUserCount(email, password string) (int, error) {
	var total int
	query := `SELECT COUNT(*) FROM users WHERE email LIKE $1 AND name LIKE $2`
	err := repo.db.QueryRow(query, "%"+email+"%", "%"+password+"%").Scan(&total)
	if err != nil {
		return 0, err
	}
	
	return total, nil
}

func (repo *userRepository) UpdateUser(id string, req userDto.Update) error {
	query := "UPDATE users SET"
	tmpQuery := []interface{}{}
	i := 1

	if req.Email != "" {
		query += fmt.Sprintf(" email = $%d,", i)
		tmpQuery = append(tmpQuery, req.Password)
		i++
	}

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id = $%d AND deleted_at IS NULL", i)
	tmpQuery = append(tmpQuery, id)

	result, err := repo.db.Exec(query, tmpQuery...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user does not exist")
	}
	return nil
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

func (repo *userRepository) UserExistById(id string) (bool, error) {
	var existId bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1 AND deleted_at IS NULL)"
	err := repo.db.QueryRow(query, id).Scan(&existId)
	if err != nil {
		return false, err
	}
	return existId, nil
}

func (repo *userRepository) GetUserByEmail(email string) (userDto.User, error) {
	var user userDto.User
	query := "SELECT id, email, password FROM users WHERE email=$1"
	err := repo.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with email %s not found", email)
		}
		return user, err
	}
	return user, nil
}