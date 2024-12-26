package usecase

import (
    "context"
    "github.com/jijimama/newspaper-app/domain"
)

// ArticleUsecase は記事関連のビジネスロジックを管理する構造体。
// データを取得するためにリポジトリを使用。
type ArticleUsecase struct {
    articleRepo ArticleRepository // 記事データを取得するためのリポジトリ
}

// ArticleRepository はデータアクセスのためのインターフェース。
// データベースや外部APIから情報を取得するための方法を定義。
type ArticleRepository interface {
    // 記事一覧を取得するメソッド
    GetArticles(ctx context.Context) ([]domain.Article, error)
}

// NewArticleUsecase は ArticleUsecase を作成するファクトリ関数。
// リポジトリを引数として渡します。
// ユースケースがリポジトリ（データ操作）を利用できるようにするため
func NewArticleUsecase(repo ArticleRepository) *ArticleUsecase {
    return &ArticleUsecase{articleRepo: repo}
}

// FetchArticles はリポジトリを使って記事一覧を取得。
// コントローラーから呼び出され、実際にデータを取得する役割を担う。
func (uc *ArticleUsecase) FetchArticles(ctx context.Context) ([]domain.Article, error) {
    // リポジトリにデータを取得するよう依頼
    return uc.articleRepo.GetArticles(ctx)
}
