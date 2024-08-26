package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/golang-jwt/jwt"
)

type UserInfo struct {
	Username string
	Id       int
	Password string
}

func generateUniqueID(existingIDs []int) int {
	for {
		id := rand.Intn(9000000) + 1000000
		if !contains(existingIDs, id) {
			return id
		}
	}
}

func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func Auth() {
	var username string
	var Id int
	var Id_collection []int
	var password string

	fmt.Println("Enter your username (something unique): ")
	fmt.Scanln(&username)

	// Generate unique ID
	Id = generateUniqueID(Id_collection)
	Id_collection = append(Id_collection, Id)

	fmt.Println("Your ID is:", Id)
	fmt.Println("Enter password: ")
	fmt.Scanln(&password)

	// Encrypting the password using the md5 hash function
	bytePassword := []byte(password)
	hashedPassword := md5.Sum(bytePassword)
	stringedPassword := hex.EncodeToString(hashedPassword[:])

	user := UserInfo{
		Username: username,
		Id:       Id,
		Password: stringedPassword,
	}

	var secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"id":       user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	fmt.Printf("Generated Token: %s\n", tokenString)
}
