// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/encero/ruian_parser/ent/addressplace"
	"github.com/encero/ruian_parser/ent/city"
	"github.com/encero/ruian_parser/ent/street"
)

// StreetCreate is the builder for creating a Street entity.
type StreetCreate struct {
	config
	mutation *StreetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sc *StreetCreate) SetName(s string) *StreetCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetID sets the "id" field.
func (sc *StreetCreate) SetID(i int32) *StreetCreate {
	sc.mutation.SetID(i)
	return sc
}

// AddCityIDs adds the "cities" edge to the City entity by IDs.
func (sc *StreetCreate) AddCityIDs(ids ...int32) *StreetCreate {
	sc.mutation.AddCityIDs(ids...)
	return sc
}

// AddCities adds the "cities" edges to the City entity.
func (sc *StreetCreate) AddCities(c ...*City) *StreetCreate {
	ids := make([]int32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCityIDs(ids...)
}

// AddAdressIDs adds the "adresses" edge to the AddressPlace entity by IDs.
func (sc *StreetCreate) AddAdressIDs(ids ...int32) *StreetCreate {
	sc.mutation.AddAdressIDs(ids...)
	return sc
}

// AddAdresses adds the "adresses" edges to the AddressPlace entity.
func (sc *StreetCreate) AddAdresses(a ...*AddressPlace) *StreetCreate {
	ids := make([]int32, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAdressIDs(ids...)
}

// Mutation returns the StreetMutation object of the builder.
func (sc *StreetCreate) Mutation() *StreetMutation {
	return sc.mutation
}

// Save creates the Street in the database.
func (sc *StreetCreate) Save(ctx context.Context) (*Street, error) {
	var (
		err  error
		node *Street
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreetCreate) SaveX(ctx context.Context) *Street {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StreetCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StreetCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StreetCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (sc *StreetCreate) sqlSave(ctx context.Context) (*Street, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (sc *StreetCreate) createSpec() (*Street, *sqlgraph.CreateSpec) {
	var (
		_node = &Street{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: street.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: street.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
		_node.Name = value
	}
	if nodes := sc.mutation.CitiesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AdressesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Street.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StreetUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (sc *StreetCreate) OnConflict(opts ...sql.ConflictOption) *StreetUpsertOne {
	sc.conflict = opts
	return &StreetUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Street.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *StreetCreate) OnConflictColumns(columns ...string) *StreetUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StreetUpsertOne{
		create: sc,
	}
}

type (
	// StreetUpsertOne is the builder for "upsert"-ing
	//  one Street node.
	StreetUpsertOne struct {
		create *StreetCreate
	}

	// StreetUpsert is the "OnConflict" setter.
	StreetUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *StreetUpsert) SetName(v string) *StreetUpsert {
	u.Set(street.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StreetUpsert) UpdateName() *StreetUpsert {
	u.SetExcluded(street.FieldName)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Street.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(street.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *StreetUpsertOne) UpdateNewValues() *StreetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(street.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Street.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *StreetUpsertOne) Ignore() *StreetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StreetUpsertOne) DoNothing() *StreetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StreetCreate.OnConflict
// documentation for more info.
func (u *StreetUpsertOne) Update(set func(*StreetUpsert)) *StreetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StreetUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StreetUpsertOne) SetName(v string) *StreetUpsertOne {
	return u.Update(func(s *StreetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StreetUpsertOne) UpdateName() *StreetUpsertOne {
	return u.Update(func(s *StreetUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *StreetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StreetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StreetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StreetUpsertOne) ID(ctx context.Context) (id int32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StreetUpsertOne) IDX(ctx context.Context) int32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StreetCreateBulk is the builder for creating many Street entities in bulk.
type StreetCreateBulk struct {
	config
	builders []*StreetCreate
	conflict []sql.ConflictOption
}

// Save creates the Street entities in the database.
func (scb *StreetCreateBulk) Save(ctx context.Context) ([]*Street, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Street, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StreetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StreetCreateBulk) SaveX(ctx context.Context) []*Street {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StreetCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StreetCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Street.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StreetUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (scb *StreetCreateBulk) OnConflict(opts ...sql.ConflictOption) *StreetUpsertBulk {
	scb.conflict = opts
	return &StreetUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Street.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *StreetCreateBulk) OnConflictColumns(columns ...string) *StreetUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StreetUpsertBulk{
		create: scb,
	}
}

// StreetUpsertBulk is the builder for "upsert"-ing
// a bulk of Street nodes.
type StreetUpsertBulk struct {
	create *StreetCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Street.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(street.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *StreetUpsertBulk) UpdateNewValues() *StreetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(street.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Street.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *StreetUpsertBulk) Ignore() *StreetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StreetUpsertBulk) DoNothing() *StreetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StreetCreateBulk.OnConflict
// documentation for more info.
func (u *StreetUpsertBulk) Update(set func(*StreetUpsert)) *StreetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StreetUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StreetUpsertBulk) SetName(v string) *StreetUpsertBulk {
	return u.Update(func(s *StreetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StreetUpsertBulk) UpdateName() *StreetUpsertBulk {
	return u.Update(func(s *StreetUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *StreetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StreetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StreetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StreetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
