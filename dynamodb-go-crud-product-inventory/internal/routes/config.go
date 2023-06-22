package routes

import (
	"net/http"
	"time"

	
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}
func (c *Config) Cors(next http.Handler) http.Handler {
	//cors.new(cors.Options{
	return cors.new(cors.Options)({
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"*"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		AllowedCredentials: true,
		MaxAge:             5,
	}).Handler(next)
}
func (c *Config) SetTimeout(timeInSeconds int) *Config {
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c
}
func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
