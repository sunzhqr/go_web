package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Id       uint   `json:"identificator"`
	Age      uint   `json:"age"`
	FullName string `json:"full_name"`
}

func main() {
	http.HandleFunc("/json", jsonResult)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func jsonResult(w http.ResponseWriter, r *http.Request) {
	result := Result{
		1, 19, "Sanzhar Myrzash",
	}
	response, err := json.Marshal(result)
	// err = errors.New("test error")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resultError := map[string]any{
			"ok":           false,
			"errorMessage": err.Error(),
		}
		json.NewEncoder(w).Encode(resultError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}
