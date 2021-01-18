// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team17/app/ent/faculty"
	"github.com/sut63/team17/app/ent/predicate"
	"github.com/sut63/team17/app/ent/prefix"
	"github.com/sut63/team17/app/ent/professor"
	"github.com/sut63/team17/app/ent/professorship"
)

// ProfessorUpdate is the builder for updating Professor entities.
type ProfessorUpdate struct {
	config
	hooks      []Hook
	mutation   *ProfessorMutation
	predicates []predicate.Professor
}

// Where adds a new predicate for the builder.
func (pu *ProfessorUpdate) Where(ps ...predicate.Professor) *ProfessorUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *ProfessorUpdate) SetName(s string) *ProfessorUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetTel sets the tel field.
func (pu *ProfessorUpdate) SetTel(s string) *ProfessorUpdate {
	pu.mutation.SetTel(s)
	return pu
}

// SetEmail sets the email field.
func (pu *ProfessorUpdate) SetEmail(s string) *ProfessorUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetProfPreID sets the prof_pre edge to Prefix by id.
func (pu *ProfessorUpdate) SetProfPreID(id int) *ProfessorUpdate {
	pu.mutation.SetProfPreID(id)
	return pu
}

// SetNillableProfPreID sets the prof_pre edge to Prefix by id if the given value is not nil.
func (pu *ProfessorUpdate) SetNillableProfPreID(id *int) *ProfessorUpdate {
	if id != nil {
		pu = pu.SetProfPreID(*id)
	}
	return pu
}

// SetProfPre sets the prof_pre edge to Prefix.
func (pu *ProfessorUpdate) SetProfPre(p *Prefix) *ProfessorUpdate {
	return pu.SetProfPreID(p.ID)
}

// SetProfFacuID sets the prof_facu edge to Faculty by id.
func (pu *ProfessorUpdate) SetProfFacuID(id int) *ProfessorUpdate {
	pu.mutation.SetProfFacuID(id)
	return pu
}

// SetNillableProfFacuID sets the prof_facu edge to Faculty by id if the given value is not nil.
func (pu *ProfessorUpdate) SetNillableProfFacuID(id *int) *ProfessorUpdate {
	if id != nil {
		pu = pu.SetProfFacuID(*id)
	}
	return pu
}

// SetProfFacu sets the prof_facu edge to Faculty.
func (pu *ProfessorUpdate) SetProfFacu(f *Faculty) *ProfessorUpdate {
	return pu.SetProfFacuID(f.ID)
}

// SetProfProsID sets the prof_pros edge to Professorship by id.
func (pu *ProfessorUpdate) SetProfProsID(id int) *ProfessorUpdate {
	pu.mutation.SetProfProsID(id)
	return pu
}

// SetNillableProfProsID sets the prof_pros edge to Professorship by id if the given value is not nil.
func (pu *ProfessorUpdate) SetNillableProfProsID(id *int) *ProfessorUpdate {
	if id != nil {
		pu = pu.SetProfProsID(*id)
	}
	return pu
}

// SetProfPros sets the prof_pros edge to Professorship.
func (pu *ProfessorUpdate) SetProfPros(p *Professorship) *ProfessorUpdate {
	return pu.SetProfProsID(p.ID)
}

// Mutation returns the ProfessorMutation object of the builder.
func (pu *ProfessorUpdate) Mutation() *ProfessorMutation {
	return pu.mutation
}

// ClearProfPre clears the prof_pre edge to Prefix.
func (pu *ProfessorUpdate) ClearProfPre() *ProfessorUpdate {
	pu.mutation.ClearProfPre()
	return pu
}

// ClearProfFacu clears the prof_facu edge to Faculty.
func (pu *ProfessorUpdate) ClearProfFacu() *ProfessorUpdate {
	pu.mutation.ClearProfFacu()
	return pu
}

