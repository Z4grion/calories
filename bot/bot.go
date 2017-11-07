package bot

import (
	"log"

	"github.com/bobheadxi/calories/facebook"
)

type Bot struct {
	API *facebook.API
}

func (b *Bot) SetApi(api *facebook.API) {
	b.API = api
	b.API.MessageHandler = b.TestMessageReceivedAndReply
}

// TestMessageReceivedAndReply : Test that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, sender facebook.Sender, msg facebook.ReceivedMessage) {
	b.API.SendTextMessage(sender.ID, "Hello!")
	log.Printf("Event: %+v", event) // {ID:2066945410258565 Time:1510063491984}
	log.Printf("Opts: %+v", sender)
	log.Printf("Msg: %+v", msg) // {ID:mid.$cAAcNE7mWyw1lyBGR51flsxJvj8_- Text:hello Seq:1028142}
}