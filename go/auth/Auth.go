package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Struct for user information
type UserInfo struct {
	Id       int    `gorm:"primarykey" json:"Id"`
	Username string `json:"username" gorm:"size:50"`
	Password string `json:"password"`
}

type UserInfoInput struct {
	Id       string `json:"Id"`
	Password string `json:"password"`
}

// Specify table name
func (UserInfo) TableName() string {
	return "userinfo"
}

// Generate a unique ID
func generateUniqueID(existingIDs []int) int {
	for {
		id := rand.Intn(9000000) + 1000000
		if !contains(existingIDs, id) {
			return id
		}
	}
}

// Check if slice contains a value
func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Authenticate user
func (user UserInfo) Auth() (bool, UserInfo, string) {
	// Connect to the database
	dsn := "host=localhost user=postgres password=123456 dbname=chatappusers port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}

	// Check if username already exists in the database
	var existingUser UserInfo
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		// Username exists
		return true, existingUser, "Username already exists"
	}

	// Username does not exist, generate a unique ID
	user.Id = generateUniqueID([]int{})

	// Encrypt the password using MD5 hash
	bytePassword := []byte(user.Password)
	hashedPassword := md5.Sum(bytePassword)
	user.Password = hex.EncodeToString(hashedPassword[:])

	// Insert the new user into the database
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	return false, user, "" // Username does not exist, so return false
}

func JWT(userId int) (string, error) {
	secretKey := "61a041114690e9212e8ee7334c58e7b3ba49449bd813f23f67e80f37adbc24dd"
	if secretKey == "" {
		return "", fmt.Errorf("SECRET_KEY environment variable is not set")
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

// Signin checks the user's credentials and returns a JWT token if successful.
func (user UserInfoInput) Signin() (bool, string, error, string) {
	dsn := "host=localhost user=postgres password=123456 dbname=chatappusers port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return false, "Database connection error", err, ""
	}

	// Convert Id from string to int
	userId, err := strconv.Atoi(user.Id)
	if err != nil {
		return false, "Invalid ID format", err, ""
	}

	// Encrypt the password using MD5 hash
	bytePassword := []byte(user.Password)
	hashedPassword := md5.Sum(bytePassword)
	hashedPasswordStr := hex.EncodeToString(hashedPassword[:]) // Convert to hex string

	// Check if the ID and hashed password match a record in the database.
	var dbUser UserInfo
	if err := db.Where("id = ? AND password = ?", userId, hashedPasswordStr).First(&dbUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, "Error: Record not found", err, ""
		}
		return false, "Database query error", err, ""
	}

	// Generate a JWT token for the authenticated user.
	token, err := JWT(userId)
	if err != nil {
		return false, "Token generation error", err, ""
	}

	return true, "Record found: Success", err, token
}
