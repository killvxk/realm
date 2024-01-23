// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"realm.pub/tavern/internal/ent/tag"
	"realm.pub/tavern/internal/ent/tome"
)

// UpdateBeaconInput represents a mutation input for updating beacons.
type UpdateBeaconInput struct {
	LastModifiedAt *time.Time
	HostID         *int
}

// Mutate applies the UpdateBeaconInput on the BeaconMutation builder.
func (i *UpdateBeaconInput) Mutate(m *BeaconMutation) {
	if v := i.LastModifiedAt; v != nil {
		m.SetLastModifiedAt(*v)
	}
	if v := i.HostID; v != nil {
		m.SetHostID(*v)
	}
}

// SetInput applies the change-set in the UpdateBeaconInput on the BeaconUpdate builder.
func (c *BeaconUpdate) SetInput(i UpdateBeaconInput) *BeaconUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateBeaconInput on the BeaconUpdateOne builder.
func (c *BeaconUpdateOne) SetInput(i UpdateBeaconInput) *BeaconUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// UpdateHostInput represents a mutation input for updating hosts.
type UpdateHostInput struct {
	LastModifiedAt   *time.Time
	ClearName        bool
	Name             *string
	ClearTags        bool
	AddTagIDs        []int
	RemoveTagIDs     []int
	ClearBeacons     bool
	AddBeaconIDs     []int
	RemoveBeaconIDs  []int
	ClearProcesses   bool
	AddProcessIDs    []int
	RemoveProcessIDs []int
}

