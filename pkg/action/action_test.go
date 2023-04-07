package action_test

import (
	"testing"

	"github.com/k3forx/ungo/pkg/action"
	"github.com/urfave/cli/v2"
)

func TestAction(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		args   []string
		isErr  bool
		errMsg string
	}{
		"success": {
			args:  []string{"", "cmd"},
			isErr: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			app := &cli.App{
				Commands: []*cli.Command{
					{Name: "cmd", Action: action.NewAction()},
				},
			}

			err := app.Run(c.args)
			if c.isErr && err == nil {
				t.Errorf("should be err, but got nil")
			}
		})
	}

}
