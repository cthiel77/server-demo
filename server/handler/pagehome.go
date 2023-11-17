package handler

import (
	_ "embed"
	"fmt"

	"github.com/cthiel77/server-demo/internal"
	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/gofiber/fiber/v2"
)

var (
	//go:embed pagehome.html
	pageHomeContent string
)

// PageHome page handler
func PageHome(c *fiber.Ctx) error {
	sdMap := initSiteData("")
	sdMap["title"] = "Server Demo"

	cntnt, _ := internal.RenderTemplate(pageHomeContent, map[string]interface{}{
		"apiURL":         fmt.Sprintf("%s/api/v1", c.BaseURL()),
		"swaggerURL":     fmt.Sprintf("%s/api/v1/swagger", c.BaseURL()),
		"devModeEnabled": cfg.DevModeEnabled(),
	})

	sdMap["content_blocks"] = []internal.ContentBlock{
		{
			Title:   "Server Demo",
			SubHead: "embedded Resources • app modes • log levels • fiber framework",
			Content: cntnt,
		},
	}

	return c.Render("web/templates/page", sdMap)
}
