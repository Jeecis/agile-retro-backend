package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeecis/goapi/internal/repository"
	"github.com/gin-gonic/gin"
)

var boardHubs = make(map[string]*Hub) // boardID -> Hub
var hubsMu sync.RWMutex

// ServeWsGin is a Gin-compatible WebSocket handler
func JoinBoard(boardRepo *repository.BoardRepository, columnRepo *repository.ColumnRepository, recordRepo *repository.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		boardID := c.Param("id")

		exists := boardRepo.BoardExists(boardID)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Board not found"})
			return
		}

		hubsMu.Lock()
		hub, exists := boardHubs[boardID]
		if !exists {
			hub = NewHub()
			boardHubs[boardID] = hub
			go hub.Run()
		}
		hubsMu.Unlock()

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}

		client := &Client{
			hub:        hub,
			conn:       conn,
			send:       make(chan Message),
			boardRepo:  boardRepo,
			columnRepo: columnRepo,
			recordRepo: recordRepo,
			mu:         sync.Mutex{},
		}

		hub.register <- client

		go client.writePump()
		go client.readPump()

	}
}
