package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var (
	addr string
	port string
)

func main() {
	// Define the port to listen on
	flag.StringVar(&addr, "address", "0.0.0.0", "Host address to listen on")
	flag.StringVar(&port, "port", "4444", "The port to listen on")

	flag.Parse()

	// Listen for incoming TCP connections
	listen := addr + ":" + port
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("TCP echo server is up and listening %s...\n\n", listen)

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		currentTime := time.Now()
		fmt.Printf("\n%s: New connection from %s was established\n", currentTime.Format(time.RFC1123Z), conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from the connection
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading data:", err)
			}
			break
		}

		// Print received data to stdout
		fmt.Println(string(buf[:n]))
	}
}
