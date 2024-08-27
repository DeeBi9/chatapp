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

type UserInfo struct {
	Username string `gorm:"size:50"`
	Id       int    `gorm:"primarykey"`
	Password string `gorm:"size:50"`
}

func (UserInfo) TableName() string {
	return "userinfo"
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

func CorrectUsername(username string, Username_collection []string) bool {
	for _, name := range Username_collection {
		if username == name {
			return false
		}
	}
	return true
}
func Auth() {
	var username string
	var Id int
	var Id_collection []int
	var Username_collection []string
	var password string

	fmt.Println("Enter your username: ")
	for {
		fmt.Scanln(&username)
		answer := CorrectUsername(username, Username_collection)
		if !answer {
			fmt.Println("Username Exists !")
			fmt.Println("Enter again :")
			continue
		} else {
			break
		}
	}
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
	fmt.Println(stringedPassword)

	/* ORM implementation to send the data to postgres server */

	dsn := "host=localhost user=postgres password=123456 dbname=chatappusers port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)

	// Insert the data into the UserInfo table
	user := UserInfo{
		Username: username,
		Id:       Id,
		Password: stringedPassword,
	}

	// Insert the record into the table
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf(result.Error.Error())
	}

	fmt.Printf("User %s inserted with ID %d\n", user.Username, user.Id)

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
