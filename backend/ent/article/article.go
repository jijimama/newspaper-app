// Code generated by ent, DO NOT EDIT.

package article

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldMonth holds the string denoting the month field in the database.
	FieldMonth = "month"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeColumn holds the string denoting the column edge name in mutations.
	EdgeColumn = "column"
	// Table holds the table name of the article in the database.
	Table = "articles"
	// ColumnTable is the table that holds the column relation/edge.
	ColumnTable = "articles"
	// ColumnInverseTable is the table name for the Column entity.
	// It exists in this package in order to avoid circular dependency with the "column" package.
	ColumnInverseTable = "columns"
	// ColumnColumn is the table column denoting the column relation/edge.
	ColumnColumn = "column_articles"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldYear,
	FieldMonth,
	FieldDay,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "articles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"column_articles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
	// MonthValidator is a validator for the "month" field. It is called by the builders before save.
	MonthValidator func(int) error
	// DayValidator is a validator for the "day" field. It is called by the builders before save.
	DayValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Article queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByMonth orders the results by the month field.
func ByMonth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMonth, opts...).ToFunc()
}

// ByDay orders the results by the day field.
func ByDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDay, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByColumnField orders the results by column field.
func ByColumnField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newColumnStep(), sql.OrderByField(field, opts...))
	}
}
func newColumnStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ColumnInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ColumnTable, ColumnColumn),
	)
}