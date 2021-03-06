// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/predicate"
	"github.com/sut63/team17/app/ent/professor"
	"github.com/sut63/team17/app/ent/professorship"
)

// ProfessorshipUpdate is the builder for updating Professorship entities.
type ProfessorshipUpdate struct {
	config
	hooks      []Hook
	mutation   *ProfessorshipMutation
	predicates []predicate.Professorship
}

// Where adds a new predicate for the builder.
func (pu *ProfessorshipUpdate) Where(ps ...predicate.Professorship) *ProfessorshipUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetProfessorship sets the professorship field.
func (pu *ProfessorshipUpdate) SetProfessorship(s string) *ProfessorshipUpdate {
	pu.mutation.SetProfessorship(s)
	return pu
}

// AddProsProfIDs adds the pros_prof edge to Professor by ids.
func (pu *ProfessorshipUpdate) AddProsProfIDs(ids ...int) *ProfessorshipUpdate {
	pu.mutation.AddProsProfIDs(ids...)
	return pu
}

// AddProsProf adds the pros_prof edges to Professor.
func (pu *ProfessorshipUpdate) AddProsProf(p ...*Professor) *ProfessorshipUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProsProfIDs(ids...)
}

// Mutation returns the ProfessorshipMutation object of the builder.
func (pu *ProfessorshipUpdate) Mutation() *ProfessorshipMutation {
	return pu.mutation
}

// RemoveProsProfIDs removes the pros_prof edge to Professor by ids.
func (pu *ProfessorshipUpdate) RemoveProsProfIDs(ids ...int) *ProfessorshipUpdate {
	pu.mutation.RemoveProsProfIDs(ids...)
	return pu
}

// RemoveProsProf removes pros_prof edges to Professor.
func (pu *ProfessorshipUpdate) RemoveProsProf(p ...*Professor) *ProfessorshipUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProsProfIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProfessorshipUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Professorship(); ok {
		if err := professorship.ProfessorshipValidator(v); err != nil {
			return 0, &ValidationError{Name: "professorship", err: fmt.Errorf("ent: validator failed for field \"professorship\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfessorshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfessorshipUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfessorshipUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfessorshipUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProfessorshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   professorship.Table,
			Columns: professorship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: professorship.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Professorship(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professorship.FieldProfessorship,
		})
	}
	if nodes := pu.mutation.RemovedProsProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professorship.ProsProfTable,
			Columns: []string{professorship.ProsProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProsProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professorship.ProsProfTable,
			Columns: []string{professorship.ProsProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{professorship.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProfessorshipUpdateOne is the builder for updating a single Professorship entity.
type ProfessorshipUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProfessorshipMutation
}

// SetProfessorship sets the professorship field.
func (puo *ProfessorshipUpdateOne) SetProfessorship(s string) *ProfessorshipUpdateOne {
	puo.mutation.SetProfessorship(s)
	return puo
}

// AddProsProfIDs adds the pros_prof edge to Professor by ids.
func (puo *ProfessorshipUpdateOne) AddProsProfIDs(ids ...int) *ProfessorshipUpdateOne {
	puo.mutation.AddProsProfIDs(ids...)
	return puo
}

// AddProsProf adds the pros_prof edges to Professor.
func (puo *ProfessorshipUpdateOne) AddProsProf(p ...*Professor) *ProfessorshipUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProsProfIDs(ids...)
}

// Mutation returns the ProfessorshipMutation object of the builder.
func (puo *ProfessorshipUpdateOne) Mutation() *ProfessorshipMutation {
	return puo.mutation
}

// RemoveProsProfIDs removes the pros_prof edge to Professor by ids.
func (puo *ProfessorshipUpdateOne) RemoveProsProfIDs(ids ...int) *ProfessorshipUpdateOne {
	puo.mutation.RemoveProsProfIDs(ids...)
	return puo
}

// RemoveProsProf removes pros_prof edges to Professor.
func (puo *ProfessorshipUpdateOne) RemoveProsProf(p ...*Professor) *ProfessorshipUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProsProfIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProfessorshipUpdateOne) Save(ctx context.Context) (*Professorship, error) {
	if v, ok := puo.mutation.Professorship(); ok {
		if err := professorship.ProfessorshipValidator(v); err != nil {
			return nil, &ValidationError{Name: "professorship", err: fmt.Errorf("ent: validator failed for field \"professorship\": %w", err)}
		}
	}

	var (
		err  error
		node *Professorship
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfessorshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfessorshipUpdateOne) SaveX(ctx context.Context) *Professorship {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProfessorshipUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfessorshipUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProfessorshipUpdateOne) sqlSave(ctx context.Context) (pr *Professorship, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   professorship.Table,
			Columns: professorship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: professorship.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Professorship.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Professorship(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professorship.FieldProfessorship,
		})
	}
	if nodes := puo.mutation.RemovedProsProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professorship.ProsProfTable,
			Columns: []string{professorship.ProsProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProsProfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professorship.ProsProfTable,
			Columns: []string{professorship.ProsProfColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Professorship{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{professorship.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
