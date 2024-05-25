package userUseCase

import (
	"fp_pinjaman_online/model/debiturFormDto"
	"fp_pinjaman_online/model/userDto"
	"fp_pinjaman_online/src/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Hasher interface {
	HashedPassword(password string) (string, error)
}

func TestLogin_Success(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(mocks.UserRepository)

	// Initialize the use case with the mock repository
	useCase := &userUC{
		userRepo: mockRepo,
	}

	// Define the input and expected output
	req := userDto.LoginRequest{Email: "test@example.com", Password: "password"}
	expectedUser := userDto.User{Id: "1", Email: "test@example.com", Roles: "user"}

	// Setup the expectations
	mockRepo.On("Login", req).Return("hashedPassword", nil)
	mockRepo.On("GetUserByEmail", req.Email).Return(expectedUser, nil)

	// Call the function
	token, err := useCase.Login(req)

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Check that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	useCase := &userUC{
		userRepo: mockRepo,
	}

	// define the input and epected output
	req := userDto.CreateRequest{Name:"",  Email: "test@example.com", Password: "password", Roles: 1}
	roleId := 1

	// setup the expectations
	mockRepo.On("UserExists", req.Email).Return(false, nil)
	mockRepo.On("CreateUser", mock.AnythingOfType("userDto.CreateRequest"), roleId).Return(nil)

	// call the function
	err := useCase.CreateUser(req, roleId)
	
	// assertions
	assert.NoError(t, err)

	// check thta the expectation were met
	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmail_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	useCase := &userUC{
		userRepo: mockRepo,
	}

	// define the input and expected output
	email := "test@example.com"
	expectedUser := userDto.User{Id: "1", Email: "test@example.com", Roles: "user"}

	// setup the expectations
	mockRepo.On("GetUserByEmail", email).Return(expectedUser, nil)

	// call the function
	user, err := useCase.GetUserByEmail(email)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// check that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestCreateDetailDebitur_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	useCase := &userUC{
		userRepo: mockRepo,
	}

	// define the input
	debitur := debiturFormDto.Debitur{
		DetailUser: debiturFormDto.DetailDebitur{
			UserID:      1,
			LimitID:     1,
			Nik:         "1234567890",
			Fullname:    "Test User",
			PhoneNumber: "1234567890",
			Address:     "Test Address",
			City:        "Test City",
			FotoKtp:     "test_ktp.jpg",
			FotoSelfie:  "test_selfie.jpg",
		},
		UserJobs: debiturFormDto.UserJobs{
			UserID:        1,
			JobName:       "Test Job",
			Salary:        5000.0,
			OfficeName:    "Test Office",
			OfficeContact: "1234567890",
			OfficeAddress: "Test Office Address",
		},
		EmergencyContact: debiturFormDto.EmergencyContact{
			UserID:      1,
			Name:        "Test Emergency Contact",
			PhoneNumber: "1234567890",
		},
	}

	// setup the expectations
	mockRepo.On("CreateDetailDebitur", debitur).Return(nil)

	// call the function
	err := useCase.CreateDetailDebitur(debitur)

	// assertions
	assert.NoError(t, err)

	// check that the expectations were met
	mockRepo.AssertExpectations(t)
}
