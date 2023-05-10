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
		Usage: "ungo adds 💩 at suffix of some text",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "count",
				Aliases:     []string{"cnt", "c"},
				Value:       1,
				Usage:       "the number of 💩",
				Destination: &emojiCount,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() == 0 {
				return errors.New("need to specify some text")
			}

			text := cCtx.Args().First()
			fmt.Printf("%s%s\n", text, strings.Repeat(defaultEmoji, emojiCount))
			return nil
		},
	}
}
