# Go Bengkel

Backend REST API untuk membantu pencatatan data pelanggan, kendaraan, dan layanan perbaikan kendaraan pada bengkel.

Aplikasi ini dibangun menggunakan Golang dengan layered architecture serta menerapkan JWT authentication dan Role-Based Access Control (RBAC) untuk mengatur hak akses pengguna.

## Features

- User registration
- User login menggunakan JWT
- User management
- Vehicle management
- Pencatatan data layanan/perbaikan kendaraan
- Change password
- Role-Based Access Control (RBAC)
  - `admin`
  - `user`
- RESTful API
- API documentation menggunakan Swagger
- PostgreSQL sebagai database
- Docker & Docker Compose untuk menjalankan aplikasi

## Tech Stack

| Technology | Description |
|---|---|
| Go | Backend programming language |
| Gin | HTTP web framework |
| GORM | ORM untuk database |
| PostgreSQL | Relational database |
| JWT | Authentication |
| bcrypt | Password hashing |
| Swagger | API documentation |
| Docker | Containerization |
| Docker Compose | Menjalankan application stack |

## Architecture

Project menggunakan layered architecture untuk memisahkan tanggung jawab setiap bagian aplikasi.

```text
Client
  │
  ▼
Handler
  │
  ▼
Service
  │
  ▼
Repository
  │
  ▼
PostgreSQL
```

### Layer

- **Handler**  
  Menangani HTTP request dan response dari client.

- **Service**  
  Menangani business logic aplikasi.

- **Repository**  
  Menangani komunikasi dan operasi terhadap database.

- **Model**  
  Mendefinisikan struktur data dan relationship database menggunakan GORM.

- **DTO**  
  Mendefinisikan struktur data untuk request dan response API.

- **Middleware**  
  Menangani authentication dan authorization.

- **Utils**  
  Berisi fungsi pendukung seperti pembuatan dan validasi JWT.

## Project Structure

```text
.
├── cmd/
│   ├── api/
│   │   └── main.go
│   ├── migrate/
│   │   └── migrate.go
│   └── seed/
│       └── seed.go
│
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── dto/
│   ├── middleware/
│   └── utils/
│
├── docs/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── .env.example
```

## Authentication & Authorization

Aplikasi menggunakan **JWT (JSON Web Token)** untuk authentication.

Setelah berhasil login, client memperoleh JWT yang digunakan untuk mengakses protected endpoint.

Request ke endpoint yang membutuhkan authentication harus menyertakan token pada header:

```http
Authorization: Bearer <JWT_TOKEN>
```

### RBAC

Aplikasi menerapkan **Role-Based Access Control** dengan dua role:

| Role | Access |
|---|---|
| `user` | Mengakses endpoint yang membutuhkan authentication |
| `admin` | Memiliki akses user serta dapat melihat seluruh user dan menghapus user |

Endpoint yang membutuhkan role `admin` dilindungi menggunakan middleware:

```text
RequiredRole("admin")
```

## Database Relationship

Relasi utama pada database:

```text
Role
 │
 │ 1:N
 ▼
User
 │
 │ 1:N
 ▼
Vehicle
 │
 │ 1:N
 ▼
DataService
```

### Entity

#### User

Menyimpan informasi pelanggan/user dan memiliki relasi dengan role serta kendaraan.

#### Role

Menyimpan role pengguna, yaitu `admin` dan `user`.

#### Vehicle

Menyimpan data kendaraan yang dimiliki oleh user.

#### DataService

Menyimpan informasi layanan atau perbaikan kendaraan, termasuk nama part dan amount yang dicatat.

## API Endpoints

### Authentication

| Method | Endpoint | Access | Description |
|---|---|---|---|
| `POST` | `/api/auth/register` | Public | Register user |
| `POST` | `/api/auth/login` | Public | Login user |

### User

| Method | Endpoint | Access | Description |
|---|---|---|---|
| `PATCH` | `/api/users/me/password` | Authenticated | Mengubah password user |
| `GET` | `/api/users` | Admin | Mendapatkan seluruh user |
| `GET` | `/api/users/:id` | Admin | Mendapatkan user berdasarkan ID |
| `DELETE` | `/api/users/:id` | Admin | Menghapus user |

### Vehicle

| Method | Endpoint | Access | Description |
|---|---|---|---|
| `POST` | `/api/vehicles` | Authenticated | Menambahkan kendaraan |
| `GET` | `/api/vehicles` | Authenticated | Mendapatkan seluruh kendaraan |

### Vehicle Data Service

| Method | Endpoint | Access | Description |
|---|---|---|---|
| `POST` | `/api/vehicles/:vehicle_id/dataservices` | Authenticated | Menambahkan data service kendaraan |
| `GET` | `/api/vehicles/:vehicle_id/dataservices` | Authenticated | Mendapatkan data service kendaraan |

## API Documentation

Dokumentasi API tersedia menggunakan Swagger.

Setelah aplikasi berjalan, Swagger dapat digunakan untuk melihat endpoint, request, response, dan mencoba API secara langsung.

```text
/swagger/index.html
```

## Environment Variables

Buat file `.env` berdasarkan `.env.example`.

```env
DB_URL=your_database_connection_string
JWT_SECRET=your_jwt_secret
```

> Jangan commit file `.env` ke repository. Gunakan `.env.example` sebagai template konfigurasi.

## Running with Docker Compose

Pastikan Docker dan Docker Compose telah terinstall.

Clone repository:

```bash
git clone <repository-url>
cd <project-directory>
```

Buat file environment:

```bash
cp .env.example .env
```

Sesuaikan konfigurasi pada `.env`, kemudian jalankan:

```bash
docker compose up -d --build
```

Untuk melihat container yang sedang berjalan:

```bash
docker compose ps
```

Untuk melihat log aplikasi:

```bash
docker compose logs -f
```

Untuk menghentikan aplikasi:

```bash
docker compose down
```

## Database Migration & Seed

Database migration dan seeding dijalankan melalui command yang tersedia pada project.

Migration digunakan untuk membuat atau memperbarui struktur tabel berdasarkan model GORM, sedangkan seed digunakan untuk memasukkan data awal yang dibutuhkan aplikasi.

## License

This project is for educational and portfolio purposes.
