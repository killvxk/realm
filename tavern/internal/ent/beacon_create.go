// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/internal/ent/beacon"
	"github.com/kcarretto/realm/tavern/internal/ent/host"
	"github.com/kcarretto/realm/tavern/internal/ent/task"
)

// BeaconCreate is the builder for creating a Beacon entity.
type BeaconCreate struct {
	config
	mutation *BeaconMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (bc *BeaconCreate) SetName(s string) *BeaconCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableName(s *string) *BeaconCreate {
	if s != nil {
		bc.SetName(*s)
	}
	return bc
}

// SetPrincipal sets the "principal" field.
func (bc *BeaconCreate) SetPrincipal(s string) *BeaconCreate {
	bc.mutation.SetPrincipal(s)
	return bc
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (bc *BeaconCreate) SetNillablePrincipal(s *string) *BeaconCreate {
	if s != nil {
		bc.SetPrincipal(*s)
	}
	return bc
}

// SetIdentifier sets the "identifier" field.
func (bc *BeaconCreate) SetIdentifier(s string) *BeaconCreate {
	bc.mutation.SetIdentifier(s)
	return bc
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableIdentifier(s *string) *BeaconCreate {
	if s != nil {
		bc.SetIdentifier(*s)
	}
	return bc
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (bc *BeaconCreate) SetAgentIdentifier(s string) *BeaconCreate {
	bc.mutation.SetAgentIdentifier(s)
	return bc
}

// SetNillableAgentIdentifier sets the "agent_identifier" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableAgentIdentifier(s *string) *BeaconCreate {
	if s != nil {
		bc.SetAgentIdentifier(*s)
	}
	return bc
}

// SetLastSeenAt sets the "last_seen_at" field.
func (bc *BeaconCreate) SetLastSeenAt(t time.Time) *BeaconCreate {
	bc.mutation.SetLastSeenAt(t)
	return bc
}

// SetNillableLastSeenAt sets the "last_seen_at" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableLastSeenAt(t *time.Time) *BeaconCreate {
	if t != nil {
		bc.SetLastSeenAt(*t)
	}
	return bc
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (bc *BeaconCreate) SetHostID(id int) *BeaconCreate {
	bc.mutation.SetHostID(id)
	return bc
}

// SetHost sets the "host" edge to the Host entity.
func (bc *BeaconCreate) SetHost(h *Host) *BeaconCreate {
	return bc.SetHostID(h.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (bc *BeaconCreate) AddTaskIDs(ids ...int) *BeaconCreate {
	bc.mutation.AddTaskIDs(ids...)
	return bc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (bc *BeaconCreate) AddTasks(t ...*Task) *BeaconCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTaskIDs(ids...)
}

// Mutation returns the BeaconMutation object of the builder.
func (bc *BeaconCreate) Mutation() *BeaconMutation {
	return bc.mutation
}

// Save creates the Beacon in the database.
func (bc *BeaconCreate) Save(ctx context.Context) (*Beacon, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BeaconCreate) SaveX(ctx context.Context) *Beacon {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BeaconCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BeaconCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BeaconCreate) defaults() {
	if _, ok := bc.mutation.Name(); !ok {
		v := beacon.DefaultName()
		bc.mutation.SetName(v)
	}
	if _, ok := bc.mutation.Identifier(); !ok {
		v := beacon.DefaultIdentifier()
		bc.mutation.SetIdentifier(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BeaconCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Beacon.name"`)}
	}
	if v, ok := bc.mutation.Name(); ok {
		if err := beacon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Beacon.name": %w`, err)}
		}
	}
	if v, ok := bc.mutation.Principal(); ok {
		if err := beacon.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Beacon.principal": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Beacon.identifier"`)}
	}
	if v, ok := bc.mutation.Identifier(); ok {
		if err := beacon.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.identifier": %w`, err)}
		}
	}
	if v, ok := bc.mutation.AgentIdentifier(); ok {
		if err := beacon.AgentIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "agent_identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.agent_identifier": %w`, err)}
		}
	}
	if _, ok := bc.mutation.HostID(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required edge "Beacon.host"`)}
	}
	return nil
}

