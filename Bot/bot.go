package bot

import (
	"CodeBot/Utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
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

func getFileName(messageStr string) (string, string, string) {
	parts := strings.SplitN(messageStr, " ", 2)
	cleanedContent := ""
	fileName := parts[0]
	fileName = strings.Replace(fileName, "!", "", 1)
	if len(parts) > 1 {
		cleanedContent = strings.TrimSpace(parts[1])
	}
	extension := getFileExtension(fileName)
	fmt.Println(extension)

	return cleanedContent, fileName, extension
}

func getFileExtension(fileName string) string {
	switch fileName {
	case "python":
		return "py"
	case "java":
		return "java"
	case "csharp":
		return "csharp"
	case "php":
		return "php"
	case "go":
		return "go"
	case "javascript":
		return "js"
	}
	return "err"
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
		discord.ChannelMessageSend(message.ChannelID, "Hello World😃")
	case strings.Contains(message.Content, "!bye"):
		discord.ChannelMessageSend(message.ChannelID, "Good Bye👋")
	case strings.Contains(message.Content, "!code"):
		cleanedData := separateMessage(message.Content, "!code")
		task, fileName, extension := getFileName(cleanedData)
		data, success := Utils.Run(task, fileName)
		if success {
			discord.ChannelFileSend(message.ChannelID, "hello."+extension, strings.NewReader(data))
		} else {
			discord.ChannelMessageSend(message.ChannelID, data)
		}

		// add more cases if required
	}

}
