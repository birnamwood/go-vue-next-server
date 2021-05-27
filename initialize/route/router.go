package route

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"go-vue-next-server/generate/wire"
	"go-vue-next-server/initialize/config"
	"go-vue-next-server/initialize/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

//Init router
func Init() {
	c := config.GetConfig()

	//依存性の注入
	userAccountHandler := wire.InitializeUserAccountHandler(database.GetDB())

	//FW echoの初期設定
	e := echo.New()
	//起動時にログにバナーを表示しない
	e.HideBanner = true
	e.HidePort = true
	// CORSの設定追加。下記のような形で設定
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//     AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	//     AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	e.Use(middleware.CORS())
	//echoのmiddleware リクエスト単位のログを出力
	filename := c.GetString("log.path") + "/http_log.%Y-%m-%d"
	rotate, err := rotatelogs.New(
		filename,
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}

	// Echoログの設定
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: rotate,
	}))
	// ログローテーション
	e.Logger.SetOutput(rotate)
	e.Use(middleware.BodyDump(bodyDumpHandler))
	//echoのmiddleware 予期せずpanic時、サーバは落とさずにエラーを返せるようにリカバリーする
	// e.Use(middleware.Recover())

	// Routing
	e1 := e.Group("/api")
	{
		e1.POST("/user-account/check", userAccountHandler.CheckUserAccount)
		e1.GET("/check", userAccountHandler.GetTest)
	}

	//App Group
	a1 := e.Group("/api/app/v1")
	{
		a1.Use(systemAuth) //認証必須
	}

	//e.start(ポート番号)でサーバースタート
	zap.S().Info("== Server Srart == Port:" + c.GetString("server.port") + " Pid:" + fmt.Sprint(os.Getpid()))

	e.Server.Addr = ":" + c.GetString("server.port")
	// gracefulシャットダウンのためgracehttpでシャットダウン命令を監視
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	c.Echo().Logger.Print("IPアドレス:", net.ParseIP(c.RealIP()), " Request Body: %v\n", string(reqBody))
	c.Echo().Logger.Print("IPアドレス:", net.ParseIP(c.RealIP()), " Response Body: %v\n", string(resBody))
}
