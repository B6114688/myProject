// Code generated by entc, DO NOT EDIT.

package faculty

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team17/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Faculty applies equality check predicate on the "faculty" field. It's identical to FacultyEQ.
func Faculty(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFaculty), v))
	})
}

// FacultyEQ applies the EQ predicate on the "faculty" field.
func FacultyEQ(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFaculty), v))
	})
}

// FacultyNEQ applies the NEQ predicate on the "faculty" field.
func FacultyNEQ(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFaculty), v))
	})
}

// FacultyIn applies the In predicate on the "faculty" field.
func FacultyIn(vs ...string) predicate.Faculty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Faculty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFaculty), v...))
	})
}

// FacultyNotIn applies the NotIn predicate on the "faculty" field.
func FacultyNotIn(vs ...string) predicate.Faculty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Faculty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFaculty), v...))
	})
}

// FacultyGT applies the GT predicate on the "faculty" field.
func FacultyGT(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFaculty), v))
	})
}

// FacultyGTE applies the GTE predicate on the "faculty" field.
func FacultyGTE(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFaculty), v))
	})
}

// FacultyLT applies the LT predicate on the "faculty" field.
func FacultyLT(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFaculty), v))
	})
}

// FacultyLTE applies the LTE predicate on the "faculty" field.
func FacultyLTE(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFaculty), v))
	})
}

// FacultyContains applies the Contains predicate on the "faculty" field.
func FacultyContains(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFaculty), v))
	})
}

// FacultyHasPrefix applies the HasPrefix predicate on the "faculty" field.
func FacultyHasPrefix(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFaculty), v))
	})
}

// FacultyHasSuffix applies the HasSuffix predicate on the "faculty" field.
func FacultyHasSuffix(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFaculty), v))
	})
}

// FacultyEqualFold applies the EqualFold predicate on the "faculty" field.
func FacultyEqualFold(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFaculty), v))
	})
}

// FacultyContainsFold applies the ContainsFold predicate on the "faculty" field.
func FacultyContainsFold(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFaculty), v))
	})
}

// HasFacuCour applies the HasEdge predicate on the "facu_cour" edge.
func HasFacuCour() predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FacuCourTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FacuCourTable, FacuCourColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFacuCourWith applies the HasEdge predicate on the "facu_cour" edge with a given conditions (other predicates).
func HasFacuCourWith(preds ...predicate.Course) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FacuCourInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FacuCourTable, FacuCourColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFacuProf applies the HasEdge predicate on the "facu_prof" edge.
func HasFacuProf() predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FacuProfTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FacuProfTable, FacuProfColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFacuProfWith applies the HasEdge predicate on the "facu_prof" edge with a given conditions (other predicates).
func HasFacuProfWith(preds ...predicate.Professor) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FacuProfInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FacuProfTable, FacuProfColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Faculty) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Faculty) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
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
func Not(p predicate.Faculty) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		p(s.Not())
	})
}
