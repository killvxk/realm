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
	"realm.pub/tavern/internal/ent/beacon"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/process"
	"realm.pub/tavern/internal/ent/tag"
)

// HostCreate is the builder for creating a Host entity.
type HostCreate struct {
	config
	mutation *HostMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (hc *HostCreate) SetCreatedAt(t time.Time) *HostCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hc *HostCreate) SetNillableCreatedAt(t *time.Time) *HostCreate {
	if t != nil {
		hc.SetCreatedAt(*t)
	}
	return hc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (hc *HostCreate) SetLastModifiedAt(t time.Time) *HostCreate {
	hc.mutation.SetLastModifiedAt(t)
	return hc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (hc *HostCreate) SetNillableLastModifiedAt(t *time.Time) *HostCreate {
	if t != nil {
		hc.SetLastModifiedAt(*t)
	}
	return hc
}

// SetIdentifier sets the "identifier" field.
func (hc *HostCreate) SetIdentifier(s string) *HostCreate {
	hc.mutation.SetIdentifier(s)
	return hc
}

// SetName sets the "name" field.
func (hc *HostCreate) SetName(s string) *HostCreate {
	hc.mutation.SetName(s)
	return hc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (hc *HostCreate) SetNillableName(s *string) *HostCreate {
	if s != nil {
		hc.SetName(*s)
	}
	return hc
}

// SetPrimaryIP sets the "primary_ip" field.
func (hc *HostCreate) SetPrimaryIP(s string) *HostCreate {
	hc.mutation.SetPrimaryIP(s)
	return hc
}

// SetNillablePrimaryIP sets the "primary_ip" field if the given value is not nil.
func (hc *HostCreate) SetNillablePrimaryIP(s *string) *HostCreate {
	if s != nil {
		hc.SetPrimaryIP(*s)
	}
	return hc
}

// SetPlatform sets the "platform" field.
func (hc *HostCreate) SetPlatform(h host.Platform) *HostCreate {
	hc.mutation.SetPlatform(h)
	return hc
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (hc *HostCreate) SetNillablePlatform(h *host.Platform) *HostCreate {
	if h != nil {
		hc.SetPlatform(*h)
	}
	return hc
}

// SetLastSeenAt sets the "last_seen_at" field.
func (hc *HostCreate) SetLastSeenAt(t time.Time) *HostCreate {
	hc.mutation.SetLastSeenAt(t)
	return hc
}

// SetNillableLastSeenAt sets the "last_seen_at" field if the given value is not nil.
func (hc *HostCreate) SetNillableLastSeenAt(t *time.Time) *HostCreate {
	if t != nil {
		hc.SetLastSeenAt(*t)
	}
	return hc
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (hc *HostCreate) AddTagIDs(ids ...int) *HostCreate {
	hc.mutation.AddTagIDs(ids...)
	return hc
}

// AddTags adds the "tags" edges to the Tag entity.
func (hc *HostCreate) AddTags(t ...*Tag) *HostCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return hc.AddTagIDs(ids...)
}

// AddBeaconIDs adds the "beacons" edge to the Beacon entity by IDs.
func (hc *HostCreate) AddBeaconIDs(ids ...int) *HostCreate {
	hc.mutation.AddBeaconIDs(ids...)
	return hc
}

// AddBeacons adds the "beacons" edges to the Beacon entity.
func (hc *HostCreate) AddBeacons(b ...*Beacon) *HostCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return hc.AddBeaconIDs(ids...)
}

// AddProcessIDs adds the "processes" edge to the Process entity by IDs.
func (hc *HostCreate) AddProcessIDs(ids ...int) *HostCreate {
	hc.mutation.AddProcessIDs(ids...)
	return hc
}

// AddProcesses adds the "processes" edges to the Process entity.
func (hc *HostCreate) AddProcesses(p ...*Process) *HostCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return hc.AddProcessIDs(ids...)
}

// Mutation returns the HostMutation object of the builder.
func (hc *HostCreate) Mutation() *HostMutation {
	return hc.mutation
}

// Save creates the Host in the database.
func (hc *HostCreate) Save(ctx context.Context) (*Host, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HostCreate) SaveX(ctx context.Context) *Host {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HostCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HostCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HostCreate) defaults() {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		v := host.DefaultCreatedAt()
		hc.mutation.SetCreatedAt(v)
	}
	if _, ok := hc.mutation.LastModifiedAt(); !ok {
		v := host.DefaultLastModifiedAt()
		hc.mutation.SetLastModifiedAt(v)
	}
	if _, ok := hc.mutation.Platform(); !ok {
		v := host.DefaultPlatform
		hc.mutation.SetPlatform(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HostCreate) check() error {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Host.created_at"`)}
	}
	if _, ok := hc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "Host.last_modified_at"`)}
	}
	if _, ok := hc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Host.identifier"`)}
	}
	if v, ok := hc.mutation.Identifier(); ok {
		if err := host.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "Host.identifier": %w`, err)}
		}
	}
	if v, ok := hc.mutation.Name(); ok {
		if err := host.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Host.name": %w`, err)}
		}
	}
	if _, ok := hc.mutation.Platform(); !ok {
		return &ValidationError{Name: "platform", err: errors.New(`ent: missing required field "Host.platform"`)}
	}
	if v, ok := hc.mutation.Platform(); ok {
		if err := host.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Host.platform": %w`, err)}
		}
	}
	return nil
}

