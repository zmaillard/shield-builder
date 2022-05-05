package handlers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"image/png"
	"net/http"
	"sign-builder/core"
)

func HandleShieldQuery(c *fiber.Ctx) error {
	pattern := c.Query("shield")
	if len(pattern) == 0 {
		c.Status(http.StatusBadRequest)
		c.WriteString("Count Not Find Parameter Shield")
		return nil
	}

	img, err := core.Build(pattern)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.WriteString("Count Not Build Shield")
		return err
	}

	var buff bytes.Buffer
	png.Encode(&buff, *img)

	c.Status(200)
	c.Type("image/png")
	c.Write(buff.Bytes())
	return nil
}
