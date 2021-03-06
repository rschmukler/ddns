package app

import (
  "io/ioutil"
  "encoding/json"

  "github.com/codegangsta/cli"
)

type DDNSApp struct {
  *cli.App
  Config map[string]map[string]string
  Updates chan DDNSUpdates
}

type DDNSUpdates struct {
  Type string
  From string
  Message string
}


func NewDDNSApp() *DDNSApp {
  app := &DDNSApp{
    App: cli.NewApp(),
    Updates: make(chan DDNSUpdates),
  }
  app.Name = "ddns"
  app.Usage = "Manage dynamic DNS"
  app.LoadConfig()
  app.SaveConfig()

  return app
}

var CONFIG_PATH = "/etc/ddns.json"

func (d *DDNSApp) LoadConfig() {

  d.Config = make(map[string]map[string]string)

  config, err := ioutil.ReadFile(CONFIG_PATH)
  if err != nil {
    check(ioutil.WriteFile(CONFIG_PATH, []byte(""), 0600))
  } else {
    json.Unmarshal(config, &d.Config)
  }
}

func (d *DDNSApp) SaveConfig() {
  str, err := json.Marshal(d.Config)
  check(err)
  check(ioutil.WriteFile(CONFIG_PATH, []byte(str), 0600))
}

func check(e error) {
  if(e != nil) {
    panic(e)
  }
}
