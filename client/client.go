package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
	"websocket/global"
	"websocket/internal/model"
	"websocket/storage"
)

func AggTrade(storage storage.Service) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	ws := "wss://stream.yshyqxx.com/stream?streams=btcusdt@aggTrade"
	fmt.Printf("websocket: %#v\n", ws)
	conn, _, err := websocket.Dial(ctx, ws, nil)
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "Internal error!")

	ctx = context.Background()
	for {
		var message model.StreamMsg
		err = wsjson.Read(ctx, conn, &message)
		if err != nil {
			log.Println("receive msg error:", err)
			continue
		}
		val, _ := json.Marshal(message)
		if err := storage.Save(val); err != nil {
			log.Println("client.Set failed", err)
		}
		global.SendMsg <- message
	}

	conn.Close(websocket.StatusNormalClosure, "")
}
