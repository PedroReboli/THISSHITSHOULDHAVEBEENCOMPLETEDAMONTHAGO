// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:           db,
		Architect:    newArchitect(db, opts...),
		Event:        newEvent(db, opts...),
		Eventmanager: newEventmanager(db, opts...),
		Manager:      newManager(db, opts...),
		Schedule:     newSchedule(db, opts...),
		User:         newUser(db, opts...),
		Volunteer:    newVolunteer(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Architect    architect
	Event        event
	Eventmanager eventmanager
	Manager      manager
	Schedule     schedule
	User         user
	Volunteer    volunteer
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		Architect:    q.Architect.clone(db),
		Event:        q.Event.clone(db),
		Eventmanager: q.Eventmanager.clone(db),
		Manager:      q.Manager.clone(db),
		Schedule:     q.Schedule.clone(db),
		User:         q.User.clone(db),
		Volunteer:    q.Volunteer.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		Architect:    q.Architect.replaceDB(db),
		Event:        q.Event.replaceDB(db),
		Eventmanager: q.Eventmanager.replaceDB(db),
		Manager:      q.Manager.replaceDB(db),
		Schedule:     q.Schedule.replaceDB(db),
		User:         q.User.replaceDB(db),
		Volunteer:    q.Volunteer.replaceDB(db),
	}
}

type queryCtx struct {
	Architect    *architectDo
	Event        *eventDo
	Eventmanager *eventmanagerDo
	Manager      *managerDo
	Schedule     *scheduleDo
	User         *userDo
	Volunteer    *volunteerDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Architect:    q.Architect.WithContext(ctx),
		Event:        q.Event.WithContext(ctx),
		Eventmanager: q.Eventmanager.WithContext(ctx),
		Manager:      q.Manager.WithContext(ctx),
		Schedule:     q.Schedule.WithContext(ctx),
		User:         q.User.WithContext(ctx),
		Volunteer:    q.Volunteer.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}