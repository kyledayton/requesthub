package main

import(
  "flag"
  "log"

  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Config struct {
  MaxRequests int
  Port        int
  Username    string
  Password    string
  YamlConfigFile  string
}

type yamlHubConfig struct {
    ForwardURL string `yaml:"forward_url,omitempty"`
}

type yamlConfig struct {
  Hubs map[string]yamlHubConfig `yaml:"hubs"`
}

func NewConfig() *Config {
  var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	var parsePort = flag.Int("p", DEFAULT_PORT, "which port to bind to")
	var parseUser = flag.String("username", "", "HTTP Basic Auth Username for accessing hub")
	var parsePass = flag.String("password", "", "HTTP Basic Auth Password for accessing hub")
  var configFile = flag.String("config", "", "YAML Configuration File")

  flag.Parse()

  return &Config{ *parseReq, *parsePort, *parseUser, *parsePass, *configFile }
}

func (c *Config) AuthEnabled() bool {
  return c.Username != "" && c.Password != ""
}

func (c *Config) HasYAMLConfig() bool {
  return c.YamlConfigFile != ""
}

func (c *Config) ApplyYAMLConfig(db *HubDatabase) error {
  config, err := ioutil.ReadFile(c.YamlConfigFile)

  if err != nil {
    return err
  }

  var cfg yamlConfig
  err = yaml.Unmarshal(config, &cfg)

  if err != nil {
    return err
  }

  for id, val := range cfg.Hubs {

    if db.Get(id) == nil {
      log.Printf("Created %s from config file\n", id)
      hub, _ := db.Create(id)

      if val.ForwardURL != "" {
        hub.ForwardURL = val.ForwardURL
      }
    }
  }

  return nil
}
