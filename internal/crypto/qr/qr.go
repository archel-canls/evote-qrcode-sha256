package qr

import (
	"encoding/base64"

	"github.com/skip2/go-qrcode"
)

// GenerateQR menghasilkan QR Code PNG dalam format base64.
// Data yang dimasukkan SEBAIKNYA berupa hash (SHA256),
// bukan data mentah seperti voter_id atau NIM.
func GenerateQR(data string) (string, error) {
	return generate(data, 256)
}

// generate adalah helper internal agar ukuran QR bisa dikontrol
func generate(data string, size int) (string, error) {
	png, err := qrcode.Encode(data, qrcode.Medium, size)
	if err != nil {
		return "", err
	}

	// prefix wajib agar bisa langsung dipakai di <img src="">
	base64Img := "data:image/png;base64," +
		base64.StdEncoding.EncodeToString(png)

	return base64Img, nil
}
