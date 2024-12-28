package services

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yusufniyi/cli-todo-app/internal/database/models"
	"github.com/yusufniyi/cli-todo-app/internal/envs"
	"github.com/yusufniyi/cli-todo-app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func (userService UserService) Login(email string, password string) {
	getUserDataFromFile(envs.ENV_VARIABLES.AuthFilepath)
}

func (userService UserService) Signup(user *models.User) models.User {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	sql := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	stmt, _ := userService.DB.Prepare(sql)
	stmt.Exec(nil, user.Name, user.Email, hashedPassword)
	defer stmt.Close()

	userDetails := fmt.Sprintf("email %s\nname %s\npassword %s\n\n", user.Email, user.Name, string(hashedPassword))

	utils.WriteToFile(envs.ENV_VARIABLES.AuthFilepath, userDetails)

	return *user
}

func (userService UserService) CheckIfUserAlreadyExistsByEmail(email string) bool {
	user := &models.User{}
	sqlStmt := `SELECT * FROM users WHERE email = ?`
	err := userService.DB.QueryRow(sqlStmt, email).Scan(user)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return false
	}
	return true
}

func getUserDataFromFile(filename string) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	// Variables to store email and password
	var email, password, name string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "email") {
			email = strings.TrimSpace(strings.TrimPrefix(line, "email:"))
		} else if strings.HasPrefix(line, "password") {
			password = strings.TrimSpace(strings.TrimPrefix(line, "password:"))
		} else if strings.HasPrefix(line, "name") {
			name = strings.TrimSpace(strings.TrimPrefix(line, "name"))
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the extracted email and password
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Password: %s\n", password)
	fmt.Println("name", name)
}
