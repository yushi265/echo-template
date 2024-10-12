package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
    _ "github.com/lib/pq"
)

func main() {
    e := echo.New()

    // 環境変数からDB接続情報を取得
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // PostgreSQLに接続
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }
    defer db.Close()

    // DB接続テスト
    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to database:", err)
    }
    fmt.Println("Successfully connected to the database!")

    // ルートエンドポイント
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World! Connected to PostgreSQL!")
    })

    e.Logger.Fatal(e.Start(":8080"))
}
