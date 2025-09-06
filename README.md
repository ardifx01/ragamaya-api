# RagaMaya API

<div align="center">
<img src="https://cdn.xann.my.id/ragamaya/59d42d65-43ee-4cc3-ba98-a1ae341d3a78.png" alt="Logo RagaMaya" width="200"/>
<h3>Temukan Makna, Hidupkan Budaya, Bersama RagaMaya</h3>

[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-latest-DC382D?style=flat-square&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-latest-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)
[![AWS](https://img.shields.io/badge/AWS-S3-232F3E?style=flat-square&logo=amazon-aws)](https://aws.amazon.com/)
[![Midtrans](https://img.shields.io/badge/Midtrans-Payment-blue?style=flat-square)](https://midtrans.com/)

</div>

## 📖 Tentang

RagaMaya adalah platform digital berbasis web yang dirancang untuk melestarikan, memperkenalkan, dan mengembangkan budaya Indonesia, khususnya batik, melalui teknologi modern. Repositori ini berisi API backend yang mendukung platform RagaMaya.

🌐 **Kunjungi platform kami: [ragamaya.space](https://ragamaya.space)**

Platform ini menggabungkan teknologi AI, pendidikan digital, dan fungsionalitas marketplace untuk menciptakan solusi inovatif dalam melestarikan batik sebagai warisan budaya sekaligus membuka peluang ekonomi kreatif bagi masyarakat.

## 🚀 Fitur

- Otentikasi dan Otorisasi
- Manajemen Pengguna
- Manajemen Produk
- Pemrosesan Pesanan
- Integrasi Pembayaran (Midtrans)
- Sistem Kuis
- Manajemen Artikel
- Manajemen Penyimpanan (AWS S3)
- Analitik
- Sistem Dompet
- Notifikasi Email
- Integrasi WhatsApp

## 🛠️ Teknologi yang Digunakan

- **Bahasa Pemrograman:** Go
- **Database:** PostgreSQL
- **Caching:** Redis
- **Penyimpanan Cloud:** AWS S3
- **Payment Gateway:** Midtrans
- **Kontainerisasi:** Docker
- **Layanan Email:** SMTP
- **Dokumentasi:** Swagger/OpenAPI

## ⚙️ Variabel Environment

Buat file `.env` di direktori root dan tambahkan variabel berikut:

```env
# Konfigurasi Database
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=

# Konfigurasi Server
PORT=

# Otentikasi
JWT_SECRET=
INTERNAL_SECRET=

# Environment
ENVIRONMENT=

# Kredensial Admin
ADMIN_USERNAME=
ADMIN_PASSWORD=

# Konfigurasi Redis
REDIS_ADDR=
REDIS_PASS=

# Konfigurasi AWS
AWS_ACCESS_KEY=
AWS_SECRET_KEY=
AWS_REGION=
AWS_BUCKET=
STORAGE_FOLDER=

# Konfigurasi Midtrans
MIDTRANS_SERVER_KEY=
MIDTRANS_ENV=

# Layanan Eksternal
FRONTEND_BASE_URL=
MLSERVICE_BASE_URL=

# Konfigurasi SMTP
SMTP_EMAIL=
SMTP_PASSWORD=
SMTP_SERVER=
SMTP_PORT=
```

## 🚀 Cara Memulai

1. Clone repositori
```bash
git clone https://github.com/RagaMaya/ragamaya-api.git
```

2. Install dependensi
```bash
go mod download
```

3. Setup variabel environment (salin dari .env.example)
```bash
cp .env.example .env
```

4. Jalankan aplikasi
```bash
make run
```

Atau menggunakan Docker:
```bash
docker-compose up -d
```

## 📁 Struktur Proyek

```
.
├── api/      # Modul API (analytics, articles, products, dll.)
├── cmd/      # Aplikasi utama
├── emails/   # Template email dan layanan
├── internal/ # Package internal
├── models/   # Model database
├── pkg/      # Package utility
└── routers/  # Route API
```

## 📄 Lisensi

Proyek ini dilisensikan di bawah ketentuan lisensi yang disediakan dalam repositori.

## 👥 Kontributor

### Tim Pengembangan
- [Rama Diaz](https://github.com/ramadiaz) - Backend Developer
- [Fahry Firdaus](https://github.com/Fahry169) - Frontend Developer
- [Kevin Sipahutar](https://github.com/vinss-droid) - Frontend Developer

---

<div align="center">
<p>© 2025 RagaMaya. Semua Hak Dilindungi.</p>
</div>