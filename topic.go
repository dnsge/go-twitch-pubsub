package twitchpubsub

import "fmt"

type TopicCallback func(MessageData)

type Topic struct {
	Name      string
	Nonce     string
	AuthToken string
	Callback  TopicCallback
}

func (t *Topic) ListenMessage() RequestMessage {
	return RequestMessage{
		BaseMessage: BaseMessage{
			Type: "LISTEN",
		},
		Nonce: t.Nonce,
		Data: ListenData{
			Topics:    []string{t.Name},
			AuthToken: t.AuthToken,
		},
	}
}

func (t *Topic) UnlistenMessage() RequestMessage {
	return RequestMessage{
		BaseMessage: BaseMessage{
			Type: "UNLISTEN",
		},
		Nonce: t.Nonce,
		Data: ListenData{
			Topics: []string{t.Name},
		},
	}
}

func (t *Topic) Identifier() string {
	return mustHashString(fmt.Sprintf("%s:%s", t.Name, t.AuthToken))
}
