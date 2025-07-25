package main

import (
	"encoding/json"
	"errors"
	"github.com/kelseyhightower/envconfig"
	errorspkg "github.com/pkg/errors"
	"net/url"
	"os"
)

type Config struct {
	TokenIntrospectionUrl          string `envconfig:"TOKEN_INTROSPECTION_URL" json:"tokenIntrospectionUrl"`
	TokenIntrospectionClientId     string `envconfig:"TOKEN_INTROSPECTION_CLIENT_ID" json:"tokenIntrospectionClientId"`
	TokenIntrospectionClientSecret string `envconfig:"TOKEN_INTROSPECTION_CLIENT_SECRET" json:"tokenIntrospectionClientSecret"`
}

func (cfg *Config) FromFile() error {
	f, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) FromEnv() error {
	return envconfig.Process("", cfg)
}

func (cfg *Config) Validate() error {
	logger.Info("validating config")

	var errs []error

	if cfg.TokenIntrospectionUrl == "" {
		errs = append(errs, errors.New("no token introspection endpoint specified in configuration"))
	} else if _, err := url.ParseRequestURI(cfg.TokenIntrospectionUrl); err != nil {
		errs = append(errs, errorspkg.Wrap(err, "the given token introspection url is not a valid url"))
	}

	if cfg.TokenIntrospectionClientId == "" {
		errs = append(errs, errors.New("no client id to use as auth for token introspection"))
	}
	if cfg.TokenIntrospectionClientSecret == "" {
		errs = append(errs, errors.New("no client secret to use as auth for token introspection"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
