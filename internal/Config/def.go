package Config

type Config struct {
	RetrieveURL string `json:"retrieveurl"`
	UpdateURL   string `json:"updateurl"`
	BasicAuth   bool   `json:"basicauth"`
	Token       string `json:"token"`

	// For basic auth. Name of parameter we're using with the token.
	KeyParam string `json:"keyparam"`
}

func (cfg *Config) SetDefaults() {
	cfg.KeyParam = "key"

	cfg.RetrieveURL = "https://mydomain.example/servers?sort=laststatupdate"
	cfg.UpdateURL = "https://mydomain.example/servers/{id}"
}
