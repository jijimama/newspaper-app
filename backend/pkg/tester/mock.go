package tester

import (
	"time"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock" // sqlmock はモックデータベースの作成に使用。
	"github.com/jijimama/newspaper-app/ent"
	"github.com/jijimama/newspaper-app/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

// モックデータベース (sqlmock) と ent.Client インスタンスを作成するヘルパー関数。
func MockDB(t *testing.T) (sqlmock.Sqlmock, *ent.Client) {
    // モックDBを初期化（sqlmock.New を使用してモックデータベースを作成）
    _, mock, err := sqlmock.New(
        // 正規表現マッチャー (QueryMatcherRegexp) を有効化。
        sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
    )
    if err != nil {
        log.Fatalf("failed to create sqlmock: %v", err)
    }

    // ent.Client をモックとして作成
    client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
    if err != nil {
        log.Fatalf("failed to create ent client: %v", err)
    }

    // モックオブジェクトとクライアントを返す。
    return mock, client
}

// モッククロック構造体
type mockClock struct {
	t time.Time
}

// モッククロックの生成
func NewMockClock(t time.Time) mockClock {
	return mockClock{t}
}

// 現在時刻を返す
func (m mockClock) Now() time.Time {
	return m.t
}