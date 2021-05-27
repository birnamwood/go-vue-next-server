package alert

import (
	"bytes"
	"encoding/base64"
	"go-vue-next-server/initialize/config"
	"net/smtp"

	"go.uber.org/zap"
)

type alertSettings struct {
	enabled     bool
	environment string
	from        string
	to          []string
}

var alert alertSettings

// Init アラートメールセットアップ
func Init() {
	c := config.GetConfig()
	// メール送信先
	alert.enabled = c.GetBool("alert.enabled")
	alert.environment = c.GetString("environment")
	alert.from = c.GetString("alert.from")
	alert.to = c.GetStringSlice("alert.to")
}

// SendMail アラートメール送信実行
func SendMail(err error, debug string, key string) {
	zap.S().Error(zap.Error(err), "key情報", key)
	if !alert.enabled {
		return
	}
	title := "error(" + alert.environment + ")"
	from := alert.from
	to := alert.to
	if len(to) < 1 {
		return
	}

	var body bytes.Buffer
	body.WriteString("【 エラー内容 】\n" + err.Error() + "\n\n")
	if key != "" {
		body.WriteString("【 キー情報 】\n" + key + "\n\n")
	} else {
		body.WriteString("【 キー情報 】\nなし\n\n")
	}
	if debug != "" {
		body.WriteString("【 スタックトレース 】\n" + debug + "\n")
	} else {
		body.WriteString("【 スタックトレース 】\nなし\n")
	}

	var header bytes.Buffer
	header.WriteString("From: " + from + "\r\n")
	header.WriteString("Subject: " + title + "\r\n")
	header.WriteString("MIME-Version: 1.0\r\n")
	header.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	header.WriteString("Content-Transfer-Encoding: base64\r\n")

	var message bytes.Buffer = header
	message.WriteString("\r\n")
	message.WriteString(base64.StdEncoding.EncodeToString(body.Bytes()))

	for _, to := range to {
		sendErr := smtp.SendMail(
			"localhost:25",
			nil,
			from,
			[]string{to},
			message.Bytes(),
		)
		if sendErr != nil {
			zap.S().Error("アラートメール送信失敗")
		}
	}
}
