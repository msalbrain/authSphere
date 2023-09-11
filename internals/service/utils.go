package service


import (
	"crypto/sha256"
	"fmt"
	"time"
	cryrand "crypto/rand"
	"math/rand"
	"encoding/hex"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	uuid "github.com/google/uuid"
)


type Claims struct {
	admin bool
	jwt.RegisteredClaims
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}



func GenerateRandomString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func HashParameter(param string) string {
	hash := sha256.Sum256([]byte(param))
	return fmt.Sprintf("%x", hash)
}

func GenerateToken(currentTime, name, email string) string {
	// TODO: implement properly
	hashedTime := HashParameter(currentTime)
	hashedName := HashParameter(name)
	hashedEmail := HashParameter(email)
	hashedUUID := HashParameter(uuid.NewString())

	combinedHash := hashedTime + hashedName + hashedEmail + hashedUUID + "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := GenerateRandomString(24, combinedHash)

	return combinedHash + randomString
}

type CustomID string

func GenerateCustomID() CustomID {
	// Generate a 12-byte random value
	randomBytes := make([]byte, 12)
	_, err := cryrand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Get the current timestamp in seconds
	timestamp := time.Now().Unix()

	// Concatenate the timestamp and random bytes
	combinedBytes := append([]byte(fmt.Sprintf("%x", timestamp)), randomBytes...)

	// Convert the combined bytes to a hexadecimal string
	idHex := hex.EncodeToString(combinedBytes)

	return CustomID(idHex)
}


