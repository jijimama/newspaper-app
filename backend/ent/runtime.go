// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/jijimama/newspaper-app/ent/article"
	"github.com/jijimama/newspaper-app/ent/column"
	"github.com/jijimama/newspaper-app/ent/newspaper"
	"github.com/jijimama/newspaper-app/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescContent is the schema descriptor for content field.
	articleDescContent := articleFields[0].Descriptor()
	// article.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	article.ContentValidator = articleDescContent.Validators[0].(func(string) error)
	// articleDescYear is the schema descriptor for year field.
	articleDescYear := articleFields[1].Descriptor()
	// article.YearValidator is a validator for the "year" field. It is called by the builders before save.
	article.YearValidator = articleDescYear.Validators[0].(func(int) error)
	// articleDescMonth is the schema descriptor for month field.
	articleDescMonth := articleFields[2].Descriptor()
	// article.MonthValidator is a validator for the "month" field. It is called by the builders before save.
	article.MonthValidator = articleDescMonth.Validators[0].(func(int) error)
	// articleDescDay is the schema descriptor for day field.
	articleDescDay := articleFields[3].Descriptor()
	// article.DayValidator is a validator for the "day" field. It is called by the builders before save.
	article.DayValidator = articleDescDay.Validators[0].(func(int) error)
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleFields[4].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleFields[5].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	// article.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	article.UpdateDefaultUpdatedAt = articleDescUpdatedAt.UpdateDefault.(func() time.Time)
	columnFields := schema.Column{}.Fields()
	_ = columnFields
	// columnDescName is the schema descriptor for name field.
	columnDescName := columnFields[0].Descriptor()
	// column.NameValidator is a validator for the "name" field. It is called by the builders before save.
	column.NameValidator = columnDescName.Validators[0].(func(string) error)
	// columnDescCreatedAt is the schema descriptor for created_at field.
	columnDescCreatedAt := columnFields[1].Descriptor()
	// column.DefaultCreatedAt holds the default value on creation for the created_at field.
	column.DefaultCreatedAt = columnDescCreatedAt.Default.(func() time.Time)
	// columnDescUpdatedAt is the schema descriptor for updated_at field.
	columnDescUpdatedAt := columnFields[2].Descriptor()
	// column.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	column.DefaultUpdatedAt = columnDescUpdatedAt.Default.(func() time.Time)
	// column.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	column.UpdateDefaultUpdatedAt = columnDescUpdatedAt.UpdateDefault.(func() time.Time)
	newspaperFields := schema.Newspaper{}.Fields()
	_ = newspaperFields
	// newspaperDescName is the schema descriptor for name field.
	newspaperDescName := newspaperFields[0].Descriptor()
	// newspaper.NameValidator is a validator for the "name" field. It is called by the builders before save.
	newspaper.NameValidator = newspaperDescName.Validators[0].(func(string) error)
	// newspaperDescCreatedAt is the schema descriptor for created_at field.
	newspaperDescCreatedAt := newspaperFields[1].Descriptor()
	// newspaper.DefaultCreatedAt holds the default value on creation for the created_at field.
	newspaper.DefaultCreatedAt = newspaperDescCreatedAt.Default.(func() time.Time)
	// newspaperDescUpdatedAt is the schema descriptor for updated_at field.
	newspaperDescUpdatedAt := newspaperFields[2].Descriptor()
	// newspaper.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	newspaper.DefaultUpdatedAt = newspaperDescUpdatedAt.Default.(func() time.Time)
	// newspaper.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	newspaper.UpdateDefaultUpdatedAt = newspaperDescUpdatedAt.UpdateDefault.(func() time.Time)
}
