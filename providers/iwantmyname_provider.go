package providers

func init() {
  RegisterProvider("iwantmyname.com", &IWantMyNameProvider {
  })
}

type IWantMyNameProvider struct {
}

func (p *IWantMyNameProvider) Update(config map[string]string) {
}
