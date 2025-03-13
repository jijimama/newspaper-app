package gateway_test

import (
    "context"
    "regexp"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/suite"
    "github.com/jijimama/newspaper-app/ent"
    "github.com/jijimama/newspaper-app/adapter/gateway"
    "github.com/jijimama/newspaper-app/pkg/tester"
)

type ArticleRepositorySuite struct {
    suite.Suite
    mock       sqlmock.Sqlmock
    client     *ent.Client
    repository gateway.ArticleRepository
}

func TestArticleRepositorySuite(t *testing.T) {
    suite.Run(t, new(ArticleRepositorySuite))
}

func (suite *ArticleRepositorySuite) SetupSuite() {
    mock, client := tester.MockDB(suite.T())
    suite.mock = mock
    suite.client = client
    suite.repository = gateway.NewArticleRepository(client)
}

func (suite *ArticleRepositorySuite) TearDownSuite() {
    suite.client.Close()
}

func (suite *ArticleRepositorySuite) TestGetArticles() {
    rows := suite.mock.NewRows([]string{"id", "year", "month", "day", "content", "newspaper_id", "name", "column_name"}).
        AddRow(1, 2023, 12, 20, "Sample Content", 1, "Daily News", "Editorial")

    suite.mock.ExpectQuery(regexp.QuoteMeta(
        `SELECT * FROM "articles" LEFT JOIN "newspapers" ON "articles"."newspaper_id" = "newspapers"."id"`)).
        WillReturnRows(rows)

    ctx := context.Background()
    articles, err := suite.repository.GetArticles(ctx)

    suite.Assert().Nil(err)
    suite.Assert().NotNil(articles)
    suite.Assert().Equal(2023, articles[0].Year)
    suite.Assert().Equal(12, articles[0].Month)
    suite.Assert().Equal(20, articles[0].Day)
    suite.Assert().Equal("Sample Content", articles[0].Content)
    suite.Assert().Equal("Daily News", articles[0].Newspaper)
    suite.Assert().Equal("Editorial", articles[0].ColumnName)
}