FROM golang:1.22-alpine

RUN apk add --no-cache git

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

# 作業ディレクトリを作成
WORKDIR /app

COPY go.mod go.sum ./

ENV TZ Asia/Tokyo

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.49.0

COPY . .

CMD ["air", "-c", ".air.toml"]
