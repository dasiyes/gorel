package router

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

type WsJsonResponse struct {
	Action        string   `json:"action"`
	Message       string   `json:"nessage"`
	MessageType   string   `json:"message_type"`
	ConnectedUser []string `json:"connected_users"`
}

var wsChan = make(chan WsPayload)
var clients = make(map[WebSocketConnection]string)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//Allow all connections by default
		return true
	},
}

// wsEndpoint provides connection point to the server's websockt handlers for the clients
func (h *routesHandler) wsEndpoint(w http.ResponseWriter, r *http.Request) {

	// upgrade the http connection to a webSocket connection
	respHeader := http.Header{}

	ws, err := upgradeConnection.Upgrade(w, r, respHeader)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	log.Println("The client connected to the server ws endpoint")

	var response WsJsonResponse
	response.Message = `<em><small>Connected to the server</small></em>`

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn)
}

// LsitenForWs is a function cycling in  a go routine to listen for any payload sent to the server over websocket connection
func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Crash ", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.Conn.ReadJSON(&payload)
		if err != nil {
			// do nothing for now
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

// ListenToWsChannel is a function to recieve the websocket payload that comes from the websocket connection listener `ListenForWs` and decide on its transformation
func ListenToWsChannel() {

	var response WsJsonResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "username":
			//get a list of all users and send it back via the broadcastToAll
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUser = users
			broadcastToAll(response)
		}
	}
}

func getUserList() []string {

	var userList []string

	for _, x := range clients {
		userList = append(userList, x)
	}
	sort.Strings(userList)
	return userList
}

func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}
