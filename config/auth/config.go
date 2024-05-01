package auth

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Log       `yaml:"logger"`
		JWT       `yaml:"jwt"`
		Transport `yaml:"transport"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// JWT -.
	JWT struct {
		SecretKey       string `mapstructure:"secret_key" yaml:"secret_key"`
		AccessTokenTTL  int64  `mapstructure:"access_token_ttl" yaml:"access_token_ttl"`
		RefreshTokenTTL int64  `mapstructure:"refresh_token_ttl" yaml:"refresh_token_ttl"`
	}
	Transport struct {
		User     UserTransport     `yaml:"user"`
		UserGrpc UserGrpcTransport `yaml:"userGrpc"`
	}
	UserTransport struct {
		Host    string        `yaml:"host" env:"USER_TRANSPORT_URL"`
		Timeout time.Duration `yaml:"timeout"`
	}
	UserGrpcTransport struct {
		Host string `env:"USER_GRPC_URL"`
	}
)

// NewConfig returns user config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config/auth/config.yml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
