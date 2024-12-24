package element

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ElementSender struct {
	baseURL     string
	accessToken string
	client      *http.Client
}

func NewElementSender(baseURL, accessToken string) *ElementSender {
	return &ElementSender{
		baseURL:     baseURL,
		accessToken: accessToken,
		client:      &http.Client{},
	}
}

type Message struct {
	RoomID  string `json:"room_id"`
	Content string `json:"content"`
}

func (s *ElementSender) SendMessage(roomID, content string) error {
	message := Message{
		RoomID:  roomID,
		Content: content,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.baseURL+"/_matrix/client/r0/rooms/"+roomID+"/send/m.room.message", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send message to Element")
	}

	return nil
}
