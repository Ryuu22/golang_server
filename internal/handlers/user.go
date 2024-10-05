package handlers

import (
	"encoding/json"
	"net/http"

	"golang_server.dankbueno.com/internal/middleware"
	"golang_server.dankbueno.com/internal/models"
	"golang_server.dankbueno.com/internal/repositories"
	"golang_server.dankbueno.com/internal/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData models.User
	_ = json.NewDecoder(r.Body).Decode(&userData)

	// Verify and encrypt password
	password, err := middleware.VerifyPassword(userData.Password)
	if err != nil {
		response := map[string]string{
			"message": "Password Error",
			"error":   err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	userData.Password = password

	err = repositories.CreateUser(userData)

	if err == nil {
		response := map[string]string{
			"message":  "User Created",
			"username": userData.Username,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{
		"message": "User Already Exists",
		"error":   err.Error(),
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userData models.User
	_ = json.NewDecoder(r.Body).Decode(&userData)

	user, err := repositories.GetUser(userData.Username)
	if err != nil {
		response := map[string]string{
			"message": "User Not Found",
			"error":   err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !middleware.ComparePassword(user.Password, userData.Password) {
		response := map[string]string{
			"message": "Invalid Password",
			"error":   err.Error(),
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := middleware.GenerateJWTFromUser(user)

	response := map[string]string{
		"message":  "User Logged In",
		"username": userData.Username,
		"token":    token,
	}
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	jwt := utils.GetJWTFromAuthHeader(r.Header.Get("Authorization"))

	payload, err := middleware.VerifyJWT(jwt)

	if err != nil {
		response := map[string]string{
			"message": "Invalid Token",
			"error":   err.Error(),
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// return user id
	json.NewEncoder(w).Encode(payload)
}
