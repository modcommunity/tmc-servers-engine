package Config

type Config struct {
	Debug uint `json:"debug"`

	RetrieveEnginesURL string `json:"reurl"`
	RetrieveURL        string `json:"retrieveurl"`
	UpdateURL          string `json:"updateurl"`
	PostHook           string `json:"posthook"`

	BasicAuth bool   `json:"basicauth"`
	Token     string `json:"token"`

	// For basic auth. Name of parameter we're using with the token.
	KeyParam string `json:"keyparam"`

	// For verifying a server.
	ClaimKeyField string `json:"claimkeyfield"`

	// Server retrieving.
	MaxServers       uint   `json:"maxservers"`
	MaxServersPerReq uint   `json:"maxserverspr"`
	Sort             string `json:"sort"`

	// Query settings.
	WaitInterval     uint `json:"waitinterval"`
	FetchInterval    uint `json:"fetchinterval"`
	PostHookInterval uint `json:"posthookinterval"`
}

func (cfg *Config) SetDefaults() {
	cfg.KeyParam = "key"
	cfg.ClaimKeyField = "claimkey"

	// Both same value for now, but still want them to be different config options in case.
	cfg.RetrieveEnginesURL = "https://mydomain.example/servers/engines"
	cfg.RetrieveURL = "https://mydomain.example/servers/servers"
	cfg.UpdateURL = "https://mydomain.example/servers/servers"
	cfg.PostHook = "https://mydomain.example/servers/stats"

	// Retrieving servers from API settings.
	cfg.MaxServers = 40
	cfg.MaxServersPerReq = 40
	cfg.Sort = "laststatupdate ASC"

	// Intervals.
	cfg.WaitInterval = 1000    // Milliseconds.
	cfg.FetchInterval = 1000   // Milliseconds.
	cfg.PostHookInterval = 300 // Seconds.
}
