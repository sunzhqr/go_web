package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(index))
	router.HandleFunc("/hello/{username}", hello)
	router.HandleFunc(`/product/{id:\d+}`, product)
	router.Handle(`/form`, http.HandlerFunc(form)).Methods(http.MethodPost, http.MethodPut)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFound(w, r)
		return
	}
	w.Write([]byte("Home"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["username"]
	io.WriteString(w, fmt.Sprintln("Hello,", user))
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	fmt.Fprintf(w, "Product ID: %d", id)
}

func form(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Form for PUT or POST")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Page Not Found")
}
