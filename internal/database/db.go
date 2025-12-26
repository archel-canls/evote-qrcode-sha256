package database

import (
	"fmt"
	"log"

	"evote-qrcode-sha256/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Mengambil variabel lingkungan menggunakan helper dari config.go
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "")
	dbname := config.GetEnv("DB_NAME", "evote_db")
	sslmode := config.GetEnv("DB_SSLMODE", "disable")

	// --- TAHAP 1: KONEKSI KE DATABASE DEFAULT 'postgres' ---
	// Digunakan untuk mengecek/membuat database target
	dsnSystem := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s TimeZone=UTC",
		host, user, password, port, sslmode)

	systemDb, err := gorm.Open(postgres.Open(dsnSystem), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal menyambung ke server PostgreSQL: %v", err)
	}

	// Periksa keberadaan database di katalog sistem PostgreSQL
	var exists int
	systemDb.Raw("SELECT 1 FROM pg_database WHERE datname = ?", dbname).Scan(&exists)

	if exists == 0 {
		log.Printf("ℹ️ Database '%s' tidak ditemukan, mencoba membuat...", dbname)
		// Jalankan perintah pembuatan database (Raw SQL)
		if err := systemDb.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname)).Error; err != nil {
			log.Fatalf("❌ Gagal membuat database otomatis: %v", err)
		}
		log.Printf("✅ Database '%s' berhasil dibuat", dbname)
	}

	// Tutup koneksi ke database sistem
	sqlDB, _ := systemDb.DB()
	sqlDB.Close()

	// --- TAHAP 2: KONEKSI KE DATABASE APLIKASI ---
	// Sekarang kita menyambung ke database target yang sudah dipastikan ada
	dsnApp := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsnApp), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal menyambung ke database '%s': %v", err)
	}

	DB = db
	log.Printf("🚀 Berhasil terhubung ke database '%s'", dbname)
}
