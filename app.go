package main

import (
  "os"
  "log"
  "github.com/codegangsta/cli"
)

const VERSION = "0.1.0"

func init() {
  log.SetFlags(0)
}

func main() {
  app := cli.NewApp()
  app.Name = "ddns"
  app.Usage = "Manage dynamic DNS"
  app.Version = VERSION

  app.Commands = []cli.Command {
    {
      Name: "update",
      Usage: "update <domain>",
      Action: update,
      Flags: []cli.Flag {
        cli.StringFlag{Name: "ip-address, i", Usage: "IP Address to report"},
        cli.StringFlag{"provider, p", "iwantmyname.com", "API provider"},
      },
    },
  }

  app.Run(os.Args)
}

func update(c *cli.Context) {
  domain := c.Args().First();
  if(len(domain) == 0) {
    log.Fatal("No domain provided. Quitting...")
  }
}
