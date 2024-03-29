package config

import (
	"errors"
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config type
type Config struct {
	Session string
	Dayth   int
	Part    int
}

// Load func
func Load() (*Config, error) {

	flag.Int("day", 0, "day of puzzle, 1-25")
	flag.Int("part", 0, "1 or 2")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	config.Session = viper.GetString("COOKIE_SESSION")
	config.Dayth = viper.GetInt("day")
	config.Part = viper.GetInt("part")
	if (config.Dayth == 0 && config.Part == 0) || (config.Dayth >= 1 && config.Dayth <= 25 && config.Part >= 0 && config.Part <= 2) {
		return &config, nil
	}
	return nil, errors.New("invalid input, run application with --help flag for more information")
}
