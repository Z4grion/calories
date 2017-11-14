package bot

import (
	"strconv"

	"github.com/bobheadxi/calories/server"
)

func (b *Bot) help(c *Context) error {
	return b.api.SendTextMessage(c.facebookID, "Hello! I'm pretty useless and can't really do anything right now, sorry :(")
}

func (b *Bot) test(c *Context) error {
	e := server.Entry{
		ID:       c.facebookID,
		Time:     c.timestamp,
		Item:     c.content,
		Calories: 5,
	}
	err := b.server.AddEntry(e)
	if err != nil {
		b.api.SendTextMessage(c.facebookID, "No new entry for you: "+err.Error())
	}
	response, err := b.server.SumCalories(c.facebookID)
	if err != nil {
		b.api.SendTextMessage(c.facebookID, "No calories for you: "+err.Error())
		return err
	}
	return b.api.SendTextMessage(c.facebookID, "your total calories are "+strconv.Itoa(response))
}
