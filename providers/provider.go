package providers

import (
  "github.com/rschmukler/ddns/app"
)

var providers map[string]Provider = make(map[string]Provider)

type Provider interface {
  GenerateConfig(config map[string]map[string]string)
  ReadConfig(config map[string]string)
  Update(ip string, updates chan app.DDNSUpdates)
}

func GetProvider(name string) (Provider, bool) {
  result, present := providers[name]
  return result, present
}

func RegisterProvider(name string, provider Provider) {
  providers[name] = provider;
}

func ListProviders() (string) {
  list := ""
  for name, _ := range providers {
    list += "\t" + name +"\n"
  }
  return list
}
