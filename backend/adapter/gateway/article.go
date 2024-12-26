package gateway

import (
    "context"

    "github.com/jijimama/newspaper-app/domain"
    "github.com/jijimama/newspaper-app/ent"
)

// ArticleGateway は記事データをデータベースから取得するゲートウェイ。
type ArticleGateway struct {
    client *ent.Client // Ent クライアントを使用してデータベースと通信
}

// NewArticleGateway は クライアントをフィールドに持つ ArticleGateway を返す。
func NewArticleGateway(client *ent.Client) *ArticleGateway {
    return &ArticleGateway{client: client}
}

// GetArticles はデータベースから記事一覧を取得。
func (g *ArticleGateway) GetArticles(ctx context.Context) ([]domain.Article, error) {
    // データベースから記事をクエリ
    articles, err := g.client.Article.Query().WithNewspaper().All(ctx)
    if err != nil {
        return nil, err // エラーが発生した場合は上位に返す
    }

    // データベースの結果をドメインモデルに変換
    result := make([]domain.Article, len(articles))
    for i, a := range articles {
        result[i] = domain.Article{
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
