// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/predicate"
	"github.com/sut63/team17/app/ent/results"
	"github.com/sut63/team17/app/ent/student"
	"github.com/sut63/team17/app/ent/subject"
	"github.com/sut63/team17/app/ent/term"
	"github.com/sut63/team17/app/ent/year"
)

// ResultsUpdate is the builder for updating Results entities.
type ResultsUpdate struct {
	config
	hooks      []Hook
	mutation   *ResultsMutation
	predicates []predicate.Results
}

// Where adds a new predicate for the builder.
func (ru *ResultsUpdate) Where(ps ...predicate.Results) *ResultsUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetGrade sets the grade field.
func (ru *ResultsUpdate) SetGrade(f float64) *ResultsUpdate {
	ru.mutation.ResetGrade()
	ru.mutation.SetGrade(f)
	return ru
}

// AddGrade adds f to grade.
func (ru *ResultsUpdate) AddGrade(f float64) *ResultsUpdate {
	ru.mutation.AddGrade(f)
	return ru
}

// SetResuYearID sets the resu_year edge to Year by id.
func (ru *ResultsUpdate) SetResuYearID(id int) *ResultsUpdate {
	ru.mutation.SetResuYearID(id)
	return ru
}

// SetNillableResuYearID sets the resu_year edge to Year by id if the given value is not nil.
func (ru *ResultsUpdate) SetNillableResuYearID(id *int) *ResultsUpdate {
	if id != nil {
		ru = ru.SetResuYearID(*id)
	}
	return ru
}

// SetResuYear sets the resu_year edge to Year.
func (ru *ResultsUpdate) SetResuYear(y *Year) *ResultsUpdate {
	return ru.SetResuYearID(y.ID)
}

// SetResuSubjID sets the resu_subj edge to Subject by id.
func (ru *ResultsUpdate) SetResuSubjID(id int) *ResultsUpdate {
	ru.mutation.SetResuSubjID(id)
	return ru
}

// SetNillableResuSubjID sets the resu_subj edge to Subject by id if the given value is not nil.
func (ru *ResultsUpdate) SetNillableResuSubjID(id *int) *ResultsUpdate {
	if id != nil {
		ru = ru.SetResuSubjID(*id)
	}
	return ru
}

// SetResuSubj sets the resu_subj edge to Subject.
func (ru *ResultsUpdate) SetResuSubj(s *Subject) *ResultsUpdate {
	return ru.SetResuSubjID(s.ID)
}

// SetResuStudID sets the resu_stud edge to Student by id.
func (ru *ResultsUpdate) SetResuStudID(id int) *ResultsUpdate {
	ru.mutation.SetResuStudID(id)
	return ru
}

// SetNillableResuStudID sets the resu_stud edge to Student by id if the given value is not nil.
func (ru *ResultsUpdate) SetNillableResuStudID(id *int) *ResultsUpdate {
	if id != nil {
		ru = ru.SetResuStudID(*id)
	}
	return ru
}

// SetResuStud sets the resu_stud edge to Student.
func (ru *ResultsUpdate) SetResuStud(s *Student) *ResultsUpdate {
	return ru.SetResuStudID(s.ID)
}

// SetResuTermID sets the resu_term edge to Term by id.
func (ru *ResultsUpdate) SetResuTermID(id int) *ResultsUpdate {
	ru.mutation.SetResuTermID(id)
	return ru
}

// SetNillableResuTermID sets the resu_term edge to Term by id if the given value is not nil.
func (ru *ResultsUpdate) SetNillableResuTermID(id *int) *ResultsUpdate {
	if id != nil {
		ru = ru.SetResuTermID(*id)
	}
	return ru
}

// SetResuTerm sets the resu_term edge to Term.
func (ru *ResultsUpdate) SetResuTerm(t *Term) *ResultsUpdate {
	return ru.SetResuTermID(t.ID)
}

// Mutation returns the ResultsMutation object of the builder.
func (ru *ResultsUpdate) Mutation() *ResultsMutation {
	return ru.mutation
}

