package main

import (
	"database/sql"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "file.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	golog := new(Golog)
	golog.Db = db

	app := cli.NewApp()
	app.Name = "ulogmeExport"
	app.Usage = "export logs to ulogme"
	app.Commands = []cli.Command{
		{
			Name:    "import",
			Aliases: []string{"i"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "logDir",
					Usage: "Directory containing ulogme's logs",
				},
				cli.IntFlag{
					Name:  "uid",
					Value: 1,
					Usage: "User ID from database",
				},
				cli.BoolFlag{
					Name:  "importKeys",
					Usage: "Import key logs",
				},
				cli.BoolFlag{
					Name:  "importWindows",
					Usage: "Import window logs",
				},
			},
			Usage: "import logs",
			Action: func(c *cli.Context) {
				if c.Bool("importKeys") {
					importKeys(
						golog,
						c.Int("uid"),
						c.String("logDir"),
					)
				}

				if c.Bool("importWindows") {
					importWindows(
						golog,
						c.Int("uid"),
						c.String("logDir"),
					)
				}
			},
		},
		{
			Name:    "export",
			Aliases: []string{"e"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "logDir",
					Usage: "Directory to output ulogme's logs",
				},
				cli.IntFlag{
					Name:  "uid",
					Value: 1,
					Usage: "User ID from database",
				},
			},
			Usage: "import logs",
			Action: func(c *cli.Context) {
				window_logs := exportWindows(
					golog,
					c.Int("uid"),
				)
                if len(window_logs) > 0 {
                    WriteStringFile(
                        "window",
                        window_logs,
                        c.String("logDir"),
                    )
                }

                key_logs := exportKeys(
                    golog,
                    c.Int("uid"),
                )
                if len(key_logs) > 0 {
                    WriteIntFile(
                        "keyfreq",
                        key_logs,
                        c.String("logDir"),
                    )
                }
			},
		},
	}
	app.Run(os.Args)
}
