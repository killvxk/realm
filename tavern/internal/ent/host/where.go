// Code generated by ent, DO NOT EDIT.

package host

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"realm.pub/tavern/internal/c2/c2pb"
	"realm.pub/tavern/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldCreatedAt, v))
}

// LastModifiedAt applies equality check predicate on the "last_modified_at" field. It's identical to LastModifiedAtEQ.
func LastModifiedAt(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldLastModifiedAt, v))
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldIdentifier, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldName, v))
}

// PrimaryIP applies equality check predicate on the "primary_ip" field. It's identical to PrimaryIPEQ.
func PrimaryIP(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldPrimaryIP, v))
}

// LastSeenAt applies equality check predicate on the "last_seen_at" field. It's identical to LastSeenAtEQ.
func LastSeenAt(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldLastSeenAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldCreatedAt, v))
}

// LastModifiedAtEQ applies the EQ predicate on the "last_modified_at" field.
func LastModifiedAtEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtNEQ applies the NEQ predicate on the "last_modified_at" field.
func LastModifiedAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtIn applies the In predicate on the "last_modified_at" field.
func LastModifiedAtIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtNotIn applies the NotIn predicate on the "last_modified_at" field.
func LastModifiedAtNotIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtGT applies the GT predicate on the "last_modified_at" field.
func LastModifiedAtGT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldLastModifiedAt, v))
}

// LastModifiedAtGTE applies the GTE predicate on the "last_modified_at" field.
func LastModifiedAtGTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldLastModifiedAt, v))
}

// LastModifiedAtLT applies the LT predicate on the "last_modified_at" field.
func LastModifiedAtLT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldLastModifiedAt, v))
}

// LastModifiedAtLTE applies the LTE predicate on the "last_modified_at" field.
func LastModifiedAtLTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldLastModifiedAt, v))
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldIdentifier, v))
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.Host {
	return predicate.Host(sql.FieldContains(FieldIdentifier, v))
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasPrefix(FieldIdentifier, v))
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasSuffix(FieldIdentifier, v))
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.Host {
	return predicate.Host(sql.FieldEqualFold(FieldIdentifier, v))
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.Host {
	return predicate.Host(sql.FieldContainsFold(FieldIdentifier, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Host {
	return predicate.Host(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Host {
	return predicate.Host(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Host {
	return predicate.Host(sql.FieldContainsFold(FieldName, v))
}

// PrimaryIPEQ applies the EQ predicate on the "primary_ip" field.
func PrimaryIPEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldPrimaryIP, v))
}

// PrimaryIPNEQ applies the NEQ predicate on the "primary_ip" field.
func PrimaryIPNEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldPrimaryIP, v))
}

// PrimaryIPIn applies the In predicate on the "primary_ip" field.
func PrimaryIPIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldPrimaryIP, vs...))
}

// PrimaryIPNotIn applies the NotIn predicate on the "primary_ip" field.
func PrimaryIPNotIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldPrimaryIP, vs...))
}

// PrimaryIPGT applies the GT predicate on the "primary_ip" field.
func PrimaryIPGT(v string) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldPrimaryIP, v))
}

// PrimaryIPGTE applies the GTE predicate on the "primary_ip" field.
func PrimaryIPGTE(v string) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldPrimaryIP, v))
}

// PrimaryIPLT applies the LT predicate on the "primary_ip" field.
func PrimaryIPLT(v string) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldPrimaryIP, v))
}

// PrimaryIPLTE applies the LTE predicate on the "primary_ip" field.
func PrimaryIPLTE(v string) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldPrimaryIP, v))
}

// PrimaryIPContains applies the Contains predicate on the "primary_ip" field.
func PrimaryIPContains(v string) predicate.Host {
	return predicate.Host(sql.FieldContains(FieldPrimaryIP, v))
}

// PrimaryIPHasPrefix applies the HasPrefix predicate on the "primary_ip" field.
func PrimaryIPHasPrefix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasPrefix(FieldPrimaryIP, v))
}

