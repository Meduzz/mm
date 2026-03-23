package config

type (
	Config struct {
		Attributes []string `json:"attributes"`
		Models     string   `json:"models"`
	}

	Runtime struct {
		Port int `json:"port"`
		PID  int `json:"pid"`
	}
)