func (hc *HostCreate) sqlSave(ctx context.Context) (*Host, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HostCreate) createSpec() (*Host, *sqlgraph.CreateSpec) {
	var (
		_node = &Host{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(host.Table, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt))
	)
	_spec.OnConflict = hc.conflict
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.SetField(host.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hc.mutation.LastModifiedAt(); ok {
		_spec.SetField(host.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if value, ok := hc.mutation.Identifier(); ok {
		_spec.SetField(host.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	if value, ok := hc.mutation.Name(); ok {
		_spec.SetField(host.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := hc.mutation.PrimaryIP(); ok {
		_spec.SetField(host.FieldPrimaryIP, field.TypeString, value)
		_node.PrimaryIP = value
	}
	if value, ok := hc.mutation.Platform(); ok {
		_spec.SetField(host.FieldPlatform, field.TypeEnum, value)
		_node.Platform = value
	}
	if value, ok := hc.mutation.LastSeenAt(); ok {
		_spec.SetField(host.FieldLastSeenAt, field.TypeTime, value)
		_node.LastSeenAt = value
	}
	if nodes := hc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   host.TagsTable,
			Columns: host.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.BeaconsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   host.BeaconsTable,
			Columns: []string{host.BeaconsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.ProcessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.ProcessesTable,
			Columns: []string{host.ProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt),
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
//	client.Host.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HostUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (hc *HostCreate) OnConflict(opts ...sql.ConflictOption) *HostUpsertOne {
	hc.conflict = opts
	return &HostUpsertOne{
		create: hc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Host.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hc *HostCreate) OnConflictColumns(columns ...string) *HostUpsertOne {
	hc.conflict = append(hc.conflict, sql.ConflictColumns(columns...))
	return &HostUpsertOne{
		create: hc,
	}
}

type (
	// HostUpsertOne is the builder for "upsert"-ing
	//  one Host node.
	HostUpsertOne struct {
		create *HostCreate
	}

	// HostUpsert is the "OnConflict" setter.
	HostUpsert struct {
		*sql.UpdateSet
	}
)

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *HostUpsert) SetLastModifiedAt(v time.Time) *HostUpsert {
	u.Set(host.FieldLastModifiedAt, v)
	return u
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *HostUpsert) UpdateLastModifiedAt() *HostUpsert {
	u.SetExcluded(host.FieldLastModifiedAt)
	return u
}

// SetIdentifier sets the "identifier" field.
func (u *HostUpsert) SetIdentifier(v string) *HostUpsert {
	u.Set(host.FieldIdentifier, v)
	return u
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *HostUpsert) UpdateIdentifier() *HostUpsert {
	u.SetExcluded(host.FieldIdentifier)
	return u
}

// SetName sets the "name" field.
func (u *HostUpsert) SetName(v string) *HostUpsert {
	u.Set(host.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HostUpsert) UpdateName() *HostUpsert {
	u.SetExcluded(host.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *HostUpsert) ClearName() *HostUpsert {
	u.SetNull(host.FieldName)
	return u
}

// SetPrimaryIP sets the "primary_ip" field.
func (u *HostUpsert) SetPrimaryIP(v string) *HostUpsert {
	u.Set(host.FieldPrimaryIP, v)
	return u
}

// UpdatePrimaryIP sets the "primary_ip" field to the value that was provided on create.
func (u *HostUpsert) UpdatePrimaryIP() *HostUpsert {
	u.SetExcluded(host.FieldPrimaryIP)
	return u
}

// ClearPrimaryIP clears the value of the "primary_ip" field.
func (u *HostUpsert) ClearPrimaryIP() *HostUpsert {
	u.SetNull(host.FieldPrimaryIP)
	return u
}

// SetPlatform sets the "platform" field.
func (u *HostUpsert) SetPlatform(v host.Platform) *HostUpsert {
	u.Set(host.FieldPlatform, v)
	return u
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *HostUpsert) UpdatePlatform() *HostUpsert {
	u.SetExcluded(host.FieldPlatform)
	return u
}

// SetLastSeenAt sets the "last_seen_at" field.
func (u *HostUpsert) SetLastSeenAt(v time.Time) *HostUpsert {
	u.Set(host.FieldLastSeenAt, v)
	return u
}

// UpdateLastSeenAt sets the "last_seen_at" field to the value that was provided on create.
func (u *HostUpsert) UpdateLastSeenAt() *HostUpsert {
	u.SetExcluded(host.FieldLastSeenAt)
	return u
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (u *HostUpsert) ClearLastSeenAt() *HostUpsert {
	u.SetNull(host.FieldLastSeenAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Host.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HostUpsertOne) UpdateNewValues() *HostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(host.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Host.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HostUpsertOne) Ignore() *HostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HostUpsertOne) DoNothing() *HostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HostCreate.OnConflict
// documentation for more info.
func (u *HostUpsertOne) Update(set func(*HostUpsert)) *HostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HostUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *HostUpsertOne) SetLastModifiedAt(v time.Time) *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *HostUpsertOne) UpdateLastModifiedAt() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *HostUpsertOne) SetIdentifier(v string) *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *HostUpsertOne) UpdateIdentifier() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.UpdateIdentifier()
	})
}

// SetName sets the "name" field.
func (u *HostUpsertOne) SetName(v string) *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HostUpsertOne) UpdateName() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *HostUpsertOne) ClearName() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.ClearName()
	})
}

