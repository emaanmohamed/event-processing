package websocket

import (
	"fmt"
	"github.com/emaanmohamed/event-processing/internal/kafka"
	"github.com/emaanmohamed/event-processing/utils"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("Failed to upgrade connection"))
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("Failed to read message: %v", err))
			return
		}

		go func(message []byte) {
			if err := kafka.SendMessage("events_topic", message); err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("Failed to send message to Kafka: %v", err))
				return
			}
		}(msg)

	}

}
