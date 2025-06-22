package storage

import (
	"crud/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

// ini nama file yang bakal jadi "database"-nya, simpen data buku ke sini
var dbFile = "data/books.json"

// LoadBooks itu fungsinya buat baca data dari file JSON
// dia buka file, terus decode isinya jadi slice of Book
// kalo file-nya ga ada atau error, bakal balik error juga
func LoadBooks() ([]model.Book, error) {
	file, err := os.Open(dbFile)
	if err != nil {
		return []model.Book{}, err // kalo error, balikin slice kosong + error
	}
	defer file.Close()

	var books []model.Book
	err = json.NewDecoder(file).Decode(&books) // decode isi file ke struct books
	if err != nil {
		return []model.Book{}, err
	}
	return books, nil
}

// SaveBooks itu fungsinya buat nyimpen data buku ke file
// dia ubah struct books jadi JSON, terus simpen ke file books.json
func SaveBooks(books []model.Book) error {
	data, err := json.MarshalIndent(books, "", "  ") // biar JSON-nya rapi
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dbFile, data, 0644) // simpen ke file
}
