// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (b *Beacon) Host(ctx context.Context) (*Host, error) {
	result, err := b.Edges.HostOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryHost().Only(ctx)
	}
	return result, err
}

func (b *Beacon) Tasks(ctx context.Context) (result []*Task, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = b.NamedTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = b.Edges.TasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = b.QueryTasks().All(ctx)
	}
	return result, err
}

func (f *File) Tomes(ctx context.Context) (result []*Tome, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = f.NamedTomes(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = f.Edges.TomesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = f.QueryTomes().All(ctx)
	}
	return result, err
}

func (h *Host) Tags(ctx context.Context) (result []*Tag, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = h.NamedTags(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = h.Edges.TagsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = h.QueryTags().All(ctx)
	}
	return result, err
}

func (h *Host) Beacons(ctx context.Context) (result []*Beacon, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = h.NamedBeacons(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = h.Edges.BeaconsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = h.QueryBeacons().All(ctx)
	}
	return result, err
}

func (h *Host) Processes(ctx context.Context) (result []*Process, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = h.NamedProcesses(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = h.Edges.ProcessesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = h.QueryProcesses().All(ctx)
	}
	return result, err
}

func (pr *Process) Host(ctx context.Context) (*Host, error) {
	result, err := pr.Edges.HostOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryHost().Only(ctx)
	}
	return result, err
}

func (pr *Process) Task(ctx context.Context) (*Task, error) {
	result, err := pr.Edges.TaskOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryTask().Only(ctx)
	}
	return result, err
}

func (q *Quest) Tome(ctx context.Context) (*Tome, error) {
	result, err := q.Edges.TomeOrErr()
	if IsNotLoaded(err) {
		result, err = q.QueryTome().Only(ctx)
	}
	return result, err
}

func (q *Quest) Bundle(ctx context.Context) (*File, error) {
	result, err := q.Edges.BundleOrErr()
	if IsNotLoaded(err) {
		result, err = q.QueryBundle().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (q *Quest) Tasks(ctx context.Context) (result []*Task, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = q.NamedTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = q.Edges.TasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = q.QueryTasks().All(ctx)
	}
	return result, err
}

func (q *Quest) Creator(ctx context.Context) (*User, error) {
	result, err := q.Edges.CreatorOrErr()
	if IsNotLoaded(err) {
		result, err = q.QueryCreator().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Tag) Hosts(ctx context.Context) (result []*Host, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedHosts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.HostsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryHosts().All(ctx)
	}
	return result, err
}

func (t *Task) Quest(ctx context.Context) (*Quest, error) {
	result, err := t.Edges.QuestOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryQuest().Only(ctx)
	}
	return result, err
}

func (t *Task) Beacon(ctx context.Context) (*Beacon, error) {
	result, err := t.Edges.BeaconOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryBeacon().Only(ctx)
	}
	return result, err
}

func (t *Task) ReportedProcesses(ctx context.Context) (result []*Process, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedReportedProcesses(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ReportedProcessesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryReportedProcesses().All(ctx)
	}
	return result, err
}

func (t *Tome) Files(ctx context.Context) (result []*File, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedFiles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.FilesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryFiles().All(ctx)
	}
	return result, err
}

func (t *Tome) Uploader(ctx context.Context) (*User, error) {
	result, err := t.Edges.UploaderOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryUploader().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Tomes(ctx context.Context) (result []*Tome, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTomes(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TomesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTomes().All(ctx)
	}
	return result, err
}
