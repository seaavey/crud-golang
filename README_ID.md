# 📚 CRUD Buku – Golang

Proyek ini merupakan aplikasi sederhana yang memfasilitasi operasi CRUD (Create, Read, Update, Delete) untuk data buku menggunakan bahasa pemrograman **Go (Golang)**. Aplikasi ini menyimpan data dalam file **JSON**, sehingga cocok sebagai media pembelajaran backend tanpa perlu mengatur basis data eksternal.

## 📁 Struktur Direktori

```
.
├── main.go            # Titik masuk utama aplikasi
├── handler/           # Handler untuk request HTTP
│   └── book.go
├── model/             # Definisi struktur data Book
│   └── book.go
├── storage/           # Fungsi baca/tulis ke file JSON
│   └── book.go
├── data/
│   └── books.json     # File penyimpanan data buku
├── static/            # File statis (HTML, CSS, JS jika dibutuhkan)
```

## 🚀 Cara Menjalankan

1. Clone repositori:

   ```bash
   git clone https://github.com/seaavey/crud-golang
   cd crud-golang
   ```

2. Jalankan aplikasi:

   ```bash
   go run main.go
   ```

3. Buka di browser:

   ```
   http://localhost:8080
   ```

## 🔗 Endpoint API

| Method | Endpoint      | Deskripsi                     |
| ------ | ------------- | ----------------------------- |
| POST   | `/books`      | Menambahkan buku baru         |
| GET    | `/books`      | Mengambil semua data          |
| GET    | `/books/{id}` | Mengambil buku berdasarkan ID |
| PUT    | `/books/{id}` | Memperbarui buku              |
| DELETE | `/books/{id}` | Menghapus buku berdasarkan ID |

## 📝 Contoh Data Buku (Format JSON)

```json
{
  "id": "1",
  "title": "Belajar Golang",
  "author": "John Doe"
}
```

## ℹ️ Catatan Tambahan

- Data buku disimpan di dalam file `data/books.json`.
- Jika folder `data/` atau file `books.json` belum tersedia, buat secara manual:
  ```bash
  mkdir data
  echo "[]" > data/books.json
  ```

## ⚙️ Teknologi yang Digunakan

- [Golang](https://golang.org)
- [Gorilla Mux](https://github.com/gorilla/mux)
- File JSON sebagai media penyimpanan data

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah [MIT License](./LICENSE) – © 2025 Muhammad Adriansyah
