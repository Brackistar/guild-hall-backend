package models

type ServerConfig struct {
	HostName string `json:"host"`
	Port     int    `json:"port"`
	HasSSL   bool   `json:"ssl"`
}
