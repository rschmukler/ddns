package providers

import (
  "fmt"
  "net/http"
  "github.com/howeyc/gopass"
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

  var username, password, domains, readUsername, readDomains string
  var myConfig map[string]string
  var present bool

  myConfig, present = config["iwantmyname.com"]

  if !present {
    myConfig = make(map[string]string)
  }

  domains, present = myConfig["domains"]
  if present {
    fmt.Printf("\tDomains: (%s) ", domains)
  } else {
    fmt.Printf("\tDomains: ")
  }
  fmt.Scanln(&readDomains)
  if len(readDomains) > 0 {
    domains = readDomains
  }


  username, present = myConfig["username"]
  if present {
    fmt.Printf("\tUsername: (%s) ", username)
  } else {
    fmt.Printf("\tUsername: ")
  }
  fmt.Scanln(&readUsername)

  if len(readUsername) > 0 {
    username = readUsername
  }

  password, present = myConfig["password"]

  if present  {
    fmt.Printf("\tPassword: (enter to keep) ")
  } else {
    fmt.Printf("\tPassword: ")
  }
  readPassword := gopass.GetPasswd()

  if(len(readPassword) > 0) {
    password = string(readPassword)
  }

  myConfig["domains"] = domains
  myConfig["username"] = username
  myConfig["password"] = string(password)
  config["iwantmyname.com"] = myConfig
}

func (p *IWantMyNameProvider) ReadConfig(config map[string]string) {
  p.username = config["username"];
  p.password = config["password"];
  p.domains = config["domains"];
}

func (p *IWantMyNameProvider) Update(ip string) {
  client := &http.Client{}

  url := fmt.Sprintf("https://iwantmyname.com/basicauth/ddns?hostname=%s", p.domains)
  if len(ip) > 0 {
    url += "&myip=" + ip
  }

  req, _ := http.NewRequest("GET", url, nil)
  req.SetBasicAuth(p.username, p.password)

  client.Do(req)
}
