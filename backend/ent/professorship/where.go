// Code generated by entc, DO NOT EDIT.

package professorship

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Professorship applies equality check predicate on the "professorship" field. It's identical to ProfessorshipEQ.
func Professorship(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipEQ applies the EQ predicate on the "professorship" field.
func ProfessorshipEQ(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipNEQ applies the NEQ predicate on the "professorship" field.
func ProfessorshipNEQ(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipIn applies the In predicate on the "professorship" field.
func ProfessorshipIn(vs ...string) predicate.Professorship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Professorship(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProfessorship), v...))
	})
}

// ProfessorshipNotIn applies the NotIn predicate on the "professorship" field.
func ProfessorshipNotIn(vs ...string) predicate.Professorship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Professorship(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProfessorship), v...))
	})
}

// ProfessorshipGT applies the GT predicate on the "professorship" field.
func ProfessorshipGT(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipGTE applies the GTE predicate on the "professorship" field.
func ProfessorshipGTE(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipLT applies the LT predicate on the "professorship" field.
func ProfessorshipLT(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipLTE applies the LTE predicate on the "professorship" field.
func ProfessorshipLTE(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipContains applies the Contains predicate on the "professorship" field.
func ProfessorshipContains(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipHasPrefix applies the HasPrefix predicate on the "professorship" field.
func ProfessorshipHasPrefix(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipHasSuffix applies the HasSuffix predicate on the "professorship" field.
func ProfessorshipHasSuffix(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipEqualFold applies the EqualFold predicate on the "professorship" field.
func ProfessorshipEqualFold(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProfessorship), v))
	})
}

// ProfessorshipContainsFold applies the ContainsFold predicate on the "professorship" field.
func ProfessorshipContainsFold(v string) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProfessorship), v))
	})
}

// HasProsProf applies the HasEdge predicate on the "pros_prof" edge.
func HasProsProf() predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProsProfTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProsProfTable, ProsProfColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProsProfWith applies the HasEdge predicate on the "pros_prof" edge with a given conditions (other predicates).
func HasProsProfWith(preds ...predicate.Professor) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProsProfInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProsProfTable, ProsProfColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Professorship) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Professorship) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
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
func Not(p predicate.Professorship) predicate.Professorship {
	return predicate.Professorship(func(s *sql.Selector) {
		p(s.Not())
	})
}
