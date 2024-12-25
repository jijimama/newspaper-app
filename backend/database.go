package main

import (
    "os"
    "context"

    "github.com/jijimama/newspaper-app/ent"
    _ "github.com/lib/pq"
)

func initDB() (*ent.Client, error) {
    // 環境変数から取得
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        // 環境変数が設定されていない場合、個別の環境変数から取得
        dbHost := os.Getenv("DB_HOST")
        dbPort := os.Getenv("DB_PORT")
        dbUser := os.Getenv("DB_USER")
        dbName := os.Getenv("DB_NAME")
        dbPassword := os.Getenv("DB_PASSWORD")

        // 接続文字列を作成
        dsn = "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"
    }
    // PostgreSQLに接続
    client, err := ent.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    // スキーマを作成
    if err := client.Schema.Create(context.Background()); err != nil {
        return nil, err
    }

    return client, nil
}
