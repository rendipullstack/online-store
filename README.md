# Create an Online Store Application BackEnd (RESTful API)

Ini adalah aplikasi BackEnd (RESTful API) untuk toko online yang memungkinkan pelanggan untuk melakukan beberapa tindakan seperti melihat daftar produk, menambahkan produk ke keranjang belanja, menghapus produk dari keranjang belanja, dan melakukan transaksi pembayaran. Pelanggan juga dapat melakukan login dan registrasi untuk mengakses fitur-fitur aplikasi.

## Fitur Aplikasi
Aplikasi ini memiliki beberapa fitur utama yang mencakup:
- Pelanggan dapat melihat daftar produk berdasarkan kategori produk
- Pelanggan dapat menambahkan produk ke keranjang belanja
- Pelanggan dapat melihat daftar produk yang telah ditambahkan ke keranjang belanja
- Pelanggan dapat menghapus produk dari keranjang belanja
- Pelanggan dapat checkout dan melakukan transaksi pembayaran
- Pelanggan dapat melakukan login dan registrasi

## Spesifikasi Teknis
Aplikasi ini dibangun dengan menggunakan teknologi sebagai berikut:
- Golang sebagai server
- Gin sebagai framework untuk pembuatan API
- Gorm sebagai library Object Relational Mapping
- MySQL sebagai database untuk menyimpan data produk dan pelanggan

## Instalasi
Untuk menjalankan aplikasi ini, Anda perlu menginstal beberapa dependensi terlebih dahulu. Berikut adalah langkah-langkah instalasinya:
- Clone repository ini ke dalam direktori lokal Anda.
- Jalankan perintah npm install untuk menginstal semua dependensi yang dibutuhkan.
- Clone file .env.example 
```
cp .env.example .env 
```
- Jalankan perintah go run main.go untuk menjalankan aplikasi.

## API Endpoints
Aplikasi ini memiliki beberapa endpoint API yang dapat diakses oleh klien. Berikut adalah daftar endpoint yang tersedia pada link berikut : 
<https://online-store-hanif.up.railway.app/swagger/index.html>

## Authentication
REST API: <https://online-store-hanif.up.railway.app>
|Role|Email|Password|
|----|-----|--------|
|Admin|admin@hanifz.com|qweqwe|
|Customer|user@hanifz.com|qweqwe|