package main

import "fmt"

type Config struct {
	host string
	port int
}

func (c *Config) SetHost(host string) *Config {
	c.host = host
	return c
}

func (c *Config) SetPort(port int) *Config {
	c.port = port
	return c
}

func (c *Config) Build() {
	fmt.Printf("Config: host=%s, port=%d\n", c.host, c.port)
}

func main() {
	config := &Config{}
	config.SetHost("localhost").SetPort(8080).Build()
}
