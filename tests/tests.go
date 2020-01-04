package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/jneubaum/honestvote/tests/logger"
)

func main() {
	ListenConn()
}

func ListenConn() {
	listen, _ := net.Listen("tcp", ":8080")
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			logger.Println("listener.go", "ListenConn", err.Error())
		}
		go HandleConn(conn)
	}
}

type Data struct {
	Message  string
	Data     []byte
	Message2 string
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	// determines if it is an http request

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Bytes()
		fmt.Println(ln)
		if strings.Fields(string(ln))[0] == "GET" {
			http.GetRegistrationCode(conn, strings.Fields(string(ln))[1])
			return
		}
		var data Data

		json.Unmarshal(ln, &data)
		fmt.Println(data.Message)

	}
}
