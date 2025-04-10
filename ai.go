package main

import (
	"io"
	"log"

	"bytes"
	"fmt"
	"net/http"
)

func Gemini(message string) string {

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + GEMINI_API_KEY

	WriteSystemInstructions()
	AppendMessage(message, "user")
	//WriteHistory()
	Write2()
	history := JSONInterface()

	resp, err := http.Post(url, "application/json", history)
	if err != nil {
		log.Fatalf("Error resp: %v", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)

	text := AppendMessage(sb, "model")

	Write2()
	//WriteHistory()
	return text
}
func Ollama(payload string) string {

	url := "http://localhost:11434/api/chat"

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))

	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println(string(body))

	return string(body)
}
