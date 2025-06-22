package model

// ini struct Book, gunanya buat nyimpen data buku
// ID itu unik, semacam primary key-nya
// Title itu judul bukunya
// Author itu nama penulis bukunya
// tag `json:"..."` itu biar waktu encode/decode ke JSON pakai nama itu
type Book struct {
	ID     string `json:"id"`     // ID buku, bentuknya string
	Title  string `json:"title"`  // Judul buku
	Author string `json:"author"` // Nama penulis
}
