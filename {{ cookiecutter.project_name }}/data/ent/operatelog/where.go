// Code generated by ent, DO NOT EDIT.

package operatelog

import (
	"{{ cookiecutter.project_name }}/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldName, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldPath, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldMethod, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldIP, v))
}

// StatusCode applies equality check predicate on the "status_code" field. It's identical to StatusCodeEQ.
func StatusCode(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldStatusCode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContainsFold(FieldName, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContainsFold(FieldPath, v))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContainsFold(FieldMethod, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldContainsFold(FieldIP, v))
}

// StatusCodeEQ applies the EQ predicate on the "status_code" field.
func StatusCodeEQ(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldEQ(FieldStatusCode, v))
}

// StatusCodeNEQ applies the NEQ predicate on the "status_code" field.
func StatusCodeNEQ(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNEQ(FieldStatusCode, v))
}

// StatusCodeIn applies the In predicate on the "status_code" field.
func StatusCodeIn(vs ...int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldIn(FieldStatusCode, vs...))
}

// StatusCodeNotIn applies the NotIn predicate on the "status_code" field.
func StatusCodeNotIn(vs ...int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldNotIn(FieldStatusCode, vs...))
}

// StatusCodeGT applies the GT predicate on the "status_code" field.
func StatusCodeGT(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGT(FieldStatusCode, v))
}

// StatusCodeGTE applies the GTE predicate on the "status_code" field.
func StatusCodeGTE(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldGTE(FieldStatusCode, v))
}

// StatusCodeLT applies the LT predicate on the "status_code" field.
func StatusCodeLT(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLT(FieldStatusCode, v))
}

// StatusCodeLTE applies the LTE predicate on the "status_code" field.
func StatusCodeLTE(v int) predicate.OperateLog {
	return predicate.OperateLog(sql.FieldLTE(FieldStatusCode, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.OperateLog {
	return predicate.OperateLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.OperateLog {
	return predicate.OperateLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OperateLog) predicate.OperateLog {
	return predicate.OperateLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OperateLog) predicate.OperateLog {
	return predicate.OperateLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OperateLog) predicate.OperateLog {
	return predicate.OperateLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
