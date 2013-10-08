package goSocketServer

type SocketServer struct {
	sockets map[int]*Socket
	uniq_id int
	onConnect func(*Socket)
	onDisconnect func(*Socket)
	onMessage func(*Socket,[]byte)
}

var Server SocketServer

func init() {
	Server.sockets = make(map[int]*Socket)
	Server.uniq_id = 0
}

func (s *SocketServer) add(socket Socket) int {
	s.sockets[s.uniq_id] = &socket
	socket.id = s.uniq_id
	s.uniq_id++
	if s.onConnect != nil {
		s.onConnect(&socket)
	}
	return Server.uniq_id-1
}

func (s *SocketServer) remove(index int) {
	socket:=s.sockets[index]
	if s.onDisconnect!=nil {
		s.onDisconnect(socket)
	}
	delete(s.sockets,index)
}

func (s *SocketServer) WriteAll(message string) {
	for _,socket := range s.sockets {
		socket.SendString(message)
	}
}

func (s *SocketServer) OnConnect(function func(*Socket)) {
	s.onConnect = function
}

func (s *SocketServer) OnDisconnect(function func(*Socket)) {
	s.onDisconnect = function
}

func (s *SocketServer) OnMessage(function func(*Socket, []byte)) {
	s.onMessage = function
}
