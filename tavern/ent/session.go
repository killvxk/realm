// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kcarretto/realm/tavern/ent/session"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// A human readable identifier for the session.
	Name string `json:"name,omitempty"`
	// The identity the session is authenticated as (e.g. 'root')
	Principal string `json:"principal,omitempty"`
	// The hostname of the system the session is running on.
	Hostname string `json:"hostname,omitempty"`
	// Unique identifier for the session. Unique to each instance of the session.
	Identifier string `json:"identifier,omitempty"`
	// Identifies the agent that the session is running as (e.g. 'imix').
	AgentIdentifier string `json:"agentIdentifier,omitempty"`
	// Unique identifier for the host the session is running on.
	HostIdentifier string `json:"hostIdentifier,omitempty"`
	// Timestamp of when a task was last claimed or updated for a target
	LastSeenAt time.Time `json:"lastSeenAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges SessionEdges `json:"edges"`
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Tags used to group the session with other sessions
	Tags []*Tag `json:"tags,omitempty"`
	// Tasks that have been assigned to the session
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTags  map[string][]*Tag
	namedTasks map[string][]*Task
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e SessionEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e SessionEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			values[i] = new(sql.NullInt64)
		case session.FieldName, session.FieldPrincipal, session.FieldHostname, session.FieldIdentifier, session.FieldAgentIdentifier, session.FieldHostIdentifier:
			values[i] = new(sql.NullString)
		case session.FieldLastSeenAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case session.FieldPrincipal:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field principal", values[i])
			} else if value.Valid {
				s.Principal = value.String
			}
		case session.FieldHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[i])
			} else if value.Valid {
				s.Hostname = value.String
			}
		case session.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				s.Identifier = value.String
			}
		case session.FieldAgentIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agentIdentifier", values[i])
			} else if value.Valid {
				s.AgentIdentifier = value.String
			}
		case session.FieldHostIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostIdentifier", values[i])
			} else if value.Valid {
				s.HostIdentifier = value.String
			}
		case session.FieldLastSeenAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastSeenAt", values[i])
			} else if value.Valid {
				s.LastSeenAt = value.Time
			}
		}
	}
	return nil
}

// QueryTags queries the "tags" edge of the Session entity.
func (s *Session) QueryTags() *TagQuery {
	return (&SessionClient{config: s.config}).QueryTags(s)
}

// QueryTasks queries the "tasks" edge of the Session entity.
func (s *Session) QueryTasks() *TaskQuery {
	return (&SessionClient{config: s.config}).QueryTasks(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("principal=")
	builder.WriteString(s.Principal)
	builder.WriteString(", ")
	builder.WriteString("hostname=")
	builder.WriteString(s.Hostname)
	builder.WriteString(", ")
	builder.WriteString("identifier=")
	builder.WriteString(s.Identifier)
	builder.WriteString(", ")
	builder.WriteString("agentIdentifier=")
	builder.WriteString(s.AgentIdentifier)
	builder.WriteString(", ")
	builder.WriteString("hostIdentifier=")
	builder.WriteString(s.HostIdentifier)
	builder.WriteString(", ")
	builder.WriteString("lastSeenAt=")
	builder.WriteString(s.LastSeenAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTags returns the Tags named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Session) NamedTags(name string) ([]*Tag, error) {
	if s.Edges.namedTags == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedTags[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Session) appendNamedTags(name string, edges ...*Tag) {
	if s.Edges.namedTags == nil {
		s.Edges.namedTags = make(map[string][]*Tag)
	}
	if len(edges) == 0 {
		s.Edges.namedTags[name] = []*Tag{}
	} else {
		s.Edges.namedTags[name] = append(s.Edges.namedTags[name], edges...)
	}
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Session) NamedTasks(name string) ([]*Task, error) {
	if s.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Session) appendNamedTasks(name string, edges ...*Task) {
	if s.Edges.namedTasks == nil {
		s.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		s.Edges.namedTasks[name] = []*Task{}
	} else {
		s.Edges.namedTasks[name] = append(s.Edges.namedTasks[name], edges...)
	}
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
