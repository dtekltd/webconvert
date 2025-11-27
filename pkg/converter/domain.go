package converter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	defaultQuality = 90
)

type ConvertReq struct {
	URL          string  `json:"url"`
	ApiKey       string  `json:"apiKey"`
	Format       string  `json:"format"`
	Size         string  `json:"size"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	MarginTop    float64 `json:"marginTop"`
	MarginBottom float64 `json:"marginBottom"`
	MarginLeft   float64 `json:"marginLeft"`
	MarginRight  float64 `json:"marginRight"`
	Quality      int     `json:"quality"`
	Landscape    bool    `json:"landscape"`
	FullPage     bool    `json:"fullpage"`
	Background   bool    `json:"background"`
}

func (r *ConvertReq) Parse(ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(r); err != nil {
		return err
	}
	if r.URL == "" {
		return fmt.Errorf("url is required")
	}

	// Set default values if not provided
	if r.Quality == 0 {
		r.Quality = defaultQuality
	}
	if r.Size != "custom" {
		r.Width, r.Height = getDefaultDimensions(r.Size, r.Landscape)
	}

	return nil
}

func getDefaultDimensions(pageSize string, landscape bool) (width, height int) {
	switch pageSize {
	case "a4":
		width, height = 1240, 1754 // A4 at 96dpi
	case "letter":
		width, height = 816, 1056 // Letter at 96dpi
	case "legal":
		width, height = 816, 1344 // Legal at 96dpi
	case "a3":
		width, height = 1754, 2480 // A3 at 96dpi
	default:
		width, height = 1024, 768
	}

	if landscape {
		width, height = height, width
	}
	return
}

// func (r *ConvertReq) parseDimensions() (width, height int, err error) {
// 	switch r.PageSize {
// 	case "a4":
// 		return 1240, 1754, nil // A4 at 96dpi
// 	case "letter":
// 		return 816, 1056, nil // Letter at 96dpi
// 	case "legal":
// 		return 816, 1344, nil // Legal at 96dpi
// 	case "a3":
// 		return 1754, 2480, nil // A3 at 96dpi
// 	case "custom":
// 		width = utils.IntWithDefault(r.Width, 1024)
// 		height = utils.IntWithDefault(r.Height, 768)
// 		if width < 100 || height < 100 {
// 			return 0, 0, fmt.Errorf("minimum dimensions are 100x100 pixels")
// 		}
// 		return width, height, nil
// 	default:
// 		return 1024, 768, nil
// 	}
// }
