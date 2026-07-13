package network

import (
	"context"
	"net"
	"time"
)

func GetIPAddr() string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localIPAddr := conn.LocalAddr().(*net.UDPAddr).IP

	return localIPAddr.String()
}
