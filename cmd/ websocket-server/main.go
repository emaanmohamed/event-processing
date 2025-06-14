package main

import (
	"fmt"
	"github.com/emaanmohamed/event-processing/configs"
	"github.com/emaanmohamed/event-processing/internal/kafka"
	"github.com/emaanmohamed/event-processing/internal/websocket"
	"log"
	"net/http"
)

func main() {

	configs.Load()

	kafka.InitProducer(configs.KafkaBroker)

	http.HandleFunc("/ws", websocket.HandleWebSocket)

	log.Println(fmt.Sprintf("WebSocket server started on %s", configs.ServerPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", configs.ServerPort), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
