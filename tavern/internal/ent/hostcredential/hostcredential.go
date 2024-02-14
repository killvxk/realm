// Code generated by ent, DO NOT EDIT.

package hostcredential

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"realm.pub/tavern/internal/c2/epb"
)

const (
	// Label holds the string label denoting the hostcredential type in the database.
	Label = "host_credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldPrincipal holds the string denoting the principal field in the database.
	FieldPrincipal = "principal"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// Table holds the table name of the hostcredential in the database.
	Table = "host_credentials"
	// HostTable is the table that holds the host relation/edge.
	HostTable = "host_credentials"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "host_credential_host"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "host_credentials"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_reported_credentials"
)

// Columns holds all SQL columns for hostcredential fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLastModifiedAt,
	FieldPrincipal,
	FieldSecret,
	FieldKind,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "host_credentials"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"host_credential_host",
	"task_reported_credentials",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "last_modified_at" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "last_modified_at" field.
	UpdateDefaultLastModifiedAt func() time.Time
	// PrincipalValidator is a validator for the "principal" field. It is called by the builders before save.
	PrincipalValidator func(string) error
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator func(string) error
)

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k epb.Credential_Kind) error {
	switch k.String() {
	case "KIND_PASSWORD", "KIND_SSH_KEY", "KIND_UNSPECIFIED":
		return nil
	default:
		return fmt.Errorf("hostcredential: invalid enum value for kind field: %q", k)
	}
}

// OrderOption defines the ordering options for the HostCredential queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLastModifiedAt orders the results by the last_modified_at field.
func ByLastModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastModifiedAt, opts...).ToFunc()
}

// ByPrincipal orders the results by the principal field.
func ByPrincipal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrincipal, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByKind orders the results by the kind field.
func ByKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKind, opts...).ToFunc()
}

// ByHostField orders the results by host field.
func ByHostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostStep(), sql.OrderByField(field, opts...))
	}
}

// ByTaskField orders the results by task field.
func ByTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskStep(), sql.OrderByField(field, opts...))
	}
}
func newHostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, HostTable, HostColumn),
	)
}
func newTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
	)
}

var (
	// epb.Credential_Kind must implement graphql.Marshaler.
	_ graphql.Marshaler = (*epb.Credential_Kind)(nil)
	// epb.Credential_Kind must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*epb.Credential_Kind)(nil)
)
