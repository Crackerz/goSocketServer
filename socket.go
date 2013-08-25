package main

import ("fmt"
	"io"
	"strconv"
	"code.google.com/p/go.net/websocket"
)

type Socket struct {
	Connection *websocket.Conn
	id int
}

func (s Socket) Handle() {
	//Register socket with server
	s.id = Server.add(s)
	io.Copy(s,s.Connection)
	Server.remove(s.id)
}

func (s Socket) Write(p []byte) (n int, err error) {
	fmt.Printf("Write Called\n")
	//fmt.Fprintf(s.Connection,s.getResp(string(p)))
	Server.WriteAll("Anon "+strconv.Itoa(s.id)+":"+string(p))
	fmt.Printf("Write Completed\n")
	return len(p), nil
}

func (s Socket) WriteString(message string) {
	fmt.Fprintf(s.Connection,message)
}

func (s Socket) getResp(req string) string {
	fmt.Printf("getResp Called\n")
	switch req {
	case "Name":
		return "GoLang"
	case "Author":
		return "William Blankenship"
	case "Univ.":
		return "SIUC"
	}
	return "Unknown Query"
}
