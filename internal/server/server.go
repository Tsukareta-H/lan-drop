package server

import (
	"fmt"
	"net/http"

	"github.com/Tsukareta-H/lan-drop/internal/file"
	"github.com/Tsukareta-H/lan-drop/internal/qr"
)

func GetPort() string {
	port := "8080"

	return port
}

func StartListen(listenAddr string) error {
	targetURL := fmt.Sprintf("http://%s", listenAddr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="zh-TW">
<head>
<meta charset="UTF-8">
<title>LAN Drop</title>
</head>

<body>

<h1>LAN Drop</h1>

<h2>上傳</h2>

<form action="/upload/" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <button>上傳</button>
</form>

<hr>

<h2>下載</h2>

<a href="/download/">查看共享檔案</a>

</body>
</html>`)
	})
	fmt.Println("服務啓動中...")
	http.Handle("/download/", http.StripPrefix("/download", file.FileDownload("./downloads")))
	http.HandleFunc("/upload/", file.FileUpload)

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
