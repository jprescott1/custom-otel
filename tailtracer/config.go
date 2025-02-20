package tailtracer

import (
	"fmt"
	"time"
)

type Config struct {
	Interval       string `mapstructure:"interval"`
	NumberOfTraces int    `mapstructure:"number_of_traces"`
}

func (cfg *Config) Validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Minutes() < 1 {
		return fmt.Errorf("when defined, the interval has to be set at least 1 minute (1m)")
	}

	if cfg.NumberOfTraces < 1 {
		return fmt.Errorf("number_of_traces must be greater than or equal to 1")
	}
	return nil
}
