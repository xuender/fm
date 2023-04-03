package pb

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

const ConfigFile = ".fm.toml"

func NewConfig() *Config {
	cfg := &Config{}

	if file, err := os.Open(ConfigFile); err == nil {
		defer file.Close()

		_ = toml.NewDecoder(file).Decode(cfg)
	}

	if !lo.Contains(cfg.Ignore, ConfigFile) {
		cfg.Ignore = append(cfg.Ignore, ConfigFile)
	}

	if cfg.Dirs == nil {
		cfg.Dirs = map[string]string{}
	}

	logs.D.Println(cfg)

	return cfg
}

func (p *Config) IsIgnore(path string) bool {
	for _, ignore := range p.Ignore {
		if path == ignore {
			return true
		}
	}

	return false
}
