package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
)

func captureAndEncodeScreenshot(url string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set a timeout for the screenshot (30 seconds)
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.CaptureScreenshot(&buf),
	})

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func main() {
	r := gin.Default()

	r.GET("/cap-image", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'url' parameter"})
			return
		}

		imageBytes, err := captureAndEncodeScreenshot(url)
		if err != nil {
			log.Printf("Error capturing screenshot: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Set response headers for an image in JPEG format
		c.Header("Content-Type", "image/jpeg")
		c.Header("Content-Disposition", "inline; filename=screenshot.jpg")

		// Respond with the image data
		c.Data(http.StatusOK, "image/jpeg", imageBytes)
	})

	r.Run(":8081")
}
