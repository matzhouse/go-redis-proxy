package main

import (
	"fmt"
	"github.com/matzhouse/go-redis-proxy/proxyer"
	"time"
)

type Server struct {
	Address   string
	Timeout   time.Duration
	Order     int
	ConnLimit int
}

func main() {

	sv1 := &Server{"127.0.0.1:6380", time.Second, 1, 10}
	sv2 := &Server{"127.0.0.1:6379", time.Second, 2, 10}

	fmt.Println("Redis proxy starting up.. \n")

	localAddr := sv1.Address
	remoteAddr := sv2.Address

	local := proxyer.Listen(localAddr)

	Pool := proxyer.Pooldigger(sv1.ConnLimit)

	for {
		conn, err := local.Accept()

		<-Pool
		fmt.Printf("Woohoo! I found a connection! \n")

		if conn == nil {
			proxyer.Fatal("accept failed: %v", err)
		}

		fmt.Println("Client connected: ", conn.RemoteAddr().String(), " \n")

		go proxyer.Forward(conn, remoteAddr)
	}
}
