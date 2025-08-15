package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", ":42069")

	if err != nil {
		log.Fatalf("Failed to start udp server %s", err)
	}

	server, err := net.DialUDP("udp", nil, udpAddr)

	defer func(server *net.UDPConn) {
		err := server.Close()
		if err != nil {
			if err != nil {
				log.Fatalf("Failed to close connection %s", err)
			}
		}
	}(server)

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println(">")

		readValue, err := reader.ReadString('\n')

		if err != nil {
			if err != nil {
				log.Fatalf("Failed to read value %s", err)
			}
		}

		_, err = server.Write([]byte(readValue))

		if err != nil {
			if err != nil {
				log.Fatalf("Failed to write to udp stream %s", err)
			}
		}
	}
}
