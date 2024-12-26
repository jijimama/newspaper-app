package presenter

import "github.com/jijimama/newspaper-app/domain"

// PresentArticles は記事のデータをAPIレスポンス用に整形。
// ユーザーに分かりやすい形式に変換。
func PresentArticles(articles []domain.Article) []map[string]interface{} {
    result := make([]map[string]interface{}, len(articles)) // 整形結果を格納するスライス
    for i, article := range articles {
        result[i] = map[string]interface{}{
            "year":       article.Year,
            "month":      article.Month,
            "day":        article.Day,
            "content":    article.Content,
            "newspaper":  article.Newspaper,
            "columnName": article.ColumnName,
        }
    }
    return result // 整形されたデータを返す
}
