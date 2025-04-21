package ws

import (
	"encoding/json"
	"log"

	"github.com/Jeecis/goapi/internal/models"
	service "github.com/Jeecis/goapi/internal/services"
)

// Action constants
const (
	ActionJoinBoard    = "join_board"
	ActionCreateRecord = "create_record"
	ActionUpdateRecord = "update_record"
	ActionDeleteRecord = "delete_record"
	ActionMoveRecord   = "move_record"
)

// Message represents any message sent via WebSocket
type Message struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

// HandlerMap maps actions to handling functions
var ActionHandlers = map[string]func(*Client, Message){
	ActionMoveRecord:   handleMoveRecord,
	ActionJoinBoard:    handleJoinBoard,
	ActionCreateRecord: handleCreateRecord,
	ActionUpdateRecord: handleUpdateRecord,
	ActionDeleteRecord: handleDeleteRecord,
}

func handleJoinBoard(c *Client, msg Message) {
	c.mu.Lock()         // Lock the mutex
	defer c.mu.Unlock() // Unlock the mutex when done
	var payload joinBoard

	payloadBytes, err := json.Marshal(msg.Payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return
	}

	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Println("unmarshal error:", err)
		return
	}

	fullBoard, err := service.GetBoard(c.boardRepo, c.columnRepo, c.recordRepo, payload.BoardID)
	if err != nil {
		log.Println("Repository error:", err)
		return
	}

	c.send <- Message{
		Action:  "join_board",
		Payload: fullBoard,
	}
}

func handleCreateRecord(c *Client, msg Message) {
	c.mu.Lock()         // Lock the mutex
	defer c.mu.Unlock() // Unlock the mutex when done
	var payload models.CreateRecord

	payloadBytes, err := json.Marshal(msg.Payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return
	}

	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Println("unmarshal error:", err)
		return
	}

	record, err := service.CreateRecord(c.boardRepo, c.columnRepo, c.recordRepo, payload)
	if err != nil {
		log.Println("Repository error:", err)
		return
	}

	c.hub.broadcast <- Message{
		Action:  "create_record",
		Payload: record,
	}

}

func handleUpdateRecord(c *Client, msg Message) {
	c.mu.Lock()         // Lock the mutex
	defer c.mu.Unlock() // Unlock the mutex when done
	var payload models.Record

	payloadBytes, err := json.Marshal(msg.Payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return
	}

	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Println("unmarshal error:", err)
		return
	}

	record, err := service.UpdateRecord(c.boardRepo, c.columnRepo, c.recordRepo, payload)
	if err != nil {
		log.Println("Repository error:", err)
		return
	}

	c.hub.broadcast <- Message{
		Action:  "update_record",
		Payload: record,
	}
}

func handleDeleteRecord(c *Client, msg Message) {
	c.mu.Lock()         // Lock the mutex
	defer c.mu.Unlock() // Unlock the mutex when done
	var payload models.DeleteRecord

	payloadBytes, err := json.Marshal(msg.Payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return
	}

	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Println("unmarshal error:", err)
		return
	}

	err = service.DeleteRecord(c.boardRepo, c.columnRepo, c.recordRepo, payload.RecordID)
	if err != nil {
		log.Println("Repository error:", err)
		return
	}

	c.hub.broadcast <- Message{
		Action:  "delete_record",
		Payload: payload,
	}
}

func handleMoveRecord(c *Client, msg Message) {
	c.mu.Lock()         // Lock the mutex
	defer c.mu.Unlock() // Unlock the mutex when done
	var payload models.MoveRecord

	payloadBytes, err := json.Marshal(msg.Payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return
	}

	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		log.Println("unmarshal error:", err)
		return
	}

	_, err = service.MoveRecord(c.recordRepo, payload)
	if err != nil {
		log.Println("Repository error:", err)
		return
	}

	c.hub.broadcast <- Message{
		Action:  "move_record",
		Payload: payload,
	}
}
