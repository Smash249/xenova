package initialize

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
)

func Initialize() {
	ch := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signal.Notify(ch, os.Interrupt)
	initViper()
	initGorm()
	go func() {
		<-ch
		cancel()
		slog.Info("系统退出")
	}()
	<-ctx.Done()
}
