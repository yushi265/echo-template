# ベースイメージとして公式のGoイメージを使用
FROM golang:1.22-alpine

# 'air'のインストールに必要なツールをインストール
RUN apk add --no-cache git

# Airインストール
# RUN curl -fL https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air -o /usr/local/bin/air \
#     && chmod +x /usr/local/bin/air

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

# 作業ディレクトリを作成
WORKDIR /app

COPY go.mod go.sum ./

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# Goモジュールをコピーし、依存関係をインストール
RUN go mod download
RUN go get github.com/lib/pq
RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@latest
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/cosmtrek/air@v1.49.0

# アプリケーションのソースコードをコピー
COPY . .

# Airを実行してホットリロードを有効にする
CMD ["air"]
