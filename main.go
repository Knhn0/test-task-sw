package main

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"test-task-sw/config"
	"test-task-sw/lib/tctx"
	"time"
)

// @title Test task Swartway
// @version 1.0
// @description API Server for application

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	var (
		loggerSugar = logger.Sugar()
		settings    = config.Read()
	)

	mainCtx, mainCtxStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer mainCtxStop()

	app, err := NewApp(loggerSugar, settings, tctx.DefaultContextProvider(mainCtx))
	if err != nil {
		loggerSugar.Error(err.Error())
		mainCtxStop()
	} else {
		app.Run()
		loggerSugar.Debug("Successful started")
	}

	select {
	case <-mainCtx.Done():
	case <-app.Notify():
	}
	loggerSugar.Debug("Shutting down")

	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCtxCancel()

	app.Stop(shutdownCtx)

	<-shutdownCtx.Done()
	loggerSugar.Debug("Successful stopped")
}
