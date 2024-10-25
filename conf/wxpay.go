package conf

type WxPay struct {
	MchID        string `mapstructure:"mch_id" yaml:"mch_id"`                     // 商户号
	APIKey       string `mapstructure:"api_key" yaml:"api_key"`                   // API密钥
	NotifyURL    string `mapstructure:"notify_url" yaml:"notify_url"`             // 支付回调地址
	ReturnURL    string `mapstructure:"return_url" yaml:"return_url"`             // 支付完成返回地址
	QRServiceURL string `mapstructure:"qr_service_url" yaml:"qr_service_url"`     // 二维码生成服务URL
}
