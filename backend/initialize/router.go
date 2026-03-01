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
	"github.com/spf13/viper"
)

func initRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	e := echo.New()
	e.Validator = validator.InitCustomValidator()
	e.Use(middleware.NewLogger())
	e.Static("/static", "/app/uploads")
	// 不需要认证的路由
	public := e.Group("/public")
	// 需要认证的路由
	private := e.Group("/private", middleware.NewAuth())
	// 需要管理员权限的路由
	admin := e.Group("/admin", middleware.NewAuth(), middleware.NewAdmin())
	router.GroupRouterHubApp.InitRouterHub(public, private, admin)
	routerPort := viper.GetString("Router.port")
	server := &http.Server{
		Addr:    ":" + routerPort,
		Handler: e,
	}
	go func() {
		slog.Info("服务器正在启动，服务地址：http://localhost:" + routerPort)
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
