package goSocketServer

import(
	"fmt"
)

type SocketServer struct {
	sockets map[int]Socket
	uniq_id int
	onConnect func(*Socket)
	onDisconnect func()
}

var Server SocketServer

func init() {
	Server.sockets = make(map[int]Socket)
	Server.uniq_id = 0
}

func (s *SocketServer) add(socket Socket) int {
	s.sockets[s.uniq_id] = socket
	s.uniq_id++
	s.onConnect(&socket)
	return Server.uniq_id-1
}

func (s *SocketServer) remove(index int) {
	(*s).onDisconnect()
	delete(s.sockets,index)
}

func printArray(name string, array map[int]Socket) {
	fmt.Println(name, " len:", len(array), " ", array)
}

func (s *SocketServer) WriteAll(message string) {
	for _,socket := range s.sockets {
		socket.WriteString(message)
	}
}

func (s *SocketServer) OnConnect(function func(*Socket)) {
	s.onConnect = function
}

func (s *SocketServer) Test() {
	fmt.Println(s.onConnect)
}

func (s *SocketServer) OnDisconnect(function func()) {
	s.onDisconnect = function
}
