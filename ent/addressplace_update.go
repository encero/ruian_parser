// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encero/ruian_parser/ent/addressplace"
	"github.com/encero/ruian_parser/ent/predicate"
	"github.com/encero/ruian_parser/ent/street"
)

// AddressPlaceUpdate is the builder for updating AddressPlace entities.
type AddressPlaceUpdate struct {
	config
	hooks    []Hook
	mutation *AddressPlaceMutation
}

// Where appends a list predicates to the AddressPlaceUpdate builder.
func (apu *AddressPlaceUpdate) Where(ps ...predicate.AddressPlace) *AddressPlaceUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetNumber sets the "number" field.
func (apu *AddressPlaceUpdate) SetNumber(i int32) *AddressPlaceUpdate {
	apu.mutation.ResetNumber()
	apu.mutation.SetNumber(i)
	return apu
}

// AddNumber adds i to the "number" field.
func (apu *AddressPlaceUpdate) AddNumber(i int32) *AddressPlaceUpdate {
	apu.mutation.AddNumber(i)
	return apu
}

// SetOrientationNumber sets the "orientation_number" field.
func (apu *AddressPlaceUpdate) SetOrientationNumber(i int32) *AddressPlaceUpdate {
	apu.mutation.ResetOrientationNumber()
	apu.mutation.SetOrientationNumber(i)
	return apu
}

// SetNillableOrientationNumber sets the "orientation_number" field if the given value is not nil.
func (apu *AddressPlaceUpdate) SetNillableOrientationNumber(i *int32) *AddressPlaceUpdate {
	if i != nil {
		apu.SetOrientationNumber(*i)
	}
	return apu
}

// AddOrientationNumber adds i to the "orientation_number" field.
func (apu *AddressPlaceUpdate) AddOrientationNumber(i int32) *AddressPlaceUpdate {
	apu.mutation.AddOrientationNumber(i)
	return apu
}

// ClearOrientationNumber clears the value of the "orientation_number" field.
func (apu *AddressPlaceUpdate) ClearOrientationNumber() *AddressPlaceUpdate {
	apu.mutation.ClearOrientationNumber()
	return apu
}

// SetOrientationNumberLetter sets the "orientation_number_letter" field.
func (apu *AddressPlaceUpdate) SetOrientationNumberLetter(s string) *AddressPlaceUpdate {
	apu.mutation.SetOrientationNumberLetter(s)
	return apu
}

// SetNillableOrientationNumberLetter sets the "orientation_number_letter" field if the given value is not nil.
func (apu *AddressPlaceUpdate) SetNillableOrientationNumberLetter(s *string) *AddressPlaceUpdate {
	if s != nil {
		apu.SetOrientationNumberLetter(*s)
	}
	return apu
}

// ClearOrientationNumberLetter clears the value of the "orientation_number_letter" field.
func (apu *AddressPlaceUpdate) ClearOrientationNumberLetter() *AddressPlaceUpdate {
	apu.mutation.ClearOrientationNumberLetter()
	return apu
}

// SetZip sets the "zip" field.
func (apu *AddressPlaceUpdate) SetZip(i int32) *AddressPlaceUpdate {
	apu.mutation.ResetZip()
	apu.mutation.SetZip(i)
	return apu
}

// AddZip adds i to the "zip" field.
func (apu *AddressPlaceUpdate) AddZip(i int32) *AddressPlaceUpdate {
	apu.mutation.AddZip(i)
	return apu
}

// AddStreetIDs adds the "streets" edge to the Street entity by IDs.
func (apu *AddressPlaceUpdate) AddStreetIDs(ids ...int32) *AddressPlaceUpdate {
	apu.mutation.AddStreetIDs(ids...)
	return apu
}

