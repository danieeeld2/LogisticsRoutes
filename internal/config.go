package internal

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/parsers/yaml"
)

var useConfigFile = false
var defaultConfig = true

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
	if(useConfigFile){
		if err := k.Load(file.Provider(confiFile), yaml.Parser()); err != nil {
			return nil, err
		}
	}
	if(defaultConfig){
		k.Set("app.Name", "LogisticsRoutes")
		k.Set("app.Debug", false)
		k.Set("log.config.enable.console", true)
		k.Set("log.config.enable.file", true)
		k.Set("log.config.logfile.path", "../logisticsroutes.log")
	}

	k.Print()

	var appConfig AppConfig
	if err := k.Unmarshal("", &appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}

