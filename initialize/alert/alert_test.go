package alert

import (
	"errors"
	"runtime/debug"
	"testing"
)

func TestSendMail(t *testing.T) {
	type args struct {
		err   error
		debug string
		key   []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "alert処理テストパラメータ",
			args: args{
				err:   errors.New("えらーです"),
				debug: string(debug.Stack()),
				key:   []byte("test"),
			},
		},
		{
			name: "alert処理テストパラメータ空",
			args: args{
				err:   nil,
				debug: "",
				key:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMail(tt.args.err, tt.args.debug, string(tt.args.key))
		})
	}
}
