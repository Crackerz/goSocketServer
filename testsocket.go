package main

import ("fmt"
	"io"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

type SocketHandler struct {
	connection *websocket.Conn
}

func (s SocketHandler) Write(p []byte) (n int, err error) {
	fmt.Fprintf(s.connection,s.getResp(string(p)))
	return len(p), nil
}

func (s SocketHandler) getResp(req string) string {
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

//Establish all http listeners
func init() {
	fmt.Printf("Configuring Server...\n")
	http.HandleFunc("/",website)
	http.Handle("/socket",websocket.Handler(socket))
}

func main() {
	fmt.Printf("Starting Server...\n")
	http.ListenAndServe(":8080",nil)
}

func website(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "website")
}

func socket(ws *websocket.Conn) {
	fmt.Printf("Received Socket Connection...\n")
	sh:=SocketHandler{ws}
	io.Copy(sh,ws)
	fmt.Printf("Handled Connection")
}
