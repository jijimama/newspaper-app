package domain

type Article struct {
    Year       int    `json:"year"`
    Month      int    `json:"month"`
    Day        int    `json:"day"`
    Content    string `json:"content"`
    Newspaper  string `json:"newspaper"`
    ColumnName string `json:"column_name"`
}
