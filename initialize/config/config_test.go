package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		env  string
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Development環境Configファイル生成",
			args: args{
				env:  "development",
				path: "./environment/",
			},
		},
		{
			name: "Production環境Configファイル生成",
			args: args{
				env:  "production",
				path: "./environment/",
			},
		},
		{
			name: "Local環境Configファイル生成",
			args: args{
				env:  "local",
				path: "./environment/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.env, tt.args.path)
		})
	}
}
