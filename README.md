# 📚 Books API

Books API adalah **RESTful API sederhana** yang dibuat menggunakan **Golang (Gin Framework)** dan **PostgreSQL**.
API ini menyediakan fitur untuk **autentikasi user**, **manajemen kategori**, dan **manajemen data buku**.
Swagger UI juga tersedia untuk mempermudah eksplorasi API.

---

## 🚀 Fitur Utama

### 1. Autentikasi User

* **Register**: Membuat akun baru.
* **Login**: Mendapatkan token JWT untuk autentikasi endpoint yang membutuhkan otorisasi.

### 2. Manajemen Kategori

* **Get Categories**: Mendapatkan semua kategori.
* **Create Category**: Menambahkan kategori baru.
* **Get Category by ID**: Mendapatkan detail kategori.
* **Delete Category**: Menghapus kategori tertentu.
* **Get Books by Category**: Mendapatkan daftar buku berdasarkan kategori.

### 3. Manajemen Buku

* **Get Books**: Mendapatkan semua buku.
* **Create Book**: Menambahkan buku baru dengan validasi `release_year` (1980-2024) dan perhitungan `thickness` berdasarkan jumlah halaman.
* **Get Book by ID**: Mendapatkan detail buku tertentu.
* **Delete Book**: Menghapus buku tertentu.

---

## ⚙️ Cara Menjalankan Project

### 1. Clone repository

```bash
git clone https://github.com/isadkang/books-api-golang-sanbercode.git
cd books-api-golang-sanbercode
```

### 2. Setup environment variables

Buat file `.env` di root project:

```env
PORT=8080
DATABASE_URL=postgres://username:password@host:port/dbname?sslmode=disable
```

> Contoh jika pakai Railway, copy DATABASE\_URL dari dashboard Railway.

---

### 3. Install dependencies

```bash
go mod download
```

### 4. Jalankan server

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

## 🧩 Swagger UI

Swagger UI tersedia untuk mempermudah eksplorasi API.

* **Endpoint Swagger:**

```
http://localhost:8080/swagger/index.html
```

* Swagger menampilkan semua endpoint, contoh **request & response JSON**, dan **autentikasi JWT**.

> Jika di-deploy ke Railway:

```
https://<your-app-name>.up.railway.app/swagger/index.html
```

---

## 📌 Contoh Endpoint

### **Auth**

#### Register User

```
POST /api/users/register
```

Request Body:

```json
{
  "username": "user123",
  "password": "123456"
}
```

Response:

```json
{
  "message": "User registered"
}
```

#### Login User

```
POST /api/users/login
```

Request Body:

```json
{
  "username": "user123",
  "password": "123456"
}
```

Response:

```json
{
  "token": "JWT_TOKEN_HERE"
}
```

---

### **Kategori**

#### Get All Categories

```
GET /api/categories
```

Headers: `Authorization: Bearer <token>`
Response:

```json
[
  {
    "id": 1,
    "name": "Novel"
  },
  {
    "id": 2,
    "name": "Komik"
  }
]
```

#### Create Category

```
POST /api/categories
```

Headers: `Authorization: Bearer <token>`
Request Body:

```json
{
  "name": "Science Fiction"
}
```

Response:

```json
{
  "message": "Category created"
}
```

#### Get Category by ID

```
GET /api/categories/{id}
```

Headers: `Authorization: Bearer <token>`
Response:

```json
{
  "id": 1,
  "name": "Novel"
}
```

#### Delete Category

```
DELETE /api/categories/{id}
```

Headers: `Authorization: Bearer <token>`
Response:

```json
{
  "message": "Category deleted"
}
```

#### Get Books by Category

```
GET /api/categories/{id}/books
```

Headers: `Authorization: Bearer <token>`
Response:

```json
[
  {
    "id": 1,
    "title": "Laskar Matahari",
    "description": "Novel karya Muzan Jackson",
    "release_year": 2000,
    "price": 80000,
    "total_page": 55,
    "thickness": "tipis"
  }
]
```

---

### **Buku**

#### Get All Books

```
GET /api/books
```

Headers: `Authorization: Bearer <token>`
Response:

```json
[
  {
    "id": 1,
    "title": "Laskar Matahari",
    "description": "Novel karya Muzan Jackson",
    "image_url": "https://example.com/matahari.jpg",
    "release_year": 2000,
    "price": 80000,
    "total_page": 55,
    "thickness": "tipis",
    "category_id": 1
  }
]
```

#### Create Book
- Release Year hanya bisa dari tahun 1980 - 2024
```
POST /api/books
```

Headers: `Authorization: Bearer <token>`
Request Body:

```json
{
  "title": "Laskar Matahari",
  "description": "Novel karya Muzan Jackson",
  "image_url": "https://example.com/matahari.jpg",
  "release_year": 2000,
  "price": 80000,
  "total_page": 55,
  "category_id": 1
}
```

Response:

```json
{
  "message": "Book created"
}
```

#### Get Book by ID

```
GET /api/books/{id}
```

Headers: `Authorization: Bearer <token>`
Response:

```json
{
  "id": 1,
  "title": "Laskar Matahari",
  "description": "Novel karya Muzan Jackson",
  "image_url": "https://example.com/matahari.jpg",
  "release_year": 2000,
  "price": 80000,
  "total_page": 55,
  "thickness": "tipis",
  "category_id": 1
}
```

#### Delete Book

```
DELETE /api/books/{id}
```

Headers: `Authorization: Bearer <token>`
Response:

```json
{
  "message": "Book deleted"
}
```

---

Ini bisa langsung kamu **copy-paste ke Notes**.

Kalau mau, aku bisa buatkan versi **lebih ringkas tapi tetap lengkap untuk developer** yang langsung bisa digunakan untuk testing Swagger / Postman juga.

Mau aku bikinkan versi itu juga?
