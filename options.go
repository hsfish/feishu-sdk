package feishu_sdk

type Options func(c *sdk)

func WithTenantProvider() Options {
	return func(c *sdk) {
		c.provider = &TenantProvider{sdk: c}
	}
}
