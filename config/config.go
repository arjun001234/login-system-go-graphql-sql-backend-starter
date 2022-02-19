package config

import "os"

type Config struct {
	DbName      string
	DbPassword  string
	DbUser      string
	Port        string
	SenderEmail string
	JwtSecret   string
	SgApiKey    string
}

func NewConfig() *Config {
	var c Config
	if c.DbName = os.Getenv("DB_NAME"); c.DbName == "" {
		c.DbName = ""
	}
	if c.DbUser = os.Getenv("DB_USER"); c.DbUser == "" {
		c.DbUser = ""
	}
	if c.DbPassword = os.Getenv("DB_PASSWORD"); c.DbPassword == "" {
		c.DbPassword = ""
	}
	if c.Port = os.Getenv("PORT"); c.Port == "" {
		c.Port = ""
	}
	if c.SenderEmail = os.Getenv("SENDER_EMAIL"); c.SenderEmail == "" {
		c.SenderEmail = ""
	}
	if c.JwtSecret = os.Getenv("JWT_SECRET"); c.JwtSecret == "" {
		c.JwtSecret = ""
	}
	if c.SgApiKey = os.Getenv("SG_API_KEY"); c.SgApiKey == "" {
		c.SgApiKey = ""
	}
	return &c
}