// ClearResuYear clears the resu_year edge to Year.
func (ru *ResultsUpdate) ClearResuYear() *ResultsUpdate {
	ru.mutation.ClearResuYear()
	return ru
}

// ClearResuSubj clears the resu_subj edge to Subject.
func (ru *ResultsUpdate) ClearResuSubj() *ResultsUpdate {
	ru.mutation.ClearResuSubj()
	return ru
}

// ClearResuStud clears the resu_stud edge to Student.
func (ru *ResultsUpdate) ClearResuStud() *ResultsUpdate {
	ru.mutation.ClearResuStud()
	return ru
}

// ClearResuTerm clears the resu_term edge to Term.
func (ru *ResultsUpdate) ClearResuTerm() *ResultsUpdate {
	ru.mutation.ClearResuTerm()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *ResultsUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.Grade(); ok {
		if err := results.GradeValidator(v); err != nil {
			return 0, &ValidationError{Name: "grade", err: fmt.Errorf("ent: validator failed for field \"grade\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResultsUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResultsUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResultsUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ResultsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   results.Table,
			Columns: results.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: results.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Grade(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: results.FieldGrade,
		})
	}
	if value, ok := ru.mutation.AddedGrade(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: results.FieldGrade,
		})
	}
	if ru.mutation.ResuYearCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuYearTable,
			Columns: []string{results.ResuYearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: year.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ResuYearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuYearTable,
			Columns: []string{results.ResuYearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: year.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ResuSubjCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuSubjTable,
			Columns: []string{results.ResuSubjColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ResuSubjIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuSubjTable,
			Columns: []string{results.ResuSubjColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ResuStudCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuStudTable,
			Columns: []string{results.ResuStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ResuStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuStudTable,
			Columns: []string{results.ResuStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ResuTermCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuTermTable,
			Columns: []string{results.ResuTermColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: term.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ResuTermIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuTermTable,
			Columns: []string{results.ResuTermColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: term.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{results.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ResultsUpdateOne is the builder for updating a single Results entity.
type ResultsUpdateOne struct {
	config
	hooks    []Hook
	mutation *ResultsMutation
}

// SetGrade sets the grade field.
func (ruo *ResultsUpdateOne) SetGrade(f float64) *ResultsUpdateOne {
	ruo.mutation.ResetGrade()
	ruo.mutation.SetGrade(f)
	return ruo
}

// AddGrade adds f to grade.
func (ruo *ResultsUpdateOne) AddGrade(f float64) *ResultsUpdateOne {
	ruo.mutation.AddGrade(f)
	return ruo
}

// SetResuYearID sets the resu_year edge to Year by id.
func (ruo *ResultsUpdateOne) SetResuYearID(id int) *ResultsUpdateOne {
	ruo.mutation.SetResuYearID(id)
	return ruo
}

// SetNillableResuYearID sets the resu_year edge to Year by id if the given value is not nil.
func (ruo *ResultsUpdateOne) SetNillableResuYearID(id *int) *ResultsUpdateOne {
	if id != nil {
		ruo = ruo.SetResuYearID(*id)
	}
	return ruo
}

// SetResuYear sets the resu_year edge to Year.
func (ruo *ResultsUpdateOne) SetResuYear(y *Year) *ResultsUpdateOne {
	return ruo.SetResuYearID(y.ID)
}

// SetResuSubjID sets the resu_subj edge to Subject by id.
func (ruo *ResultsUpdateOne) SetResuSubjID(id int) *ResultsUpdateOne {
	ruo.mutation.SetResuSubjID(id)
	return ruo
}

// SetNillableResuSubjID sets the resu_subj edge to Subject by id if the given value is not nil.
func (ruo *ResultsUpdateOne) SetNillableResuSubjID(id *int) *ResultsUpdateOne {
	if id != nil {
		ruo = ruo.SetResuSubjID(*id)
	}
	return ruo
}

// SetResuSubj sets the resu_subj edge to Subject.
func (ruo *ResultsUpdateOne) SetResuSubj(s *Subject) *ResultsUpdateOne {
	return ruo.SetResuSubjID(s.ID)
}

// SetResuStudID sets the resu_stud edge to Student by id.
func (ruo *ResultsUpdateOne) SetResuStudID(id int) *ResultsUpdateOne {
	ruo.mutation.SetResuStudID(id)
	return ruo
}

// SetNillableResuStudID sets the resu_stud edge to Student by id if the given value is not nil.
func (ruo *ResultsUpdateOne) SetNillableResuStudID(id *int) *ResultsUpdateOne {
	if id != nil {
		ruo = ruo.SetResuStudID(*id)
	}
	return ruo
}

// SetResuStud sets the resu_stud edge to Student.
func (ruo *ResultsUpdateOne) SetResuStud(s *Student) *ResultsUpdateOne {
	return ruo.SetResuStudID(s.ID)
}

// SetResuTermID sets the resu_term edge to Term by id.
func (ruo *ResultsUpdateOne) SetResuTermID(id int) *ResultsUpdateOne {
	ruo.mutation.SetResuTermID(id)
	return ruo
}

// SetNillableResuTermID sets the resu_term edge to Term by id if the given value is not nil.
func (ruo *ResultsUpdateOne) SetNillableResuTermID(id *int) *ResultsUpdateOne {
	if id != nil {
		ruo = ruo.SetResuTermID(*id)
	}
	return ruo
}

// SetResuTerm sets the resu_term edge to Term.
func (ruo *ResultsUpdateOne) SetResuTerm(t *Term) *ResultsUpdateOne {
	return ruo.SetResuTermID(t.ID)
}

// Mutation returns the ResultsMutation object of the builder.
func (ruo *ResultsUpdateOne) Mutation() *ResultsMutation {
	return ruo.mutation
}

// ClearResuYear clears the resu_year edge to Year.
func (ruo *ResultsUpdateOne) ClearResuYear() *ResultsUpdateOne {
	ruo.mutation.ClearResuYear()
	return ruo
}

// ClearResuSubj clears the resu_subj edge to Subject.
func (ruo *ResultsUpdateOne) ClearResuSubj() *ResultsUpdateOne {
	ruo.mutation.ClearResuSubj()
	return ruo
}

// ClearResuStud clears the resu_stud edge to Student.
func (ruo *ResultsUpdateOne) ClearResuStud() *ResultsUpdateOne {
	ruo.mutation.ClearResuStud()
	return ruo
}

// ClearResuTerm clears the resu_term edge to Term.
func (ruo *ResultsUpdateOne) ClearResuTerm() *ResultsUpdateOne {
	ruo.mutation.ClearResuTerm()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *ResultsUpdateOne) Save(ctx context.Context) (*Results, error) {
	if v, ok := ruo.mutation.Grade(); ok {
		if err := results.GradeValidator(v); err != nil {
			return nil, &ValidationError{Name: "grade", err: fmt.Errorf("ent: validator failed for field \"grade\": %w", err)}
		}
	}

	var (
		err  error
		node *Results
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResultsUpdateOne) SaveX(ctx context.Context) *Results {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *ResultsUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResultsUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ResultsUpdateOne) sqlSave(ctx context.Context) (r *Results, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   results.Table,
			Columns: results.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: results.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Results.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Grade(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: results.FieldGrade,
		})
	}
	if value, ok := ruo.mutation.AddedGrade(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: results.FieldGrade,
		})
	}
	if ruo.mutation.ResuYearCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuYearTable,
			Columns: []string{results.ResuYearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: year.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ResuYearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuYearTable,
			Columns: []string{results.ResuYearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: year.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ResuSubjCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuSubjTable,
			Columns: []string{results.ResuSubjColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ResuSubjIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuSubjTable,
			Columns: []string{results.ResuSubjColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ResuStudCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuStudTable,
			Columns: []string{results.ResuStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ResuStudIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuStudTable,
			Columns: []string{results.ResuStudColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ResuTermCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuTermTable,
			Columns: []string{results.ResuTermColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: term.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ResuTermIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   results.ResuTermTable,
			Columns: []string{results.ResuTermColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: term.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Results{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{results.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}