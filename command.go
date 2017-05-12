package xkcdnews

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"strings"
)

func (b *Bot) Command(args []string) {
	if len(args) == 0 {
		b.log.Println("No args provided")
	} else {
		switch strings.ToLower(args[0]) {
		case "tweet":
			msg := fmt.Sprintf("%s", args[1:])
			params := &twitter.StatusUpdateParams{
				TrimUser: twitter.Bool(true),
			}
			tweet, _, err := b.client.Statuses.Update(msg, params)
			if err != nil {
				b.log.Println("Error posting tweet:", err)
				return
			}
			b.log.Println("Tweet Posted: https://twitter.com/XKCDNews/status/" + tweet.IDStr)
		default:
			b.log.Println("Command unknown:", args[0])
		}
	}
}
