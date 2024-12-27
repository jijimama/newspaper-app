package gateway

import (
    "context"

    "github.com/jijimama/newspaper-app/domain"
    "github.com/jijimama/newspaper-app/ent"
)

// ArticleRepository は記事データをデータベースから取得するリポジトリインターフェース。
type ArticleRepository interface {
    GetArticles(ctx context.Context) ([]*domain.Article, error)
}

type articleRepository struct {
    client *ent.Client // Ent クライアントを使用してデータベースと通信
}

// NewArticleRepository は クライアントをフィールドに持つ ArticleRepository を返す。
func NewArticleRepository(client *ent.Client) ArticleRepository {
    return &articleRepository{client: client}
}

// GetArticles はデータベースから記事一覧を取得。
func (r *articleRepository) GetArticles(ctx context.Context) ([]*domain.Article, error) {
    // データベースから記事をクエリ
    articles, err := r.client.Article.Query().WithNewspaper().All(ctx)
    if err != nil {
        return nil, err // エラーが発生した場合は上位に返す
    }

    // データベースの結果をドメインモデルに変換
    result := make([]*domain.Article, len(articles))
    for i, a := range articles {
        result[i] = &domain.Article{
            Year:       a.Year,
            Month:      a.Month,
            Day:        a.Day,
            Content:    a.Content,
            Newspaper:  a.Edges.Newspaper.Name,
            ColumnName: a.Edges.Newspaper.ColumnName,
        }
    }

    return result, nil // ドメインモデルを返す
}