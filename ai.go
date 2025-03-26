package main

import (
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"bytes"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
)

func GeminiAI(message string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(GEMINI_API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(0.9)
	model.SetTopP(0.5)
	model.SetTopK(20)
	model.SetMaxOutputTokens(100)
	model.SystemInstruction = genai.NewUserContent(genai.Text("You are Yoda from Star Wars."))
	model.ResponseMIMEType = "application/json"

	cs := model.StartChat()

	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text("Hello, I have 2 dogs in my house."),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text("Great to meet you. What would you like to know?"),
			},
			Role: "model",
		},
	}

    WriteJSON(cs.History)
    JSONInterface()
	iter := cs.SendMessageStream(ctx, genai.Text("How many paws are in my house?"))
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//printResponse(resp)
	}
}

func OpenAI(payload string) string {
	url := "http://localhost:11434/api/chat"

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "Bearer "+OPENAI)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println(string(body))

	return string(body)
}
