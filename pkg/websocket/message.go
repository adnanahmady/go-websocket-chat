package websocket

type Message struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Sender  Sender `json:"sender"`
	Payload any    `json:"payload"`
}

type Sender struct {
	Name string `json:"name"`
}
