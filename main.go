package main

import (
	"demo/protocol"
	"log"
	"net"
	"os"
)

const (
	Address = "127.0.0.1:6379"
	Network = "tcp"
)

func Conn(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		conn = nil
	}

	return conn, nil
}

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatalf("Os.args < 0")
	}

	redisConn, err := Conn(Network, Address)
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	defer redisConn.Close()

	reqCommand := protocol.GetRequest(args)
	_, err = redisConn.Write(reqCommand)
	if err != nil {
		log.Fatalf("Conn Write err: %v", err)
	}

	command := make([]byte, 1024)
	n, err := redisConn.Read(command)
	log.Printf("n=%v, cmd=%v", n, string(command[:n]))
	if err != nil {
		log.Fatalf("Conn Read err: %v", err)
	}

	reply, err := protocol.GetReply(command[:n])
	if err != nil {
		log.Fatalf("GetReply err: %v", err)
	}

	log.Printf("Reply=%v", reply)
}