// SetPrimaryIP sets the "primary_ip" field.
func (u *HostUpsertOne) SetPrimaryIP(v string) *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.SetPrimaryIP(v)
	})
}

// UpdatePrimaryIP sets the "primary_ip" field to the value that was provided on create.
func (u *HostUpsertOne) UpdatePrimaryIP() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.UpdatePrimaryIP()
	})
}

// ClearPrimaryIP clears the value of the "primary_ip" field.
func (u *HostUpsertOne) ClearPrimaryIP() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.ClearPrimaryIP()
	})
}

// SetPlatform sets the "platform" field.
func (u *HostUpsertOne) SetPlatform(v host.Platform) *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.SetPlatform(v)
	})
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *HostUpsertOne) UpdatePlatform() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.UpdatePlatform()
	})
}

// SetLastSeenAt sets the "last_seen_at" field.
func (u *HostUpsertOne) SetLastSeenAt(v time.Time) *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.SetLastSeenAt(v)
	})
}

// UpdateLastSeenAt sets the "last_seen_at" field to the value that was provided on create.
func (u *HostUpsertOne) UpdateLastSeenAt() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.UpdateLastSeenAt()
	})
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (u *HostUpsertOne) ClearLastSeenAt() *HostUpsertOne {
	return u.Update(func(s *HostUpsert) {
		s.ClearLastSeenAt()
	})
}

