package slack

import (
	"log"

	"github.com/slack-go/slack"
)

//webhookUrl := "https://hooks.slack.com/services/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"

func SlackNotify(alertMessage string) {
	oauthToken := "xoxb-6401933870261-6405502276643-5f0k2jecjvHo77juDnKnoYJ8"
	channelID := "C06CKJZJT6U"

	if oauthToken == "" {
		panic("Missing slack token")
	}
	if channelID == "" {
		panic("Missing slack channel id")
	}

	//rtm := api.NewRTM()
	//go rtm.ManageConnection()

	api := slack.New(
		oauthToken,
		slack.OptionDebug(true),
	)
	attachment := slack.Attachment{
		Color:   "danger",
		Pretext: "Alertname",
		Text:    alertMessage,
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}
	channelID, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText(alertMessage, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		log.Fatalf("Failed to post message: %s", err)
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
