// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                  = new(Query)
	Audit              *audit
	Commit             *commit
	ConfigItem         *configItem
	Content            *content
	IDGenerator        *iDGenerator
	ReleasedConfigItem *releasedConfigItem
	Template           *template
	TemplateRelease    *templateRelease
	TemplateSpace      *templateSpace
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Audit = &Q.Audit
	Commit = &Q.Commit
	ConfigItem = &Q.ConfigItem
	Content = &Q.Content
	IDGenerator = &Q.IDGenerator
	ReleasedConfigItem = &Q.ReleasedConfigItem
	Template = &Q.Template
	TemplateRelease = &Q.TemplateRelease
	TemplateSpace = &Q.TemplateSpace
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		Audit:              newAudit(db, opts...),
		Commit:             newCommit(db, opts...),
		ConfigItem:         newConfigItem(db, opts...),
		Content:            newContent(db, opts...),
		IDGenerator:        newIDGenerator(db, opts...),
		ReleasedConfigItem: newReleasedConfigItem(db, opts...),
		Template:           newTemplate(db, opts...),
		TemplateRelease:    newTemplateRelease(db, opts...),
		TemplateSpace:      newTemplateSpace(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Audit              audit
	Commit             commit
	ConfigItem         configItem
	Content            content
	IDGenerator        iDGenerator
	ReleasedConfigItem releasedConfigItem
	Template           template
	TemplateRelease    templateRelease
	TemplateSpace      templateSpace
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Audit:              q.Audit.clone(db),
		Commit:             q.Commit.clone(db),
		ConfigItem:         q.ConfigItem.clone(db),
		Content:            q.Content.clone(db),
		IDGenerator:        q.IDGenerator.clone(db),
		ReleasedConfigItem: q.ReleasedConfigItem.clone(db),
		Template:           q.Template.clone(db),
		TemplateRelease:    q.TemplateRelease.clone(db),
		TemplateSpace:      q.TemplateSpace.clone(db),
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
		db:                 db,
		Audit:              q.Audit.replaceDB(db),
		Commit:             q.Commit.replaceDB(db),
		ConfigItem:         q.ConfigItem.replaceDB(db),
		Content:            q.Content.replaceDB(db),
		IDGenerator:        q.IDGenerator.replaceDB(db),
		ReleasedConfigItem: q.ReleasedConfigItem.replaceDB(db),
		Template:           q.Template.replaceDB(db),
		TemplateRelease:    q.TemplateRelease.replaceDB(db),
		TemplateSpace:      q.TemplateSpace.replaceDB(db),
	}
}

type queryCtx struct {
	Audit              IAuditDo
	Commit             ICommitDo
	ConfigItem         IConfigItemDo
	Content            IContentDo
	IDGenerator        IIDGeneratorDo
	ReleasedConfigItem IReleasedConfigItemDo
	Template           ITemplateDo
	TemplateRelease    ITemplateReleaseDo
	TemplateSpace      ITemplateSpaceDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Audit:              q.Audit.WithContext(ctx),
		Commit:             q.Commit.WithContext(ctx),
		ConfigItem:         q.ConfigItem.WithContext(ctx),
		Content:            q.Content.WithContext(ctx),
		IDGenerator:        q.IDGenerator.WithContext(ctx),
		ReleasedConfigItem: q.ReleasedConfigItem.WithContext(ctx),
		Template:           q.Template.WithContext(ctx),
		TemplateRelease:    q.TemplateRelease.WithContext(ctx),
		TemplateSpace:      q.TemplateSpace.WithContext(ctx),
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
