package service

import (
	"url-shortener/internal/api"
	"url-shortener/internal/element"
)

type MessageProcessor struct {
	apiClient     *api.APIClient
	elementSender *element.ElementSender
}

func NewMessageProcessor(apiClient *api.APIClient, elementSender *element.ElementSender) *MessageProcessor {
	return &MessageProcessor{
		apiClient:     apiClient,
		elementSender: elementSender,
	}
}

func (p *MessageProcessor) Process() error {

	var apiResponse struct {
		RoomID  string `json:"room_id"`
		Message string `json:"message"`
	}
	err := p.apiClient.FetchData("/data-endpoint", &apiResponse)
	if err != nil {
		return err
	}
	// Отправка сообщения в Element
	return p.elementSender.SendMessage(apiResponse.RoomID, apiResponse.Message)
}
