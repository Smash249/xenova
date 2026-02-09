package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/sirupsen/logrus"
)

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	cyan   = "\033[36m"
	blue   = "\033[34m"
	white  = "\033[37m"
)

func NewLogger() echo.MiddlewareFunc {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		DisableQuote:           true,
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx *echo.Context) error {
			start := time.Now()

			res, err := echo.UnwrapResponse(ctx.Response())
			if err != nil {
				return next(ctx)
			}

			res.After(func() {
				stop := time.Now()
				req := ctx.Request()
				latency := stop.Sub(start)
				timestamp := stop.Format(time.RFC3339)

				ip := ctx.RealIP()
				method := req.Method
				path := req.URL.Path
				query := req.URL.RawQuery
				status := res.Status
				size := res.Size
				userAgent := req.UserAgent()
				referer := req.Referer()

				statusColor := func(code int) string {
					switch {
					case code >= 500:
						return red
					case code >= 400:
						return yellow
					case code >= 300:
						return cyan
					default:
						return green
					}
				}

				methodColor := func(m string) string {
					switch m {
					case "GET":
						return blue
					case "POST":
						return green
					case "PUT":
						return yellow
					case "DELETE":
						return red
					case "PATCH":
						return cyan
					default:
						return white
					}
				}

				fullPath := path
				if query != "" {
					fullPath = path + "?" + query
				}

				msg := fmt.Sprintf("%s%3d%s | %-15s | %s%-6s%s | %-30s | %10v | %6d bytes | %s",
					statusColor(status), status, reset,
					ip,
					methodColor(method), method, reset,
					fullPath,
					latency,
					size,
					timestamp,
				)

				fields := logrus.Fields{
					"user_agent": userAgent,
				}
				if referer != "" {
					fields["referer"] = referer
				}

				switch {
				case status >= 500:
					log.WithFields(fields).Errorf("🚨 %s", msg)
				case status >= 400:
					log.WithFields(fields).Warnf("⚠️ %s", msg)
				case status >= 300:
					log.WithFields(fields).Infof("➡️ %s", msg)
				default:
					log.WithFields(fields).Infof("✅ %s", msg)
				}
			})
			return next(ctx)
		}
	}
}
