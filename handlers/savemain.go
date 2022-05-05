package handlers

import (
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"path"
	"sign-builder/core"
)

var largeHeight = 50
var smallHeight = 20

func HandleShieldPostQuery(c *fiber.Ctx) error {
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

	smallImage := imaging.Resize(*img, 0, smallHeight, imaging.Linear)
	smallFileName := path.Join(os.TempDir(), pattern+".png")
	smallKey := path.Join("Shields", "20x", pattern+".png")
	err = imaging.Save(smallImage, smallFileName)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.WriteString("Count Not Save Shield")
		return err
	}

	err = core.UploadS3(smallKey, smallFileName)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.WriteString("Count Not Upload Shield")
		return err
	}

	largeImage := imaging.Resize(*img, 0, largeHeight, imaging.Linear)
	largeFileName := path.Join(os.TempDir(), pattern+".png")
	largeKey := path.Join("Shields", pattern+".png")
	err = imaging.Save(largeImage, largeFileName)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.WriteString("Count Not Resize Shield")
		return err
	}
	err = core.UploadS3(largeKey, largeFileName)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.WriteString("Count Not Upload Resized Shield")
		return err
	}

	c.Status(200)
	return nil
}
