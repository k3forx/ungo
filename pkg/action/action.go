package action

import (
	"github.com/urfave/cli/v2"
)

func NewAction() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		// if cCtx.Args().Len() == 0 {
		// 	return nil
		// }

		// text := cCtx.Args().First()
		// // emoji := strings.Repeat(defaultEmoji, ungo.GetNum)
		// fmt.Printf("%s%v\n", text, emoji)
		return nil
	}
}
