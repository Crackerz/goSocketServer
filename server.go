package goSocketServer

import(
	"fmt"
	"io/ioutil"
)

type SocketServer struct {
	sockets map[int]Socket
	uniq_id int
	program string
}

var Server SocketServer

func init() {
	Server.sockets = make(map[int]Socket)
	Server.uniq_id = 0
}

func SetProgram(filename string) {
	Server.SetProgram(filename)
}

func (s *SocketServer) SetProgram(filename string) {
	clientProgram,err:=ioutil.ReadFile(filename)
	if err!=nil {
		panic(err.Error())
	}
	(*s).program = string(clientProgram)
}

func (s *SocketServer) add(socket Socket) int {
	s.sockets[s.uniq_id] = socket
	s.uniq_id++
	printArray("Server", (*s).sockets)
	socket.WriteString("obj ="+ (*s).program)
	return Server.uniq_id-1
}

func (s *SocketServer) remove(index int) {
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
