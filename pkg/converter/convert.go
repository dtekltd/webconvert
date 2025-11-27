package converter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
	"github.com/tphan267/webconvert/pkg/utils"
)

func convertHandler(c *fiber.Ctx) error {
	req := &ConvertReq{}
	if err := req.Parse(c); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Create Chrome context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		emulation.SetDeviceMetricsOverride(int64(req.Width), int64(req.Height), 1.0, false),
		chromedp.Navigate(req.URL),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			switch req.Format {
			case "pdf":
				buf, _, err = page.PrintToPDF().
					WithPrintBackground(req.Background).
					WithPaperWidth(float64(req.Width) / 96). // Convert pixels to inches (96dpi)
					WithPaperHeight(float64(req.Height) / 96).
					WithMarginTop(req.MarginTop).WithMarginBottom(req.MarginBottom).
					WithMarginLeft(req.MarginLeft).WithMarginRight(req.MarginRight).
					WithLandscape(req.Landscape).
					Do(ctx)
			case "jpeg", "png":
				var qualityParam int64
				if req.Format == "jpeg" {
					qualityParam = int64(req.Quality)
				}
				buf, err = page.CaptureScreenshot().
					WithFormat(page.CaptureScreenshotFormat(req.Format)).
					WithQuality(qualityParam).
					WithCaptureBeyondViewport(req.FullPage).
					Do(ctx)
			default:
				return fmt.Errorf("unsupported output format: %s", req.Format)
			}
			return err
		}),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("Conversion failed: " + err.Error())
	}

	// Set response headers
	filename, _ := utils.UrlToFilename(req.URL)
	contentType, ext := getContentType(req.Format)

	c.Set("Content-Type", contentType)
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.%s", filename, ext))
	c.Set("Content-Length", strconv.Itoa(len(buf)))

	return c.Send(buf)
}

func getContentType(format string) (contentType, ext string) {
	switch format {
	case "pdf":
		return "application/pdf", "pdf"
	case "jpeg":
		return "image/jpeg", "jpg"
	case "png":
		return "image/png", "png"
	default:
		return "application/octet-stream", "bin"
	}
}
