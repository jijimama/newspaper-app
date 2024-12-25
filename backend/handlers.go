package main

import (
    "net/http"
    "context"

    "github.com/gin-gonic/gin"
    "github.com/jijimama/newspaper-app/ent"
)

func getArticlesHandler(c *gin.Context, client *ent.Client) {
    articles, err := getArticles(client)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed querying articles"})
        return
    }
    c.JSON(http.StatusOK, articles)
}

func getArticles(client *ent.Client) ([]Article, error) {
    articles, err := client.Article.Query().WithNewspaper().All(context.Background())
    if err != nil {
        return nil, err
    }

    result := make([]Article, len(articles))
    for i, a := range articles {
        result[i] = Article{
            Year:       a.Year,
            Month:      a.Month,
            Day:        a.Day,
            Content:    a.Content,
            Newspaper:  a.Edges.Newspaper.Name,
            ColumnName: a.Edges.Newspaper.ColumnName,
        }
    }

    return result, nil
}
