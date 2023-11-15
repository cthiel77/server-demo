package handler

import (
	_ "embed"

	"github.com/cthiel77/server-demo/config"
	"github.com/cthiel77/server-demo/internal"
	"github.com/cthiel77/server-demo/internal/db"
	"github.com/gofiber/fiber/v2"
)

// ApiHeroGetByID func gets hero by given ID or 404 error.
// @Description Get hero by given ID.
// @Summary get hero by given ID
// @Tags Hero
// @Accept json
// @Produce json
// @Param id path int true "Hero ID" default(2)
// @Success 200 {object} db.HeroSet
// @Router /v1/hero/{id} [get]
func ApiHeroGetByID(c *fiber.Ctx) error {

	querystack := internal.GetLocal[db.HeroQueryStack](c, "db")
	id, _ := c.ParamsInt("id", 0)
	hero, err := querystack.GetByID(uint64(id))
	if err != nil {
		config.Logger.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "hero with the given ID is not found",
		})
	}
	return c.JSON(*hero)
}

// ApiHeroGetAll func gets all defined heroes or 404 error.
// @Description Get all heroes.
// @Summary get all heroes
// @Tags Hero
// @Accept json
// @Produce json
// @Success 200 {object} []db.HeroSet
// @Router /v1/hero [get]
func ApiHeroGetAll(c *fiber.Ctx) error {

	querystack := internal.GetLocal[db.HeroQueryStack](c, "db")

	list, err := querystack.GetAll()
	if err != nil {
		config.Logger.Error(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "hero with the given ID is not found",
		})
	}
	return c.JSON(list)
}
