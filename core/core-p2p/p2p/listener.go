package p2p

import (
	"log"
	"net"
)

<<<<<<< HEAD
func ListenConn(args string) {
	portString := ":" + args
	listen, err := net.Listen("tcp", portString)

=======
func ListenConn(port string) {
	listen, err := net.Listen("tcp", port)
>>>>>>> 7d7e46143029a19e3b5a657fca370736af2ead58
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}
}