// ClearProfPros clears the prof_pros edge to Professorship.
func (pu *ProfessorUpdate) ClearProfPros() *ProfessorUpdate {
	pu.mutation.ClearProfPros()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProfessorUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfessorMutation)
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
func (pu *ProfessorUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfessorUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfessorUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProfessorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   professor.Table,
			Columns: professor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: professor.FieldID,
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
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professor.FieldName,
		})
	}
	if value, ok := pu.mutation.Tel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professor.FieldTel,
		})
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professor.FieldEmail,
		})
	}
	if pu.mutation.ProfPreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfPreTable,
			Columns: []string{professor.ProfPreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProfPreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfPreTable,
			Columns: []string{professor.ProfPreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProfFacuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfFacuTable,
			Columns: []string{professor.ProfFacuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProfFacuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfFacuTable,
			Columns: []string{professor.ProfFacuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProfProsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfProsTable,
			Columns: []string{professor.ProfProsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professorship.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProfProsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfProsTable,
			Columns: []string{professor.ProfProsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professorship.FieldID,
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
			err = &NotFoundError{professor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProfessorUpdateOne is the builder for updating a single Professor entity.
type ProfessorUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProfessorMutation
}

// SetName sets the name field.
func (puo *ProfessorUpdateOne) SetName(s string) *ProfessorUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetTel sets the tel field.
func (puo *ProfessorUpdateOne) SetTel(s string) *ProfessorUpdateOne {
	puo.mutation.SetTel(s)
	return puo
}

// SetEmail sets the email field.
func (puo *ProfessorUpdateOne) SetEmail(s string) *ProfessorUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetProfPreID sets the prof_pre edge to Prefix by id.
func (puo *ProfessorUpdateOne) SetProfPreID(id int) *ProfessorUpdateOne {
	puo.mutation.SetProfPreID(id)
	return puo
}

// SetNillableProfPreID sets the prof_pre edge to Prefix by id if the given value is not nil.
func (puo *ProfessorUpdateOne) SetNillableProfPreID(id *int) *ProfessorUpdateOne {
	if id != nil {
		puo = puo.SetProfPreID(*id)
	}
	return puo
}

// SetProfPre sets the prof_pre edge to Prefix.
func (puo *ProfessorUpdateOne) SetProfPre(p *Prefix) *ProfessorUpdateOne {
	return puo.SetProfPreID(p.ID)
}

// SetProfFacuID sets the prof_facu edge to Faculty by id.
func (puo *ProfessorUpdateOne) SetProfFacuID(id int) *ProfessorUpdateOne {
	puo.mutation.SetProfFacuID(id)
	return puo
}

// SetNillableProfFacuID sets the prof_facu edge to Faculty by id if the given value is not nil.
func (puo *ProfessorUpdateOne) SetNillableProfFacuID(id *int) *ProfessorUpdateOne {
	if id != nil {
		puo = puo.SetProfFacuID(*id)
	}
	return puo
}

// SetProfFacu sets the prof_facu edge to Faculty.
func (puo *ProfessorUpdateOne) SetProfFacu(f *Faculty) *ProfessorUpdateOne {
	return puo.SetProfFacuID(f.ID)
}

// SetProfProsID sets the prof_pros edge to Professorship by id.
func (puo *ProfessorUpdateOne) SetProfProsID(id int) *ProfessorUpdateOne {
	puo.mutation.SetProfProsID(id)
	return puo
}

// SetNillableProfProsID sets the prof_pros edge to Professorship by id if the given value is not nil.
func (puo *ProfessorUpdateOne) SetNillableProfProsID(id *int) *ProfessorUpdateOne {
	if id != nil {
		puo = puo.SetProfProsID(*id)
	}
	return puo
}

// SetProfPros sets the prof_pros edge to Professorship.
func (puo *ProfessorUpdateOne) SetProfPros(p *Professorship) *ProfessorUpdateOne {
	return puo.SetProfProsID(p.ID)
}

// Mutation returns the ProfessorMutation object of the builder.
func (puo *ProfessorUpdateOne) Mutation() *ProfessorMutation {
	return puo.mutation
}

// ClearProfPre clears the prof_pre edge to Prefix.
func (puo *ProfessorUpdateOne) ClearProfPre() *ProfessorUpdateOne {
	puo.mutation.ClearProfPre()
	return puo
}

// ClearProfFacu clears the prof_facu edge to Faculty.
func (puo *ProfessorUpdateOne) ClearProfFacu() *ProfessorUpdateOne {
	puo.mutation.ClearProfFacu()
	return puo
}

// ClearProfPros clears the prof_pros edge to Professorship.
func (puo *ProfessorUpdateOne) ClearProfPros() *ProfessorUpdateOne {
	puo.mutation.ClearProfPros()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *ProfessorUpdateOne) Save(ctx context.Context) (*Professor, error) {

	var (
		err  error
		node *Professor
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfessorMutation)
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
func (puo *ProfessorUpdateOne) SaveX(ctx context.Context) *Professor {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProfessorUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfessorUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProfessorUpdateOne) sqlSave(ctx context.Context) (pr *Professor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   professor.Table,
			Columns: professor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: professor.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Professor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professor.FieldName,
		})
	}
	if value, ok := puo.mutation.Tel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professor.FieldTel,
		})
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: professor.FieldEmail,
		})
	}
	if puo.mutation.ProfPreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfPreTable,
			Columns: []string{professor.ProfPreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProfPreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfPreTable,
			Columns: []string{professor.ProfPreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prefix.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProfFacuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfFacuTable,
			Columns: []string{professor.ProfFacuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProfFacuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfFacuTable,
			Columns: []string{professor.ProfFacuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: faculty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProfProsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfProsTable,
			Columns: []string{professor.ProfProsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professorship.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProfProsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   professor.ProfProsTable,
			Columns: []string{professor.ProfProsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: professorship.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Professor{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{professor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}