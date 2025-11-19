// models/message.go
package models

import (
	"database/sql"
	"fmt"

	"forum/db"
	"forum/utils"
)

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
	CreatedAt  int    `json:"created_at"`
	IsRead     bool   `json:"is_read"`
}

type WSMessage struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Message `json:"message"`
	Members []Members `json:"members"`
	Typing  bool      `json:"is_typing"`
}

// GetMessageHistory retrieves a paginated list of messages between sender
// and receiver that were created before the provided `time` (unix seconds).
// It returns at most 10 messages ordered by creation time descending.
func GetMessageHistory(sender, receiver, time int) (messages []Message, err error) {
	rows, err := db.DB.Query(`
	    SELECT * FROM messages
	    WHERE created_at < ? AND ((sender_id = ? AND receiver_id = ?)
	    OR (sender_id = ? AND receiver_id = ?))
	    ORDER BY created_at DESC
	    LIMIT 10`, time, sender, receiver, receiver, sender)
	// Process rows
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return
	}
	// fmt.Println("rows", rows)

	for rows.Next() {
		var message Message
		if err = rows.Scan(utils.GetScanFields(&message)...); err != nil {
			fmt.Println(err)
			return
		}
		messages = append(messages, message)
	}
	return
}

// StoreMessage inserts the message into the database, using all fields
// except `ID`. Returns any execution error.
func (message *Message) StoreMessage() error {
	// message.CreatedAt = int(time.Now().Unix())
	_, err := db.DB.Exec("INSERT INTO messages VALUES(NULL ,?,?,?,?,?)", utils.GetExecFields(message, "ID")...)
	return err
}

// UpdateRead marks messages between the message's sender and receiver as read.
func (message *Message) UpdateRead() error {
	_, err := db.DB.Exec("UPDATE messages SET is_read = true WHERE sender_id = ? AND receiver_id = ? AND is_read = false", message.ReceiverID, message.SenderID)
	return err
}