// Mutate applies the UpdateHostInput on the HostMutation builder.
func (i *UpdateHostInput) Mutate(m *HostMutation) {
	if v := i.LastModifiedAt; v != nil {
		m.SetLastModifiedAt(*v)
	}
	if i.ClearName {
		m.ClearName()
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearTags {
		m.ClearTags()
	}
	if v := i.AddTagIDs; len(v) > 0 {
		m.AddTagIDs(v...)
	}
	if v := i.RemoveTagIDs; len(v) > 0 {
		m.RemoveTagIDs(v...)
	}
	if i.ClearBeacons {
		m.ClearBeacons()
	}
	if v := i.AddBeaconIDs; len(v) > 0 {
		m.AddBeaconIDs(v...)
	}
	if v := i.RemoveBeaconIDs; len(v) > 0 {
		m.RemoveBeaconIDs(v...)
	}
	if i.ClearProcesses {
		m.ClearProcesses()
	}
	if v := i.AddProcessIDs; len(v) > 0 {
		m.AddProcessIDs(v...)
	}
	if v := i.RemoveProcessIDs; len(v) > 0 {
		m.RemoveProcessIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateHostInput on the HostUpdate builder.
func (c *HostUpdate) SetInput(i UpdateHostInput) *HostUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateHostInput on the HostUpdateOne builder.
func (c *HostUpdateOne) SetInput(i UpdateHostInput) *HostUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateQuestInput represents a mutation input for creating quests.
type CreateQuestInput struct {
	Name       string
	Parameters *string
	TomeID     int
}

// Mutate applies the CreateQuestInput on the QuestMutation builder.
func (i *CreateQuestInput) Mutate(m *QuestMutation) {
	m.SetName(i.Name)
	if v := i.Parameters; v != nil {
		m.SetParameters(*v)
	}
	m.SetTomeID(i.TomeID)
}

// SetInput applies the change-set in the CreateQuestInput on the QuestCreate builder.
func (c *QuestCreate) SetInput(i CreateQuestInput) *QuestCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateTagInput represents a mutation input for creating tags.
type CreateTagInput struct {
	Name    string
	Kind    tag.Kind
	HostIDs []int
}

// Mutate applies the CreateTagInput on the TagMutation builder.
func (i *CreateTagInput) Mutate(m *TagMutation) {
	m.SetName(i.Name)
	m.SetKind(i.Kind)
	if v := i.HostIDs; len(v) > 0 {
		m.AddHostIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTagInput on the TagCreate builder.
func (c *TagCreate) SetInput(i CreateTagInput) *TagCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTagInput represents a mutation input for updating tags.
type UpdateTagInput struct {
	Name          *string
	Kind          *tag.Kind
	ClearHosts    bool
	AddHostIDs    []int
	RemoveHostIDs []int
}

// Mutate applies the UpdateTagInput on the TagMutation builder.
func (i *UpdateTagInput) Mutate(m *TagMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if i.ClearHosts {
		m.ClearHosts()
	}
	if v := i.AddHostIDs; len(v) > 0 {
		m.AddHostIDs(v...)
	}
	if v := i.RemoveHostIDs; len(v) > 0 {
		m.RemoveHostIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateTagInput on the TagUpdate builder.
func (c *TagUpdate) SetInput(i UpdateTagInput) *TagUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTagInput on the TagUpdateOne builder.
func (c *TagUpdateOne) SetInput(i UpdateTagInput) *TagUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTomeInput represents a mutation input for creating tomes.
type CreateTomeInput struct {
	Name         string
	Description  string
	Author       string
	SupportModel *tome.SupportModel
	Tactic       *tome.Tactic
	ParamDefs    *string
	Eldritch     string
	FileIDs      []int
}

// Mutate applies the CreateTomeInput on the TomeMutation builder.
func (i *CreateTomeInput) Mutate(m *TomeMutation) {
	m.SetName(i.Name)
	m.SetDescription(i.Description)
	m.SetAuthor(i.Author)
	if v := i.SupportModel; v != nil {
		m.SetSupportModel(*v)
	}
	if v := i.Tactic; v != nil {
		m.SetTactic(*v)
	}
	if v := i.ParamDefs; v != nil {
		m.SetParamDefs(*v)
	}
	m.SetEldritch(i.Eldritch)
	if v := i.FileIDs; len(v) > 0 {
		m.AddFileIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTomeInput on the TomeCreate builder.
func (c *TomeCreate) SetInput(i CreateTomeInput) *TomeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTomeInput represents a mutation input for updating tomes.
type UpdateTomeInput struct {
	LastModifiedAt *time.Time
	Name           *string
	Description    *string
	Author         *string
	SupportModel   *tome.SupportModel
	Tactic         *tome.Tactic
	ClearParamDefs bool
	ParamDefs      *string
	Eldritch       *string
	ClearFiles     bool
	AddFileIDs     []int
	RemoveFileIDs  []int
}

// Mutate applies the UpdateTomeInput on the TomeMutation builder.
func (i *UpdateTomeInput) Mutate(m *TomeMutation) {
	if v := i.LastModifiedAt; v != nil {
		m.SetLastModifiedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.Author; v != nil {
		m.SetAuthor(*v)
	}
	if v := i.SupportModel; v != nil {
		m.SetSupportModel(*v)
	}
	if v := i.Tactic; v != nil {
		m.SetTactic(*v)
	}
	if i.ClearParamDefs {
		m.ClearParamDefs()
	}
	if v := i.ParamDefs; v != nil {
		m.SetParamDefs(*v)
	}
	if v := i.Eldritch; v != nil {
		m.SetEldritch(*v)
	}
	if i.ClearFiles {
		m.ClearFiles()
	}
	if v := i.AddFileIDs; len(v) > 0 {
		m.AddFileIDs(v...)
	}
	if v := i.RemoveFileIDs; len(v) > 0 {
		m.RemoveFileIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateTomeInput on the TomeUpdate builder.
func (c *TomeUpdate) SetInput(i UpdateTomeInput) *TomeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTomeInput on the TomeUpdateOne builder.
func (c *TomeUpdateOne) SetInput(i UpdateTomeInput) *TomeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name          *string
	PhotoURL      *string
	IsActivated   *bool
	IsAdmin       *bool
	ClearTomes    bool
	AddTomeIDs    []int
	RemoveTomeIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.PhotoURL; v != nil {
		m.SetPhotoURL(*v)
	}
	if v := i.IsActivated; v != nil {
		m.SetIsActivated(*v)
	}
	if v := i.IsAdmin; v != nil {
		m.SetIsAdmin(*v)
	}
	if i.ClearTomes {
		m.ClearTomes()
	}
	if v := i.AddTomeIDs; len(v) > 0 {
		m.AddTomeIDs(v...)
	}
	if v := i.RemoveTomeIDs; len(v) > 0 {
		m.RemoveTomeIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
