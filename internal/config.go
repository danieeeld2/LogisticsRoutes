package internal

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"strings"
	"fmt"
)

type AppConfig struct {
	App struct {
		Name string `koanf:"name"`
		Debug bool `koanf:"debug"`
	} `koanf:"app"`

	Log struct {
		Config struct {
			Enable struct {
				Console bool `koanf:"console"`
				File bool `koanf:"file"`
				Color bool `koanf:"color"`
			} `koanf:"enable"`
			
			Logfile struct {
				Path string `koanf:"path"`
			} `koanf:"logfile"`
		} `koanf:"config"`
	} `koanf:"log"`
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

	fmt.Println("name is = ", k)

	var appConfig AppConfig
	if err := k.Unmarshal("", &appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}

