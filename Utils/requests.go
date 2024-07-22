package Utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func writeData(data map[string]interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	err = ioutil.WriteFile("respons.json", jsonData, 0644)
	if err != nil {
		fmt.Println("error writing JSON to file: %v", err)
	}
}

func separateContent(data string, language string) [3]string {
	parts := strings.Split(data, "```\n")
	fmt.Println(len(parts))

	codeSnippet := ""
	description := ""

	if len(parts) <= 1 {
		codeSnippet = data
	} else if len(parts) == 2 {
		codeSnippet = strings.TrimPrefix(parts[0], language+"\n")
		description = parts[1]
	} else if len(parts) == 3 {
		description = parts[0]
		codeSnippet = strings.TrimPrefix(parts[2], language+"\n")
		description = description + parts[2]
	}
	sections := [3]string{codeSnippet, description}
	return sections

}

func SendRequest(task string, language string, typeOfJob string) (string, bool) {

	//return "Unable to Generate Code", false
	payload := ""
	if typeOfJob == "write" {
		payload = "Create " + language + " code to do the following task: " + task + ", Give me the code as a string with comments"
	} else if typeOfJob == "fix" {
		payload = "fix the error in the following " + language + " code: " + task + ", please state what changes were made covered by triple backticks"
	}

	// Use your API KEY here
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_TOKEN")
	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": payload}},
			"max_tokens": 1000,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)

	writeData(data)

	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return "Error while decoding JSON response:", false
	}

	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(content)

	return content, true

}
