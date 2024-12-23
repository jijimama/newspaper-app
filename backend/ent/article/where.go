// Code generated by ent, DO NOT EDIT.

package article

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jijimama/newspaper-app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldContent, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldYear, v))
}

// Month applies equality check predicate on the "month" field. It's identical to MonthEQ.
func Month(v int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldMonth, v))
}

// Day applies equality check predicate on the "day" field. It's identical to DayEQ.
func Day(v int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldDay, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldUpdatedAt, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldContent, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldYear, v))
}

// MonthEQ applies the EQ predicate on the "month" field.
func MonthEQ(v int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldMonth, v))
}

// MonthNEQ applies the NEQ predicate on the "month" field.
func MonthNEQ(v int) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldMonth, v))
}

// MonthIn applies the In predicate on the "month" field.
func MonthIn(vs ...int) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldMonth, vs...))
}

// MonthNotIn applies the NotIn predicate on the "month" field.
func MonthNotIn(vs ...int) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldMonth, vs...))
}

// MonthGT applies the GT predicate on the "month" field.
func MonthGT(v int) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldMonth, v))
}

// MonthGTE applies the GTE predicate on the "month" field.
func MonthGTE(v int) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldMonth, v))
}

// MonthLT applies the LT predicate on the "month" field.
func MonthLT(v int) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldMonth, v))
}

// MonthLTE applies the LTE predicate on the "month" field.
func MonthLTE(v int) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldMonth, v))
}

// DayEQ applies the EQ predicate on the "day" field.
func DayEQ(v int) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldDay, v))
}

// DayNEQ applies the NEQ predicate on the "day" field.
func DayNEQ(v int) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldDay, v))
}

// DayIn applies the In predicate on the "day" field.
func DayIn(vs ...int) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldDay, vs...))
}

// DayNotIn applies the NotIn predicate on the "day" field.
func DayNotIn(vs ...int) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldDay, vs...))
}

// DayGT applies the GT predicate on the "day" field.
func DayGT(v int) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldDay, v))
}

// DayGTE applies the GTE predicate on the "day" field.
func DayGTE(v int) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldDay, v))
}

// DayLT applies the LT predicate on the "day" field.
func DayLT(v int) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldDay, v))
}

// DayLTE applies the LTE predicate on the "day" field.
func DayLTE(v int) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldDay, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasColumn applies the HasEdge predicate on the "column" edge.
func HasColumn() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ColumnTable, ColumnColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasColumnWith applies the HasEdge predicate on the "column" edge with a given conditions (other predicates).
func HasColumnWith(preds ...predicate.Column) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := newColumnStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Article) predicate.Article {
	return predicate.Article(sql.NotPredicates(p))
}
