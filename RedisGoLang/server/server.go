package server

import (
	"RedisGoLang/config"
	"io"
	"log"
	"net"
	"strconv"
)

func clientRead(conn net.Conn) (string, error) {
	// Capped at 512bytes
	var buffer = make([]byte, 512)
	n, err := conn.Read(buffer[:])
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func clientWrite(data string, conn net.Conn) error {
	if _, err := conn.Write([]byte(data)); err != nil {
		return err
	}
	return nil
}

func RunSyncTCPServer(serverConfig config.ServerConfig) {
	log.Println("Starting the RedisGoLang Server")

	lstnr, err := net.Listen("tcp", serverConfig.ServerHost+":"+strconv.Itoa(serverConfig.ServerPort))
	if err != nil {
		panic(err)
	}

	var clientConnections = 0

	for {
		// Blocking call waiting to connect
		connection, err := lstnr.Accept()
		if err != nil {
			panic(err)
		}

		clientConnections += 1
		log.Println("Client connected. Address : ", connection.RemoteAddr(),
			". Total number of Clients:", clientConnections)

		for {
			// send data over the socket
			command, err := clientRead(connection)
			if err != nil {
				connection.Close()
				clientConnections -= 1
				log.Println("Client Disconnected. Address: ", connection.RemoteAddr(),
					". Number of Clients: ", clientConnections)
				if err == io.EOF {
					// graceful termination
					break
				}
				log.Println("err", err)
			}

			log.Println("Command from Client:", connection.RemoteAddr(), " is: ", command)
			if err = clientWrite(command, connection); err != nil {
				log.Println(err)
			}
		}

	}

}
