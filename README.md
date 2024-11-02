# msvc-syahril-app

**Golang Boilerplate with Gin and PostgreSQL**

A boilerplate project for building scalable and maintainable web applications using Golang, the Gin web framework, and PostgreSQL as the database.

**msvc-syahril-app**

Sebuah proyek boilerplate untuk membangun aplikasi web yang skalabel dan mudah dipelihara menggunakan Golang, framework web Gin, dan PostgreSQL sebagai database.

---

## Table of Contents / Daftar Isi

- [Description / Deskripsi](#description-deskripsi)
- [Features / Fitur](#features-fitur)
- [Project Structure / Struktur Proyek](#project-structure-struktur-proyek)
- [Prerequisites / Prasyarat](#prerequisites-prasyarat)
- [Installation / Instalasi](#installation-instalasi)
- [Configuration / Konfigurasi](#configuration-konfigurasi)
- [Database Migration / Migrasi Database](#database-migration-migrasi-database)
- [Running the Application / Menjalankan Aplikasi](#running-the-application-menjalankan-aplikasi)
- [API Documentation / Dokumentasi API](#api-documentation-dokumentasi-api)
- [API Endpoints / Endpoint API](#api-endpoints-endpoint-api)
- [Response Format / Format Respons](#response-format-format-respons)
- [Example Responses / Contoh Respons](#example-responses-contoh-respons)
- [Contributing / Kontribusi](#contributing-kontribusi)
- [License / Lisensi](#license-lisensi)

---

## Description / Deskripsi

**msvc-syahril-app** is a Golang boilerplate project using the Gin web framework and PostgreSQL as the database. This project is designed to simplify the development of web applications with an organized structure and best practices.

**msvc-syahril-app** adalah proyek boilerplate Golang yang menggunakan framework web Gin dan PostgreSQL sebagai database. Proyek ini dirancang untuk memudahkan pengembangan aplikasi web dengan struktur yang terorganisir dan praktik terbaik.

---

## Features / Fitur

### English

- **Gin Web Framework**: Fast and efficient web development.
- **GORM ORM**: Simplifies database interactions with PostgreSQL.
- **CRUD Operations**: Provides Create, Read, Update, and Delete operations for the `Project` model.
- **Organized Project Structure**: Separates code by responsibility (MVC).
- **API Documentation with Swagger**: Facilitates API testing and documentation.
- **Configuration with `.env`**: Easily manage environment variables.
- **Repository and Interface Pattern**: Implements best practices for data access.
- **MVC Architecture**: Separates business logic, data, and user interface.

### Indonesian

- **Framework Web Gin**: Pengembangan web yang cepat dan efisien.
- **GORM ORM**: Mempermudah interaksi database dengan PostgreSQL.
- **Operasi CRUD**: Menyediakan operasi Create, Read, Update, dan Delete untuk model `Project`.
- **Struktur Proyek Terorganisir**: Memisahkan kode berdasarkan tanggung jawab (MVC).
- **Dokumentasi API dengan Swagger**: Memfasilitasi pengujian dan dokumentasi API.
- **Konfigurasi dengan `.env`**: Mengelola variabel lingkungan dengan mudah.
- **Pola Repository dan Interface**: Menerapkan praktik terbaik untuk akses data.
- **Arsitektur MVC**: Memisahkan logika bisnis, data, dan antarmuka pengguna.

---

## Project Structure / Struktur Proyek


### Structure Explanation / Penjelasan Struktur

#### English

- **app/**: Contains the main application code.
  - **controllers/**: Handles HTTP requests and responses.
    - **project_controller.go**: Controller for `Project` endpoints.
  - **helpers/**: Utility functions and standard response formats.
    - **response.go**: Defines standard API response structures.
  - **interfaces/**: Defines contracts for repositories.
    - **project_repository_interface.go**: Interface for `ProjectRepository`.
  - **models/**: Defines data structures and database models.
    - **project.go**: Defines the `Project` model.
  - **repositories/**: Interacts directly with the database.
    - **project_repository.go**: Implementation of `ProjectRepository` interface.
  - **services/**: Contains business logic.
    - **project_service.go**: Business logic for `Project` operations.
- **config/**: Application configuration settings.
  - **config.go**: Loads configurations from `.env`.
  - **database.go**: Sets up database connection using GORM.
- **docs/**: Contains Swagger documentation files.
  - **docs.go**: Generated Swagger documentation.
- **routes/**: Defines API routes and connects them to controllers.
  - **routes.go**: Sets up routes and connects them to controllers.
- **.env**: Environment variables file (not included in version control).
- **go.mod & go.sum**: Go module files.
- **main.go**: The main entry point of the application.

#### Indonesian

- **app/**: Berisi kode utama aplikasi.
  - **controllers/**: Menangani permintaan HTTP dan respons.
    - **project_controller.go**: Controller untuk endpoint `Project`.
  - **helpers/**: Fungsi utilitas dan format respons standar.
    - **response.go**: Mendefinisikan struktur respons API standar.
  - **interfaces/**: Mendefinisikan kontrak untuk repository.
    - **project_repository_interface.go**: Interface untuk `ProjectRepository`.
  - **models/**: Mendefinisikan struktur data dan model database.
    - **project.go**: Mendefinisikan model `Project`.
  - **repositories/**: Berinteraksi langsung dengan database.
    - **project_repository.go**: Implementasi interface `ProjectRepository`.
  - **services/**: Berisi logika bisnis.
    - **project_service.go**: Logika bisnis untuk operasi `Project`.
- **config/**: Pengaturan konfigurasi aplikasi.
  - **config.go**: Memuat konfigurasi dari `.env`.
  - **database.go**: Mengatur koneksi database menggunakan GORM.
- **docs/**: Berisi file dokumentasi Swagger.
  - **docs.go**: Dokumentasi Swagger yang dihasilkan.
- **routes/**: Mendefinisikan rute API dan menghubungkannya dengan controller.
  - **routes.go**: Mengatur rute dan menghubungkannya dengan controller.
- **.env**: File variabel lingkungan (tidak termasuk dalam kontrol versi).
- **go.mod & go.sum**: File modul Go.
- **main.go**: Titik masuk utama aplikasi.

---

## Prerequisites / Prasyarat

### English

- **Go** version 1.16 or later
- **PostgreSQL** installed and running
- **Git** for cloning the repository

### Indonesian

- **Go** versi 1.16 atau lebih baru
- **PostgreSQL** terinstal dan berjalan
- **Git** untuk mengkloning repositori

---

## Installation / Instalasi

### English

1. **Clone the repository**

    ```bash
    git clone https://github.com/yourusername/msvc-syahril-app.git
    cd msvc-syahril-app
    ```

2. **Install dependencies**

    ```bash
    go mod tidy
    ```

### Indonesian

1. **Kloning repositori**

    ```bash
    git clone https://github.com/yourusername/msvc-syahril-app.git
    cd msvc-syahril-app
    ```

2. **Instal dependensi**

    ```bash
    go mod tidy
    ```

---

## Configuration / Konfigurasi

### English

Create a `.env` file in the root directory and add the following configuration:

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
DB_SSLMODE=disable

APP_PORT=8080