// Exec executes the query.
func (u *HostUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HostCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HostUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HostUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HostUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HostCreateBulk is the builder for creating many Host entities in bulk.
type HostCreateBulk struct {
	config
	err      error
	builders []*HostCreate
	conflict []sql.ConflictOption
}

// Save creates the Host entities in the database.
func (hcb *HostCreateBulk) Save(ctx context.Context) ([]*Host, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Host, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HostCreateBulk) SaveX(ctx context.Context) []*Host {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HostCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HostCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Host.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HostUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (hcb *HostCreateBulk) OnConflict(opts ...sql.ConflictOption) *HostUpsertBulk {
	hcb.conflict = opts
	return &HostUpsertBulk{
		create: hcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Host.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hcb *HostCreateBulk) OnConflictColumns(columns ...string) *HostUpsertBulk {
	hcb.conflict = append(hcb.conflict, sql.ConflictColumns(columns...))
	return &HostUpsertBulk{
		create: hcb,
	}
}

// HostUpsertBulk is the builder for "upsert"-ing
// a bulk of Host nodes.
type HostUpsertBulk struct {
	create *HostCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Host.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HostUpsertBulk) UpdateNewValues() *HostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(host.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Host.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HostUpsertBulk) Ignore() *HostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HostUpsertBulk) DoNothing() *HostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HostCreateBulk.OnConflict
// documentation for more info.
func (u *HostUpsertBulk) Update(set func(*HostUpsert)) *HostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HostUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *HostUpsertBulk) SetLastModifiedAt(v time.Time) *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *HostUpsertBulk) UpdateLastModifiedAt() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *HostUpsertBulk) SetIdentifier(v string) *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *HostUpsertBulk) UpdateIdentifier() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.UpdateIdentifier()
	})
}

// SetName sets the "name" field.
func (u *HostUpsertBulk) SetName(v string) *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *HostUpsertBulk) UpdateName() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *HostUpsertBulk) ClearName() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.ClearName()
	})
}

// SetPrimaryIP sets the "primary_ip" field.
func (u *HostUpsertBulk) SetPrimaryIP(v string) *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.SetPrimaryIP(v)
	})
}

// UpdatePrimaryIP sets the "primary_ip" field to the value that was provided on create.
func (u *HostUpsertBulk) UpdatePrimaryIP() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.UpdatePrimaryIP()
	})
}

// ClearPrimaryIP clears the value of the "primary_ip" field.
func (u *HostUpsertBulk) ClearPrimaryIP() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.ClearPrimaryIP()
	})
}

// SetPlatform sets the "platform" field.
func (u *HostUpsertBulk) SetPlatform(v host.Platform) *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.SetPlatform(v)
	})
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *HostUpsertBulk) UpdatePlatform() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.UpdatePlatform()
	})
}

// SetLastSeenAt sets the "last_seen_at" field.
func (u *HostUpsertBulk) SetLastSeenAt(v time.Time) *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.SetLastSeenAt(v)
	})
}

// UpdateLastSeenAt sets the "last_seen_at" field to the value that was provided on create.
func (u *HostUpsertBulk) UpdateLastSeenAt() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.UpdateLastSeenAt()
	})
}

// ClearLastSeenAt clears the value of the "last_seen_at" field.
func (u *HostUpsertBulk) ClearLastSeenAt() *HostUpsertBulk {
	return u.Update(func(s *HostUpsert) {
		s.ClearLastSeenAt()
	})
}

// Exec executes the query.
func (u *HostUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HostCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HostCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HostUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
