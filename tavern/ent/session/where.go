// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/realm/tavern/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Principal applies equality check predicate on the "principal" field. It's identical to PrincipalEQ.
func Principal(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrincipal), v))
	})
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// AgentIdentifier applies equality check predicate on the "agentIdentifier" field. It's identical to AgentIdentifierEQ.
func AgentIdentifier(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgentIdentifier), v))
	})
}

// HostIdentifier applies equality check predicate on the "hostIdentifier" field. It's identical to HostIdentifierEQ.
func HostIdentifier(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostIdentifier), v))
	})
}

// LastSeenAt applies equality check predicate on the "lastSeenAt" field. It's identical to LastSeenAtEQ.
func LastSeenAt(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeenAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PrincipalEQ applies the EQ predicate on the "principal" field.
func PrincipalEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrincipal), v))
	})
}

// PrincipalNEQ applies the NEQ predicate on the "principal" field.
func PrincipalNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrincipal), v))
	})
}

// PrincipalIn applies the In predicate on the "principal" field.
func PrincipalIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrincipal), v...))
	})
}

// PrincipalNotIn applies the NotIn predicate on the "principal" field.
func PrincipalNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrincipal), v...))
	})
}

// PrincipalGT applies the GT predicate on the "principal" field.
func PrincipalGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrincipal), v))
	})
}

// PrincipalGTE applies the GTE predicate on the "principal" field.
func PrincipalGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrincipal), v))
	})
}

// PrincipalLT applies the LT predicate on the "principal" field.
func PrincipalLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrincipal), v))
	})
}

// PrincipalLTE applies the LTE predicate on the "principal" field.
func PrincipalLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrincipal), v))
	})
}

// PrincipalContains applies the Contains predicate on the "principal" field.
func PrincipalContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrincipal), v))
	})
}

// PrincipalHasPrefix applies the HasPrefix predicate on the "principal" field.
func PrincipalHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrincipal), v))
	})
}

// PrincipalHasSuffix applies the HasSuffix predicate on the "principal" field.
func PrincipalHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrincipal), v))
	})
}

// PrincipalIsNil applies the IsNil predicate on the "principal" field.
func PrincipalIsNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrincipal)))
	})
}

// PrincipalNotNil applies the NotNil predicate on the "principal" field.
func PrincipalNotNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrincipal)))
	})
}

// PrincipalEqualFold applies the EqualFold predicate on the "principal" field.
func PrincipalEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrincipal), v))
	})
}

// PrincipalContainsFold applies the ContainsFold predicate on the "principal" field.
func PrincipalContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrincipal), v))
	})
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostname), v))
	})
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHostname), v...))
	})
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHostname), v...))
	})
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostname), v))
	})
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostname), v))
	})
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostname), v))
	})
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostname), v))
	})
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostname), v))
	})
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostname), v))
	})
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostname), v))
	})
}

// HostnameIsNil applies the IsNil predicate on the "hostname" field.
func HostnameIsNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHostname)))
	})
}

// HostnameNotNil applies the NotNil predicate on the "hostname" field.
func HostnameNotNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHostname)))
	})
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostname), v))
	})
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostname), v))
	})
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIdentifier), v...))
	})
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIdentifier), v...))
	})
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdentifier), v))
	})
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdentifier), v))
	})
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdentifier), v))
	})
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdentifier), v))
	})
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdentifier), v))
	})
}

// AgentIdentifierEQ applies the EQ predicate on the "agentIdentifier" field.
func AgentIdentifierEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierNEQ applies the NEQ predicate on the "agentIdentifier" field.
func AgentIdentifierNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierIn applies the In predicate on the "agentIdentifier" field.
func AgentIdentifierIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAgentIdentifier), v...))
	})
}

// AgentIdentifierNotIn applies the NotIn predicate on the "agentIdentifier" field.
func AgentIdentifierNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAgentIdentifier), v...))
	})
}

