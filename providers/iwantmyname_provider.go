package providers

func init() {
  RegisterProvider("iwantmyname.com", &IWantMyNameProvider {
  })
}

type IWantMyNameProvider struct {
  username string
  password string
}

func (p *IWantMyNameProvider) LoadConfig(config map[string]string) {
}

func (p *IWantMyNameProvider) Update() {
}
