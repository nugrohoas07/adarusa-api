package userRepository

import (
	"database/sql"
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