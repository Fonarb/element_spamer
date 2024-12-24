package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"url-shortener/internal/api"
)

func FetchCommand() *cobra.Command {
	var endpoint string

	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch JSON data from API",
		Run: func(cmd *cobra.Command, args []string) {

			apiClient := api.NewAPIClient("https://example.com/api")

			var data interface{}
			if err := apiClient.FetchData(endpoint, &data); err != nil {
				log.Fatalf("Error fetching data: %v", err)
			}

			fmt.Printf("Fetched data: %+v\n", data)
		},
	}

	// Добавляем флаг для указания endpoint
	cmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "API endpoint to fetch data from (required)")
	cmd.MarkFlagRequired("endpoint")

	return cmd
}
