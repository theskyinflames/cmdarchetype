package config

import "errors"

type (
	Config struct {
		SourceData         string             `mapstructure:"source-data"`
		ResultReceivers    []string           `mapstructure:"result-receivers"`
		DBConnectionParams DBConnectionParams `mapstructure:"db-connection-params"`
		DoAsynchronously   bool               `mapstructure:"do-async"`
	}

	DBConnectionParams struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBURL    string `mapstructure:"db-url"`
	}
)

func (c *Config) Validate() error {

	if len(c.SourceData) == 0 {
		return errors.New("'source-data' parameter can't be empty")
	}
	if len(c.ResultReceivers) == 0 {
		return errors.New("'result-receivers' parameter can't be empty")
	}

	return c.DBConnectionParams.Validate()
}

func (u DBConnectionParams) Validate() error {
	if len(u.User) == 0 {
		return errors.New("'user' flag can't be empty")
	}
	if len(u.Password) == 0 {
		return errors.New("'password' flag can't be empty")
	}
	if len(u.DBURL) == 0 {
		return errors.New("'db-url' flag can't be empty")
	}
	return nil
}
