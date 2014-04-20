package providers

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "github.com/rschmukler/go-prompt"
  "github.com/rschmukler/ddns/app"
)

func init() {
  RegisterProvider("iwantmyname.com", &IWantMyNameProvider {
  })
}

type IWantMyNameProvider struct {
  username string
  password string
  domains  string
}

func (p *IWantMyNameProvider) GenerateConfig(config map[string]map[string]string) {
  myConfig, present := config["iwantmyname.com"]

  if !present {
    myConfig = make(map[string]string)
  }

  myConfig["domains"] = prompt.StringOrKeep("\tDomains", myConfig["domains"])
  myConfig["username"] = prompt.StringOrKeep("\tUsername", myConfig["username"])
  myConfig["password"] = prompt.PasswordOrKeep("\tPassword", myConfig["password"])

  config["iwantmyname.com"] = myConfig
}

func (p *IWantMyNameProvider) ReadConfig(config map[string]string) {
  p.username = config["username"];
  p.password = config["password"];
  p.domains = config["domains"];
}

func (p *IWantMyNameProvider) Update(ip string, updates chan app.DDNSUpdates) {
  client := &http.Client{}

  url := fmt.Sprintf("https://iwantmyname.com/basicauth/ddns?hostname=%s", p.domains)
  if len(ip) > 0 {
    url += "&myip=" + ip
  }

  req, _ := http.NewRequest("GET", url, nil)
  req.SetBasicAuth(p.username, p.password)

  resp, err := client.Do(req)
  if err != nil {
    updates <- update("error", err.Error())
  } else {
    defer resp.Body.Close()
    if(resp.StatusCode == 200) {
      updates <- update("success", fmt.Sprintf("updated domains %s", p.domains))
    } else {
      body, _ := ioutil.ReadAll(resp.Body)
      updates <- update("error", string(body))
    }
  }
}

func update(of, msg string) app.DDNSUpdates {
  return app.DDNSUpdates{
    of,
    "iwantmyname.com",
    msg,
  }
}
