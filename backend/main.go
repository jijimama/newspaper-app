package main

import (
	"net/http"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jijimama/newspaper-app/ent"
	_ "github.com/lib/pq"
)

type Article struct {
    Year      int    `json:"year"`
    Month     int    `json:"month"`
    Day       int    `json:"day"`
    Content   string `json:"content"`
	Newspaper string `json:"newspaper"`
	ColumnName string `json:"column_name"`
}

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
	
  	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
      		"message": "Hello, World!",
		})
  	})

	router.GET("/articles", func(c *gin.Context) {
        articles, err := getArticles(client)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed querying articles"})
            return
        }
        c.JSON(http.StatusOK, articles)
    })

  	router.Run(":8080")
}

func getArticles(client *ent.Client) ([]Article, error) {
    articles, err := client.Article.Query().WithNewspaper().All(context.Background())
    if err != nil {
        return nil, err
    }

    result := make([]Article, len(articles))
    for i, a := range articles {
        result[i] = Article{
            Year:       a.Year,
            Month:      a.Month,
            Day:        a.Day,
            Content:    a.Content,
            Newspaper:  a.Edges.Newspaper.Name,
			ColumnName: a.Edges.Newspaper.ColumnName,
        }
    }

    return result, nil
}