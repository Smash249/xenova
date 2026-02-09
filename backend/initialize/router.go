package initialize

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Smash249/xenova/backend/pkg/middleware"
	"github.com/Smash249/xenova/backend/pkg/validator"
	"github.com/Smash249/xenova/backend/router"
	"github.com/labstack/echo/v5"
)

func initRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	e := echo.New()
	e.Validator = validator.InitCustomValidator()
	e.Use(middleware.NewLogger())
	public := e.Group("/public")
	private := e.Group("/private", middleware.NewAuth())
	router.GroupRouterHubApp.InitRouterHub(public, private)
	server := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	go func() {
		slog.Info("服务器正在启动，服务地址：http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("服务器启动失败", "error", err)
			stop()
		}
	}()
	<-ctx.Done()
	slog.Info("接收到停止信号，正在关闭服务器...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("服务器关闭失败", "error", err)
	} else {
		slog.Info("系统已安全退出")
	}
}
