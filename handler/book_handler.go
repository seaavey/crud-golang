package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"crud/model"
	"crud/storage"
)

// CreateBook menambahkan buku baru dari request JSON
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Request tidak valid", http.StatusBadRequest)
		return
	}

	// Validasi sederhana
	if book.Title == "" || book.Author == "" {
		http.Error(w, "Judul dan penulis wajib diisi", http.StatusBadRequest)
		return
	}

	books, err := storage.LoadBooks()
	if err != nil {
		http.Error(w, "Gagal memuat data buku", http.StatusInternalServerError)
		return
	}

	// Cari ID terbesar
	maxID := 0
	for _, b := range books {
		var id int
		fmt.Sscanf(b.ID, "%d", &id)
		if id > maxID {
			maxID = id
		}
	}
	book.ID = fmt.Sprintf("%d", maxID+1)

	books = append(books, book)
	if err := storage.SaveBooks(books); err != nil {
		http.Error(w, "Gagal menyimpan data buku", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBooks mengembalikan semua data buku
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := storage.LoadBooks()
	if err != nil {
		http.Error(w, "Gagal memuat data buku", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

// GetBook mengembalikan satu buku berdasarkan ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	books, err := storage.LoadBooks()
	if err != nil {
		http.Error(w, "Gagal memuat data buku", http.StatusInternalServerError)
		return
	}

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.NotFound(w, r)
}

// UpdateBook memperbarui data buku berdasarkan ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var updated model.Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Request tidak valid", http.StatusBadRequest)
		return
	}

	if updated.Title == "" || updated.Author == "" {
		http.Error(w, "Judul dan penulis wajib diisi", http.StatusBadRequest)
		return
	}

	books, err := storage.LoadBooks()
	if err != nil {
		http.Error(w, "Gagal memuat data buku", http.StatusInternalServerError)
		return
	}

	for i, book := range books {
		if book.ID == id {
			updated.ID = id // pastikan ID tetap sama
			books[i] = updated
			if err := storage.SaveBooks(books); err != nil {
				http.Error(w, "Gagal menyimpan data buku", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.NotFound(w, r)
}

// DeleteBook menghapus data buku berdasarkan ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	books, err := storage.LoadBooks()
	if err != nil {
		http.Error(w, "Gagal memuat data buku", http.StatusInternalServerError)
		return
	}

	newBooks := []model.Book{}
	found := false
	for _, book := range books {
		if book.ID != id {
			newBooks = append(newBooks, book)
		} else {
			found = true
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	if err := storage.SaveBooks(newBooks); err != nil {
		http.Error(w, "Gagal menyimpan data buku", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