// PrimaryIPHasSuffix applies the HasSuffix predicate on the "primary_ip" field.
func PrimaryIPHasSuffix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasSuffix(FieldPrimaryIP, v))
}

// PrimaryIPIsNil applies the IsNil predicate on the "primary_ip" field.
func PrimaryIPIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldPrimaryIP))
}

// PrimaryIPNotNil applies the NotNil predicate on the "primary_ip" field.
func PrimaryIPNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldPrimaryIP))
}

// PrimaryIPEqualFold applies the EqualFold predicate on the "primary_ip" field.
func PrimaryIPEqualFold(v string) predicate.Host {
	return predicate.Host(sql.FieldEqualFold(FieldPrimaryIP, v))
}

// PrimaryIPContainsFold applies the ContainsFold predicate on the "primary_ip" field.
func PrimaryIPContainsFold(v string) predicate.Host {
	return predicate.Host(sql.FieldContainsFold(FieldPrimaryIP, v))
}

// PlatformEQ applies the EQ predicate on the "platform" field.
func PlatformEQ(v c2pb.Host_Platform) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldPlatform, v))
}

// PlatformNEQ applies the NEQ predicate on the "platform" field.
func PlatformNEQ(v c2pb.Host_Platform) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldPlatform, v))
}

// PlatformIn applies the In predicate on the "platform" field.
func PlatformIn(vs ...c2pb.Host_Platform) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldPlatform, vs...))
}

// PlatformNotIn applies the NotIn predicate on the "platform" field.
func PlatformNotIn(vs ...c2pb.Host_Platform) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldPlatform, vs...))
}

// LastSeenAtEQ applies the EQ predicate on the "last_seen_at" field.
func LastSeenAtEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldLastSeenAt, v))
}

// LastSeenAtNEQ applies the NEQ predicate on the "last_seen_at" field.
func LastSeenAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldLastSeenAt, v))
}

// LastSeenAtIn applies the In predicate on the "last_seen_at" field.
func LastSeenAtIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldLastSeenAt, vs...))
}

// LastSeenAtNotIn applies the NotIn predicate on the "last_seen_at" field.
func LastSeenAtNotIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldLastSeenAt, vs...))
}

// LastSeenAtGT applies the GT predicate on the "last_seen_at" field.
func LastSeenAtGT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldLastSeenAt, v))
}

// LastSeenAtGTE applies the GTE predicate on the "last_seen_at" field.
func LastSeenAtGTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldLastSeenAt, v))
}

// LastSeenAtLT applies the LT predicate on the "last_seen_at" field.
func LastSeenAtLT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldLastSeenAt, v))
}

// LastSeenAtLTE applies the LTE predicate on the "last_seen_at" field.
func LastSeenAtLTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldLastSeenAt, v))
}

// LastSeenAtIsNil applies the IsNil predicate on the "last_seen_at" field.
func LastSeenAtIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldLastSeenAt))
}

// LastSeenAtNotNil applies the NotNil predicate on the "last_seen_at" field.
func LastSeenAtNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldLastSeenAt))
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBeacons applies the HasEdge predicate on the "beacons" edge.
func HasBeacons() predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BeaconsTable, BeaconsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBeaconsWith applies the HasEdge predicate on the "beacons" edge with a given conditions (other predicates).
func HasBeaconsWith(preds ...predicate.Beacon) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := newBeaconsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.HostFile) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProcesses applies the HasEdge predicate on the "processes" edge.
func HasProcesses() predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProcessesTable, ProcessesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProcessesWith applies the HasEdge predicate on the "processes" edge with a given conditions (other predicates).
func HasProcessesWith(preds ...predicate.HostProcess) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := newProcessesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCredentials applies the HasEdge predicate on the "credentials" edge.
func HasCredentials() predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CredentialsTable, CredentialsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCredentialsWith applies the HasEdge predicate on the "credentials" edge with a given conditions (other predicates).
func HasCredentialsWith(preds ...predicate.HostCredential) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := newCredentialsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Host) predicate.Host {
	return predicate.Host(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Host) predicate.Host {
	return predicate.Host(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Host) predicate.Host {
	return predicate.Host(sql.NotPredicates(p))
}
