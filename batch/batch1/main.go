package main

// go build -o batch1 ./batch/batch1/*.go

import (
	"flag"
	"go-vue-next-server/initialize/alert"
	"go-vue-next-server/initialize/config"
	"go-vue-next-server/initialize/database"
	"go-vue-next-server/initialize/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	env := flag.String("e", "local", "動作環境名")
	flag.Parse()
	config.Init(*env, "environment")

	logger := logger.Init("batch_log")
	zap.ReplaceGlobals(logger)
	alert.Init()
	database.Init()

	zap.S().Info("プログラム起動")
	stop := signalHandler()
	initializeRoop(stop)
}

// graceful shutdownのためシグナルを監視
func signalHandler() <-chan struct{} {
	stop := make(chan struct{}, 0)
	go func() {
		quit := make(chan os.Signal, 2)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		close(stop)
		<-quit
		os.Exit(1)
	}()
	return stop
}

// 終了指示が来るまでループ処理を行う
func initializeRoop(ch <-chan struct{}) {
	for {
		select {
		case <-ch:
			zap.S().Info("プログラム終了")
			return
		default:
			mainLoop()
			print("10秒待機\n")
			time.Sleep(time.Second * 10)
		}
	}
}

//mainLoop notificationテーブルからPush通知対象を取得
func mainLoop() {
	zap.S().Info("バッチメインループ")
}
