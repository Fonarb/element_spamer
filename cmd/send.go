package cmd

import (
	"log"
	"url-shortener/internal/element"

	"github.com/spf13/cobra"
)

func SendCommand() *cobra.Command {
	var roomID, message string

	cmd := &cobra.Command{
		Use:   "send",
		Short: "Send a message to Element",
		Run: func(cmd *cobra.Command, args []string) {
			// Инициализация отправителя Element
			elementSender := element.NewElementSender("https://element-server.com", "YOUR_ACCESS_TOKEN")

			// Отправка сообщения
			if err := elementSender.SendMessage(roomID, message); err != nil {
				log.Fatalf("Error sending message: %v", err)
			}

			log.Println("Message sent successfully!")
		},
	}

	// Добавляем флаги для указания roomID и message
	cmd.Flags().StringVarP(&roomID, "room", "r", "", "Room ID to send message to (required)")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Message content to send (required)")
	cmd.MarkFlagRequired("room")
	cmd.MarkFlagRequired("message")

	return cmd
}
