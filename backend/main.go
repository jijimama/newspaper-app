package main

import (
	"github.com/gin-gonic/gin"
	"context"
	"encoding/csv"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jijimama/newspaper-app/ent"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 環境変数から取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// 接続文字列を作成
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

    //PostgreSQLに接続
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// CSVファイルからデータを読み込み、データベースに挿入
    err = loadCSVData(client, "data/newspapers.csv")
    if err != nil {
        log.Fatalf("failed loading CSV data: %v", err)
    }
	
  	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
      		"message": "Hello, World!",
		})
  	})

  	router.Run(":8080")
}

func loadCSVData(client *ent.Client, filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for _, record := range records {
        name := record[0]
        columnName := record[1]

        _, err := client.Newspaper.
            Create().
            SetName(name).
            SetColumnName(columnName).
            Save(context.Background())
        if err != nil {
            return err
        }
    }

    return nil
}