package main

import (
	"log"

	"github.com/Andhika-GIT/wild_oasis_be/internal/infrastructure/config"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func main() {
	app := config.Bootstrap()

	tokenAuth = jwtauth.New("HS256", []byte(app.Global.GetString("JWT_SECRET")), nil)

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"user_id": 123})

	if err != nil {
		log.Fatalf("error while encoding token : %v", err)
	} else {
		log.Printf("token result : %s", tokenString)
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("error while hashing password : %v", err)
	} else {
		log.Printf("hashed password : %s", string(hashedBytes))
	}

}
