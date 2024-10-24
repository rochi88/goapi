package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rochi88/goapi/app/helpers"
	"github.com/rochi88/goapi/app/resources"
)

func DefaultMiddleware(a *fiber.App) {
	// Custom File Writer
	file, err := os.OpenFile("./storage/logs/"+time.Now().Format("2006-01-01")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	a.Use(
		cors.New(),
		etag.New(),
		limiter.New(limiter.Config{
			Max:        20,
			Expiration: 15 * time.Minute,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).
					JSON(resources.ApiResponse{
						Success:    false,
						StatusCode: fiber.StatusTooManyRequests,
						Data:       "error",
						Message:    "Too many requests, please try again later",
					})
			},
			SkipSuccessfulRequests: true,
		}),
		helmet.New(),
		logger.New(logger.Config{
			Output:     file,
			Format:     "${time} ${method} ${status} ${path} in ${latency}\n",
			TimeFormat: "15:04:05.00",
			TimeZone:   helpers.GetEnv("APP_TIMEZONE", "Asia/Dhaka"),
		}),
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
		recover.New(recover.Config{
			EnableStackTrace: true,
		}),
		func(c *fiber.Ctx) error {
			// Custom middleware here
			return c.Next()
		},
	)
}
