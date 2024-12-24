package main

import (
    "context"
    "encoding/csv"
    "flag"
    "log"
    "os"
	"fmt"
    "strconv"
	"github.com/joho/godotenv"

    "github.com/jijimama/newspaper-app/ent"
	"github.com/jijimama/newspaper-app/ent/newspaper"
    _ "github.com/lib/pq"
)

func main() {
    // コマンドライン引数を解析
    newspaperName := flag.String("newspaper", "", "Name of the newspaper")
    csvFilePath := flag.String("csv", "data/articles/mainichi.csv", "Path to the CSV file")
    flag.Parse()

    if *newspaperName == "" {
        log.Fatal("newspaper name is required")
    }

    // .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 環境変数から取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// 接続文字列を作成
	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

    //PostgreSQLに接続
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// すべての新聞レコードを表示
    allNewspapers, err := client.Newspaper.Query().All(context.Background())
    if err != nil {
        log.Fatalf("failed querying newspapers: %v", err)
    }
    fmt.Println("All newspapers:")
    for _, n := range allNewspapers {
        fmt.Printf("Name:%s: %s\n",
            n.Name)
    }

    // 新聞を取得
    newspaper, err := client.Newspaper.
        Query().
        Where(newspaper.NameEQ(*newspaperName)).
        Only(context.Background())
    if err != nil {
        log.Fatalf("failed searching newspaper: %v", err)
    }

    // CSVファイルからデータを読み込み、データベースに挿入
    err = LoadCSVData(client, *csvFilePath, newspaper)
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
