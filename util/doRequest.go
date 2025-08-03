package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

func Execute(cmd *cobra.Command, baseUrl *url.URL, param map[string]interface{}) *http.Response {
	queryParams := baseUrl.Query()

	token, _ := cmd.Flags().GetString("token")
	bearer := fmt.Sprintf("Bearer %s", token)

	for key, value := range param {
		if value != nil && value != "" {
			queryParams.Add(key, fmt.Sprintf("%v", value))
		}
	}

	baseUrl.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	return resp
}
