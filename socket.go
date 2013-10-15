package goSocketServer

import ("fmt"
	"io"
	"code.google.com/p/go.net/websocket"
)

/*
Socket is a simple structure that wraps the functionality of a websocket.Conn, allowing a websocket to be registered with a server in this library.
*/
type Socket struct {
	Connection *websocket.Conn
	id int
}

/*
Handle() is a blocking function that handles all communication with the websocket as defined by the callback functions provided to the server. When calling Handle() as a thread, take care to ensure the websocket is not closed by your application.
*/
func (s *Socket) Handle() {
	s.Register()
	io.Copy(s,s.Connection) //Blocking function to handle all communication
	s.Disconnect()
}

/*
GetID() returns the unique id of the websocket.
*/
func (s *Socket) GetId() int {
	return s.id
}

/*
Register() causes the default server in this library to start monitoring the Socket
*/
func (s *Socket) Register() {
	s.id = Server.add(*s)
}

/*
Disconnect() causes the default server in this library to stop monitoring the Socket
*/
func (s *Socket) Disconnect() {
	Server.remove(s.id)
}

/*
Write() is called when a message is received from the websocket. Write() DOES NOT send a message to the websocket. To send a message to the websocket use SendBytes() or SendString().
*/
func (s *Socket) Write(p []byte) (n int, err error) {
	if Server.onMessage != nil {
		Server.onMessage(s,p)
	}
	return len(p), nil
}

/*
SendString() sends a message to the websocket.
*/
func (s *Socket) SendString(message string) {
	fmt.Fprintf(s.Connection,message)
}

/*
SendBytes() sends a message to the websocket.
*/
func (s *Socket) SendBytes(message []byte) {
	fmt.Fprintf(s.Connection,"%s",message)
}

/*
NewSocket() creates a new Socket structure that wraps the websocket. The unique ID will be 0 and is assigned when the Socket is registered with a server.
*/
func NewSocket(ws *websocket.Conn) Socket {
	return Socket{ws,0}
}
