package usecase

import (
    "context"
    "testing"

    "github.com/jijimama/newspaper-app/domain"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
)

// mockArticleRepository は ArticleRepository インターフェースを模擬したモック構造体。
type mockArticleRepository struct {
    mock.Mock
}

// NewMockArticleRepository は mockArticleRepository のインスタンスを生成するヘルパー関数。
func NewMockArticleRepository() *mockArticleRepository {
    return &mockArticleRepository{}
}

// GetArticles はモックの GetArticles 実装。
// 引数と戻り値を記録し、期待される値を返す。
func (m *mockArticleRepository) GetArticles(ctx context.Context) ([]domain.Article, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]domain.Article), args.Error(1)
}

// ArticleUsecaseTestSuite は ArticleUsecase のテストスイートを定義。
type ArticleUsecaseTestSuite struct {
    suite.Suite
    articleUsecase      *ArticleUsecase        // テスト対象のユースケース
    mockArticleRepo     *mockArticleRepository // モックリポジトリ
}

// SetupTest は各テストケースの前に実行されるセットアップ処理。
func (suite *ArticleUsecaseTestSuite) SetupTest() {
    suite.mockArticleRepo = NewMockArticleRepository() // モックリポジトリを初期化
    suite.articleUsecase = NewArticleUsecase(suite.mockArticleRepo) // ユースケースにモックを注入
}

// TestFetchArticles_Success は FetchArticles が成功する場合のテスト。
func (suite *ArticleUsecaseTestSuite) TestFetchArticles_Success() {
    // テスト用のデータを準備
    articles := []domain.Article{
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

    // モックの挙動を設定
    suite.mockArticleRepo.On("GetArticles", mock.Anything).Return(articles, nil)

    // テスト対象の関数を実行
    result, err := suite.articleUsecase.FetchArticles(context.Background())

    // 結果を検証
    suite.Assert().NoError(err, "エラーが発生していないことを確認")
    suite.Assert().Equal(len(articles), len(result), "取得した記事の件数が一致することを確認")
    suite.Assert().Equal(articles[0].Content, result[0].Content, "記事の内容が一致することを確認")
}

// TestFetchArticles_Failure は FetchArticles が失敗する場合のテスト。
func (suite *ArticleUsecaseTestSuite) TestFetchArticles_Failure() {
    // モックの挙動を設定
    suite.mockArticleRepo.On("GetArticles", mock.Anything).Return([]domain.Article{}, context.DeadlineExceeded)

    // テスト対象の関数を実行
    result, err := suite.articleUsecase.FetchArticles(context.Background())

    // 結果を検証
    suite.Assert().Error(err, "エラーが発生することを確認")
    suite.Assert().Equal(context.DeadlineExceeded, err, "期待されるエラーが返ることを確認")
    suite.Assert().Equal([]domain.Article{}, result, "結果が空のスライスであることを確認")
}

// テストスイートを実行するためのエントリーポイント。
func TestArticleUsecaseTestSuite(t *testing.T) {
    suite.Run(t, new(ArticleUsecaseTestSuite)) // テストスイートを実行
}
