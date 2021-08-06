package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/slack-go/slack"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "slacker",
		Usage: "Slack API implementation",
		Commands: []cli.Command{
			{
				Name:        "set",
				Description: "Set Slack token application",
				Action:      SetToken,
			},
			{
				Name:        "sendto",
				Description: "Send a message to a user by email",
				Action:      SendTo,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func SetToken(ctx *cli.Context) {

	if err := ioutil.WriteFile("token", []byte(ctx.Args()[0]), 0644); err != nil {
		fmt.Printf("Internal error: %s\n", err)
		return
	}

}

func SendTo(ctx *cli.Context) {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Println("No token available, to register a token, make: 'slacker set-token [token]'")
		return
	}

	app := slack.New(string(token))
	if ok := ctx.Args()[0]; ok == "" {
		fmt.Println("You have to indicate and email and message as: \n slacker sendto 'dupont@example.com' 'You message'")
		return
	}
	user, err := app.GetUserByEmail(ctx.Args()[0])
	if err != nil {
		fmt.Printf("User for email %s not found \n", ctx.Args()[0])
		return
	}
	conv := &slack.OpenConversationParameters{
		ReturnIM: true,
		Users: []string{
			user.ID,
		},
	}
	channel, _, _, err := app.OpenConversation(conv)
	if err != nil {
		panic(err)
	}

	_, _, _, err = app.SendMessage(channel.ID, slack.MsgOptionText(ctx.Args()[1], false))
	if err != nil {
		panic(err)
	}

}
