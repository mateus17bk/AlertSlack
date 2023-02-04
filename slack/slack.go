package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendSlack(textAlert string) {

	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		panic("SLACK_AUTH_TOKEN não foi configurado!")
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		panic("SLACK_CHANNEL_ID não foi configur")
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   "danger",
		Pretext: "Alerta de servidor down",
		Text:    textAlert,
	}
	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("Mensagem enviado com sucesso! %s as %s", channelID, timestamp)
	}

}
