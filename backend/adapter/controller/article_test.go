package controller

import (
    "context"
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
    "github.com/jijimama/newspaper-app/domain"
)

// MockArticleUseCase は ArticleUseCase をモック化したもの
type MockArticleUseCase struct {
	mock.Mock
}

// NewMockArticleUseCase は MockArticleUseCase の新しいインスタンスを作成する
func NewMockArticleUseCase() *MockArticleUseCase {
	return &MockArticleUseCase{}
}

// GetArticles はユースケースのモックメソッド
func (m *MockArticleUseCase) GetArticles(ctx context.Context) ([]*domain.Article, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]*domain.Article), args.Error(1)
}

// ArticleControllerSuite は ArticleController のテストスイート
type ArticleControllerSuite struct {
	suite.Suite
	mockUseCase      *MockArticleUseCase
	articleController *ArticleController
}

// TestArticleControllerTestSuite はテストスイートを実行する
func TestArticleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleControllerSuite))
}

// SetupTest は各テストの前に実行されるセットアップメソッド
func (suite *ArticleControllerSuite) SetupTest() {
    suite.mockUseCase = NewMockArticleUseCase()
    suite.articleController = NewArticleController(suite.mockUseCase)
}

// TestGetArticles は GetArticles メソッドのテスト
func (suite *ArticleControllerSuite) TestGetArticles() {
	// モックの返却値を設定
	articles := []*domain.Article{
        {
            Year:       2023,
            Month:      12,
            Day:        1,
            Content:    "Article Content 1",
            Newspaper:  "Daily News",
            ColumnName: "Opinion",
        },
        {
            Year:       2023,
            Month:      12,
            Day:        2,
            Content:    "Article Content 2",
            Newspaper:  "Daily News",
            ColumnName: "Sports",
        },
    }
	suite.mockUseCase.On("GetArticles", mock.Anything).Return(articles, nil)

	// HTTPリクエストを作成
	req, _ := http.NewRequest(http.MethodGet, "/articles", nil)
	w := httptest.NewRecorder()

    ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = req

    // ハンドラメソッドを直接呼び出し
	suite.articleController.GetArticles(ginContext)

	// レスポンスを確認
	suite.Assert().Equal(http.StatusOK, w.Code)

    // レスポンスボディを読み取る
    bodyBytes, _ := io.ReadAll(w.Body)
    var articlesResponse []domain.Article
    err := json.Unmarshal(bodyBytes, &articlesResponse)
    suite.Assert().Nil(err)

    // レスポンスの内容を検証
    suite.Assert().Equal(2, len(articlesResponse))
    suite.Assert().Equal(2023, articlesResponse[0].Year)
    suite.Assert().Equal(12, articlesResponse[0].Month)
    suite.Assert().Equal(1, articlesResponse[0].Day)
    suite.Assert().Equal("Article Content 1", articlesResponse[0].Content)
    suite.Assert().Equal("Daily News", articlesResponse[0].Newspaper)
    suite.Assert().Equal("Opinion", articlesResponse[0].ColumnName)

    suite.Assert().Equal(2023, articlesResponse[1].Year)
    suite.Assert().Equal(12, articlesResponse[1].Month)
    suite.Assert().Equal(2, articlesResponse[1].Day)
    suite.Assert().Equal("Article Content 2", articlesResponse[1].Content)
    suite.Assert().Equal("Daily News", articlesResponse[1].Newspaper)
    suite.Assert().Equal("Sports", articlesResponse[1].ColumnName)
}
