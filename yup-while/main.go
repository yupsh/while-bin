package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/while"
)

func main() {
	app := &cli.App{
		Name:  "while",
		Usage: "read from stdin and process line by line",
		UsageText: `while [FILE...]

   Read from standard input and process line by line.`,
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "while: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments if provided
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Create and execute the while command
	cmd := While(params...)
	return yup.Run(cmd)
}
