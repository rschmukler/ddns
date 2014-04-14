package providers

import (
  "fmt"
  "http"
  "github.com/howeyc/gopass"
)

func init() {
  RegisterProvider("iwantmyname.com", &IWantMyNameProvider {
  })
}

type IWantMyNameProvider struct {
  username string
  password string
}

func (p *IWantMyNameProvider) GenerateConfig(config map[string]map[string]string) {
  var username string

  fmt.Printf("\tUsername: ")
  fmt.Scanln(&username)

  fmt.Printf("\tPassword: ")
  password := gopass.GetPasswd()

  myConfig := make(map[string]string)

  myConfig["username"] = username
  myConfig["password"] = string(password)
  config["iwantmyname.com"] = myConfig
}

func (p *IWantMyNameProvider) ReadConfig(config map[string]string) {
  p.username = config["username"];
  p.password = config["password"];
}

func (p *IWantMyNameProvider) Update(domain, ip string) {
  client := &http.Client

  url := fmt.Stringf("https://iwantmyname.com/basicauth/ddns?hostname=%s&myip=%s", domain, ip)
  req := http.NewRequest("GET", url)
  req.SetBasicAuth(p.username, p.password)

  resp := client.Do(req)
}
