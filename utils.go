package main

import (
	"bytes"
	"log"
	"os"

	"encoding/json"
	"fmt"
)


func WriteParameters() {
}

func WriteHistory() {
// 1. Read main config file
	mainConfig, err := os.ReadFile("payload.json")
	if err != nil {
		fmt.Printf("Error reading main config: %v\n", err)
		return
	}

	constsConfig, err := os.ReadFile("history.json")
	if err != nil {
		fmt.Printf("Error reading consts config: %v\n", err)
		return
	}

	// 3. Parse both JSON files
	var mainData map[string]any
	var constsData map[string]any

	if err := json.Unmarshal(mainConfig, &mainData); err != nil {
		fmt.Printf("Error parsing main config: %v\n", err)
		return
	}

	if err := json.Unmarshal(constsConfig, &constsData); err != nil {
		fmt.Printf("Error parsing consts config: %v\n", err)
		return
	}

	// 4. Merge the consts into main config
	mainData["contents"] = constsData

	// 5. Write merged result back to file
	mergedData, err := json.MarshalIndent(mainData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling merged config: %v\n", err)
		return
	}

	if err := os.WriteFile("payload.json", mergedData, 0644); err != nil {
		fmt.Printf("Error writing merged config: %v\n", err)
		return
	}

	fmt.Println("Successfully merged config files!")
}
func WriteSystemInstructions() {
	data, err := os.ReadFile("sys.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var person Persona
	if err := json.Unmarshal(data, &person); err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	file, err := os.Create("payload.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(person); err != nil {
		log.Println(err)
	}

}

func GetText(response string) string {
	var iot Response
	err := json.Unmarshal([]byte(response), &iot)
	if err != nil {
		panic(err)
	}

	candidates := iot.Candidates

	if len(candidates) == 0 {
		log.Println("No Value")
		return ""
	}

	if len(candidates[0].Content.Parts) == 0 {
		log.Println("No Value 2")
		return ""
	}

	text := candidates[0].Content.Parts[0].Text

	return text
}

func AppendMessage(response string, role string) string {

	var text string
	if role == "model" {
		text = GetText(response)
	} else {
		text = response
	}

	file, err := os.ReadFile("history.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var d History
	err = json.Unmarshal(file, &d)

	if err != nil {
		fmt.Println(err)
	}
	d.Contents = append(d.Contents,
		struct {
			Role  string `json:"role"`
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			Role: role,
			Parts: []struct {
				Text string `json:"text"`
			}{
				{Text: text},
			},
		})

	if role == "model" {
		text = GetText(response)
	} else {
		text = response
	}
	WriteJSON(d)
	return text

}

func WriteJSON(contents History) {
	// Marshal the data to JSON with indentation
	dataBytes, err := json.MarshalIndent(contents, "", "  ")
	if err != nil {
		log.Println(err)
	}

	// Open the file with create/truncate permissions
	// Use 0644 permissions (owner read/write, group/others read)
	f, err := os.OpenFile("history.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close() // Ensure file is closed when function exits

	// Write the JSON data
	_, err = f.Write(dataBytes)
	if err != nil {
		log.Println(err)
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
