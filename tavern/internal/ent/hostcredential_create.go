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
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/hostcredential"
	"realm.pub/tavern/internal/ent/task"
)

// HostCredentialCreate is the builder for creating a HostCredential entity.
type HostCredentialCreate struct {
	config
	mutation *HostCredentialMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (hcc *HostCredentialCreate) SetCreatedAt(t time.Time) *HostCredentialCreate {
	hcc.mutation.SetCreatedAt(t)
	return hcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hcc *HostCredentialCreate) SetNillableCreatedAt(t *time.Time) *HostCredentialCreate {
	if t != nil {
		hcc.SetCreatedAt(*t)
	}
	return hcc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (hcc *HostCredentialCreate) SetLastModifiedAt(t time.Time) *HostCredentialCreate {
	hcc.mutation.SetLastModifiedAt(t)
	return hcc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (hcc *HostCredentialCreate) SetNillableLastModifiedAt(t *time.Time) *HostCredentialCreate {
	if t != nil {
		hcc.SetLastModifiedAt(*t)
	}
	return hcc
}

// SetPrincipal sets the "principal" field.
func (hcc *HostCredentialCreate) SetPrincipal(s string) *HostCredentialCreate {
	hcc.mutation.SetPrincipal(s)
	return hcc
}

// SetSecret sets the "secret" field.
func (hcc *HostCredentialCreate) SetSecret(s string) *HostCredentialCreate {
	hcc.mutation.SetSecret(s)
	return hcc
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (hcc *HostCredentialCreate) SetHostID(id int) *HostCredentialCreate {
	hcc.mutation.SetHostID(id)
	return hcc
}

// SetHost sets the "host" edge to the Host entity.
func (hcc *HostCredentialCreate) SetHost(h *Host) *HostCredentialCreate {
	return hcc.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (hcc *HostCredentialCreate) SetTaskID(id int) *HostCredentialCreate {
	hcc.mutation.SetTaskID(id)
	return hcc
}

// SetTask sets the "task" edge to the Task entity.
func (hcc *HostCredentialCreate) SetTask(t *Task) *HostCredentialCreate {
	return hcc.SetTaskID(t.ID)
}

// Mutation returns the HostCredentialMutation object of the builder.
func (hcc *HostCredentialCreate) Mutation() *HostCredentialMutation {
	return hcc.mutation
}

// Save creates the HostCredential in the database.
func (hcc *HostCredentialCreate) Save(ctx context.Context) (*HostCredential, error) {
	hcc.defaults()
	return withHooks(ctx, hcc.sqlSave, hcc.mutation, hcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hcc *HostCredentialCreate) SaveX(ctx context.Context) *HostCredential {
	v, err := hcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcc *HostCredentialCreate) Exec(ctx context.Context) error {
	_, err := hcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcc *HostCredentialCreate) ExecX(ctx context.Context) {
	if err := hcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hcc *HostCredentialCreate) defaults() {
	if _, ok := hcc.mutation.CreatedAt(); !ok {
		v := hostcredential.DefaultCreatedAt()
		hcc.mutation.SetCreatedAt(v)
	}
	if _, ok := hcc.mutation.LastModifiedAt(); !ok {
		v := hostcredential.DefaultLastModifiedAt()
		hcc.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hcc *HostCredentialCreate) check() error {
	if _, ok := hcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HostCredential.created_at"`)}
	}
	if _, ok := hcc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "HostCredential.last_modified_at"`)}
	}
	if _, ok := hcc.mutation.Principal(); !ok {
		return &ValidationError{Name: "principal", err: errors.New(`ent: missing required field "HostCredential.principal"`)}
	}
	if v, ok := hcc.mutation.Principal(); ok {
		if err := hostcredential.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "HostCredential.principal": %w`, err)}
		}
	}
	if _, ok := hcc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`ent: missing required field "HostCredential.secret"`)}
	}
	if v, ok := hcc.mutation.Secret(); ok {
		if err := hostcredential.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "HostCredential.secret": %w`, err)}
		}
	}
	if _, ok := hcc.mutation.HostID(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required edge "HostCredential.host"`)}
	}
	if _, ok := hcc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "HostCredential.task"`)}
	}
	return nil
}

func (hcc *HostCredentialCreate) sqlSave(ctx context.Context) (*HostCredential, error) {
	if err := hcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hcc.mutation.id = &_node.ID
	hcc.mutation.done = true
	return _node, nil
}

func (hcc *HostCredentialCreate) createSpec() (*HostCredential, *sqlgraph.CreateSpec) {
	var (
		_node = &HostCredential{config: hcc.config}
		_spec = sqlgraph.NewCreateSpec(hostcredential.Table, sqlgraph.NewFieldSpec(hostcredential.FieldID, field.TypeInt))
	)
	_spec.OnConflict = hcc.conflict
	if value, ok := hcc.mutation.CreatedAt(); ok {
		_spec.SetField(hostcredential.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hcc.mutation.LastModifiedAt(); ok {
		_spec.SetField(hostcredential.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if value, ok := hcc.mutation.Principal(); ok {
		_spec.SetField(hostcredential.FieldPrincipal, field.TypeString, value)
		_node.Principal = value
	}
	if value, ok := hcc.mutation.Secret(); ok {
		_spec.SetField(hostcredential.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if nodes := hcc.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostcredential.HostTable,
			Columns: []string{hostcredential.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.host_credential_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hcc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostcredential.TaskTable,
			Columns: []string{hostcredential.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_reported_credentials = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HostCredential.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HostCredentialUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (hcc *HostCredentialCreate) OnConflict(opts ...sql.ConflictOption) *HostCredentialUpsertOne {
	hcc.conflict = opts
	return &HostCredentialUpsertOne{
		create: hcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HostCredential.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hcc *HostCredentialCreate) OnConflictColumns(columns ...string) *HostCredentialUpsertOne {
	hcc.conflict = append(hcc.conflict, sql.ConflictColumns(columns...))
	return &HostCredentialUpsertOne{
		create: hcc,
	}
}

type (
	// HostCredentialUpsertOne is the builder for "upsert"-ing
	//  one HostCredential node.
	HostCredentialUpsertOne struct {
		create *HostCredentialCreate
	}

	// HostCredentialUpsert is the "OnConflict" setter.
	HostCredentialUpsert struct {
		*sql.UpdateSet
	}
)

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *HostCredentialUpsert) SetLastModifiedAt(v time.Time) *HostCredentialUpsert {
	u.Set(hostcredential.FieldLastModifiedAt, v)
	return u
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *HostCredentialUpsert) UpdateLastModifiedAt() *HostCredentialUpsert {
	u.SetExcluded(hostcredential.FieldLastModifiedAt)
	return u
}

// SetPrincipal sets the "principal" field.
func (u *HostCredentialUpsert) SetPrincipal(v string) *HostCredentialUpsert {
	u.Set(hostcredential.FieldPrincipal, v)
	return u
}

// UpdatePrincipal sets the "principal" field to the value that was provided on create.
func (u *HostCredentialUpsert) UpdatePrincipal() *HostCredentialUpsert {
	u.SetExcluded(hostcredential.FieldPrincipal)
	return u
}

// SetSecret sets the "secret" field.
func (u *HostCredentialUpsert) SetSecret(v string) *HostCredentialUpsert {
	u.Set(hostcredential.FieldSecret, v)
	return u
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *HostCredentialUpsert) UpdateSecret() *HostCredentialUpsert {
	u.SetExcluded(hostcredential.FieldSecret)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.HostCredential.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HostCredentialUpsertOne) UpdateNewValues() *HostCredentialUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(hostcredential.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HostCredential.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HostCredentialUpsertOne) Ignore() *HostCredentialUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HostCredentialUpsertOne) DoNothing() *HostCredentialUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HostCredentialCreate.OnConflict
// documentation for more info.
func (u *HostCredentialUpsertOne) Update(set func(*HostCredentialUpsert)) *HostCredentialUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HostCredentialUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *HostCredentialUpsertOne) SetLastModifiedAt(v time.Time) *HostCredentialUpsertOne {
	return u.Update(func(s *HostCredentialUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *HostCredentialUpsertOne) UpdateLastModifiedAt() *HostCredentialUpsertOne {
	return u.Update(func(s *HostCredentialUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetPrincipal sets the "principal" field.
func (u *HostCredentialUpsertOne) SetPrincipal(v string) *HostCredentialUpsertOne {
	return u.Update(func(s *HostCredentialUpsert) {
		s.SetPrincipal(v)
	})
}

// UpdatePrincipal sets the "principal" field to the value that was provided on create.
func (u *HostCredentialUpsertOne) UpdatePrincipal() *HostCredentialUpsertOne {
	return u.Update(func(s *HostCredentialUpsert) {
		s.UpdatePrincipal()
	})
}

// SetSecret sets the "secret" field.
func (u *HostCredentialUpsertOne) SetSecret(v string) *HostCredentialUpsertOne {
	return u.Update(func(s *HostCredentialUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *HostCredentialUpsertOne) UpdateSecret() *HostCredentialUpsertOne {
	return u.Update(func(s *HostCredentialUpsert) {
		s.UpdateSecret()
	})
}

// Exec executes the query.
func (u *HostCredentialUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HostCredentialCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HostCredentialUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HostCredentialUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HostCredentialUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HostCredentialCreateBulk is the builder for creating many HostCredential entities in bulk.
type HostCredentialCreateBulk struct {
	config
	err      error
	builders []*HostCredentialCreate
	conflict []sql.ConflictOption
}

// Save creates the HostCredential entities in the database.
func (hccb *HostCredentialCreateBulk) Save(ctx context.Context) ([]*HostCredential, error) {
	if hccb.err != nil {
		return nil, hccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hccb.builders))
	nodes := make([]*HostCredential, len(hccb.builders))
	mutators := make([]Mutator, len(hccb.builders))
	for i := range hccb.builders {
		func(i int, root context.Context) {
			builder := hccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostCredentialMutation)
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
					_, err = mutators[i+1].Mutate(root, hccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hccb *HostCredentialCreateBulk) SaveX(ctx context.Context) []*HostCredential {
	v, err := hccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hccb *HostCredentialCreateBulk) Exec(ctx context.Context) error {
	_, err := hccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hccb *HostCredentialCreateBulk) ExecX(ctx context.Context) {
	if err := hccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.HostCredential.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HostCredentialUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (hccb *HostCredentialCreateBulk) OnConflict(opts ...sql.ConflictOption) *HostCredentialUpsertBulk {
	hccb.conflict = opts
	return &HostCredentialUpsertBulk{
		create: hccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.HostCredential.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hccb *HostCredentialCreateBulk) OnConflictColumns(columns ...string) *HostCredentialUpsertBulk {
	hccb.conflict = append(hccb.conflict, sql.ConflictColumns(columns...))
	return &HostCredentialUpsertBulk{
		create: hccb,
	}
}

// HostCredentialUpsertBulk is the builder for "upsert"-ing
// a bulk of HostCredential nodes.
type HostCredentialUpsertBulk struct {
	create *HostCredentialCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.HostCredential.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *HostCredentialUpsertBulk) UpdateNewValues() *HostCredentialUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(hostcredential.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.HostCredential.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HostCredentialUpsertBulk) Ignore() *HostCredentialUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HostCredentialUpsertBulk) DoNothing() *HostCredentialUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HostCredentialCreateBulk.OnConflict
// documentation for more info.
func (u *HostCredentialUpsertBulk) Update(set func(*HostCredentialUpsert)) *HostCredentialUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HostCredentialUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *HostCredentialUpsertBulk) SetLastModifiedAt(v time.Time) *HostCredentialUpsertBulk {
	return u.Update(func(s *HostCredentialUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *HostCredentialUpsertBulk) UpdateLastModifiedAt() *HostCredentialUpsertBulk {
	return u.Update(func(s *HostCredentialUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetPrincipal sets the "principal" field.
func (u *HostCredentialUpsertBulk) SetPrincipal(v string) *HostCredentialUpsertBulk {
	return u.Update(func(s *HostCredentialUpsert) {
		s.SetPrincipal(v)
	})
}

// UpdatePrincipal sets the "principal" field to the value that was provided on create.
func (u *HostCredentialUpsertBulk) UpdatePrincipal() *HostCredentialUpsertBulk {
	return u.Update(func(s *HostCredentialUpsert) {
		s.UpdatePrincipal()
	})
}

// SetSecret sets the "secret" field.
func (u *HostCredentialUpsertBulk) SetSecret(v string) *HostCredentialUpsertBulk {
	return u.Update(func(s *HostCredentialUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *HostCredentialUpsertBulk) UpdateSecret() *HostCredentialUpsertBulk {
	return u.Update(func(s *HostCredentialUpsert) {
		s.UpdateSecret()
	})
}

// Exec executes the query.
func (u *HostCredentialUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HostCredentialCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HostCredentialCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HostCredentialUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
