package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type Result struct {
	Id       uint   `json:"identificator"`
	Age      uint   `json:"age"`
	FullName string `json:"full_name"`
}

type User struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Grade    string `json:"grade"`
	Language string `json:"lang"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.FullName, validation.Required, is.Alpha),
		validation.Field(&u.Grade, validation.Required, validation.NotNil, is.Alpha),
		validation.Field(&u.Email, validation.Required, validation.NotNil, is.Email),
		validation.Field(&u.Phone, validation.Required, is.E164),
	)
}

func main() {
	http.HandleFunc("/json", JsonHandler)
	http.HandleFunc("/user", UserHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func WriteJson(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		100, 19, "Golang FastAPI",
	}
	err := WriteJson(w, http.StatusOK, result)
	// err = errors.New("test")
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]interface{}{
			"ok":           false,
			"errorMessage": err.Error(),
		})
		return
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteJson(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"ok":    false,
			"error": "Method Not Allowed",
		})
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Decoding Error:", err.Error())
		WriteJson(w, http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": "Internal Server Error",
		})
		return
	}
	err = user.Validate()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]interface{}{
			"ok":    false,
			"error": "Bad Request",
		})
		fmt.Fprintln(w, "Validation Error:", err.Error())
		fmt.Println("Validation Error:", err.Error())
		return
	}
	fmt.Printf("data from request %v\n", user)
	WriteJson(w, http.StatusOK, map[string]interface{}{
		"ok":      true,
		"message": "Succesfully decoded!",
	})
}
