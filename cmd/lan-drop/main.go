package main

import (
	"fmt"

	"github.com/Tsukareta-H/lan-drop/internal/server"
)

func main() {
	fmt.Println("LAN Drop started")
	localIP := server.GetIPAddr()
	fmt.Println(localIP)
}
