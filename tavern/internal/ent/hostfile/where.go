// Code generated by ent, DO NOT EDIT.

package hostfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"realm.pub/tavern/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldCreatedAt, v))
}

// LastModifiedAt applies equality check predicate on the "last_modified_at" field. It's identical to LastModifiedAtEQ.
func LastModifiedAt(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldLastModifiedAt, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldPath, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldOwner, v))
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldGroup, v))
}

// Permissions applies equality check predicate on the "permissions" field. It's identical to PermissionsEQ.
func Permissions(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldPermissions, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldSize, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldHash, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldCreatedAt, v))
}

// LastModifiedAtEQ applies the EQ predicate on the "last_modified_at" field.
func LastModifiedAtEQ(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtNEQ applies the NEQ predicate on the "last_modified_at" field.
func LastModifiedAtNEQ(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtIn applies the In predicate on the "last_modified_at" field.
func LastModifiedAtIn(vs ...time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtNotIn applies the NotIn predicate on the "last_modified_at" field.
func LastModifiedAtNotIn(vs ...time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtGT applies the GT predicate on the "last_modified_at" field.
func LastModifiedAtGT(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldLastModifiedAt, v))
}

// LastModifiedAtGTE applies the GTE predicate on the "last_modified_at" field.
func LastModifiedAtGTE(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldLastModifiedAt, v))
}

// LastModifiedAtLT applies the LT predicate on the "last_modified_at" field.
func LastModifiedAtLT(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldLastModifiedAt, v))
}

// LastModifiedAtLTE applies the LTE predicate on the "last_modified_at" field.
func LastModifiedAtLTE(v time.Time) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldLastModifiedAt, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContainsFold(FieldPath, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldOwner, v))
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContains(FieldOwner, v))
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasPrefix(FieldOwner, v))
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasSuffix(FieldOwner, v))
}

// OwnerIsNil applies the IsNil predicate on the "owner" field.
func OwnerIsNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldIsNull(FieldOwner))
}

// OwnerNotNil applies the NotNil predicate on the "owner" field.
func OwnerNotNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldNotNull(FieldOwner))
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEqualFold(FieldOwner, v))
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContainsFold(FieldOwner, v))
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldGroup, v))
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContains(FieldGroup, v))
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasPrefix(FieldGroup, v))
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasSuffix(FieldGroup, v))
}

// GroupIsNil applies the IsNil predicate on the "group" field.
func GroupIsNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldIsNull(FieldGroup))
}

// GroupNotNil applies the NotNil predicate on the "group" field.
func GroupNotNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldNotNull(FieldGroup))
}

// GroupEqualFold applies the EqualFold predicate on the "group" field.
func GroupEqualFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEqualFold(FieldGroup, v))
}

// GroupContainsFold applies the ContainsFold predicate on the "group" field.
func GroupContainsFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContainsFold(FieldGroup, v))
}

// PermissionsEQ applies the EQ predicate on the "permissions" field.
func PermissionsEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldPermissions, v))
}

// PermissionsNEQ applies the NEQ predicate on the "permissions" field.
func PermissionsNEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldPermissions, v))
}

// PermissionsIn applies the In predicate on the "permissions" field.
func PermissionsIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldPermissions, vs...))
}

// PermissionsNotIn applies the NotIn predicate on the "permissions" field.
func PermissionsNotIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldPermissions, vs...))
}

// PermissionsGT applies the GT predicate on the "permissions" field.
func PermissionsGT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldPermissions, v))
}

// PermissionsGTE applies the GTE predicate on the "permissions" field.
func PermissionsGTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldPermissions, v))
}

// PermissionsLT applies the LT predicate on the "permissions" field.
func PermissionsLT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldPermissions, v))
}

// PermissionsLTE applies the LTE predicate on the "permissions" field.
func PermissionsLTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldPermissions, v))
}

// PermissionsContains applies the Contains predicate on the "permissions" field.
func PermissionsContains(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContains(FieldPermissions, v))
}

// PermissionsHasPrefix applies the HasPrefix predicate on the "permissions" field.
func PermissionsHasPrefix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasPrefix(FieldPermissions, v))
}

// PermissionsHasSuffix applies the HasSuffix predicate on the "permissions" field.
func PermissionsHasSuffix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasSuffix(FieldPermissions, v))
}

// PermissionsIsNil applies the IsNil predicate on the "permissions" field.
func PermissionsIsNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldIsNull(FieldPermissions))
}

// PermissionsNotNil applies the NotNil predicate on the "permissions" field.
func PermissionsNotNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldNotNull(FieldPermissions))
}

// PermissionsEqualFold applies the EqualFold predicate on the "permissions" field.
func PermissionsEqualFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEqualFold(FieldPermissions, v))
}

// PermissionsContainsFold applies the ContainsFold predicate on the "permissions" field.
func PermissionsContainsFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContainsFold(FieldPermissions, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v uint64) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldSize, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldHasSuffix(FieldHash, v))
}

// HashIsNil applies the IsNil predicate on the "hash" field.
func HashIsNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldIsNull(FieldHash))
}

// HashNotNil applies the NotNil predicate on the "hash" field.
func HashNotNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldNotNull(FieldHash))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.HostFile {
	return predicate.HostFile(sql.FieldContainsFold(FieldHash, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...[]byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...[]byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v []byte) predicate.HostFile {
	return predicate.HostFile(sql.FieldLTE(FieldContent, v))
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldIsNull(FieldContent))
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.HostFile {
	return predicate.HostFile(sql.FieldNotNull(FieldContent))
}

// HasHost applies the HasEdge predicate on the "host" edge.
func HasHost() predicate.HostFile {
	return predicate.HostFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostTable, HostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostWith applies the HasEdge predicate on the "host" edge with a given conditions (other predicates).
func HasHostWith(preds ...predicate.Host) predicate.HostFile {
	return predicate.HostFile(func(s *sql.Selector) {
		step := newHostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.HostFile {
	return predicate.HostFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.HostFile {
	return predicate.HostFile(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HostFile) predicate.HostFile {
	return predicate.HostFile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HostFile) predicate.HostFile {
	return predicate.HostFile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.HostFile) predicate.HostFile {
	return predicate.HostFile(sql.NotPredicates(p))
}
