package app

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpServer HttpServerConfig
	Postgres   PostgresConfig
	Auth       AccessToken
	Swagger    SwaggerConfig
}

type HttpServerConfig struct {
	Port            int           `env:"HTTP_PORT" default:"8080"`
	ShutdownTimeout time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT" default:"1s"`
	ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT" default:"30s"`
	WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"10s"`
}

type AccessToken struct {
	Key      string        `env:"JWT_ACCESS_KEY" default:"stone"`
	Duration time.Duration `env:"JWT_ACCESS_DURATION" default:"30m"`
}

type SwaggerConfig struct {
	SwaggerHost string `env: "SWAGGER_HOST" default:"localhost:8080"`
}

type PostgresConfig struct {
	DatabaseName string `env:"DB_NAME" default:"postgres"`
	User         string `env:"DB_USER" default:"postgres"`
	Password     string `env:"DB_PASS" default:"postgres"`
	Host         string `env:"DB_HOST" default:"localhost"`
	Port         string `env:"DB_PORT" default:"5432"`
	PoolMinSize  string `env:"DB_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize  string `env:"DB_POOL_MAX_SIZE" default:"10"`
	SSLMode      string `env:"DB_SSL_MODE" default:"disable"`
}

func ReadConfigFromEnv() *Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error reading env")
	}

	return &cfg
}

func ReadConfigFromFile(filename string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(filename, &cfg)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error reading file")
	}

	return &cfg
}

func ReadConfig(filename string) *Config {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Warn().Msgf("File not found %s", filename)
		return ReadConfigFromEnv()
	}

	return ReadConfigFromFile(filename)
}

func (pg PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		pg.Host,
		pg.Port,
		pg.DatabaseName,
		pg.User,
		pg.Password,
		pg.SSLMode,
	)
}

func (pg PostgresConfig) URL() string {
	if pg.SSLMode == "" {
		pg.SSLMode = "disable"
	}

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pg.User, pg.Password, pg.Host, pg.Port, pg.DatabaseName, pg.SSLMode)

	return connectString
}
