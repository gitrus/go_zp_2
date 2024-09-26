package main

import (
	"go_zp_2/pkg/multiplyr"
	"net"
	"net/rpc"
)

func main() {

	multiplyService := multiplyr.NewMultiplyService()

	err := rpc.Register(multiplyService)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1121") // OSI
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// Accept incoming connections and serve requests.
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
