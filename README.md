# 📚 Book CRUD – Golang

This project is a simple application that facilitates CRUD operations (Create, Read, Update, Delete) for book data using the **Go (Golang)** programming language. It uses a **JSON file** for data storage, making it a great resource for learning backend development without the complexity of setting up a database.

## 📁 Project Structure

```
.
├── main.go            # Main entry point of the application
├── handler/           # HTTP request handlers
│   └── book.go
├── model/             # Data structure definition for Book
│   └── book.go
├── storage/           # Functions to read/write from JSON file
│   └── book.go
├── data/
│   └── books.json     # Book data storage file
├── static/            # Static files (HTML, CSS, JS if needed)
```

## 🚀 Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/seaavey/crud-golang
   cd crud-golang
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. Open your browser and go to:

   ```
   http://localhost:8080
   ```

## 🔗 API Endpoints

| Method | Endpoint      | Description                    |
| ------ | ------------- | ------------------------------ |
| POST   | `/books`      | Add a new book                 |
| GET    | `/books`      | Retrieve all book data         |
| GET    | `/books/{id}` | Retrieve a specific book by ID |
| PUT    | `/books/{id}` | Update a book                  |
| DELETE | `/books/{id}` | Delete a book by ID            |

## 📝 Sample Book Data (JSON Format)

```json
{
  "id": "1",
  "title": "Learn Golang",
  "author": "John Doe"
}
```

## ℹ️ Additional Notes

- Book data is stored in the `data/books.json` file.
- If the `data/` folder or `books.json` file does not exist, create it manually:
  ```bash
  mkdir data
  echo "[]" > data/books.json
  ```

## ⚙️ Technologies Used

- [Golang](https://golang.org)
- [Gorilla Mux](https://github.com/gorilla/mux)
- JSON file-based data storage

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE) – © 2025 Muhammad Adriansyah
