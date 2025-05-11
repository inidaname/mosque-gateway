package types

import "time"

type Config struct {
	Addr    string
	Env     string
	ApiURL  string
	Auth    AuthConfig
	Mailer  MailConfig
	Dospace DigitalOceanSpace
	// DbURL        string
	// SandBoxDbURL string
}

type AuthConfig struct {
	Basic BasicConfig
	Token TokenConfig
}

type BasicConfig struct {
	User string
	Pass string
}

type TokenConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
}

type MailConfig struct {
	FromEmail string
	Resend    ResendConfig
}

type ResendConfig struct {
	ApiKey string
}

type DigitalOceanSpace struct {
	AccessKey string
	SecretKey string
	SpaceName string
	Endpoint  string
}
