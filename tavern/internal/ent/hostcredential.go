// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/hostcredential"
	"realm.pub/tavern/internal/ent/task"
)

// HostCredential is the model entity for the HostCredential schema.
type HostCredential struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Timestamp of when this ent was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Timestamp of when this ent was last updated
	LastModifiedAt time.Time `json:"last_modified_at,omitempty"`
	// Identity associated with this credential (e.g. username).
	Principal string `json:"principal,omitempty"`
	// Secret for this credential (e.g. password).
	Secret string `json:"secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostCredentialQuery when eager-loading is set.
	Edges                     HostCredentialEdges `json:"edges"`
	host_credential_host      *int
	task_reported_credentials *int
	selectValues              sql.SelectValues
}

// HostCredentialEdges holds the relations/edges for other nodes in the graph.
type HostCredentialEdges struct {
	// Host the credential was reported on.
	Host *Host `json:"host,omitempty"`
	// Task that reported this credential.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// HostOrErr returns the Host value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostCredentialEdges) HostOrErr() (*Host, error) {
	if e.loadedTypes[0] {
		if e.Host == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.Host, nil
	}
	return nil, &NotLoadedError{edge: "host"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostCredentialEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[1] {
		if e.Task == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HostCredential) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hostcredential.FieldID:
			values[i] = new(sql.NullInt64)
		case hostcredential.FieldPrincipal, hostcredential.FieldSecret:
			values[i] = new(sql.NullString)
		case hostcredential.FieldCreatedAt, hostcredential.FieldLastModifiedAt:
			values[i] = new(sql.NullTime)
		case hostcredential.ForeignKeys[0]: // host_credential_host
			values[i] = new(sql.NullInt64)
		case hostcredential.ForeignKeys[1]: // task_reported_credentials
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HostCredential fields.
func (hc *HostCredential) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hostcredential.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hc.ID = int(value.Int64)
		case hostcredential.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hc.CreatedAt = value.Time
			}
		case hostcredential.FieldLastModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_modified_at", values[i])
			} else if value.Valid {
				hc.LastModifiedAt = value.Time
			}
		case hostcredential.FieldPrincipal:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field principal", values[i])
			} else if value.Valid {
				hc.Principal = value.String
			}
		case hostcredential.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				hc.Secret = value.String
			}
		case hostcredential.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field host_credential_host", value)
			} else if value.Valid {
				hc.host_credential_host = new(int)
				*hc.host_credential_host = int(value.Int64)
			}
		case hostcredential.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_reported_credentials", value)
			} else if value.Valid {
				hc.task_reported_credentials = new(int)
				*hc.task_reported_credentials = int(value.Int64)
			}
		default:
			hc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HostCredential.
// This includes values selected through modifiers, order, etc.
func (hc *HostCredential) Value(name string) (ent.Value, error) {
	return hc.selectValues.Get(name)
}

// QueryHost queries the "host" edge of the HostCredential entity.
func (hc *HostCredential) QueryHost() *HostQuery {
	return NewHostCredentialClient(hc.config).QueryHost(hc)
}

// QueryTask queries the "task" edge of the HostCredential entity.
func (hc *HostCredential) QueryTask() *TaskQuery {
	return NewHostCredentialClient(hc.config).QueryTask(hc)
}

// Update returns a builder for updating this HostCredential.
// Note that you need to call HostCredential.Unwrap() before calling this method if this HostCredential
// was returned from a transaction, and the transaction was committed or rolled back.
func (hc *HostCredential) Update() *HostCredentialUpdateOne {
	return NewHostCredentialClient(hc.config).UpdateOne(hc)
}

// Unwrap unwraps the HostCredential entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hc *HostCredential) Unwrap() *HostCredential {
	_tx, ok := hc.config.driver.(*txDriver)
	if !ok {
		panic("ent: HostCredential is not a transactional entity")
	}
	hc.config.driver = _tx.drv
	return hc
}

// String implements the fmt.Stringer.
func (hc *HostCredential) String() string {
	var builder strings.Builder
	builder.WriteString("HostCredential(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(hc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_modified_at=")
	builder.WriteString(hc.LastModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("principal=")
	builder.WriteString(hc.Principal)
	builder.WriteString(", ")
	builder.WriteString("secret=")
	builder.WriteString(hc.Secret)
	builder.WriteByte(')')
	return builder.String()
}

// HostCredentials is a parsable slice of HostCredential.
type HostCredentials []*HostCredential
