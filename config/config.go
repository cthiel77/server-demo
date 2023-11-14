package config

import (
	"embed"
	"encoding/json"

	"github.com/spf13/pflag"

	"github.com/cthiel77/server-demo/internal/cfg/loader"
	"github.com/cthiel77/server-demo/internal/cli/flags"
	defaultError "github.com/cthiel77/server-demo/internal/error"
	"github.com/cthiel77/server-demo/internal/log"
	"github.com/cthiel77/server-demo/internal/meta"
	metaModel "github.com/cthiel77/server-demo/internal/meta/model"
	"github.com/sirupsen/logrus"
)

// Config application config options
type Config struct {
	Meta   metaModel.AppMeta `json:"-"`
	Server ServerCfg         `json:"server"`
}

var (
	// Logger the app logger
	Logger *logrus.Logger

	//go:embed config.json
	configJSON []byte

	//StaticFilesEmbedded embed the static asset files
	//go:embed web/static/*
	StaticFilesEmbedded embed.FS
	//TemplateFilesEmbedded embed the template files
	//go:embed web/templates/*
	TemplateFilesEmbedded embed.FS

	//CmdNameWebServer command name for starting the api
	CmdNameWebServer = "runApi"

	//App init app config from config struct
	App = Config{
		Meta: meta.App,
	}
)

// Init the init method
func Init() (flagsProcessed bool) {
	flags.Init()
	pflag.Parse()

	flagsProcessed = meta.ProcessFlags()

	Logger = log.NewLogger(App.Meta)
	// load config file if flag is set. Override embedded default config data
	c, loadErr := loader.LoadConfigData()
	if loadErr != nil {
		if loadErr.Code != defaultError.StatusConfigFileNameNotSet {
			Logger.Error(loadErr)
			return
		}
	} else {
		configJSON = *c
	}
	//create app config from json payload
	if e := json.Unmarshal(configJSON, &App); e != nil {
		Logger.Error(e)
	}
	Logger.Info("loaded config data")

	return
}

// ServerCfg the server configuration
type ServerCfg struct {
	Port string `json:"port"`
}
