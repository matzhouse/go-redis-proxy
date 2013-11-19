package proxyer

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func Forward(local net.Conn, remoteAddr string) (fan chan int) {

	timeout := time.Second
	fmt.Println("Connecting to ", remoteAddr)

	remote, err := net.DialTimeout("tcp", remoteAddr, timeout)

	var buffer []byte

	if Canread(remote, buffer) == false {
		log.Fatal("Redis server not available")
		Closer(remote)
	}

	defer Closer(remote)

	if err != nil {
		fmt.Println(err)
		Closer(local)
	}

	fan = make(chan int)

	if remote == nil {
		fmt.Fprintf(os.Stderr, "remote dial failed: %v\n", err)
		return
	}
	go io.Copy(local, remote)
	go io.Copy(remote, local)

	return fan

}

func Closer(conn net.Conn) {

	conn.Close()

}

func Canread(c net.Conn, buffer []byte) bool {

	var ping []byte

	ping = []byte("PING")

	c.Write(ping)

	bytesRead, err := c.Read(buffer)
	if err != nil {
		c.Close()
		log.Println(err)
		return false
	}
	log.Println("Read ", bytesRead, " bytes")
	return true
}
