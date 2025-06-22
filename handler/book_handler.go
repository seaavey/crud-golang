// ini buat menakses file ini "handler/nama_file"
package handler

// import module-module do butuhkan kaya json, fmt, http, mux dari gorilla
// dan saya juga import file lain seperti dari folder model dan storage
// nah folder model itu buat tau struktur data Book itu kek mana
// sedangkan storage itu buat akses dan ubah data ke "database" (bisa file, dsb)

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"crud/model"
	"crud/storage"
)

// ini fungsi buat nambahin buku baru
// dia bakal baca body request dari client (format JSON)
// terus nyari ID paling gede dari data yang udah ada, trus tambahin 1
// abis itu datanya disimpan, dan ngirim balik data buku yang baru
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)

	books, _ := storage.LoadBooks()

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
	storage.SaveBooks(books)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// ini fungsi buat ngambil semua data buku
// dia langsung ngeload dari storage dan ngirim dalam bentuk JSON
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := storage.LoadBooks()
	json.NewEncoder(w).Encode(books)
}

// ini fungsi buat ngambil 1 buku aja berdasarkan ID
// jadi ambil ID-nya dari URL, trus cek datanya
// kalo ada, kirim datanya. kalo ga ada, kirim 404
func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	books, _ := storage.LoadBooks()

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

// ini buat update data buku berdasarkan ID
// ambil ID dari URL, terus baca data baru dari body request
// cari ID yang cocok, kalo ketemu update datanya dan simpan
// kalo ga ketemu, kirim 404
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated model.Book
	json.NewDecoder(r.Body).Decode(&updated)

	books, _ := storage.LoadBooks()
	for i, book := range books {
		if book.ID == id {
			books[i] = updated
			storage.SaveBooks(books)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

// ini buat hapus data buku berdasarkan ID
// ambil ID dari URL, terus filter data yang bukan ID itu
// simpan lagi datanya, trus kirim status sukses tanpa isi (204)
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	books, _ := storage.LoadBooks()

	newBooks := []model.Book{}
	for _, book := range books {
		if book.ID != id {
			newBooks = append(newBooks, book)
		}
	}
	storage.SaveBooks(newBooks)
	w.WriteHeader(http.StatusNoContent)
}
