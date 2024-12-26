package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config はアプリケーション全体の設定を格納する構造体
type Config struct {
    GinMode string
}

// LoadConfig は設定を初期化して返す
func LoadConfig() *Config {
    ginMode := getEnv("GIN_MODE", "debug")

    // GIN_MODE が release でない場合に .env ファイルを読み込む
    if ginMode != "release" {
        err := godotenv.Load()
        if err != nil {
            log.Printf("Error loading .env file: %v", err)
        }
    }

    return &Config{
        GinMode: ginMode,
    }
}

// getEnv は環境変数の値を取得し、デフォルト値を返す
func getEnv(key string, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
