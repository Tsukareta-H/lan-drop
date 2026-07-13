package server

import (
	"fmt"
	"net/http"

	"github.com/Tsukareta-H/lan-drop/internal/qr"
)

func GetPort() string {
	port := "8080"

	return port
}

func StartListen(ipAddr string, port string) error {
	listenAddr := fmt.Sprintf("%s:%s", ipAddr, port)
	targetURL := fmt.Sprintf("http://%s", listenAddr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
	})
	fmt.Println("服務啓動中...")
	fmt.Printf("請輸入 URL: %s\n或掃描下方 QR Code\n", targetURL)

	qrString, err := qr.QRCode(targetURL)
	if err != nil {
		return fmt.Errorf("生成 QR Code 失敗: %w", err)
	}
	fmt.Println(qrString)

	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		return fmt.Errorf("服務啓動失敗: %w", err)
	}
	return nil
}
