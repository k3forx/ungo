package action

import "github.com/urfave/cli/v2"

func NewTestAction() func(cCtx *cli.Context) error {
	return NewAction()
}
