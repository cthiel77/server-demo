package handler

import (
	_ "embed"
	"fmt"

	"github.com/cthiel77/server-demo/internal"
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

	cntnt := internal.RenderTemplate(pageHomeContent, map[string]interface{}{
		"api_url":     fmt.Sprintf("%s/api/v1", c.BaseURL()),
		"swagger_url": fmt.Sprintf("%s/api/v1/swagger", c.BaseURL()),
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
