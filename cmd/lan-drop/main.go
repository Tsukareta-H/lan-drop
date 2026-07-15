package main

import (
	"fmt"
	"log"

	"github.com/Tsukareta-H/lan-drop/internal/network"
	"github.com/Tsukareta-H/lan-drop/internal/server"
)

func main() {
	fmt.Println("LAN Drop started")
	ip := network.GetIPAddr()
	port := server.GetPort()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	err := server.StartListen(listenAddr)
	if err != nil {
		log.Fatalln(err)
	}
}
