package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

const (
	defaultLayout = "2006/1/2 15:04:05"
)

func main() {
	app := cli.NewApp()
	app.Name = "unix2date"
	app.Version = fmt.Sprintf("0.1.0\ndate: %s\ncommit: %s", Date, Commit)
	app.Usage = "Convert UNIX timestamp to date string."
	app.UsageText = "unix2date [global options] [arguments...]"
	app.Authors = []cli.Author{
		cli.Author{Name: "voidnt",
			Email: "voidint@126.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "layout, l",
			Value: defaultLayout,
			Usage: "date pattern",
		},
		cli.BoolFlag{
			Name:  "pretty, p",
			Usage: "pretty output",
		},
		cli.BoolFlag{
			Name:  "utc",
			Usage: "use UTC time zone",
		},
	}

	app.Action = func(ctx *cli.Context) {
		sVal := ctx.Args().First()
		if len(sVal) == 0 {
			bVal, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			sVal = strings.TrimSpace(string(bVal))
		}

		val, err := strconv.ParseInt(sVal, 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		var loc *time.Location
		if ctx.Bool("utc") {
			loc = time.UTC
		} else {
			loc = time.Local
		}

		t := time.Unix(val, 0).In(loc)
		layout := ctx.String("layout")

		if ctx.Bool("pretty") {
			fmt.Printf("%d => %s\n", val, t.Format(layout))
		} else {
			fmt.Println(t.Format(layout))
		}
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
