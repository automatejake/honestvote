package http

import (
	"fmt"
	"net"
	"strings"
)

func GetRegistrationCode(conn net.Conn, message string) {

	fmt.Println(message)

}

func router(conn net.Conn, ln string) {
	fmt.Println("blah")
	// request line
	m := strings.Fields(ln)[0] // method
	// u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	// fmt/.Println("***URI", u)

	// // multiplexer
	// if m == "GET" && u == "/" {
	// 	index(conn)
	// }
	// if m == "GET" && u == "/about" {
	// 	about(conn)
	// }
	// if m == "GET" && u == "/contact" {
	// 	contact(conn)
	// }
	// if m == "GET" && u == "/apply" {
	// 	apply(conn)
	// }
	// if m == "POST" && u == "/apply" {
	// 	applyProcess(conn)
	// }
}
