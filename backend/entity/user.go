package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	AuthToken string
}

func NewUser(name, email, password string) *User {
	// UserID
	uid, err := generateUserID()
	if err != nil {
		log.Println(err)
		return nil
	}

	// Password Hash
	sha256Password := sha256.Sum256([]byte(password))
	hash := hex.EncodeToString(sha256Password[:])

	// AuthToken
	authToken, err := generateAuthToken()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &User{
		ID:        uid,
		Name:      name,
		Email:     email,
		Password:  hash,
		AuthToken: authToken,
	}
}

func generateUserID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("error: generateUserID")
	}
	return u.String(), nil
}

func generateAuthToken() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("error: generateAuthToken")
	}
	return u.String(), nil
}
