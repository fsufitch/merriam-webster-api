package main

import (
	"errors"
	"fmt"
	"os"

	mwapi "github.com/fsufitch/merriam-webster-api"
	"github.com/urfave/cli/v2"
)

var userAgent = fmt.Sprintf("merriam-webster CLI / mwapi %s", mwapi.Version)

type printfFunc func(string, ...interface{}) (int, error)

func debugf(format string, other ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, "[DEBUG] "+format, other...)
}

type dictionary struct {
	CommandName    string
	CommandAliases []string
	Name           string
	URL            string
	Lookup         func(terms []string, url string, key string, json bool, verbose bool, debugf printfFunc) error
}

var collegiate = dictionary{
	CommandName:    "collegiate",
	CommandAliases: []string{"c", "cd"},
	Name:           "Merriam-Webster's Collegiate® Dictionary with Audio",
	URL:            "https://www.dictionaryapi.com/api/v3/references/collegiate/json",
	Lookup:         collegiateLookup,
}

func command(d dictionary) *cli.Command {
	return &cli.Command{
		Name:    d.CommandName,
		Aliases: d.CommandAliases,
		Usage:   fmt.Sprintf("use the %s", d.Name),
		// Category:  "dictionary",
		ArgsUsage:              "[search terms]",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "api-key",
				Aliases:  []string{"k"},
				Usage:    "use the key `API_KEY` to access the dictionary",
				Required: true,
				EnvVars:  []string{"MERRIAM_WEBSTER_KEY", "MW_KEY", "API_KEY"},
			},
			&cli.StringFlag{
				Name:        "url",
				Usage:       "use a different `BASE_URL` for the query",
				Value:       d.URL,
				DefaultText: d.URL,
				EnvVars:     []string{"MERRIAM_WEBSTER_URL", "MW_URL", "DICTIONARY_URL"},
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "output the raw JSON response of the query",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "output more details; if --json was specified, pretty-print the entire output",
			},
		},
		Action: func(c *cli.Context) error {
			if len(c.Args().Slice()) == 0 {
				return errors.New("no search terms specified")
			}
			if c.String("url") == "" {
				return errors.New("API URL is empty")
			}
			if c.String("api-key") == "" {
				return errors.New("API key is empty")
			}

			debugf := debugf
			if !c.Bool("debug") {
				debugf = func(string, ...interface{}) (int, error) { return 0, nil }
			}

			return d.Lookup(
				c.Args().Slice(),
				c.String("url"),
				c.String("api-key"),
				c.Bool("json"),
				c.Bool("verbose"),
				debugf,
			)
		},
	}
}

var app = &cli.App{
	Name:                 "merriam-webster",
	Usage:                "perform lookups in the Merriam-Webster dictionary",
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "print debug messages to stderr",
		},
	},
	Commands: []*cli.Command{
		command(collegiate),
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error during execution: %v\n", err)
		os.Exit(1)
	}
}
