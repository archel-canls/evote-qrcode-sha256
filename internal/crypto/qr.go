package crypto

import (
	"encoding/base64"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(nim string) (string, error) {

	png, err := qrcode.Encode(nim, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	// convert ke base64 biar bisa langsung ditampilkeun di HTML <img>
	base64Img := base64.StdEncoding.EncodeToString(png)

	return base64Img, nil
}
