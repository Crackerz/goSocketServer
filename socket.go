package goSocketServer

import ("fmt"
	"io"
	"code.google.com/p/go.net/websocket"
)

type Socket struct {
	Connection *websocket.Conn
	Id int
}

func (s *Socket) Handle() {
	(*s).Register()
	io.Copy(*s,(*s).Connection) //Blocking function to handle all communication
	(*s).Disconnect()
}

//Register the socket with the server
func (s *Socket) Register() {
	(*s).Id = Server.add(*s)
}

//Cleanup server after loosing connection with socket
func (s *Socket) Disconnect() {
	Server.remove((*s).Id)
}

func (s Socket) Write(p []byte) (n int, err error) {
	fmt.Printf("Write Called\n")
	Server.WriteAll(string(p))
	fmt.Printf("Write Completed\n")
	return len(p), nil
}

func (s *Socket) WriteString(message string) {
	fmt.Fprintf((*s).Connection,message)
}
