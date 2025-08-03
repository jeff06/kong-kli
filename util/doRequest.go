package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

func Execute(cmd *cobra.Command, baseUrl *url.URL, verbs string, queryParameters map[string]interface{}, body map[string]string) *http.Response {
	queryParams := baseUrl.Query()

	token, _ := cmd.Flags().GetString("token")
	bearer := fmt.Sprintf("Bearer %s", token)

	for key, value := range queryParameters {
		if value != nil && value != "" {
			queryParams.Add(key, fmt.Sprintf("%v", value))
		}
	}

	baseUrl.RawQuery = queryParams.Encode()

	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil
	}

	req, err := http.NewRequest(verbs, baseUrl.String(), bytes.NewBuffer(jsonData))
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
