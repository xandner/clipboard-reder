package server

import (
	"clip/logger"
	"clip/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type server struct {
	l     logger.Logger
	u     usecase.Clipboard
	conns map[*websocket.Conn]bool
}
type Server interface {
	Main()
}

func NewServer(l logger.Logger, u usecase.Clipboard) Server {
	return &server{
		l,
		u,
		make(map[*websocket.Conn]bool),
	}
}

func (s *server) Main() {
	http.HandleFunc("/clipboard", s.websocketHandler)
	// http.HandleFunc("/search", s.websocketHandler)
	fmt.Println("server started")
	http.ListenAndServe(":9999", nil)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s *server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	log.Println("connected")
	err = conn.WriteMessage(websocket.TextMessage, []byte(s.getLastClipboardData()))
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (s *server) getLastClipboardData() string {
	err, data := s.u.GetLast10()
	if err != nil {
		s.l.Error(fmt.Sprintf("Error while getting last 10 data %v", err))
	}
	// buf := new(bytes.Buffer)
	// err = binary.Write(buf, binary.LittleEndian, data)
	// if err != nil {
	// 	log.Println(err)
	// 	s.l.Error(fmt.Sprintf("Error while marshalling data %v", err))
	// }
	// return buf.Bytes()
	stringData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		s.l.Error(fmt.Sprintf("Error while marshalling data %v", err))
	}
	return string(stringData)
}
