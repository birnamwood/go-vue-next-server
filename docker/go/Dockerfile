FROM golang:1.16

# コンテナ側作業ディレクトリ
WORKDIR /go/src/go-vue-next-server

# プロジェクトのソースをコンテナに追加
COPY ./ /go/src/go-vue-next-server

# 日本語セット
RUN apt-get update \
    && apt-get install -y locales \
    && locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL=ja_JP.UTF-8
RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8
RUN ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

#マイグレーションツール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/bin/migrate

# Package追加
RUN go get golang.org/x/tools/cmd/goimports
RUN go get golang.org/x/lint/golint
RUN go get github.com/securego/gosec
RUN go get -u github.com/kisielk/errcheck
RUN go get -u github.com/sonatard/noctx/cmd/noctx
## DI TOOL
RUN go get github.com/google/wire/cmd/wire@v0.4.0
## アセットをバイナリ化
RUN go get github.com/rakyll/statik@v0.1.7
## Lint TOOL
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.32.2

# go module非対応のものを入れる
ENV GO111MODULE=off
## タスクランナー
RUN go get github.com/oxequa/realize
ENV GO111MODULE=on
