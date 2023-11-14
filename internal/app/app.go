package app

import (
	//_ use native embed
	_ "embed"
	"fmt"

	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/cthiel77/server-demo/internal/log"
	"github.com/cthiel77/server-demo/internal/meta"
	flag "github.com/spf13/pflag"
)

var (
	//go:embed usage.txt
	flagModeUsage string
)

// InitFlags initializes application flags
func InitFlags() {
	flag.StringVarP(&cfg.Mode, "mode", "m", cfg.ModeDefault, flagModeUsage)
}

// PrintConfig prints the config
func PrintConfig() {
	fmt.Printf("AppName: %v\n", meta.GetAppName())
	fmt.Printf("AppMode: %v\n", cfg.Mode)
	fmt.Printf("logLevel: %v\n", log.LogLevel)
	fmt.Printf("TimestampFormat: %v\n", log.TimestampFormat)
	fmt.Printf("DisableTimestamp: %v\n", log.DisableTimestamp)
	fmt.Printf("DisableColors: %v\n", log.DisableColors)
	fmt.Printf("FullTimestamp: %v\n", log.FullTimestamp)
}
