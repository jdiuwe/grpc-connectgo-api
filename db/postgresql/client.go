package postgresql

import (
	"database/sql"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	_ "github.com/lib/pq"
)

const (
	connStr = "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
)

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func (config Config) Validate() error {
	err := validation.ValidateStruct(&config,
		validation.Field(&config.User, validation.Required, is.Alpha),
		validation.Field(&config.User, validation.Required, validation.Length(2, 100)),
		validation.Field(&config.Password, validation.Required, is.ASCII),
		validation.Field(&config.Password, validation.Required, validation.Length(2, 100)),
		validation.Field(&config.Name, validation.Required, is.Alphanumeric),
		validation.Field(&config.Name, validation.Required, validation.Length(2, 100)),
		validation.Field(&config.Host, validation.Required, is.Host),
		validation.Field(&config.Port, validation.Required, is.Port),
	)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

func NewPostgresqlClient(cfg *Config) (*sql.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("db configuration error: %w", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf(connStr, cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password))
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping error: %w", err)
	}

	return db, nil
}
