// socket-client project main.go
package main
import (
        //"fmt"
        "net"
        "os"
	"log"
)
const (
        //SERVER_HOST = "localhost"
        SERVER_HOST = "0.0.0.0"
        SERVER_PORT = "9988"
        SERVER_TYPE = "tcp"
	MESSAGE_1 = "ping"
	MESSAGE_2 = "pong"
)

func SocketClient() {
	conn, err := net.Dial(SERVER_TYPE,SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	
	defer conn.Close()

	conn.Write([]byte(MESSAGE_1))
	conn.Write([]byte(MESSAGE_2))
	log.Printf("Send: %s,%s", MESSAGE_1,MESSAGE_2)
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {

	SocketClient()

}
