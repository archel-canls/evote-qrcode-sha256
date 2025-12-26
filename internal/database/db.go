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
	// 1. PRIORITAS: Menggunakan DATABASE_URL (Sangat disarankan untuk Railway)
	// Masukkan URL ini di tab Variables Railway dengan Key: DATABASE_URL
	dsn := config.GetEnv("DATABASE_URL", "")

	if dsn != "" {
		log.Println("🌐 Menyambung ke database menggunakan DATABASE_URL...")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			// Perbaikan: Jumlah placeholder harus sama dengan jumlah variabel
			log.Fatalf("❌ Gagal menyambung via DATABASE_URL: %v", err)
		}
		DB = db
		log.Println("🚀 Berhasil terhubung ke database Railway via URL")
		return
	}

	// 2. FALLBACK: Menggunakan parameter individual (Hanya jika DATABASE_URL kosong/Lokal)
	log.Println("🏠 DATABASE_URL tidak ditemukan, beralih ke konfigurasi manual...")
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "")
	dbname := config.GetEnv("DB_NAME", "evote_db")
	sslmode := config.GetEnv("DB_SSLMODE", "disable")

	// --- TAHAP 1: KONEKSI KE DATABASE SISTEM 'postgres' (UNTUK AUTO-CREATE DB DI LOKAL) ---
	dsnSystem := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s TimeZone=UTC",
		host, user, password, port, sslmode)

	systemDb, err := gorm.Open(postgres.Open(dsnSystem), &gorm.Config{})
	if err == nil {
		var exists int
		systemDb.Raw("SELECT 1 FROM pg_database WHERE datname = ?", dbname).Scan(&exists)

		if exists == 0 {
			log.Printf("ℹ️ Database '%s' tidak ditemukan, mencoba membuat...", dbname)
			if err := systemDb.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname)).Error; err != nil {
				log.Printf("⚠️ Gagal membuat database otomatis: %v", err)
			} else {
				log.Printf("✅ Database '%s' berhasil dibuat", dbname)
			}
		}
		sqlDB, _ := systemDb.DB()
		sqlDB.Close()
	}

	// --- TAHAP 2: KONEKSI UTAMA KE DATABASE APLIKASI ---
	dsnApp := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsnApp), &gorm.Config{})
	if err != nil {
		// PERBAIKAN: Ditambahkan variabel 'dbname' agar sesuai dengan placeholder '%s'
		log.Fatalf("❌ Gagal menyambung ke database '%s': %v", dbname, err)
	}

	DB = db
	log.Printf("🚀 Berhasil terhubung ke database '%s'", dbname)
}
