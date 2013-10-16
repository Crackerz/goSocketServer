/*
goSocketServer is a simple wrapper for the go.net/websocket library
*/
package goSocketServer

/*
SocketServer manages connecting, disconnecting, and handling messages from all websockets registered with it.
*/
type SocketServer struct {
	sockets      map[int]*Socket //a map of all Sockets the server is monitoring w/ their IDs as keys.
	uniq_id      int             //Used for generating unique IDs
	onConnect    func(*Socket)
	onDisconnect func(*Socket)
	onMessage    func(*Socket, []byte)
}

/*
Server is the default SocketServer for this library. Unless you know what you are doing, it is recommended you use this default server for all websockets in your application.
*/
var Server SocketServer

func init() {
	Server.sockets = make(map[int]*Socket)
	Server.uniq_id = 0
}

/*
add() causes the server to start monitoring a websocket wrapped in this library's Socket struct.
*/
func (s *SocketServer) add(socket Socket) int {
	s.sockets[s.uniq_id] = &socket
	socket.id = s.uniq_id
	s.uniq_id++
	if s.onConnect != nil {
		s.onConnect(&socket)
	}
	return Server.uniq_id - 1
}

/*
remove() stops the server from monitoring a websocket with the specified.
*/
func (s *SocketServer) remove(id int) {
	socket := s.sockets[id]
	if s.onDisconnect != nil {
		s.onDisconnect(socket)
	}
	delete(s.sockets, id)
}

/*
WriteAll() sends the string "message" to all sockets currently connected to the server.
*/
func (s *SocketServer) WriteAll(message string) {
	for _, socket := range s.sockets {
		socket.SendString(message)
	}
}

/*
OnConnect() sets a callback function that will be called whenever a new websocket connects to the server. The connecting websocket will be wrapped in this library's Socket struct and passed to the callback function as a parameter.
*/
func (s *SocketServer) OnConnect(function func(*Socket)) {
	s.onConnect = function
}

/*
OnDisconnect() sets a callback function that will be called whenever a websocket disconnects from the server. The Socket struct mapped to this websocket will be passed to the callback function as a parameter.
*/
func (s *SocketServer) OnDisconnect(function func(*Socket)) {
	s.onDisconnect = function
}

/*
OnMessage() sets a callback function that will be called whenever the server recieves a message from the websocket. The Socket struct mapped to the websocket that sent a message and the []byte array representing the message will be passed to the callback function as parameters.
*/
func (s *SocketServer) OnMessage(function func(*Socket, []byte)) {
	s.onMessage = function
}
