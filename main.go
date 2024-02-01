package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func eventshandler(analytics <-chan *slacker.CommandEvent) {
	for events := range analytics {
		fmt.Println("command events")
		fmt.Println(events.Command)
		fmt.Println(events.Timestamp)
		fmt.Println(events.Parameters)
		fmt.Println(events.Event)
	}
}

func main() {
	os.Setenv("Slack_bot_token", "SlackBotToken")

	os.Setenv("Slack_App_Token", "AppToken")
	bot := slacker.NewClient(os.Getenv("Slack_bot_token"), os.Getenv("Slack_App_Token"))

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Age calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(botctx slacker.BotContext, request slacker.Request, w slacker.ResponseWriter) {
			year := request.Param("year")
			yob, _ := strconv.Atoi(year)

			age := 2024 - yob
			r := fmt.Sprintf("Age is %d", age)
			w.Reply(r)

		},
	})
	go eventshandler(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
