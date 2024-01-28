package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error 405: Method is not allowed!", http.StatusMethodNotAllowed)
	}
	_, err := w.Write([]byte("Hello dude!"))
	if err != nil {
		fmt.Println("Error!")
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Problem with form parsing!", http.StatusInternalServerError)
		return
	}
	_, err := fmt.Fprintf(w, "Success!\n")
	if err != nil {
		fmt.Println("Error!")
	}
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Surname: %s\n", surname)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main() {
	srv := http.Server{
		Addr: ":8080",
	}
	serverFiles := http.FileServer(http.Dir("./serverFiles"))
	http.Handle("/", serverFiles)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Starting server...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
