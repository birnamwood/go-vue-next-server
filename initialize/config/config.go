package config

import (
	"bytes"
	"io/ioutil"

	"github.com/rakyll/statik/fs"
	"github.com/spf13/viper"

	//statik バイナリ化したassetファイルのPATH
	_ "go-vue-next-server/generate/statik"
)

var c *viper.Viper

// Init Configの初期化
// ライブラリStatikでビルド時にassetのyamlをバイナリに含め、
// それをもとにConfigを生成する
func Init(env string, path string) {
	//Viperという設定ファイル用ライブラリを使う
	c = viper.New()
	c.SetConfigType("YAML")
	//
	fileSystem, _ := fs.New()
	f, err := fileSystem.Open("/" + path + "/" + env + ".yaml")
	if err != nil {
		panic(err)
	}
	r, _ := ioutil.ReadAll(f)
	if err := c.ReadConfig(bytes.NewBuffer(r)); err != nil {
		panic(err)
	}
}

// GetConfig returns config
func GetConfig() *viper.Viper {
	return c
}