// AddStreets adds the "streets" edges to the Street entity.
func (apu *AddressPlaceUpdate) AddStreets(s ...*Street) *AddressPlaceUpdate {
	ids := make([]int32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apu.AddStreetIDs(ids...)
}

// Mutation returns the AddressPlaceMutation object of the builder.
func (apu *AddressPlaceUpdate) Mutation() *AddressPlaceMutation {
	return apu.mutation
}

// ClearStreets clears all "streets" edges to the Street entity.
func (apu *AddressPlaceUpdate) ClearStreets() *AddressPlaceUpdate {
	apu.mutation.ClearStreets()
	return apu
}

// RemoveStreetIDs removes the "streets" edge to Street entities by IDs.
func (apu *AddressPlaceUpdate) RemoveStreetIDs(ids ...int32) *AddressPlaceUpdate {
	apu.mutation.RemoveStreetIDs(ids...)
	return apu
}

// RemoveStreets removes "streets" edges to Street entities.
func (apu *AddressPlaceUpdate) RemoveStreets(s ...*Street) *AddressPlaceUpdate {
	ids := make([]int32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apu.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AddressPlaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(apu.hooks) == 0 {
		affected, err = apu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressPlaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			apu.mutation = mutation
			affected, err = apu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(apu.hooks) - 1; i >= 0; i-- {
			if apu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = apu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AddressPlaceUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AddressPlaceUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AddressPlaceUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (apu *AddressPlaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   addressplace.Table,
			Columns: addressplace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: addressplace.FieldID,
			},
		},
	}
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldNumber,
		})
	}
	if value, ok := apu.mutation.AddedNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldNumber,
		})
	}
	if value, ok := apu.mutation.OrientationNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldOrientationNumber,
		})
	}
	if value, ok := apu.mutation.AddedOrientationNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldOrientationNumber,
		})
	}
	if apu.mutation.OrientationNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: addressplace.FieldOrientationNumber,
		})
	}
	if value, ok := apu.mutation.OrientationNumberLetter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: addressplace.FieldOrientationNumberLetter,
		})
	}
	if apu.mutation.OrientationNumberLetterCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: addressplace.FieldOrientationNumberLetter,
		})
	}
	if value, ok := apu.mutation.Zip(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldZip,
		})
	}
	if value, ok := apu.mutation.AddedZip(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldZip,
		})
	}
	if apu.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addressplace.StreetsTable,
			Columns: addressplace.StreetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: street.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !apu.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addressplace.StreetsTable,
			Columns: addressplace.StreetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addressplace.StreetsTable,
			Columns: addressplace.StreetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{addressplace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AddressPlaceUpdateOne is the builder for updating a single AddressPlace entity.
type AddressPlaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AddressPlaceMutation
}

// SetNumber sets the "number" field.
func (apuo *AddressPlaceUpdateOne) SetNumber(i int32) *AddressPlaceUpdateOne {
	apuo.mutation.ResetNumber()
	apuo.mutation.SetNumber(i)
	return apuo
}

// AddNumber adds i to the "number" field.
func (apuo *AddressPlaceUpdateOne) AddNumber(i int32) *AddressPlaceUpdateOne {
	apuo.mutation.AddNumber(i)
	return apuo
}

// SetOrientationNumber sets the "orientation_number" field.
func (apuo *AddressPlaceUpdateOne) SetOrientationNumber(i int32) *AddressPlaceUpdateOne {
	apuo.mutation.ResetOrientationNumber()
	apuo.mutation.SetOrientationNumber(i)
	return apuo
}

// SetNillableOrientationNumber sets the "orientation_number" field if the given value is not nil.
func (apuo *AddressPlaceUpdateOne) SetNillableOrientationNumber(i *int32) *AddressPlaceUpdateOne {
	if i != nil {
		apuo.SetOrientationNumber(*i)
	}
	return apuo
}

// AddOrientationNumber adds i to the "orientation_number" field.
func (apuo *AddressPlaceUpdateOne) AddOrientationNumber(i int32) *AddressPlaceUpdateOne {
	apuo.mutation.AddOrientationNumber(i)
	return apuo
}

// ClearOrientationNumber clears the value of the "orientation_number" field.
func (apuo *AddressPlaceUpdateOne) ClearOrientationNumber() *AddressPlaceUpdateOne {
	apuo.mutation.ClearOrientationNumber()
	return apuo
}

// SetOrientationNumberLetter sets the "orientation_number_letter" field.
func (apuo *AddressPlaceUpdateOne) SetOrientationNumberLetter(s string) *AddressPlaceUpdateOne {
	apuo.mutation.SetOrientationNumberLetter(s)
	return apuo
}

// SetNillableOrientationNumberLetter sets the "orientation_number_letter" field if the given value is not nil.
func (apuo *AddressPlaceUpdateOne) SetNillableOrientationNumberLetter(s *string) *AddressPlaceUpdateOne {
	if s != nil {
		apuo.SetOrientationNumberLetter(*s)
	}
	return apuo
}

