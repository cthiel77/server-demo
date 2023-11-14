package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/cthiel77/server-demo/config"
	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/cthiel77/server-demo/internal/meta"
	"github.com/cthiel77/server-demo/server/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

// RunAPIServer runs the server
func RunAPIServer(args []string) {

	engine := html.NewFileSystem(http.FS(config.TemplateFilesEmbedded), ".gohtml")
	engine.Layout("embed")    // Optional. Default: "embed"
	engine.Delims("{{", "}}") // Optional. Default: engine delimiters
	//extend template functions
	engine.AddFuncMap(map[string]any{
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
	})
	engine.AddFunc("unescape", func(s string) template.HTML {
		return template.HTML(s)
	})

	fbCfg := fiber.Config{
		Views:   engine,
		AppName: meta.App.Name,
	}

	//dev settings
	if cfg.DevModeEnabled() {
		engine.Reload(true) // Optional. Default: false
		engine.Debug(true)  // Optional. Default: false
	} else {
		fbCfg.Prefork = true
		fbCfg.DisableStartupMessage = true
	}

	app := fiber.New(fbCfg)

	// Middleware
	app.Use(recover.New())
	if cfg.DevModeEnabled() {
		app.Use(logger.New())
	}

	//▽
	//▽  static files
	//▽
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(config.StaticFilesEmbedded),
		PathPrefix: "web/static/assets",
		Browse:     false,
	}))

	//▽
	//▽  dynamic files
	//▽
	app.Get("/", handler.PageHome)

	// Handle not founds
	app.Use(handler.Page404)
	//▽
	//▽  run server
	//▽
	config.Logger.Fatal(app.Listen(fmt.Sprintf(":%s", config.App.Server.Port)))
}
