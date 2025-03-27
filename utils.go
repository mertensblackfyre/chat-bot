package main

import (
	"bytes"
	"log"
	"os"

	"encoding/json"
	"fmt"
)

func GetText(response string) {

	var iot Response
	err := json.Unmarshal([]byte(response), &iot)
	if err != nil {
		panic(err)
	}

	candidates := iot.Candidates
	text := candidates[0].Content.Parts[0].Text

	data := &HistotryItem{
		Role: "user",
		Parts: []struct {
			Text string `json:"text"`
		}{
			{Text: text},
		},
	}
/*
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
    fmt.Println(b)
*/
	file, err := os.ReadFile("history.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var d History
	err = json.Unmarshal(file, &d)

	if err != nil {
		fmt.Println(err)
	}

    d = append(d,data...)


}
func AppendHistory() {

}
func WriteJSON(contents History) {
	dataBytes, err := json.MarshalIndent(contents, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("history.json", dataBytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func JSONInterface() *bytes.Buffer {
	file, err := os.ReadFile("history.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var data History
	err = json.Unmarshal(file, &data)

	if err != nil {
		fmt.Println(err)
	}

	history, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	bytes := bytes.NewBuffer(history)

	return bytes
}

func MapToJSONString(inputMap map[string]any) (string, error) {
	jsonBytes, err := json.Marshal(inputMap)

	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