func (bc *BeaconCreate) sqlSave(ctx context.Context) (*Beacon, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BeaconCreate) createSpec() (*Beacon, *sqlgraph.CreateSpec) {
	var (
		_node = &Beacon{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(beacon.Table, sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt))
	)
	_spec.OnConflict = bc.conflict
	if value, ok := bc.mutation.Name(); ok {
		_spec.SetField(beacon.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bc.mutation.Principal(); ok {
		_spec.SetField(beacon.FieldPrincipal, field.TypeString, value)
		_node.Principal = value
	}
	if value, ok := bc.mutation.Identifier(); ok {
		_spec.SetField(beacon.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	if value, ok := bc.mutation.AgentIdentifier(); ok {
		_spec.SetField(beacon.FieldAgentIdentifier, field.TypeString, value)
		_node.AgentIdentifier = value
	}
	if value, ok := bc.mutation.LastSeenAt(); ok {
		_spec.SetField(beacon.FieldLastSeenAt, field.TypeTime, value)
		_node.LastSeenAt = value
	}
	if nodes := bc.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   beacon.HostTable,
			Columns: []string{beacon.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.beacon_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
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
//	client.Beacon.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BeaconUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (bc *BeaconCreate) OnConflict(opts ...sql.ConflictOption) *BeaconUpsertOne {
	bc.conflict = opts
	return &BeaconUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Beacon.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bc *BeaconCreate) OnConflictColumns(columns ...string) *BeaconUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BeaconUpsertOne{
		create: bc,
	}
}

type (
	// BeaconUpsertOne is the builder for "upsert"-ing
	//  one Beacon node.
	BeaconUpsertOne struct {
		create *BeaconCreate
	}

	// BeaconUpsert is the "OnConflict" setter.
	BeaconUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *BeaconUpsert) SetName(v string) *BeaconUpsert {
	u.Set(beacon.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *BeaconUpsert) UpdateName() *BeaconUpsert {
	u.SetExcluded(beacon.FieldName)
	return u
}

// SetPrincipal sets the "principal" field.
func (u *BeaconUpsert) SetPrincipal(v string) *BeaconUpsert {
	u.Set(beacon.FieldPrincipal, v)
	return u
}

// UpdatePrincipal sets the "principal" field to the value that was provided on create.
func (u *BeaconUpsert) UpdatePrincipal() *BeaconUpsert {
	u.SetExcluded(beacon.FieldPrincipal)
	return u
}

// ClearPrincipal clears the value of the "principal" field.
func (u *BeaconUpsert) ClearPrincipal() *BeaconUpsert {
	u.SetNull(beacon.FieldPrincipal)
	return u
}

// SetIdentifier sets the "identifier" field.
func (u *BeaconUpsert) SetIdentifier(v string) *BeaconUpsert {
	u.Set(beacon.FieldIdentifier, v)
	return u
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *BeaconUpsert) UpdateIdentifier() *BeaconUpsert {
	u.SetExcluded(beacon.FieldIdentifier)
	return u
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (u *BeaconUpsert) SetAgentIdentifier(v string) *BeaconUpsert {
	u.Set(beacon.FieldAgentIdentifier, v)
	return u
}

// UpdateAgentIdentifier sets the "agent_identifier" field to the value that was provided on create.
func (u *BeaconUpsert) UpdateAgentIdentifier() *BeaconUpsert {
	u.SetExcluded(beacon.FieldAgentIdentifier)
	return u
}

// ClearAgentIdentifier clears the value of the "agent_identifier" field.
func (u *BeaconUpsert) ClearAgentIdentifier() *BeaconUpsert {
	u.SetNull(beacon.FieldAgentIdentifier)
	return u
}

// SetLastSeenAt sets the "last_seen_at" field.
func (u *BeaconUpsert) SetLastSeenAt(v time.Time) *BeaconUpsert {
	u.Set(beacon.FieldLastSeenAt, v)
	return u
}

// UpdateLastSeenAt sets the "last_seen_at" field to the value that was provided on create.
func (u *BeaconUpsert) UpdateLastSeenAt() *BeaconUpsert {
	u.SetExcluded(beacon.FieldLastSeenAt)
	return u
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (u *BeaconUpsert) ClearLastSeenAt() *BeaconUpsert {
	u.SetNull(beacon.FieldLastSeenAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Beacon.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *BeaconUpsertOne) UpdateNewValues() *BeaconUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Beacon.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BeaconUpsertOne) Ignore() *BeaconUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BeaconUpsertOne) DoNothing() *BeaconUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BeaconCreate.OnConflict
// documentation for more info.
func (u *BeaconUpsertOne) Update(set func(*BeaconUpsert)) *BeaconUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BeaconUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *BeaconUpsertOne) SetName(v string) *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *BeaconUpsertOne) UpdateName() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateName()
	})
}

// SetPrincipal sets the "principal" field.
func (u *BeaconUpsertOne) SetPrincipal(v string) *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.SetPrincipal(v)
	})
}

// UpdatePrincipal sets the "principal" field to the value that was provided on create.
func (u *BeaconUpsertOne) UpdatePrincipal() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdatePrincipal()
	})
}

// ClearPrincipal clears the value of the "principal" field.
func (u *BeaconUpsertOne) ClearPrincipal() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.ClearPrincipal()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *BeaconUpsertOne) SetIdentifier(v string) *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *BeaconUpsertOne) UpdateIdentifier() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateIdentifier()
	})
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (u *BeaconUpsertOne) SetAgentIdentifier(v string) *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.SetAgentIdentifier(v)
	})
}

