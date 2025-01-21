package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yusufniyi/cli-todo-app/internal/config"
	file_service "github.com/yusufniyi/cli-todo-app/internal/service/file"
	"github.com/yusufniyi/cli-todo-app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Unable to hash password")
	}
	return string(hashedPassword)
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func generateJwtToken(email string, userId int) (string, error) {
	// Create token claims
	claims := jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Expiry time (24 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.TokenEncryptionKey)
	if err != nil {
		return "", fmt.Errorf("Fatal: failed to sign token: %w", err)
	}

	return tokenString, nil
}

func decodeJwtoken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Fatal: unexpected signing method: %v", token.Header["alg"])
		}
		return config.TokenEncryptionKey, nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid")
		fmt.Println("Claims:", claims)
		return claims, nil
	} else {
		return jwt.MapClaims{}, errors.New("Fatal: Invalid jwt claim")
	}
}

func isUserAuthenticated(email string) bool {
	// Read encrypted token
	// Decrypt the token
	// Decode the token
	// Compare emails
	encryptedToken, err := file_service.ReadFromFile(config.AuthFileName)
	if err != nil {
		log.Println("Warning: error reading authentication data:", err)
		return false
	}

	// Decrypt the data
	decryptedToken, err := utils.Decrypt(config.TokenEncryptionKey, encryptedToken)
	if err != nil {
		fmt.Println("Warning: failed to decrypt data:", err)
		return false
	}

	claims, err := decodeJwtoken(decryptedToken)
	if err != nil {
		fmt.Println("claims", claims)
	}

	return claims["email"] == email
}
