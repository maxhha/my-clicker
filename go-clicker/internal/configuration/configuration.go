package configuration

type Configuartion struct {
	addr       string
	server_url string
}

func (c *Configuartion) Addr() string {
	return c.addr
}

func (c *Configuartion) ServerUrl() string {
	return c.server_url
}

type ConfigurationBuilder struct {
	addr       string
	server_url string
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	return &ConfigurationBuilder{}
}

func (b *ConfigurationBuilder) SetAddr(addr string) *ConfigurationBuilder {
	b.addr = addr
	return b
}

func (b *ConfigurationBuilder) SetServerUrl(url string) *ConfigurationBuilder {
	b.server_url = url
	return b
}

func (b *ConfigurationBuilder) Build() Configuartion {
	var c = Configuartion{
		addr:       ":3000",
		server_url: "http://localhost:3000",
	}

	if len(b.addr) > 0 {
		c.addr = b.addr
	}

	if len(b.server_url) > 0 {
		c.server_url = b.server_url
	}

	return c
}
