package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fourtf/studyhub/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//CreateUser Creates a new user in the database
func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		json.NewDecoder(r.Body).Decode(user)
		if isUserValid(user) {
			pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Println(err)
				json.NewEncoder(w).Encode(err)
			}

			user.Password = string(pass)

			createdUser := db.Create(user)
			var errMessage = createdUser.Error

			if createdUser.Error != nil {
				log.Println(errMessage)
			}
			resp := findOne(db, user.Name, user.Password)
			json.NewEncoder(w).Encode(resp)
		} else {
			json.NewEncoder(w).Encode(models.AuthResponse{Message: "Invalid credentials"})
		}
	}
}

//Login sends the user a valid jwt token to access authed paths
func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			var resp = models.AuthResponse{Message: "Invalid request"}
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := findOne(db, user.Name, user.Password)
		json.NewEncoder(w).Encode(resp)
	}
}

func findOne(db *gorm.DB, name, password string) models.AuthResponse {
	user := &models.User{}

	if err := db.Where("name = ?", name).First(user).Error; err != nil {
		var resp = models.AuthResponse{Message: "User not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = models.AuthResponse{Message: "Invalid login credentials. Please try again"}
		return resp
	}

	claims := &models.Claims{UserID: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expiresAt}}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("tokenSigningKey")))
	if err != nil {
		log.Println(err)
	}

	var resp = models.AuthResponse{Message: "Logged in", Token: tokenString}
	return resp
}

var (
	nameRegEx         = regexp.MustCompile(`^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$`)
	emailRegEx        = regexp.MustCompile(`^\S+@\S+\.\S+$`)
	specialCharacters = regexp.MustCompile(`[@$!%*#?&]`)
)

func isUserValid(user *models.User) bool {
	var isNameValid = nameRegEx.MatchString(user.Name)
	var isEmailValid = emailRegEx.MatchString(user.Email)
	var isPasswordValid = validatePassword(user.Password)
	return isNameValid && isEmailValid && isPasswordValid
}

func validatePassword(password string) bool {
	containsSpecialCharacter := specialCharacters.MatchString(password)
	isLongEnough := len(password) > 7
	return isLongEnough && containsSpecialCharacter
}
