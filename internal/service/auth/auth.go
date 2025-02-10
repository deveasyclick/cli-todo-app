package authservice

import (
	"fmt"
	"log"
	"os"

	"github.com/yusufniyi/cli-todo-app/internal/config"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	"github.com/yusufniyi/cli-todo-app/internal/helpers/aesutils"
	"github.com/yusufniyi/cli-todo-app/internal/helpers/file"
)

type Service struct {
	userRepository repositories.User
}

func (service *Service) Login(email string, password string) {
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

	user, err := service.userRepository.FindUser(email)
	if err != nil {
		log.Fatalf("Fatal: Unable to fetch user with email %s from database: %s", email, err)
	}

	if user.Email == "" {
		log.Fatalf("Fatal: User with email %s does not exist", email)
	}

	hasCorrectPassword := comparePassword(user.Password, password)

	if !hasCorrectPassword {
		log.Fatalln("Fatal: Incorrect login details. Please try again")
	}

	token, err := generateJwtToken(email, user.ID)
	if err != nil {
		log.Fatalln("Fatal: Error generating jwt token for user")
	}
	aes := aesutils.NewAESUtil()
	aes.SetKey([]byte(config.TokenEncryptionKey))
	encryptedToken, err := aes.Encrypt(token)
	if err != nil {
		fmt.Println("Fatal: Failed to encrypt user data")
		os.Exit(1)
	}
	if err = file.Save(config.AuthFileName, encryptedToken); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("User login successfully")
}

func (service *Service) Signup(user *models.User) models.User {
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

	dbUser, err = service.userRepository.FindUser(user.Email)
	if err != nil {
		log.Fatalln(err)
	}

	if dbUser.Email == user.Email {
		log.Fatalln("User already exists, please sign in.")
	}

	userId, err = service.userRepository.AddUser(user)
	if err != nil {
		log.Fatalln("Error adding user to the database", err)
	}

	token, err = generateJwtToken(user.Email, userId)
	if err != nil {
		log.Fatalln(err)
	}

	if err := file.Save(config.AuthFileName, token); err != nil {
		fmt.Println("Unable to write login data to file")
	}

	fmt.Printf("User with email %s registered successfully\n", user.Email)
	return dbUser
}

func (service *Service) Logout() {
	if err := file.Remove(config.AuthFileName); err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	fmt.Println("User logout successfully")
}

func (service *Service) Authenticate() Token {
	isAuthenicated, token := authenticate()
	if !isAuthenicated {
		log.Fatalln("Unauthorized command, please login.")
	}
	return token
}

func New(userRepository repositories.User) *Service {
	return &Service{userRepository: userRepository}
}
