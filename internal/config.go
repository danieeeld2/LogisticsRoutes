package internal

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"strings"
)

type AppConfig struct {
	App struct {
		Name string `koanf:"name"`
		Debug bool `koanf:"debug"`
	} `koanf:"app"`

	LogConfig struct {
		EnableConsole 	bool   `koanf:"enable_console"`
		EnableFile    	bool   `koanf:"enable_file"`
		EnableColor   	bool   `koanf:"enable_color"`
		LogFilePath   	string `koanf:"log_file_path"`
	} `koanf:"log_config"`
}

func loadConfig(confiFile string) (*AppConfig, error) {
	k := koanf.New(".")
	if err := k.Load(file.Provider(confiFile), yaml.Parser()); err != nil {
		if err := k.Load(env.Provider("LogisticsRoutes_", ".", func(s string) string {
			return strings.Replace(strings.ToLower(
				strings.TrimPrefix(s, "LogisticsRoutes_")), "_", ".", -1)
		}), nil); err != nil {
			return nil, err
		}
	}

	var appConfig AppConfig
	if err := k.Unmarshal("", &appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}

