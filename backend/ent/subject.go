// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team17/app/ent/subject"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code int `json:"code,omitempty"`
	// Subjects holds the value of the "subjects" field.
	Subjects string `json:"subjects,omitempty"`
	// Creditpiont holds the value of the "creditpiont" field.
	Creditpiont int `json:"creditpiont,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges SubjectEdges `json:"edges"`
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// SubjResu holds the value of the subj_resu edge.
	SubjResu []*Results
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubjResuOrErr returns the SubjResu value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) SubjResuOrErr() ([]*Results, error) {
	if e.loadedTypes[0] {
		return e.SubjResu, nil
	}
	return nil, &NotLoadedError{edge: "subj_resu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // code
		&sql.NullString{}, // subjects
		&sql.NullInt64{},  // creditpiont
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(values ...interface{}) error {
	if m, n := len(values), len(subject.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field code", values[0])
	} else if value.Valid {
		s.Code = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field subjects", values[1])
	} else if value.Valid {
		s.Subjects = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field creditpiont", values[2])
	} else if value.Valid {
		s.Creditpiont = int(value.Int64)
	}
	return nil
}

// QuerySubjResu queries the subj_resu edge of the Subject.
func (s *Subject) QuerySubjResu() *ResultsQuery {
	return (&SubjectClient{config: s.config}).QuerySubjResu(s)
}

// Update returns a builder for updating this Subject.
// Note that, you need to call Subject.Unwrap() before calling this method, if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return (&SubjectClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", code=")
	builder.WriteString(fmt.Sprintf("%v", s.Code))
	builder.WriteString(", subjects=")
	builder.WriteString(s.Subjects)
	builder.WriteString(", creditpiont=")
	builder.WriteString(fmt.Sprintf("%v", s.Creditpiont))
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject

func (s Subjects) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
