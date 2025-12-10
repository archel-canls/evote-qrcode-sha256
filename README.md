# E-Voting QR Code + Hash SHA-256

Sistem e-voting berbasis QR Code, Golang Gin, PostgreSQL, dan SHA-256 untuk memastikan integritas suara.

## 🚀 Fitur Utama

- Registrasi pemilih + QR Code (base64)
- Voting menggunakan NIM dan CandidateID
- Hash SHA-256 untuk verifikasi suara
- Anti manipulasi data (hash verifiable)
- QR Scanner berbasis HTML5
- API backend Golang (Gin)
- PostgreSQL database
- Docker ready

---

## 📂 Struktur Proyek

evote-qrcode-sha256/
- cmd/server/main.go  
- internal/config  
- internal/controllers  
- internal/crypto  
- internal/database  
- internal/models  
- internal/routes  
- frontend  
- migrations  
- Dockerfile  
- docker-compose.yml  

---

## ▶️ Menjalankan Project

### 1. Copy Env
