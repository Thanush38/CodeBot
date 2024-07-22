package Utils

const (
	apiEndpoint    = "https://api.openai.com/v1/chat/completions"
	addFunc        = "def add_numbers(a, b):\n    return a + b"
	pythonEven     = "def is_even_or_odd(number):\n    if number % 2 == 0:\n        return \"even\"\n    else:\n        return \"odd\""
	javaScriptEven = "function isEvenOrOdd(number) {\n    if (number % 2 === 0) {\n        return \"even\";\n    } else {\n        return \"odd\";\n    }\n}\n\nconst testNumber = 4;\nconst result = isEvenOrOdd(testNumber);\nconsole.log(\\`The number \\${testNumber} is \\${result}.\\`);"
)

func pythonCode(task string) (string, bool) {
	switch task {
	case "add":
		return addFunc, true
	case "even or odd":
		return pythonEven, true
	}
	return "Unable to Generate Code", false
}

func javascriptCode(task string) (string, bool) {
	switch task {
	case "even or odd":
		return javaScriptEven, true
	}
	return "unable to generate code", false
}

func Run(task string, language string) (string, bool) {
	switch language {
	case "python":
		return pythonCode(task)
	case "javascript":
		return javascriptCode(task)
	}
	return "Unable to Generate Code", false

	// Use your API KEY here
	//apiKey :=
	//client := resty.New()
	//
	//response, err := client.R().
	//	SetAuthToken(apiKey).
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(map[string]interface{}{
	//		"model":      "gpt-3.5-turbo",
	//		"messages":   []interface{}{map[string]interface{}{"role": "system", "content": "Hi can you tell me what is the factorial of 10?"}},
	//		"max_tokens": 50,
	//	}).
	//	Post(apiEndpoint)
	//
	//if err != nil {
	//	log.Fatalf("Error while sending send the request: %v", err)
	//}
	//
	//body := response.Body()
	//
	//var data map[string]interface{}
	//err = json.Unmarshal(body, &data)
	//if err != nil {
	//	fmt.Println("Error while decoding JSON response:", err)
	//	return
	//}
	//
	//// Extract the content from the JSON response
	//content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	//fmt.Println(content)

}
