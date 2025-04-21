package ws

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Jeecis/goapi/internal/repository"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub        *Hub
	conn       *websocket.Conn
	send       chan Message
	boardRepo  *repository.BoardRepository
	columnRepo *repository.ColumnRepository
	recordRepo *repository.RecordRepository
	mu         sync.Mutex
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: apply real origin check in production
	},
}

const pongWait = 60 * time.Second

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bye"))
		c.conn.Close()
	}()

	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("read error:", err)
			break
		}

		if handler, ok := ActionHandlers[msg.Action]; ok {
			handler(c, msg)
		} else {
			log.Printf("Unknown action: %s", msg.Action)
		}
	}
}

const pingPeriod = 54 * time.Second
const writeWait = 10 * time.Second

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bye"))
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteJSON(msg); err != nil {
				log.Println("write error:", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("ping error:", err)
				return
			}
		}
	}
}
