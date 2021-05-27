package database

//
import (
	"go-vue-next-server/initialize/config"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

//Init database接続
func Init() {
	c := config.GetConfig()
	dbLogger := newDBLogger(c)
	//configからデータベースのプロバイダとパスを取得しOpenする
	var err error
	db, err = gorm.Open(postgres.Open(getDsn(c)), &gorm.Config{
		Logger: dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // trueで`Users` のテーブルは `user` になります。
		},
	})
	if err != nil {
		panic("データベースへの接続失敗")
	}

	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		panic("データベースへの接続失敗")
	}
	pingErr := sqlDB.Ping()
	if pingErr != nil {
		panic("データベースへの接続失敗")
	}
	sqlDB.SetMaxIdleConns(c.GetInt("max-idle-connection"))                //待機コネクション数
	sqlDB.SetMaxOpenConns(c.GetInt("max-open-connection"))                //接続コネクション数
	sqlDB.SetConnMaxLifetime(time.Second * c.GetDuration("max-lifetime")) //コネクションを再利用する時間
}

//GetDB return db connection pool
func GetDB() *gorm.DB {
	return db
}

// Close Database connection
func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Error("データベースのCloseに失敗しました。")
	}
	if err = sqlDB.Close(); err != nil {
		zap.S().Error("データベースのCloseに失敗しました。")
	}
}

func getDsn(c *viper.Viper) string {
	//configの内容取得
	dsn := "host=" + c.GetString("db.host") +
		" port=" + c.GetString("db.port") +
		" dbname=" + c.GetString("db.dbname") +
		" user=" + c.GetString("db.user") +
		" password=" + c.GetString("db.password")
	return dsn
}

// newDBLogger O/R マッパー用ロガー　プロダクション以外はInfoレベル以上全て出力
func newDBLogger(c *viper.Viper) logger.Interface {
	dblog := log.New(os.Stdout, "", log.LstdFlags)

	environment := c.GetString("environment")

	if environment != "production" {
		return logger.New(
			dblog,
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // color
			},
		)
	}
	return logger.New(
		dblog,
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // color
		},
	)

}
