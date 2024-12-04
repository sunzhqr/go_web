package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Id       uint   `json:"identificator"`
	Age      uint   `json:"age"`
	FullName string `json:"full_name"`
}

func main() {
	http.HandleFunc("/json", UserHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func WriteJson(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
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
