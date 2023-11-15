package internal

import "github.com/gofiber/fiber/v2"

// SetLocal sets local resources
func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}

// GetLocal rerturs local resources
func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}
