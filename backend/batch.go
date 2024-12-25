package main

import (
    "context"
    "encoding/csv"
    "log"
    "os"
    "strconv"
	"github.com/joho/godotenv"

    "github.com/jijimama/newspaper-app/ent"
	"github.com/jijimama/newspaper-app/ent/newspaper"
    _ "github.com/lib/pq"
)

func RunBatch(newspaperName string, csvFilePath string) {
    if newspaperName == "" {
        log.Fatal("newspaper name is required")
    }

    // .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// データベース接続を初期化
    client, err := initDB()
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()

    // 新聞を取得
    newspaper, err := client.Newspaper.
        Query().
        Where(newspaper.NameEQ(newspaperName)).
        Only(context.Background())
    if err != nil {
        log.Fatalf("failed searching newspaper: %v", err)
    }

    // CSVファイルからデータを読み込み、データベースに挿入
    err = LoadCSVData(client, csvFilePath, newspaper)
    if err != nil {
        log.Fatalf("failed loading CSV data: %v", err)
    }
}

// LoadCSVData は、指定されたCSVファイルからデータを読み込み、データベースに挿入します。
func LoadCSVData(client *ent.Client, filePath string, newspaper *ent.Newspaper) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for _, record := range records {
        year, err := strconv.Atoi(record[0])
        if err != nil {
            return err
        }
        month, err := strconv.Atoi(record[1])
        if err != nil {
            return err
        }
        day, err := strconv.Atoi(record[2])
        if err != nil {
            return err
        }
        content := record[3]

        _, err = client.Article.
            Create().
            SetYear(year).
            SetMonth(month).
            SetDay(day).
            SetContent(content).
            SetNewspaper(newspaper).
            Save(context.Background())
        if err != nil {
            return err
        }
    }

    return nil
}
