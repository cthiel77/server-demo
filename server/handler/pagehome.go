package handler

import (
	_ "embed"

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

	sdMap["content_blocks"] = []internal.ContentBlock{
		{
			Title:   "Server Demo",
			SubHead: "embedded Resources • app modes • log levels • fiber framework",
			Content: pageHomeContent,
		},
	}

	return c.Render("web/templates/page", sdMap)
}
