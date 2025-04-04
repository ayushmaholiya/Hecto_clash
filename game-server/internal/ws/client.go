package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
}

type MessageType string

const (
	MESSAGE_TYPE_JOIN   MessageType = "join"
	MESSAGE_TYPE_LEAVE  MessageType = "leave"
	MESSAGE_TYPE_SUBMIT MessageType = "submit"
	MESSAGE_TYPE_END    MessageType = "end"
	MESSAGE_TYPE_ERROR  MessageType = "error"
)

type Message struct {
	Type      MessageType `json:"type"`
	Content   string      `json:"content"`
	RoomID    string      `json:"roomId"`
	SenderID  string      `json:"senderId"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(m, &msg); err != nil {
			log.Printf("error: %v", err)
			continue
		}

		switch msg.Type {
		case MESSAGE_TYPE_SUBMIT:
			hub.handleSubmission(c, msg)
		default:
			hub.Broadcast <- &msg
		}
	}
}