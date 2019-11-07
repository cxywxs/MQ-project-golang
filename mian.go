package main

import (
	"example.com/m/v2/clientpkg"
	"example.com/m/v2/serverpkg"
	"time"
)

func main() {
	go serverpkg.ServerConn()
	go clientpkg.ClientConn()
	for {
		time.Sleep(5 * time.Second)
	}
}
