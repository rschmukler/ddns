package main

import (
  "os"
  "log"
  "time"
  "github.com/codegangsta/cli"
  "github.com/rschmukler/go-ip-checker"
  "github.com/rschmukler/ddns/app"
  "github.com/rschmukler/ddns/providers"
)

const VERSION = "0.1.3"


func init() {
  log.SetFlags(0)
}

func main() {
  app := app.NewDDNSApp()
  app.Version = VERSION
  app.Commands = []cli.Command {
    {
      Name: "update",
      Usage: "update <domain>",
      Action: update(app),
      Flags: []cli.Flag {
        cli.StringFlag{Name: "ip-address, i", Usage: "IP Address to report"},
        cli.StringFlag{"provider, p", "iwantmyname.com", "API provider"},
        cli.BoolFlag{"configure, c", "Reconfigure"},
      },
    },
    {
      Name: "run",
      Usage: "run",
      Action: run(app),
      Flags: []cli.Flag {
        cli.StringFlag{Name: "ip-address, i", Usage: "IP Address to report"},
        cli.StringFlag{"provider, p", "iwantmyname.com", "API provider"},
        cli.IntFlag{"every, e", 10, "Check every"},
      },
    },
  }
  app.Run(os.Args)

}

func update(app *app.DDNSApp) func(c *cli.Context) {
  return func(c *cli.Context) {
    provider, ip := setupProvider(app, c)
    go provider.Update(ip, app.Updates)
    printUpdate(<- app.Updates)
  }
}

func run(myApp *app.DDNSApp) func(c *cli.Context) {
  return func(c *cli.Context) {
    provider, ip := setupProvider(myApp, c)
    runEvery := time.Minute * time.Duration(c.Int("every"))

    ipChanged := ipchecker.Poll(runEvery)

    for {
      select {
        case ip = <-ipChanged:
          printUpdate(app.DDNSUpdates{"Info", "App", "Updating DDNS to ip " + ip})
          go provider.Update(ip, myApp.Updates)
        case update := <-myApp.Updates:
          go printUpdate(update)
      }
    }
  }
}

func setupProvider(app *app.DDNSApp, c *cli.Context) (providers.Provider, string) {
    provider, providerPresent := providers.GetProvider(c.String("provider"))

    if !providerPresent {
      log.Fatalf("Could not find provider specified (%s)\nValid Providers are:\n%s", c.String("provider"), providers.ListProviders())
    }

    config, configPresent := app.Config[c.String("provider")]

    reconfigure := c.Bool("configure")
    if !configPresent || reconfigure {
      log.Printf("Please enter the following information: ")
      provider.GenerateConfig(app.Config)
      app.SaveConfig()
      config = app.Config[c.String("provider")]
    }
    ip := c.String("ip")
    if len(ip) == 0  {
      ip = ipchecker.Check()
    }
    provider.ReadConfig(config)
    return provider, ip
}

func printUpdate(update app.DDNSUpdates) {
  log.Printf("[%s] (%s) %s\n", update.Type, update.From, update.Message)
}