// AgentIdentifierGT applies the GT predicate on the "agentIdentifier" field.
func AgentIdentifierGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierGTE applies the GTE predicate on the "agentIdentifier" field.
func AgentIdentifierGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierLT applies the LT predicate on the "agentIdentifier" field.
func AgentIdentifierLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierLTE applies the LTE predicate on the "agentIdentifier" field.
func AgentIdentifierLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierContains applies the Contains predicate on the "agentIdentifier" field.
func AgentIdentifierContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierHasPrefix applies the HasPrefix predicate on the "agentIdentifier" field.
func AgentIdentifierHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierHasSuffix applies the HasSuffix predicate on the "agentIdentifier" field.
func AgentIdentifierHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierIsNil applies the IsNil predicate on the "agentIdentifier" field.
func AgentIdentifierIsNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAgentIdentifier)))
	})
}

// AgentIdentifierNotNil applies the NotNil predicate on the "agentIdentifier" field.
func AgentIdentifierNotNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAgentIdentifier)))
	})
}

// AgentIdentifierEqualFold applies the EqualFold predicate on the "agentIdentifier" field.
func AgentIdentifierEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAgentIdentifier), v))
	})
}

// AgentIdentifierContainsFold applies the ContainsFold predicate on the "agentIdentifier" field.
func AgentIdentifierContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAgentIdentifier), v))
	})
}

// HostIdentifierEQ applies the EQ predicate on the "hostIdentifier" field.
func HostIdentifierEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierNEQ applies the NEQ predicate on the "hostIdentifier" field.
func HostIdentifierNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierIn applies the In predicate on the "hostIdentifier" field.
func HostIdentifierIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHostIdentifier), v...))
	})
}

// HostIdentifierNotIn applies the NotIn predicate on the "hostIdentifier" field.
func HostIdentifierNotIn(vs ...string) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHostIdentifier), v...))
	})
}

// HostIdentifierGT applies the GT predicate on the "hostIdentifier" field.
func HostIdentifierGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierGTE applies the GTE predicate on the "hostIdentifier" field.
func HostIdentifierGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierLT applies the LT predicate on the "hostIdentifier" field.
func HostIdentifierLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierLTE applies the LTE predicate on the "hostIdentifier" field.
func HostIdentifierLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierContains applies the Contains predicate on the "hostIdentifier" field.
func HostIdentifierContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierHasPrefix applies the HasPrefix predicate on the "hostIdentifier" field.
func HostIdentifierHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierHasSuffix applies the HasSuffix predicate on the "hostIdentifier" field.
func HostIdentifierHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierIsNil applies the IsNil predicate on the "hostIdentifier" field.
func HostIdentifierIsNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHostIdentifier)))
	})
}

// HostIdentifierNotNil applies the NotNil predicate on the "hostIdentifier" field.
func HostIdentifierNotNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHostIdentifier)))
	})
}

// HostIdentifierEqualFold applies the EqualFold predicate on the "hostIdentifier" field.
func HostIdentifierEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostIdentifier), v))
	})
}

// HostIdentifierContainsFold applies the ContainsFold predicate on the "hostIdentifier" field.
func HostIdentifierContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostIdentifier), v))
	})
}

// LastSeenAtEQ applies the EQ predicate on the "lastSeenAt" field.
func LastSeenAtEQ(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtNEQ applies the NEQ predicate on the "lastSeenAt" field.
func LastSeenAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtIn applies the In predicate on the "lastSeenAt" field.
func LastSeenAtIn(vs ...time.Time) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastSeenAt), v...))
	})
}

// LastSeenAtNotIn applies the NotIn predicate on the "lastSeenAt" field.
func LastSeenAtNotIn(vs ...time.Time) predicate.Session {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastSeenAt), v...))
	})
}

// LastSeenAtGT applies the GT predicate on the "lastSeenAt" field.
func LastSeenAtGT(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtGTE applies the GTE predicate on the "lastSeenAt" field.
func LastSeenAtGTE(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtLT applies the LT predicate on the "lastSeenAt" field.
func LastSeenAtLT(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtLTE applies the LTE predicate on the "lastSeenAt" field.
func LastSeenAtLTE(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSeenAt), v))
	})
}

// LastSeenAtIsNil applies the IsNil predicate on the "lastSeenAt" field.
func LastSeenAtIsNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastSeenAt)))
	})
}

// LastSeenAtNotNil applies the NotNil predicate on the "lastSeenAt" field.
func LastSeenAtNotNil() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastSeenAt)))
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		p(s.Not())
	})
}
