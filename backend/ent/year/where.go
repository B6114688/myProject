// Code generated by entc, DO NOT EDIT.

package year

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Years applies equality check predicate on the "years" field. It's identical to YearsEQ.
func Years(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldYears), v))
	})
}

// YearsEQ applies the EQ predicate on the "years" field.
func YearsEQ(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldYears), v))
	})
}

// YearsNEQ applies the NEQ predicate on the "years" field.
func YearsNEQ(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldYears), v))
	})
}

// YearsIn applies the In predicate on the "years" field.
func YearsIn(vs ...int) predicate.Year {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Year(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldYears), v...))
	})
}

// YearsNotIn applies the NotIn predicate on the "years" field.
func YearsNotIn(vs ...int) predicate.Year {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Year(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldYears), v...))
	})
}

// YearsGT applies the GT predicate on the "years" field.
func YearsGT(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldYears), v))
	})
}

// YearsGTE applies the GTE predicate on the "years" field.
func YearsGTE(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldYears), v))
	})
}

// YearsLT applies the LT predicate on the "years" field.
func YearsLT(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldYears), v))
	})
}

// YearsLTE applies the LTE predicate on the "years" field.
func YearsLTE(v int) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldYears), v))
	})
}

// HasYearTerm applies the HasEdge predicate on the "year_term" edge.
func HasYearTerm() predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YearTermTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, YearTermTable, YearTermColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasYearTermWith applies the HasEdge predicate on the "year_term" edge with a given conditions (other predicates).
func HasYearTermWith(preds ...predicate.Term) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YearTermInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, YearTermTable, YearTermColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasYearResu applies the HasEdge predicate on the "year_resu" edge.
func HasYearResu() predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YearResuTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, YearResuTable, YearResuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasYearResuWith applies the HasEdge predicate on the "year_resu" edge with a given conditions (other predicates).
func HasYearResuWith(preds ...predicate.Results) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YearResuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, YearResuTable, YearResuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasYearActi applies the HasEdge predicate on the "year_acti" edge.
func HasYearActi() predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YearActiTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, YearActiTable, YearActiColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasYearActiWith applies the HasEdge predicate on the "year_acti" edge with a given conditions (other predicates).
func HasYearActiWith(preds ...predicate.Activity) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(YearActiInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, YearActiTable, YearActiColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Year) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Year) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
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
func Not(p predicate.Year) predicate.Year {
	return predicate.Year(func(s *sql.Selector) {
		p(s.Not())
	})
}