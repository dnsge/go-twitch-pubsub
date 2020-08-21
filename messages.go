package twitchpubsub

type BaseMessage struct {
	Type string `json:"type"`
}

type ListenData struct {
	Topics    []string `json:"topics"`
	AuthToken string   `json:"auth_token,omitempty"`
}

type RequestMessage struct {
	BaseMessage
	Nonce string     `json:"nonce,omitempty"`
	Data  ListenData `json:"data"`
}

type ResponseMessage struct {
	BaseMessage
	Nonce string `json:"nonce,omitempty"`
	Error string `json:"error"`
}

type MessageData struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

type MessageMessage struct {
	BaseMessage
	Data MessageData `json:"data"`
}
