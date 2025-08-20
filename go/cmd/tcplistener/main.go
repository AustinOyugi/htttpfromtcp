package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	server, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatalf("Error starting server! %s", err)
	}

	log.Println("Server started successfully at port tcp:42069")

	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			log.Fatalf("Failed to close connection %s", err)
		}
	}(server)

	for {
		connection, err := server.Accept()

		if err != nil {
			log.Fatalf("Server failed to accept connections %s", err)
		}

		log.Println("Connection was accepted successfully")

		lineChannel := getLinesChannel(connection)

		for msg := range lineChannel {
			fmt.Println(msg)
		}
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {

	lineChannel := make(chan string)

	go func() {

		var lineContents []string

		for {
			fileContents := make([]byte, 8)

			_, err := file.Read(fileContents)

			if err == io.EOF {
				lineValue := strings.Join(lineContents, "")
				lineChannel <- lineValue
				close(lineChannel)
				break
			}

			parts := strings.Split(string(fileContents), "\n")
			lineContents = append(lineContents, parts[0])

			if len(parts) == 1 {
				continue
			}

			lineValue := strings.Join(lineContents, "")
			lineChannel <- lineValue
			lineContents = []string{parts[1]}
		}
	}()

	return lineChannel
}
