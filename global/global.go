package global

import "websocket/internal/model"

var (
	MessageQueueLen = 1024
	SendMsg         chan model.StreamMsg
)

func init() {
	SendMsg = make(chan model.StreamMsg, 10)
}
