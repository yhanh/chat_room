package ws

type Room struct {
	Id      string             `json:"id"`
	Name    string             `json:"name"`
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
			if _, ok := h.Rooms[cl.RoomId]; ok {
				r := h.Rooms[cl.RoomId]

				if _, ok := r.Clients[cl.Id]; !ok {
					r.Clients[cl.Id] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				if _, ok := h.Rooms[cl.RoomId].Clients[cl.Id]; ok {
					// Broadcast a message saying that the client has left the room
					if len(h.Rooms[cl.RoomId].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomId:   cl.RoomId,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomId].Clients, cl.Id)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomId]; ok {
				for _, cl := range h.Rooms[m.RoomId].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
