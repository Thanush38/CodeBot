package bot

import (
	"CodeBot/Utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {

	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// add a event handler
	discord.AddHandler(newMessage)

	// open session
	discord.Open()
	defer discord.Close() // close session, after function termination

	// keep bot running untill there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func separateMessage(messageStr string, removeStr string) string {
	cleanedContent := strings.Replace(messageStr, removeStr, "", 1)
	cleanedContent = strings.TrimSpace(cleanedContent)
	fmt.Println(cleanedContent)
	return cleanedContent
}
func separateContent(data string, language string) [3]string {
	parts := strings.Split(data, "```\n")
	fmt.Println("parts in separation: " + fmt.Sprintf(strconv.Itoa(len(parts))))

	codeSnippet := ""
	description := ""

	startIndex := strings.Index(data, "```"+language)
	endIndex := strings.LastIndex(data, "```")

	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		fmt.Println("Data does not contain the expected code block.")

		return [3]string{"Data does not contain the expected code block."}
	}

	// Extract and clean the code snippet
	codeSnippet = data[startIndex+len("```javascript\n") : endIndex]
	description = data[:startIndex] + data[endIndex+len("```"):]

	fmt.Println("Code Snippet:\n", codeSnippet)
	fmt.Println("\nDescription:\n", description)

	sections := [3]string{codeSnippet, description}
	return sections

}

func getFileName(messageStr string, typeOfJob string) (string, string, string, bool) {
	parts := strings.SplitN(messageStr, " ", 2)
	cleanedContent := ""
	fileName := parts[0]
	fileName = strings.Replace(fileName, "!", "", 1)
	if len(parts) > 1 {
		cleanedContent = strings.TrimSpace(parts[1])
	}

	data := ""
	if fileName == "" {
		data = "No language was entered!\n To view how to use this bot visit: https://github.com/Thanush38/CodeBot"
		return fileName, ".txt", data, false
	}
	if cleanedContent == "" {
		data = "No task was entered!\n To view how to use this bot visit: https://github.com/Thanush38/CodeBot"
		return fileName, ".txt", data, false
	}
	extension := getFileExtension(fileName)

	if extension == "err" {
		data = "Entered Language is not supported! \n To view a list of all languages available visit: https://github.com/Thanush38/CodeBot"
		return fileName, ".txt", data, false
	}

	fmt.Println(extension)

	content, success := Utils.SendRequest(cleanedContent, fileName, typeOfJob)

	//if success {
	//	discord.ChannelFileSend(message.ChannelID, "hello."+extension, strings.NewReader(data))
	//} else {
	//	discord.ChannelMessageSend(message.ChannelID, data)
	//}

	return fileName, extension, content, success
}

func getFileExtension(fileName string) string {
	switch fileName {
	case "python":
		return "py"
	case "java":
		return "java"
	case "csharp":
		return "cs"
	case "php":
		return "php"
	case "go":
		return "go"
	case "javascript":
		return "js"
	case "typescript":
		return "ts"
	case "ruby":
		return "rb"
	case "swift":
		return "swift"
	case "kotlin":
		return "kt"
	case "rust":
		return "rs"
	case "scala":
		return "scala"
	case "perl":
		return "pl"
	case "r":
		return "r"
	case "shell":
		return "sh"
	case "html":
		return "html"
	case "css":
		return "css"
	case "sql":
		return "sql"
	case "dart":
		return "dart"
	case "haskell":
		return "hs"
	case "lua":
		return "lua"
	case "objective-c":
		return "m"
	case "c":
		return "c"
	case "cpp":
		return "cpp"
	default:
		return "err"
	}
}
func getFileTitle(language string) string {

	now := time.Now()
	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	Name := fmt.Sprintf("CodeBot %s Code generated at %d-%02d-%02d %02d-%02d-%02d", language, year, month, day, hour, minute, second)

	return Name

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	/* prevent bot responding to its own message
	this is achived by looking into the message author id
	if message.author.id is same as bot.author.id then just return
	*/
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// respond to user message if it contains `!help` or `!bye`

	switch {
	case strings.Contains(message.Content, "!help"):
		discord.ChannelMessageSend(message.ChannelID, "Type: !code !language instructions to generate code \n Type: !fix !language and code to fix")
	case strings.Contains(message.Content, "!code"):
		cleanedData := separateMessage(message.Content, "!code")
		language, extension, data, success := getFileName(cleanedData, "write")
		if success {
			sections := separateContent(data, language)
			discord.ChannelFileSend(message.ChannelID, getFileTitle(language)+"."+extension, strings.NewReader(sections[0]))
			if sections[1] != "" {
				discord.ChannelMessageSend(message.ChannelID, sections[1])
			}
		} else {
			discord.ChannelMessageSend(message.ChannelID, data)
		}
	case strings.Contains(message.Content, "!fix"):
		cleanedData := separateMessage(message.Content, "!fix")
		language, extension, data, success := getFileName(cleanedData, "fix")
		if success {
			discord.ChannelFileSend(message.ChannelID, getFileTitle(language)+"."+extension, strings.NewReader(data))
		} else {
			discord.ChannelMessageSend(message.ChannelID, data)
		}
	}

}
