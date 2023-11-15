package handler

import (
	_ "embed"

	"github.com/cthiel77/server-demo/internal"
	"github.com/gofiber/fiber/v2"
)

var (
	//go:embed page404.html
	page404Content string
)

// Page404 page handler
func Page404(c *fiber.Ctx) error {
	sdMap := initSiteData("404")
	sdMap["title"] = "404"

	sdMap["content_blocks"] = []internal.ContentBlock{
		{
			Title:   "Ups die Seite wurde nicht gefunden",
			SubHead: "",
			Content: page404Content,
			Link: internal.ContentLink{
				Label: "zurück zur Startseite",
				URL:   "/",
			},
		},
	}

	return c.Status(404).Render("web/templates/page", sdMap)
}
