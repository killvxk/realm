// Code generated by ent, DO NOT EDIT.

package tome

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tome type in the database.
	Label = "tome"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldParamDefs holds the string denoting the param_defs field in the database.
	FieldParamDefs = "param_defs"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldEldritch holds the string denoting the eldritch field in the database.
	FieldEldritch = "eldritch"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// Table holds the table name of the tome in the database.
	Table = "tomes"
	// FilesTable is the table that holds the files relation/edge. The primary key declared below.
	FilesTable = "tome_files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
)

// Columns holds all SQL columns for tome fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLastModifiedAt,
	FieldName,
	FieldDescription,
	FieldParamDefs,
	FieldHash,
	FieldEldritch,
}

var (
	// FilesPrimaryKey and FilesColumn2 are the table columns denoting the
	// primary key for the files relation (M2M).
	FilesPrimaryKey = []string{"tome_id", "file_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "realm.pub/tavern/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "last_modified_at" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "last_modified_at" field.
	UpdateDefaultLastModifiedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ParamDefsValidator is a validator for the "param_defs" field. It is called by the builders before save.
	ParamDefsValidator func(string) error
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
)

// OrderOption defines the ordering options for the Tome queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByParamDefs orders the results by the param_defs field.
func ByParamDefs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParamDefs, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByEldritch orders the results by the eldritch field.
func ByEldritch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEldritch, opts...).ToFunc()
}

// ByFilesCount orders the results by files count.
func ByFilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFilesStep(), opts...)
	}
}

// ByFiles orders the results by files terms.
func ByFiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FilesTable, FilesPrimaryKey...),
	)
}