// ClearOrientationNumberLetter clears the value of the "orientation_number_letter" field.
func (apuo *AddressPlaceUpdateOne) ClearOrientationNumberLetter() *AddressPlaceUpdateOne {
	apuo.mutation.ClearOrientationNumberLetter()
	return apuo
}

// SetZip sets the "zip" field.
func (apuo *AddressPlaceUpdateOne) SetZip(i int32) *AddressPlaceUpdateOne {
	apuo.mutation.ResetZip()
	apuo.mutation.SetZip(i)
	return apuo
}

// AddZip adds i to the "zip" field.
func (apuo *AddressPlaceUpdateOne) AddZip(i int32) *AddressPlaceUpdateOne {
	apuo.mutation.AddZip(i)
	return apuo
}

// AddStreetIDs adds the "streets" edge to the Street entity by IDs.
func (apuo *AddressPlaceUpdateOne) AddStreetIDs(ids ...int32) *AddressPlaceUpdateOne {
	apuo.mutation.AddStreetIDs(ids...)
	return apuo
}

// AddStreets adds the "streets" edges to the Street entity.
func (apuo *AddressPlaceUpdateOne) AddStreets(s ...*Street) *AddressPlaceUpdateOne {
	ids := make([]int32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apuo.AddStreetIDs(ids...)
}

// Mutation returns the AddressPlaceMutation object of the builder.
func (apuo *AddressPlaceUpdateOne) Mutation() *AddressPlaceMutation {
	return apuo.mutation
}

// ClearStreets clears all "streets" edges to the Street entity.
func (apuo *AddressPlaceUpdateOne) ClearStreets() *AddressPlaceUpdateOne {
	apuo.mutation.ClearStreets()
	return apuo
}

// RemoveStreetIDs removes the "streets" edge to Street entities by IDs.
func (apuo *AddressPlaceUpdateOne) RemoveStreetIDs(ids ...int32) *AddressPlaceUpdateOne {
	apuo.mutation.RemoveStreetIDs(ids...)
	return apuo
}

// RemoveStreets removes "streets" edges to Street entities.
func (apuo *AddressPlaceUpdateOne) RemoveStreets(s ...*Street) *AddressPlaceUpdateOne {
	ids := make([]int32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return apuo.RemoveStreetIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AddressPlaceUpdateOne) Select(field string, fields ...string) *AddressPlaceUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AddressPlace entity.
func (apuo *AddressPlaceUpdateOne) Save(ctx context.Context) (*AddressPlace, error) {
	var (
		err  error
		node *AddressPlace
	)
	if len(apuo.hooks) == 0 {
		node, err = apuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressPlaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			apuo.mutation = mutation
			node, err = apuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(apuo.hooks) - 1; i >= 0; i-- {
			if apuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = apuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AddressPlaceUpdateOne) SaveX(ctx context.Context) *AddressPlace {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AddressPlaceUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AddressPlaceUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (apuo *AddressPlaceUpdateOne) sqlSave(ctx context.Context) (_node *AddressPlace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   addressplace.Table,
			Columns: addressplace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: addressplace.FieldID,
			},
		},
	}
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AddressPlace.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, addressplace.FieldID)
		for _, f := range fields {
			if !addressplace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != addressplace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldNumber,
		})
	}
	if value, ok := apuo.mutation.AddedNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldNumber,
		})
	}
	if value, ok := apuo.mutation.OrientationNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldOrientationNumber,
		})
	}
	if value, ok := apuo.mutation.AddedOrientationNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldOrientationNumber,
		})
	}
	if apuo.mutation.OrientationNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: addressplace.FieldOrientationNumber,
		})
	}
	if value, ok := apuo.mutation.OrientationNumberLetter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: addressplace.FieldOrientationNumberLetter,
		})
	}
	if apuo.mutation.OrientationNumberLetterCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: addressplace.FieldOrientationNumberLetter,
		})
	}
	if value, ok := apuo.mutation.Zip(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldZip,
		})
	}
	if value, ok := apuo.mutation.AddedZip(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: addressplace.FieldZip,
		})
	}
	if apuo.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addressplace.StreetsTable,
			Columns: addressplace.StreetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: street.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !apuo.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addressplace.StreetsTable,
			Columns: addressplace.StreetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addressplace.StreetsTable,
			Columns: addressplace.StreetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AddressPlace{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{addressplace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
