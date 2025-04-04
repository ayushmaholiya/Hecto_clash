package ws

type Room struct {
	ID      string             `json:"id"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
    for {
        select {
        case cl := <-h.Register:
            // Check if the room exists
            if room, ok := h.Rooms[cl.RoomID]; ok {
                // Check if the room already has 2 clients
                if len(room.Clients) >= 2 {
                    // Notify the client that the room is full
                    cl.Message <- &Message{
                        Type:    MESSAGE_TYPE_ERROR,
                        Content: "Room is full",
                        RoomID:  cl.RoomID,
                    }
                    close(cl.Message)
                    continue
                }

                // Add the client to the room
                room.Clients[cl.ID] = cl
                cl.Message <- &Message{
                    Type:     MESSAGE_TYPE_JOIN,
                    Content:  "User joined the room",
                    RoomID:   cl.RoomID,
                    SenderID: cl.ID,
                }
            } else {
                // If the room doesn't exist, create it and add the client
                h.Rooms[cl.RoomID] = &Room{
                    ID:      cl.RoomID,
                    Clients: map[string]*Client{cl.ID: cl},
                }
                cl.Message <- &Message{
                    Type:     MESSAGE_TYPE_JOIN,
                    Content:  "User joined the room",
                    RoomID:   cl.RoomID,
                    SenderID: cl.ID,
                }
            }

        case cl := <-h.Unregister:
            // Handle client unregistration
            if room, ok := h.Rooms[cl.RoomID]; ok {
                if _, ok := room.Clients[cl.ID]; ok {
                    delete(room.Clients, cl.ID)
                    close(cl.Message)

                    // If the room is empty, delete it
                    if len(room.Clients) == 0 {
                        delete(h.Rooms, cl.RoomID)
                    } else {
                        // Notify other clients in the room
                        h.Broadcast <- &Message{
                            Type:     MESSAGE_TYPE_LEAVE,
                            Content:  "User left the room",
                            RoomID:   cl.RoomID,
                            SenderID: cl.ID,
                        }
                    }
                }
            }

        case m := <-h.Broadcast:
            // Broadcast messages to all clients in the room except the sender
            if room, ok := h.Rooms[m.RoomID]; ok {
                for _, cl := range room.Clients {
                    if cl.ID == m.SenderID {
                        continue
                    }
                    cl.Message <- m
                }
            }
        }
    }
}

