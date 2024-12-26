package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jijimama/newspaper-app/adapter/controller"
)

// NewRouter はアプリケーションのルーティングを設定
func NewRouter(articleController *controller.ArticleController) *gin.Engine {
    r := gin.Default()

    // ヘルスチェック用エンドポイント
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })

    // 記事一覧のエンドポイント
    r.GET("/articles", articleController.GetArticles)

    return r
}
