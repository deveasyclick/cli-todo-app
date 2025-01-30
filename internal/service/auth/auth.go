package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/yusufniyi/cli-todo-app/internal/config"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	file_service "github.com/yusufniyi/cli-todo-app/internal/service/file"
	"github.com/yusufniyi/cli-todo-app/internal/utils"
)

type AuthService struct {
	UserRepository repositories.UserRepository
}

func (authService *AuthService) Login(email string, password string) {
	// Check if user is already authenticated and return if user is already authenticated

	// fetch user from database
	// compare user password
	// generate jwt token
	// Encrypt the token
	//save the token in a file
	// Print user login successfully

	// Read encrypted data from the file
	isAlreadyAuthenticated := isUserEmailAuthenticated(email)

	if isAlreadyAuthenticated {
		fmt.Println("User is already authenticated")
		return
	}

	user, err := authService.UserRepository.FindUser(email)
	if err != nil {
		log.Fatalf("Fatal: Unable to fetch user with email %s from database", email)
	}

	hasCorrectPassword := comparePassword(user.Password, password)

	if !hasCorrectPassword {
		log.Fatalln("Fatal: Incorrect login details. Please try again")
	}

	token, err := generateJwtToken(email, user.ID)
	if err != nil {
		log.Fatalln("Fatal: Error generating jwt token for user")
	}
	encryptedToken, err := utils.Encrypt(config.TokenEncryptionKey, token)
	if err != nil {
		fmt.Println("Fatal: Failed to encrypt user data")
		os.Exit(1)
	}
	if err = file_service.SaveToFile(config.AuthFileName, encryptedToken); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("User login successfully")
}

func (authService *AuthService) Signup(user *models.User) models.User {
	// hash the password
	// save in db
	// generate token
	// encrypt token
	// Save token in file
	hashedPassword := hashPassword(user.Password)
	user.Password = hashedPassword
	var err error
	var userId int
	var dbUser models.User
	var token string

	dbUser, err = authService.UserRepository.FindUser(user.Email)
	if err != nil {
		log.Fatalln(err)
	}

	if dbUser.Email == user.Email {
		log.Fatalln("User already exists, please sign in.")
	}

	userId, err = authService.UserRepository.AddUser(user)
	if err != nil {
		log.Fatalln("Error adding user to the database", err)
	}

	token, err = generateJwtToken(user.Email, userId)
	if err != nil {
		log.Fatalln(err)
	}

	if err := file_service.SaveToFile(config.AuthFileName, token); err != nil {
		fmt.Println("Unable to write login data to file")
	}

	fmt.Printf("User with id %d registered successfully\n", userId)
	return dbUser
}

func (authService *AuthService) Authenticate(email string) {
	if !isUserEmailAuthenticated(email) {
		log.Fatalln("Authentication required")
	}
}

func (authService *AuthService) Logout() {
	if err := file_service.DeleteFile(config.AuthFileName); err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	fmt.Println("User logout successfully")
}
