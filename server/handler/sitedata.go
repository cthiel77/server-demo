package handler

import (
	"github.com/cthiel77/server-demo/internal/cfg"
	"github.com/cthiel77/server-demo/internal/log"
	"github.com/cthiel77/server-demo/internal/meta"
	"github.com/gofiber/fiber/v2"
)

func initSiteData(currentPageKey string) (data fiber.Map) {
	data = fiber.Map{}
	data["author_name"] = meta.GetAuthorName()
	data["author_email"] = meta.GetAuthorEmail()
	data["company"] = meta.GetAuthorCompany()
	data["page_menu"] = getNavigationData()
	data["current_page"] = currentPageKey
	data["version"] = meta.GetVersionID()
	data["buildTime"] = meta.GetBuildTime()
	data["appMode"] = cfg.GetAppModeKey()
	data["logLevel"] = log.GetLogLevel()
	return
}
