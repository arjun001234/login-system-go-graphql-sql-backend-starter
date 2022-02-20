package config

import "os"

type Config struct {
	DbName         string
	DbPassword     string
	DbUser         string
	Port           string
	SenderEmail    string
	JwtSecret      string
	SgApiKey       string
	GoogleClientId string
}

func NewConfig() *Config {
	var c Config
	if c.DbName = os.Getenv("DB_NAME"); c.DbName == "" {
		c.DbName = "login_system"
	}
	if c.DbUser = os.Getenv("DB_USER"); c.DbUser == "" {
		c.DbUser = "arjun"
	}
	if c.DbPassword = os.Getenv("DB_PASSWORD"); c.DbPassword == "" {
		c.DbPassword = "arjun"
	}
	if c.Port = os.Getenv("PORT"); c.Port == "" {
		c.Port = "4000"
	}
	if c.SenderEmail = os.Getenv("SENDER_EMAIL"); c.SenderEmail == "" {
		c.SenderEmail = "arjunkanojia001@outlook.com"
	}
	if c.JwtSecret = os.Getenv("JWT_SECRET"); c.JwtSecret == "" {
		c.JwtSecret = ""
	}
	if c.SgApiKey = os.Getenv("SG_API_KEY"); c.SgApiKey == "" {
		c.SgApiKey = ""
	}
	if c.GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID"); c.GoogleClientId == "" {
		c.GoogleClientId = ""
	}
	return &c
}
