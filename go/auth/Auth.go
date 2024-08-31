package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"

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

func (user UserInfo) Signin() string {
	// Connect to the database
	dsn := "host=localhost user=postgres password=123456 dbname=chatappusers port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}

	// Check if Id and password are in the database
	if err := db.Where("id = ?", user.Id).First(&user).Error; err == nil {
		// ID exists
	}

	// Encrypt the password using MD5 hash
	bytePassword := []byte(user.Password)
	hashedPassword := md5.Sum(bytePassword)

	i
}

func (user UserInfo) JWT() {
	// Generating the JWT token to authenticate the user to be used in authencation
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
