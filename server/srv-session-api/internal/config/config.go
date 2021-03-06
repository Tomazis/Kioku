package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	ServiceName string `yaml:"serviceName"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
	ImportDB    bool `yaml:"importDB"`
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	Host              string `yaml:"host"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Timeout           int64  `yaml:"timeout"`
}

type GrpcDBA struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Timeout int64  `yaml:"timeout"`
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

// Jaeger - contains all parameters metrics information.
type Jaeger struct {
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

// Status config for service.
type Status struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	VersionPath   string `yaml:"versionPath"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
}

// Telemetry config for service
type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project   Project   `yaml:"project"`
	Grpc      Grpc      `yaml:"grpc"`
	Rest      Rest      `yaml:"rest"`
	Metrics   Metrics   `yaml:"metrics"`
	Jaeger    Jaeger    `yaml:"jaeger"`
	Status    Status    `yaml:"status"`
	Telemetry Telemetry `yaml:"telemetry"`
	GrpcDBA   GrpcDBA   `yaml:"grpc_dba"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
