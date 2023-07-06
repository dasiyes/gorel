package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Title     string        `yaml:"title"`
	ServerCfg ServiceConfig `yaml:"http_srv_cfg"`
}

type Config interface {
	GetTitle() string
	GetSrvCfg() *ServiceConfig
}

type ServiceConfig struct {
	Name              string        `yaml:"name"`
	Port              int           `yaml:"port"`
	Host              []string      `yaml:"host"`
	TLS               *TLS          `yaml:"tls"`
	Timeout           time.Duration `yaml:"timeout"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
	Version           string        `yaml:"version"`
}

type TLS struct {
	IsDisabled               bool     `yaml:"disabled"`
	PublicKey                string   `yaml:"public_key"`
	PrivateKey               string   `yaml:"private_key"`
	MinVersion               string   `yaml:"min_version"`
	MaxVersion               string   `yaml:"max_version"`
	CurvePreferences         []uint16 `yaml:"curve_preferences"`
	PreferServerCipherSuites bool     `yaml:"prefer_server_cipher_suites"`
	CipherSuites             []uint16 `yaml:"cipher_suites"`
	EnableMTLS               bool     `yaml:"enable_mtls"`
}

func InitConfig(conf string) (Config, error) {

	c := config{}

	err := c.loadConfig(conf)
	if err != nil {
		log.Printf("Error loading config file: %s", err)
		return nil, err
	}

	return &c, nil
}

func (c *config) loadConfig(cf string) error {

	confContent, err := os.ReadFile(cf)
	if err != nil {
		return err
	}

	confContent = []byte(os.ExpandEnv(string(confContent)))

	if err := yaml.Unmarshal(confContent, &c); err != nil {
		return err
	}

	log.Printf("port: %v", c.ServerCfg.Port)
	if c.ServerCfg.Port == 0 {
		c.ServerCfg.Port = 8080
	}

	return nil
}

func (c *config) GetTitle() string {
	return c.Title
}

func (c *config) GetSrvCfg() *ServiceConfig {
	return &c.ServerCfg
}
