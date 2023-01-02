package config

type Config struct {
	Host     string            `json:"host"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Sites    map[string]string `json:"sites"`
}
