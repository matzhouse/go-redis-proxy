package proxyer

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func Forward(local net.Conn, remoteAddr string) (fan chan int) {

	timeout := time.Second
	remote, err := net.DialTimeout("tcp", remoteAddr, timeout)

	fan = make(chan int)

	defer Closer(remote)

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
