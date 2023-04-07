package action

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewAction() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Printf("Hello %q", cCtx.Args().Get(0))
		return nil
	}
}
