package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jijimama/newspaper-app/usecase"
    "github.com/jijimama/newspaper-app/adapter/presenter"
)

// ArticleController は記事関連のHTTPリクエストを処理するコントローラー。
// ユースケースに処理を依頼し、その結果を整形して返す。
type ArticleController struct {
    articleUsecase *usecase.ArticleUsecase // ユースケースを通じてビジネスロジックを実行
}

// NewArticleController は ArticleController を作成するファクトリ関数。
func NewArticleController(uc *usecase.ArticleUsecase) *ArticleController {
    return &ArticleController{articleUsecase: uc}
}

// GetArticles は /articles エンドポイントへのGETリクエストを処理。
// コントローラーは、ユースケースを使ってビジネスロジックを実行し、結果をレスポンスとして返す。
func (ctrl *ArticleController) GetArticles(c *gin.Context) {
    // ユースケースを呼び出して記事一覧を取得
    articles, err := ctrl.articleUsecase.FetchArticles(c.Request.Context())
    if err != nil {
        // エラーが発生した場合は500エラーを返す
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch articles"})
        return
    }

    // データをプレゼンターを使って整形
    response := presenter.PresentArticles(articles)
    // 整形したデータをJSONとして返す
    c.JSON(http.StatusOK, response)
}
