package conf

type Wechat struct {
	AppID          string `mapstructure:"app_id" yaml:"app_id"`                     //appid
	AppSecret      string `mapstructure:"app_secret" yaml:"app_secret"`             //appsecret
	Token          string `mapstructure:"token" yaml:"token"`                       //token
	EncodingAESKey string `mapstructure:"encoding_aes_key" yaml:"encoding_aes_key"` //EncodingAESKey
	BaseURL        string `mapstructure:"base_url" yaml:"base_url"`                 //base url
	RedirectURI    string `mapstructure:"redirect_uri" yaml:"redirect_uri"`         //redirect uri
}