// UpdateAgentIdentifier sets the "agent_identifier" field to the value that was provided on create.
func (u *BeaconUpsertOne) UpdateAgentIdentifier() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateAgentIdentifier()
	})
}

// ClearAgentIdentifier clears the value of the "agent_identifier" field.
func (u *BeaconUpsertOne) ClearAgentIdentifier() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.ClearAgentIdentifier()
	})
}

// SetLastSeenAt sets the "last_seen_at" field.
func (u *BeaconUpsertOne) SetLastSeenAt(v time.Time) *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.SetLastSeenAt(v)
	})
}

// UpdateLastSeenAt sets the "last_seen_at" field to the value that was provided on create.
func (u *BeaconUpsertOne) UpdateLastSeenAt() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateLastSeenAt()
	})
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (u *BeaconUpsertOne) ClearLastSeenAt() *BeaconUpsertOne {
	return u.Update(func(s *BeaconUpsert) {
		s.ClearLastSeenAt()
	})
}

// Exec executes the query.
func (u *BeaconUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BeaconCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BeaconUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BeaconUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BeaconUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BeaconCreateBulk is the builder for creating many Beacon entities in bulk.
type BeaconCreateBulk struct {
	config
	err      error
	builders []*BeaconCreate
	conflict []sql.ConflictOption
}

// Save creates the Beacon entities in the database.
func (bcb *BeaconCreateBulk) Save(ctx context.Context) ([]*Beacon, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Beacon, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BeaconMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BeaconCreateBulk) SaveX(ctx context.Context) []*Beacon {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BeaconCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BeaconCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Beacon.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BeaconUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (bcb *BeaconCreateBulk) OnConflict(opts ...sql.ConflictOption) *BeaconUpsertBulk {
	bcb.conflict = opts
	return &BeaconUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Beacon.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bcb *BeaconCreateBulk) OnConflictColumns(columns ...string) *BeaconUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BeaconUpsertBulk{
		create: bcb,
	}
}

// BeaconUpsertBulk is the builder for "upsert"-ing
// a bulk of Beacon nodes.
type BeaconUpsertBulk struct {
	create *BeaconCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Beacon.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *BeaconUpsertBulk) UpdateNewValues() *BeaconUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Beacon.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BeaconUpsertBulk) Ignore() *BeaconUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BeaconUpsertBulk) DoNothing() *BeaconUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BeaconCreateBulk.OnConflict
// documentation for more info.
func (u *BeaconUpsertBulk) Update(set func(*BeaconUpsert)) *BeaconUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BeaconUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *BeaconUpsertBulk) SetName(v string) *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *BeaconUpsertBulk) UpdateName() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateName()
	})
}

// SetPrincipal sets the "principal" field.
func (u *BeaconUpsertBulk) SetPrincipal(v string) *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.SetPrincipal(v)
	})
}

// UpdatePrincipal sets the "principal" field to the value that was provided on create.
func (u *BeaconUpsertBulk) UpdatePrincipal() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdatePrincipal()
	})
}

// ClearPrincipal clears the value of the "principal" field.
func (u *BeaconUpsertBulk) ClearPrincipal() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.ClearPrincipal()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *BeaconUpsertBulk) SetIdentifier(v string) *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *BeaconUpsertBulk) UpdateIdentifier() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateIdentifier()
	})
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (u *BeaconUpsertBulk) SetAgentIdentifier(v string) *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.SetAgentIdentifier(v)
	})
}

// UpdateAgentIdentifier sets the "agent_identifier" field to the value that was provided on create.
func (u *BeaconUpsertBulk) UpdateAgentIdentifier() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateAgentIdentifier()
	})
}

// ClearAgentIdentifier clears the value of the "agent_identifier" field.
func (u *BeaconUpsertBulk) ClearAgentIdentifier() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.ClearAgentIdentifier()
	})
}

// SetLastSeenAt sets the "last_seen_at" field.
func (u *BeaconUpsertBulk) SetLastSeenAt(v time.Time) *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.SetLastSeenAt(v)
	})
}

// UpdateLastSeenAt sets the "last_seen_at" field to the value that was provided on create.
func (u *BeaconUpsertBulk) UpdateLastSeenAt() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.UpdateLastSeenAt()
	})
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (u *BeaconUpsertBulk) ClearLastSeenAt() *BeaconUpsertBulk {
	return u.Update(func(s *BeaconUpsert) {
		s.ClearLastSeenAt()
	})
}

// Exec executes the query.
func (u *BeaconUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BeaconCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BeaconCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BeaconUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}