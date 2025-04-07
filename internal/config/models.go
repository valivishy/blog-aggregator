package config

import "fmt"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config Config) String() string {
	return fmt.Sprintf("{DbUrl: %s, CurrentUserName: %s}", config.DbUrl, config.CurrentUserName)
}
