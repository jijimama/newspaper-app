package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jijimama/newspaper-app/usecase"
)

// ArticleController は記事関連のHTTPリクエストを処理するコントローラー。
// ユースケースに処理を依頼し、その結果を整形して返す。
type ArticleController struct {
    articleUseCase usecase.ArticleUseCase // ユースケースを通じてビジネスロジックを実行
}

// NewArticleController は ArticleController を作成するファクトリ関数。
func NewArticleController(articleUseCase usecase.ArticleUseCase) *ArticleController {
    return &ArticleController{
		articleUseCase: articleUseCase,
	}
}

// GetArticles は /articles エンドポイントへのGETリクエストを処理。
// コントローラーは、ユースケースを使ってビジネスロジックを実行し、結果をレスポンスとして返す。
func (ctrl *ArticleController) GetArticles(c *gin.Context) {
    // ユースケースを呼び出して記事一覧を取得
    articles, err := ctrl.articleUseCase.GetArticles(c.Request.Context())
    if err != nil {
        // エラーが発生した場合は500エラーを返す
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch articles"})
        return
    }

    c.JSON(http.StatusOK, articles)
}
