package ungo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	defaultEmoji = "💩"
)

func NewApp() *cli.App {
	var (
		emojiCount int
	)
	return &cli.App{
		Name:  "ungo",
		Usage: "ungo adds emoji of shit at suffix",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "count",
				Aliases:     []string{"cnt", "c"},
				Value:       1,
				Usage:       "specify the number of emoji",
				Destination: &emojiCount,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() == 0 {
				return errors.New("need to specify some text")
			}

			text := cCtx.Args().First()
			emoji := strings.Repeat(defaultEmoji, emojiCount)
			fmt.Printf("%s%v\n", text, emoji)
			return nil
		},
	}
}
