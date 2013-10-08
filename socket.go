package goSocketServer

import ("fmt"
	"io"
	"code.google.com/p/go.net/websocket"
)

type Socket struct {
	Connection *websocket.Conn
	id int
}

func (s *Socket) Handle() {
	s.Register()
	io.Copy(s,s.Connection) //Blocking function to handle all communication
	s.Disconnect()
}

func (s *Socket) GetId() int {
	return s.id
}

//Register the socket with the server
func (s *Socket) Register() {
	s.id = Server.add(*s)
}

//Cleanup server after loosing connection with socket
func (s *Socket) Disconnect() {
	Server.remove(s.id)
}

func (s *Socket) Write(p []byte) (n int, err error) {
	if Server.onMessage != nil {
		Server.onMessage(s,p)
	}
	return len(p), nil
}

func (s *Socket) SendString(message string) {
	fmt.Fprintf(s.Connection,message)
}

func (s *Socket) SendBytes(message []byte) {
	fmt.Fprintf(s.Connection,"%s",message)
}

func NewSocket(ws *websocket.Conn) Socket {
	return Socket{ws,0}
}
