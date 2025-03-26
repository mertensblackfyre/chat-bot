package main

import (
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"

	"encoding/json"
	"fmt"
)


type CustomContent struct {
	Parts []string `json:"parts"`
	Role  string       `json:"role"` // "user", "model", "system", etc.
}

func WriteJSON(contents []*genai.Content) {

	// Marshal with indentation for better readability
	dataBytes, err := json.MarshalIndent(contents, "", "  ")
	if err != nil {

		log.Fatalln(err)
	}
	err = os.WriteFile("history.json", dataBytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func JSONInterface() {
	file, err := os.ReadFile("history.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data []*CustomContent
	err = json.Unmarshal(file, &data)

	// Accessing the first item's Parts
	if len(data) > 0 {
		firstItem := data[0]
		fmt.Println(firstItem.Parts) // Access Parts directly, not with ["Parts"]
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}

	fmt.Println("---")
}

func MapToJSONString(inputMap map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(inputMap)

	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func Parser(body string) string {

	var response Response

	err := json.Unmarshal([]byte(body), &response)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	if len(response.Message.Content) > 0 {
		content := response.Message.Content
		return content
	}

	return ""
}
