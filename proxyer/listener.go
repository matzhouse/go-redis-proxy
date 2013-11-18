package proxyer

import (
	"net"
)

func Listen(localAddr string) (local net.Listener) {

	local, err := net.Listen("tcp", localAddr)
	if local == nil {
		Fatal("cannot listen: %v", err)
	}

	return

}
