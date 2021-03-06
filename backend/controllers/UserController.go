package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fourtf/studyhub/models"
	"github.com/fourtf/studyhub/utils"
	"golang.org/x/crypto/bcrypt"
)

var db = utils.ConnectToDB()

//CreateUser Creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}

//Login sends the user a valid jwt token to access authed paths
func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := findOne(user.Name, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func findOne(name, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Name = ?", name).First(user).Error; err != nil {
		var resp = map[string]interface{}{"message": "User not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk.StandardClaims)
	utils.LoadEnvironmentVariables()
	tokenString, error := token.SignedString([]byte(os.Getenv("tokenSigningKey")))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"message": "logged in"}
	resp["Token"] = tokenString //Store the token in the response
	return resp
}
