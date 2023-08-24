package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server started, listening on localhost:8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
		}

		data := buffer[:n]
		fmt.Println("Received data:", string(data))

		// Отправка ответа клиенту
		response := []byte("response from server: get data: " + string(data[:3]) + "...\n")
		_, err = conn.Write(response)
		if err != nil {
			fmt.Println("Error writing response:", err)
			break
		}
	}
}
