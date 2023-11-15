package flags

import (
	"github.com/cthiel77/server-demo/internal/app"
	"github.com/cthiel77/server-demo/internal/log"
	"github.com/cthiel77/server-demo/internal/meta"
)

var initialized = false

// Init initialize flags
func Init() {
	//skip this, if all flags are already initialized
	if initialized {
		return
	}
	app.InitFlags()
	meta.InitFlags()
	log.InitFlags()

	initialized = true
}

// Process process flags
func Process() {}
