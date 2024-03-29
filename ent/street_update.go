// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encero/ruian_parser/ent/addressplace"
	"github.com/encero/ruian_parser/ent/city"
	"github.com/encero/ruian_parser/ent/predicate"
	"github.com/encero/ruian_parser/ent/street"
)

// StreetUpdate is the builder for updating Street entities.
type StreetUpdate struct {
	config
	hooks    []Hook
	mutation *StreetMutation
}

// Where appends a list predicates to the StreetUpdate builder.
func (su *StreetUpdate) Where(ps ...predicate.Street) *StreetUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StreetUpdate) SetName(s string) *StreetUpdate {
	su.mutation.SetName(s)
	return su
}

// AddCityIDs adds the "cities" edge to the City entity by IDs.
func (su *StreetUpdate) AddCityIDs(ids ...int32) *StreetUpdate {
	su.mutation.AddCityIDs(ids...)
	return su
}

// AddCities adds the "cities" edges to the City entity.
func (su *StreetUpdate) AddCities(c ...*City) *StreetUpdate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCityIDs(ids...)
}

// AddAdressIDs adds the "adresses" edge to the AddressPlace entity by IDs.
func (su *StreetUpdate) AddAdressIDs(ids ...int32) *StreetUpdate {
	su.mutation.AddAdressIDs(ids...)
	return su
}

// AddAdresses adds the "adresses" edges to the AddressPlace entity.
func (su *StreetUpdate) AddAdresses(a ...*AddressPlace) *StreetUpdate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddAdressIDs(ids...)
}

// Mutation returns the StreetMutation object of the builder.
func (su *StreetUpdate) Mutation() *StreetMutation {
	return su.mutation
}

// ClearCities clears all "cities" edges to the City entity.
func (su *StreetUpdate) ClearCities() *StreetUpdate {
	su.mutation.ClearCities()
	return su
}

// RemoveCityIDs removes the "cities" edge to City entities by IDs.
func (su *StreetUpdate) RemoveCityIDs(ids ...int32) *StreetUpdate {
	su.mutation.RemoveCityIDs(ids...)
	return su
}

// RemoveCities removes "cities" edges to City entities.
func (su *StreetUpdate) RemoveCities(c ...*City) *StreetUpdate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCityIDs(ids...)
}

// ClearAdresses clears all "adresses" edges to the AddressPlace entity.
func (su *StreetUpdate) ClearAdresses() *StreetUpdate {
	su.mutation.ClearAdresses()
	return su
}

// RemoveAdressIDs removes the "adresses" edge to AddressPlace entities by IDs.
func (su *StreetUpdate) RemoveAdressIDs(ids ...int32) *StreetUpdate {
	su.mutation.RemoveAdressIDs(ids...)
	return su
}

// RemoveAdresses removes "adresses" edges to AddressPlace entities.
func (su *StreetUpdate) RemoveAdresses(a ...*AddressPlace) *StreetUpdate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveAdressIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StreetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StreetUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StreetUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StreetUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StreetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   street.Table,
			Columns: street.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: street.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
	}
	if su.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   street.CitiesTable,
			Columns: street.CitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCitiesIDs(); len(nodes) > 0 && !su.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   street.CitiesTable,
			Columns: street.CitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   street.CitiesTable,
			Columns: street.CitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.AdressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   street.AdressesTable,
			Columns: street.AdressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: addressplace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedAdressesIDs(); len(nodes) > 0 && !su.mutation.AdressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   street.AdressesTable,
			Columns: street.AdressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: addressplace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AdressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   street.AdressesTable,
			Columns: street.AdressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: addressplace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{street.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StreetUpdateOne is the builder for updating a single Street entity.
type StreetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StreetMutation
}

// SetName sets the "name" field.
func (suo *StreetUpdateOne) SetName(s string) *StreetUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// AddCityIDs adds the "cities" edge to the City entity by IDs.
func (suo *StreetUpdateOne) AddCityIDs(ids ...int32) *StreetUpdateOne {
	suo.mutation.AddCityIDs(ids...)
	return suo
}

// AddCities adds the "cities" edges to the City entity.
func (suo *StreetUpdateOne) AddCities(c ...*City) *StreetUpdateOne {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCityIDs(ids...)
}

// AddAdressIDs adds the "adresses" edge to the AddressPlace entity by IDs.
func (suo *StreetUpdateOne) AddAdressIDs(ids ...int32) *StreetUpdateOne {
	suo.mutation.AddAdressIDs(ids...)
	return suo
}

// AddAdresses adds the "adresses" edges to the AddressPlace entity.
func (suo *StreetUpdateOne) AddAdresses(a ...*AddressPlace) *StreetUpdateOne {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddAdressIDs(ids...)
}

// Mutation returns the StreetMutation object of the builder.
func (suo *StreetUpdateOne) Mutation() *StreetMutation {
	return suo.mutation
}

// ClearCities clears all "cities" edges to the City entity.
func (suo *StreetUpdateOne) ClearCities() *StreetUpdateOne {
	suo.mutation.ClearCities()
	return suo
}

// RemoveCityIDs removes the "cities" edge to City entities by IDs.
func (suo *StreetUpdateOne) RemoveCityIDs(ids ...int32) *StreetUpdateOne {
	suo.mutation.RemoveCityIDs(ids...)
	return suo
}

// RemoveCities removes "cities" edges to City entities.
func (suo *StreetUpdateOne) RemoveCities(c ...*City) *StreetUpdateOne {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCityIDs(ids...)
}

// ClearAdresses clears all "adresses" edges to the AddressPlace entity.
func (suo *StreetUpdateOne) ClearAdresses() *StreetUpdateOne {
	suo.mutation.ClearAdresses()
	return suo
}

// RemoveAdressIDs removes the "adresses" edge to AddressPlace entities by IDs.
func (suo *StreetUpdateOne) RemoveAdressIDs(ids ...int32) *StreetUpdateOne {
	suo.mutation.RemoveAdressIDs(ids...)
	return suo
}

// RemoveAdresses removes "adresses" edges to AddressPlace entities.
func (suo *StreetUpdateOne) RemoveAdresses(a ...*AddressPlace) *StreetUpdateOne {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveAdressIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StreetUpdateOne) Select(field string, fields ...string) *StreetUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Street entity.
func (suo *StreetUpdateOne) Save(ctx context.Context) (*Street, error) {
	var (
		err  error
		node *Street
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StreetUpdateOne) SaveX(ctx context.Context) *Street {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StreetUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StreetUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StreetUpdateOne) sqlSave(ctx context.Context) (_node *Street, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   street.Table,
			Columns: street.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: street.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Street.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, street.FieldID)
		for _, f := range fields {
			if !street.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != street.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
	}
	if suo.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   street.CitiesTable,
			Columns: street.CitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: city.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCitiesIDs(); len(nodes) > 0 && !suo.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   street.CitiesTable,
			Columns: street.CitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   street.CitiesTable,
			Columns: street.CitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.AdressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   street.AdressesTable,
			Columns: street.AdressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: addressplace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedAdressesIDs(); len(nodes) > 0 && !suo.mutation.AdressesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   street.AdressesTable,
			Columns: street.AdressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: addressplace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AdressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   street.AdressesTable,
			Columns: street.AdressesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: addressplace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Street{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{street.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
