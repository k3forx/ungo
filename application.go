package ungo

import (
	"github.com/k3forx/ungo/pkg/action"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:   "boom",
		Usage:  "make an explosive entrance",
		Action: action.NewAction(),
	}
}
