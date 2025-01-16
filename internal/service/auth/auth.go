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
	isAlreadyAuthenticated := isUserAuthenticated(email)

	if isAlreadyAuthenticated {
		return
	}

	user, err := authService.UserRepository.FindUser(email)
	if err != nil {
		log.Fatalf("Unable to fetch user with email %s from database", email)
	}

	hasCorrectPassword := comparePassword(user.Password, password)

	if !hasCorrectPassword {
		log.Fatalln("Incorrect login details. Please try again")
	}

	token, err := generateJwtToken(email, user.ID)
	if err != nil {
		log.Fatalln("Error generating jwt token for user")
	}
	encryptedToken, err := utils.Encrypt(config.TokenEncryptionKey, token)
	if err != nil {
		fmt.Println("Failed to encrypt user data")
		os.Exit(1)
	}
	if err = file_service.SaveToFile(config.AuthFileName, encryptedToken); err != nil {
		log.Fatalln("Unable to save login data", err)
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
	userId, err = authService.UserRepository.AddUser(user)
	if err != nil {
		log.Fatalln("Error adding user to the database")
	}
	dbUser, err = authService.UserRepository.FindUser(user.Email)
	if err != nil {
		log.Fatalln("Error fetching user from the database")
	}
	log.Printf("User with id %d added to database", userId)

	token, err = generateJwtToken(user.Email, userId)
	if err != nil {
		log.Fatalln("Error generating jwt token")
	}

	file_service.SaveToFile(config.AuthFileName, token)

	fmt.Println("User signup successfully")
	return dbUser
}

func (userService *AuthService) Logout(user *models.User) {
	// Remove token from file
}

func (userService *AuthService) Authenticate(email string) {
	if !isUserAuthenticated(email) {
		log.Fatalln("Authentication required")
	}
}
