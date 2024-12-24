package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"url-shortener/internal/api"
	"url-shortener/internal/element"
)

func main() {
	fmt.Println("Спам сообщений в Element")
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println()
		fmt.Println("Доступные варианты действий:")
		fmt.Println("[1] Получить запросы из текущего релиза")
		fmt.Println("[2] Отправка сообщений")
		fmt.Print("Для выбора нажмите цифру на клавиатуре")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			// Обработка fetch
			fmt.Print("Enter API endpoint: ")
			endpoint, _ := reader.ReadString('\n')
			endpoint = strings.TrimSpace(endpoint)

			apiClient := api.NewAPIClient("https://example.com/api")

			var data interface{}
			err := apiClient.FetchData(endpoint, &data)
			if err != nil {
				fmt.Printf("Error fetching data: %v\n", err)
				return
			}
			fmt.Printf("Fetched data: %+v\n", data)

		case "2":
			// Обработка send
			fmt.Print("  [1] Отправить сообщение в комнату Сборка Проектов РЦ")
			roomID, _ := reader.ReadString('\n')
			roomID = strings.TrimSpace(roomID)

			fmt.Print("  [2] Редактировать страндартное сообщение")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			elementSender := element.NewElementSender("https://element-server.com", "YOUR_ACCESS_TOKEN")
			err := elementSender.SendMessage(roomID, message)
			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
				return
			}
			fmt.Println("Message sent successfully!")

		default:
			fmt.Println("ошибка выбора действия")
		}
	}
}
