package feishu_sdk

type Options func(c *Sdk)

func WithTenantProvider() Options {
	return func(c *Sdk) {
		c.provider = &TenantProvider{Sdk: c}
	}
}
