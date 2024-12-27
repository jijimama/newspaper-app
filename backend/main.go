package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jijimama/newspaper-app/config"
	"github.com/jijimama/newspaper-app/infrastructure/database"
	"github.com/jijimama/newspaper-app/adapter/controller"
    "github.com/jijimama/newspaper-app/adapter/gateway"
	"github.com/jijimama/newspaper-app/adapter/router"
    "github.com/jijimama/newspaper-app/usecase"
)

func main() {
	// 設定を読み込む
    cfg := config.LoadConfig()
	// GIN_MODE を設定
    gin.SetMode(cfg.GinMode)
	// データベース接続を初期化
    client, err := database.NewDatabaseSQLFactory(database.InstancePostgres)
    if err != nil {
        log.Fatalf("failed opening connection to database: %v", err)
    }
    defer client.Close()

	// ゲートウェイ、ユースケース、コントローラーを初期化
    articleRepository := gateway.NewArticleRepository(client)
    articleUsecase := usecase.NewArticleUseCase(articleRepository)
    articleController := controller.NewArticleController(articleUsecase)

	// ルーターを設定
    r := router.NewRouter(articleController)

	// サーバーを起動
	r.Run(":8080")
}
