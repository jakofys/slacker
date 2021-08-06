package main

import (
	"errors"
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
			cli.Command{
				Name:        "set",
				Description: "Set Slack token application",
				Action:      SetToken,
			},
			cli.Command{
				Name:        "sendto",
				Description: "Send a message to a user by email",
				Action:      SendTo,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "notifier",
						Value: "info",
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func SetToken(ctx *cli.Context) {

	if ioutil.WriteFile("token", []byte(ctx.Args()[0]), 0644) != nil {
		panic(errors.New("File problem"))
	}

}

func SendTo(ctx *cli.Context) {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		panic(errors.New("File problem"))
	}

	app := slack.New(string(token))
	if ok := ctx.Args()[0]; ok == "" {
		panic(errors.New("Data don't give"))
	}
	user, err := app.GetUserByEmail(ctx.Args()[0])
	if err != nil {
		panic(errors.New("User not found"))
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
	app.SendMessage(channel.ID, slack.MsgOptionText(ctx.Args()[1], false))

}
