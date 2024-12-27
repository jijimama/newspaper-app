package usecase

import (
    "context"
    "github.com/jijimama/newspaper-app/adapter/gateway"
    "github.com/jijimama/newspaper-app/domain"
)

// ArticleUseCase はデータアクセスのためのインターフェース。
// データベースや外部APIから情報を取得するための方法を定義。
type ArticleUseCase interface {
    // 記事一覧を取得するメソッド
    GetArticles(ctx context.Context) ([]*domain.Article, error)
}

// articleUseCase は記事関連のビジネスロジックを管理する構造体。
// データを取得するためにリポジトリを使用。
type articleUseCase struct {
    articleRepository gateway.ArticleRepository // 記事データを取得するためのリポジトリ
}

// NewArticleUseCase は articleUseCase を作成するファクトリ関数。
// リポジトリを引数として渡します。
// ユースケースがリポジトリ（データ操作）を利用できるようにするため
func NewArticleUseCase(articleRepository gateway.ArticleRepository) *articleUseCase {
    return &articleUseCase{
        articleRepository: articleRepository,
    }
}

// GetArticles はリポジトリを使って記事一覧を取得。
// コントローラーから呼び出され、実際にデータを取得する役割を担う。
func (uc *articleUseCase) GetArticles(ctx context.Context) ([]*domain.Article, error) {
    // リポジトリにデータを取得するよう依頼
    return uc.articleRepository.GetArticles(ctx)
}
