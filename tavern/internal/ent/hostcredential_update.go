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
	"realm.pub/tavern/internal/c2/epb"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/hostcredential"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/task"
)

// HostCredentialUpdate is the builder for updating HostCredential entities.
type HostCredentialUpdate struct {
	config
	hooks    []Hook
	mutation *HostCredentialMutation
}

// Where appends a list predicates to the HostCredentialUpdate builder.
func (hcu *HostCredentialUpdate) Where(ps ...predicate.HostCredential) *HostCredentialUpdate {
	hcu.mutation.Where(ps...)
	return hcu
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (hcu *HostCredentialUpdate) SetLastModifiedAt(t time.Time) *HostCredentialUpdate {
	hcu.mutation.SetLastModifiedAt(t)
	return hcu
}

// SetPrincipal sets the "principal" field.
func (hcu *HostCredentialUpdate) SetPrincipal(s string) *HostCredentialUpdate {
	hcu.mutation.SetPrincipal(s)
	return hcu
}

// SetSecret sets the "secret" field.
func (hcu *HostCredentialUpdate) SetSecret(s string) *HostCredentialUpdate {
	hcu.mutation.SetSecret(s)
	return hcu
}

// SetKind sets the "kind" field.
func (hcu *HostCredentialUpdate) SetKind(ek epb.Credential_Kind) *HostCredentialUpdate {
	hcu.mutation.SetKind(ek)
	return hcu
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (hcu *HostCredentialUpdate) SetHostID(id int) *HostCredentialUpdate {
	hcu.mutation.SetHostID(id)
	return hcu
}

// SetHost sets the "host" edge to the Host entity.
func (hcu *HostCredentialUpdate) SetHost(h *Host) *HostCredentialUpdate {
	return hcu.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (hcu *HostCredentialUpdate) SetTaskID(id int) *HostCredentialUpdate {
	hcu.mutation.SetTaskID(id)
	return hcu
}

// SetTask sets the "task" edge to the Task entity.
func (hcu *HostCredentialUpdate) SetTask(t *Task) *HostCredentialUpdate {
	return hcu.SetTaskID(t.ID)
}

// Mutation returns the HostCredentialMutation object of the builder.
func (hcu *HostCredentialUpdate) Mutation() *HostCredentialMutation {
	return hcu.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (hcu *HostCredentialUpdate) ClearHost() *HostCredentialUpdate {
	hcu.mutation.ClearHost()
	return hcu
}

// ClearTask clears the "task" edge to the Task entity.
func (hcu *HostCredentialUpdate) ClearTask() *HostCredentialUpdate {
	hcu.mutation.ClearTask()
	return hcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hcu *HostCredentialUpdate) Save(ctx context.Context) (int, error) {
	hcu.defaults()
	return withHooks(ctx, hcu.sqlSave, hcu.mutation, hcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hcu *HostCredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := hcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hcu *HostCredentialUpdate) Exec(ctx context.Context) error {
	_, err := hcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcu *HostCredentialUpdate) ExecX(ctx context.Context) {
	if err := hcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hcu *HostCredentialUpdate) defaults() {
	if _, ok := hcu.mutation.LastModifiedAt(); !ok {
		v := hostcredential.UpdateDefaultLastModifiedAt()
		hcu.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hcu *HostCredentialUpdate) check() error {
	if v, ok := hcu.mutation.Principal(); ok {
		if err := hostcredential.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "HostCredential.principal": %w`, err)}
		}
	}
	if v, ok := hcu.mutation.Secret(); ok {
		if err := hostcredential.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "HostCredential.secret": %w`, err)}
		}
	}
	if v, ok := hcu.mutation.Kind(); ok {
		if err := hostcredential.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "HostCredential.kind": %w`, err)}
		}
	}
	if _, ok := hcu.mutation.HostID(); hcu.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostCredential.host"`)
	}
	if _, ok := hcu.mutation.TaskID(); hcu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostCredential.task"`)
	}
	return nil
}

func (hcu *HostCredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hostcredential.Table, hostcredential.Columns, sqlgraph.NewFieldSpec(hostcredential.FieldID, field.TypeInt))
	if ps := hcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hcu.mutation.LastModifiedAt(); ok {
		_spec.SetField(hostcredential.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := hcu.mutation.Principal(); ok {
		_spec.SetField(hostcredential.FieldPrincipal, field.TypeString, value)
	}
	if value, ok := hcu.mutation.Secret(); ok {
		_spec.SetField(hostcredential.FieldSecret, field.TypeString, value)
	}
	if value, ok := hcu.mutation.Kind(); ok {
		_spec.SetField(hostcredential.FieldKind, field.TypeEnum, value)
	}
	if hcu.mutation.HostCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hcu.mutation.HostIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hcu.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hcu.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hostcredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hcu.mutation.done = true
	return n, nil
}

// HostCredentialUpdateOne is the builder for updating a single HostCredential entity.
type HostCredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostCredentialMutation
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (hcuo *HostCredentialUpdateOne) SetLastModifiedAt(t time.Time) *HostCredentialUpdateOne {
	hcuo.mutation.SetLastModifiedAt(t)
	return hcuo
}

// SetPrincipal sets the "principal" field.
func (hcuo *HostCredentialUpdateOne) SetPrincipal(s string) *HostCredentialUpdateOne {
	hcuo.mutation.SetPrincipal(s)
	return hcuo
}

// SetSecret sets the "secret" field.
func (hcuo *HostCredentialUpdateOne) SetSecret(s string) *HostCredentialUpdateOne {
	hcuo.mutation.SetSecret(s)
	return hcuo
}

// SetKind sets the "kind" field.
func (hcuo *HostCredentialUpdateOne) SetKind(ek epb.Credential_Kind) *HostCredentialUpdateOne {
	hcuo.mutation.SetKind(ek)
	return hcuo
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (hcuo *HostCredentialUpdateOne) SetHostID(id int) *HostCredentialUpdateOne {
	hcuo.mutation.SetHostID(id)
	return hcuo
}

// SetHost sets the "host" edge to the Host entity.
func (hcuo *HostCredentialUpdateOne) SetHost(h *Host) *HostCredentialUpdateOne {
	return hcuo.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (hcuo *HostCredentialUpdateOne) SetTaskID(id int) *HostCredentialUpdateOne {
	hcuo.mutation.SetTaskID(id)
	return hcuo
}

// SetTask sets the "task" edge to the Task entity.
func (hcuo *HostCredentialUpdateOne) SetTask(t *Task) *HostCredentialUpdateOne {
	return hcuo.SetTaskID(t.ID)
}

// Mutation returns the HostCredentialMutation object of the builder.
func (hcuo *HostCredentialUpdateOne) Mutation() *HostCredentialMutation {
	return hcuo.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (hcuo *HostCredentialUpdateOne) ClearHost() *HostCredentialUpdateOne {
	hcuo.mutation.ClearHost()
	return hcuo
}

// ClearTask clears the "task" edge to the Task entity.
func (hcuo *HostCredentialUpdateOne) ClearTask() *HostCredentialUpdateOne {
	hcuo.mutation.ClearTask()
	return hcuo
}

// Where appends a list predicates to the HostCredentialUpdate builder.
func (hcuo *HostCredentialUpdateOne) Where(ps ...predicate.HostCredential) *HostCredentialUpdateOne {
	hcuo.mutation.Where(ps...)
	return hcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hcuo *HostCredentialUpdateOne) Select(field string, fields ...string) *HostCredentialUpdateOne {
	hcuo.fields = append([]string{field}, fields...)
	return hcuo
}

// Save executes the query and returns the updated HostCredential entity.
func (hcuo *HostCredentialUpdateOne) Save(ctx context.Context) (*HostCredential, error) {
	hcuo.defaults()
	return withHooks(ctx, hcuo.sqlSave, hcuo.mutation, hcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hcuo *HostCredentialUpdateOne) SaveX(ctx context.Context) *HostCredential {
	node, err := hcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hcuo *HostCredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := hcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcuo *HostCredentialUpdateOne) ExecX(ctx context.Context) {
	if err := hcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hcuo *HostCredentialUpdateOne) defaults() {
	if _, ok := hcuo.mutation.LastModifiedAt(); !ok {
		v := hostcredential.UpdateDefaultLastModifiedAt()
		hcuo.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hcuo *HostCredentialUpdateOne) check() error {
	if v, ok := hcuo.mutation.Principal(); ok {
		if err := hostcredential.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "HostCredential.principal": %w`, err)}
		}
	}
	if v, ok := hcuo.mutation.Secret(); ok {
		if err := hostcredential.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`ent: validator failed for field "HostCredential.secret": %w`, err)}
		}
	}
	if v, ok := hcuo.mutation.Kind(); ok {
		if err := hostcredential.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "HostCredential.kind": %w`, err)}
		}
	}
	if _, ok := hcuo.mutation.HostID(); hcuo.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostCredential.host"`)
	}
	if _, ok := hcuo.mutation.TaskID(); hcuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostCredential.task"`)
	}
	return nil
}

func (hcuo *HostCredentialUpdateOne) sqlSave(ctx context.Context) (_node *HostCredential, err error) {
	if err := hcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hostcredential.Table, hostcredential.Columns, sqlgraph.NewFieldSpec(hostcredential.FieldID, field.TypeInt))
	id, ok := hcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HostCredential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hostcredential.FieldID)
		for _, f := range fields {
			if !hostcredential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hostcredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hcuo.mutation.LastModifiedAt(); ok {
		_spec.SetField(hostcredential.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := hcuo.mutation.Principal(); ok {
		_spec.SetField(hostcredential.FieldPrincipal, field.TypeString, value)
	}
	if value, ok := hcuo.mutation.Secret(); ok {
		_spec.SetField(hostcredential.FieldSecret, field.TypeString, value)
	}
	if value, ok := hcuo.mutation.Kind(); ok {
		_spec.SetField(hostcredential.FieldKind, field.TypeEnum, value)
	}
	if hcuo.mutation.HostCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hcuo.mutation.HostIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hcuo.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hcuo.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HostCredential{config: hcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hostcredential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hcuo.mutation.done = true
	return _node, nil
}
