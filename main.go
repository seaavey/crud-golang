package main

import (
	"crud/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// bikin router baru dari gorilla mux
	r := mux.NewRouter()

	// ini daftar route yang dipake buat operasi CRUD buku
	r.HandleFunc("/books", handler.CreateBook).Methods("POST")     // buat nambah buku
	r.HandleFunc("/books", handler.GetBooks).Methods("GET")        // ambil semua buku
	r.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")    // ambil satu buku by ID
	r.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT") // update buku by ID
	r.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE") // hapus buku by ID

	// ini buat handle file statis, misal HTML/CSS/JS di folder static
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/").Handler(fs)

	// nyalain server di port 8080
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
