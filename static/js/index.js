const api = 'http://localhost:8080/books';
let booksData = [];
let selectedIdToDelete = null;

function renderBooks(books) {
  const list = document.getElementById('book-list');
  list.innerHTML = '';
  books.forEach((book) => {
    list.innerHTML += `
          <tr>
            <td>${book.id}</td>
            <td>${book.title}</td>
            <td>${book.author}</td>
            <td>
              <div class="d-flex gap-1">
                  <button class="btn btn-sm btn-primary" onclick="showEditModal('${book.id}', '${book.title.replace(/'/g, "\\'")}', '${book.author.replace(/'/g, "\\'")}')">Edit</button>
                  <button class="btn btn-sm btn-danger" onclick="showDeleteModal('${book.id}', '${book.title.replace(/'/g, "\\'")}')">Hapus</button>
                </div>
              </td>
          </tr>`;
  });
}

function loadBooks() {
  fetch(api)
    .then((res) => res.json())
    .then((data) => {
      booksData = data;
      renderBooks(data);
    });
}

function showDeleteModal(id, title) {
  selectedIdToDelete = id;
  document.getElementById('bookInfo').innerText = `"${title}" (ID: ${id})`;
  new bootstrap.Modal(document.getElementById('deleteModal')).show();
}

function showEditModal(id, title, author) {
  document.getElementById('edit-id').value = id;
  document.getElementById('edit-title').value = title;
  document.getElementById('edit-author').value = author;
  new bootstrap.Modal(document.getElementById('editModal')).show();
}

document.getElementById('confirmDelete').addEventListener('click', function () {
  if (selectedIdToDelete) {
    fetch(`${api}/${selectedIdToDelete}`, { method: 'DELETE' }).then(() => {
      selectedIdToDelete = null;
      bootstrap.Modal.getInstance(document.getElementById('deleteModal')).hide();
      loadBooks();
    });
  }
});

document.getElementById('add-form').addEventListener('submit', function (e) {
  e.preventDefault();
  const title = document.getElementById('title').value;
  const author = document.getElementById('author').value;

  fetch(api, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title, author }),
  }).then(() => {
    bootstrap.Modal.getInstance(document.getElementById('addModal')).hide();
    this.reset();
    loadBooks();
  });
});

document.getElementById('edit-form').addEventListener('submit', function (e) {
  e.preventDefault();
  const id = document.getElementById('edit-id').value;
  const title = document.getElementById('edit-title').value;
  const author = document.getElementById('edit-author').value;

  fetch(`${api}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id, title, author }),
  }).then(() => {
    bootstrap.Modal.getInstance(document.getElementById('editModal')).hide();
    loadBooks();
  });
});

document.getElementById('search').addEventListener('input', function () {
  const keyword = this.value.toLowerCase();
  const filtered = booksData.filter((book) => book.title.toLowerCase().includes(keyword) || book.author.toLowerCase().includes(keyword));
  renderBooks(filtered);
});

loadBooks();
