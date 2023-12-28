package main

import (
	"log"
	"net"
	"time"
)

func processConn(conn net.Conn) {
	buf := make([]byte, 1024) //max size of request is 1mb
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Request Processing")

	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\n Hello Terminal!\r\n"))
	log.Println("Request DONE")

	conn.Close()
}
func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	//to Accept all the time not to close after single request keep it in infinite loop
	for {
		log.Println("Accepting Request")

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		//adding go routine to handle concurrent requests (threads specifically)
		go processConn(conn)
	}
}
