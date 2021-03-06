// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/course"
	"github.com/sut63/team17/app/ent/institution"
	"github.com/sut63/team17/app/ent/predicate"
)

// InstitutionUpdate is the builder for updating Institution entities.
type InstitutionUpdate struct {
	config
	hooks      []Hook
	mutation   *InstitutionMutation
	predicates []predicate.Institution
}

// Where adds a new predicate for the builder.
func (iu *InstitutionUpdate) Where(ps ...predicate.Institution) *InstitutionUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// SetInstitution sets the institution field.
func (iu *InstitutionUpdate) SetInstitution(s string) *InstitutionUpdate {
	iu.mutation.SetInstitution(s)
	return iu
}

// AddInstCourIDs adds the inst_cour edge to Course by ids.
func (iu *InstitutionUpdate) AddInstCourIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.AddInstCourIDs(ids...)
	return iu
}

// AddInstCour adds the inst_cour edges to Course.
func (iu *InstitutionUpdate) AddInstCour(c ...*Course) *InstitutionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iu.AddInstCourIDs(ids...)
}

// Mutation returns the InstitutionMutation object of the builder.
func (iu *InstitutionUpdate) Mutation() *InstitutionMutation {
	return iu.mutation
}

// RemoveInstCourIDs removes the inst_cour edge to Course by ids.
func (iu *InstitutionUpdate) RemoveInstCourIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.RemoveInstCourIDs(ids...)
	return iu
}

// RemoveInstCour removes inst_cour edges to Course.
func (iu *InstitutionUpdate) RemoveInstCour(c ...*Course) *InstitutionUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iu.RemoveInstCourIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *InstitutionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := iu.mutation.Institution(); ok {
		if err := institution.InstitutionValidator(v); err != nil {
			return 0, &ValidationError{Name: "institution", err: fmt.Errorf("ent: validator failed for field \"institution\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstitutionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InstitutionUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InstitutionUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InstitutionUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *InstitutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   institution.Table,
			Columns: institution.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: institution.FieldID,
			},
		},
	}
	if ps := iu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Institution(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: institution.FieldInstitution,
		})
	}
	if nodes := iu.mutation.RemovedInstCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.InstCourTable,
			Columns: []string{institution.InstCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.InstCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.InstCourTable,
			Columns: []string{institution.InstCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{institution.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InstitutionUpdateOne is the builder for updating a single Institution entity.
type InstitutionUpdateOne struct {
	config
	hooks    []Hook
	mutation *InstitutionMutation
}

// SetInstitution sets the institution field.
func (iuo *InstitutionUpdateOne) SetInstitution(s string) *InstitutionUpdateOne {
	iuo.mutation.SetInstitution(s)
	return iuo
}

// AddInstCourIDs adds the inst_cour edge to Course by ids.
func (iuo *InstitutionUpdateOne) AddInstCourIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.AddInstCourIDs(ids...)
	return iuo
}

// AddInstCour adds the inst_cour edges to Course.
func (iuo *InstitutionUpdateOne) AddInstCour(c ...*Course) *InstitutionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iuo.AddInstCourIDs(ids...)
}

// Mutation returns the InstitutionMutation object of the builder.
func (iuo *InstitutionUpdateOne) Mutation() *InstitutionMutation {
	return iuo.mutation
}

// RemoveInstCourIDs removes the inst_cour edge to Course by ids.
func (iuo *InstitutionUpdateOne) RemoveInstCourIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.RemoveInstCourIDs(ids...)
	return iuo
}

// RemoveInstCour removes inst_cour edges to Course.
func (iuo *InstitutionUpdateOne) RemoveInstCour(c ...*Course) *InstitutionUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iuo.RemoveInstCourIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (iuo *InstitutionUpdateOne) Save(ctx context.Context) (*Institution, error) {
	if v, ok := iuo.mutation.Institution(); ok {
		if err := institution.InstitutionValidator(v); err != nil {
			return nil, &ValidationError{Name: "institution", err: fmt.Errorf("ent: validator failed for field \"institution\": %w", err)}
		}
	}

	var (
		err  error
		node *Institution
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstitutionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InstitutionUpdateOne) SaveX(ctx context.Context) *Institution {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *InstitutionUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InstitutionUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *InstitutionUpdateOne) sqlSave(ctx context.Context) (i *Institution, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   institution.Table,
			Columns: institution.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: institution.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Institution.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.Institution(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: institution.FieldInstitution,
		})
	}
	if nodes := iuo.mutation.RemovedInstCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.InstCourTable,
			Columns: []string{institution.InstCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.InstCourIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.InstCourTable,
			Columns: []string{institution.InstCourColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	i = &Institution{config: iuo.config}
	_spec.Assign = i.assignValues
	_spec.ScanValues = i.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{institution.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}
