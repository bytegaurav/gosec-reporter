package main

import (
	"Reporter/messenger"
	"Reporter/models"
	"Reporter/processors"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)


func main() {
	file := flag.String("log-file", "", "location of gosec (https://github.com/securego/gosec) log file to be processed")
	space := flag.String("space-id", "", "id of chat room")
	key := flag.String("key", "", "key param of web hook")
	token := flag.String("token", "", "token param of web hook")
	flag.Parse()
	url := fmt.Sprintf("https://chat.googleapis.com/v1/spaces/%s/messages?key=%s&token=%s", *space, *key, *token)

	if *file == "" {
		log.Fatal("input file missing")
	}
	fileData, err := ioutil.ReadFile(*file)

	if err != nil {
		log.Fatal(err)
		return
	}
	logObject := models.GoSec{}
	json.Unmarshal(fileData, &logObject)

	firstMessage := models.Chat{
		Text: "SAST Result from recent pipeline build",
	}

	threadCreateMessage, err :=messenger.SendMessage(url, firstMessage)

	if err != nil {
		log.Fatal(err)
	}

	severity := processors.SortIssuesBySeverity(logObject)
	message := models.Chat{
		Text: fmt.Sprintf("```HIGH   : %d \nMEDIUM : %d \nLOW    : %d	\n```", severity.High, severity.Medium, severity.Low),
		Thread: threadCreateMessage.Thread,
	}
	messenger.SendMessage(url, message)

	byType, largest := processors.SortIssuesByMessage(logObject)

	messageString := ""
	for k, v := range byType {

		temp := k

		for i := len(k) - 2; i < largest; i++ {
			temp += " "

		}
		temp += fmt.Sprintf(":	%d issue(s)", v)
		messageString += "\n" + temp
	}

	messenger.SendMessage(url, models.Chat{
		Text:   fmt.Sprintf("```\n%s\n```", messageString),
		Thread: threadCreateMessage.Thread,
	})

	if err != nil {
		log.Fatal(err)
	}

	return

}

