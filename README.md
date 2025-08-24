# 📚 Books API

Books API adalah RESTful API sederhana yang dibuat menggunakan **Golang (Gin Framework)**.  
API ini menyediakan fitur untuk **autentikasi user**, **mengelola kategori**, dan **mengelola data buku**.

---

## 🚀 Fitur Utama
- **Autentikasi User**
  - Register & Login dengan JWT.
- **Manajemen Kategori**
  - CRUD kategori.
  - Mendapatkan semua buku dari kategori tertentu.
- **Manajemen Buku**
  - CRUD buku.
  - Perhitungan otomatis `thickness` berdasarkan total halaman.
  - Validasi `release_year` antara 1980 - 2024.

---

## ⚙️ Cara Menjalankan Project

### 1. Clone repository
```bash
git clone https://github.com/username/books-api.git
cd books-api
