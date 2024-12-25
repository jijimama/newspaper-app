package main

import (
	"log"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// コマンドライン引数を解析
    runBatch := flag.Bool("batch", false, "Run batch process")
	newspaperName := flag.String("newspaper", "", "Name of the newspaper")
    csvFilePath := flag.String("csv", "data/articles/mainichi.csv", "Path to the CSV file")
    flag.Parse()

    if *runBatch {
        RunBatch(*newspaperName, *csvFilePath)
        return
    }

	// GIN_MODE 環境変数を設定
    ginMode := os.Getenv("GIN_MODE")
    if ginMode == "" {
        ginMode = "debug" // デフォルトは "debug"
    }
    gin.SetMode(ginMode)

	router := gin.Default()

	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// データベース接続を初期化
    client, err := initDB()
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
	defer client.Close()

  	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
      		"message": "Hello, World!",
		})
  	})

	router.GET("/articles", func(c *gin.Context) {
        getArticlesHandler(c, client)
    })

  	router.Run(":8080")
}
