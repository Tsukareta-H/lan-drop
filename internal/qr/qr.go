package qr

import (
	"github.com/skip2/go-qrcode"
)

func QRCode(targetURL string) (string, error) {
	qrc, err := qrcode.New(targetURL, qrcode.Medium)
	if err != nil {
		return "", err
	}

	return qrc.ToString(false), nil

}
