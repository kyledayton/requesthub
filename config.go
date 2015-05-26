package main

import "flag"

type Config struct {
  MaxRequests int
  Port        int
  Username    string
  Password    string
}

func NewConfig() *Config {
  var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	var parsePort = flag.Int("p", DEFAULT_PORT, "which port to bind to")
	var parseUser = flag.String("username", "", "HTTP Basic Auth Username for accessing hub")
	var parsePass = flag.String("password", "", "HTTP Basic Auth Password for accessing hub")

  flag.Parse()

  return &Config{ *parseReq, *parsePort, *parseUser, *parsePass }
}

func (c *Config) AuthEnabled() bool {
  return c.Username != "" && c.Password != ""
}
