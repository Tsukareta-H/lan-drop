package main

import (
	"fmt"
	"net/http"

	"github.com/Tsukareta-H/lan-drop/internal/server"
)

func main() {
	fmt.Println("LAN Drop started")
	ip := server.GetIPAddr()
	port := server.GetPort()

	targetURL := fmt.Sprintf("http://%s:%s", ip, port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("請在手機瀏覽器輸入這個網址: %s\n", targetURL)
	fmt.Println("伺服器啟動中... ")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("伺服器啟動失敗:", err)
	}
}
