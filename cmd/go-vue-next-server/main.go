package main

import (
	"flag"

	"go-vue-next-server/initialize/alert"
	"go-vue-next-server/initialize/config"
	"go-vue-next-server/initialize/database"
	"go-vue-next-server/initialize/logger"
	"go-vue-next-server/initialize/route"

	"go.uber.org/zap"
)

func main() {
	// 環境設定取得 flag.String(<パラメータ名>, <デフォルト値>, <パラメータの説明>)
	env := flag.String("e", "local", "動作環境名")
	//変数宣言のあとに、flag.Parseを実行することでコマンドラインのパラメータがパースされ、各変数に値が格納されます
	flag.Parse()
	//パラメータを渡してconfigの初期化を行う
	config.Init(*env, "environment")

	//Loggerの初期化
	logger := logger.Init("app_log")
	zap.ReplaceGlobals(logger)
	logger.Info("Logger Initialize")
	alert.Init()

	//Database初期化
	database.Init()
	logger.Info("DB Initialize")

	route.Init()
	database.Close()
}
