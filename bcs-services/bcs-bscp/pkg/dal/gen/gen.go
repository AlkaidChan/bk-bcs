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
	App                *app
	ArchivedApp        *archivedApp
	Audit              *audit
	Commit             *commit
	ConfigHook         *configHook
	ConfigItem         *configItem
	Content            *content
	Credential         *credential
	CredentialScope    *credentialScope
	Event              *event
	Group              *group
	GroupAppBind       *groupAppBind
	Hook               *hook
	HookRelease        *hookRelease
	IDGenerator        *iDGenerator
	Release            *release
	ReleasedConfigItem *releasedConfigItem
	ReleasedGroup      *releasedGroup
	ResourceLock       *resourceLock
	Strategy           *strategy
	Template           *template
	TemplateRelease    *templateRelease
	TemplateSet        *templateSet
	TemplateSpace      *templateSpace
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	App = &Q.App
	ArchivedApp = &Q.ArchivedApp
	Audit = &Q.Audit
	Commit = &Q.Commit
	ConfigHook = &Q.ConfigHook
	ConfigItem = &Q.ConfigItem
	Content = &Q.Content
	Credential = &Q.Credential
	CredentialScope = &Q.CredentialScope
	Event = &Q.Event
	Group = &Q.Group
	GroupAppBind = &Q.GroupAppBind
	Hook = &Q.Hook
	HookRelease = &Q.HookRelease
	IDGenerator = &Q.IDGenerator
	Release = &Q.Release
	ReleasedConfigItem = &Q.ReleasedConfigItem
	ReleasedGroup = &Q.ReleasedGroup
	ResourceLock = &Q.ResourceLock
	Strategy = &Q.Strategy
	Template = &Q.Template
	TemplateRelease = &Q.TemplateRelease
	TemplateSet = &Q.TemplateSet
	TemplateSpace = &Q.TemplateSpace
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		App:                newApp(db, opts...),
		ArchivedApp:        newArchivedApp(db, opts...),
		Audit:              newAudit(db, opts...),
		Commit:             newCommit(db, opts...),
		ConfigHook:         newConfigHook(db, opts...),
		ConfigItem:         newConfigItem(db, opts...),
		Content:            newContent(db, opts...),
		Credential:         newCredential(db, opts...),
		CredentialScope:    newCredentialScope(db, opts...),
		Event:              newEvent(db, opts...),
		Group:              newGroup(db, opts...),
		GroupAppBind:       newGroupAppBind(db, opts...),
		Hook:               newHook(db, opts...),
		HookRelease:        newHookRelease(db, opts...),
		IDGenerator:        newIDGenerator(db, opts...),
		Release:            newRelease(db, opts...),
		ReleasedConfigItem: newReleasedConfigItem(db, opts...),
		ReleasedGroup:      newReleasedGroup(db, opts...),
		ResourceLock:       newResourceLock(db, opts...),
		Strategy:           newStrategy(db, opts...),
		Template:           newTemplate(db, opts...),
		TemplateRelease:    newTemplateRelease(db, opts...),
		TemplateSet:        newTemplateSet(db, opts...),
		TemplateSpace:      newTemplateSpace(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	App                app
	ArchivedApp        archivedApp
	Audit              audit
	Commit             commit
	ConfigHook         configHook
	ConfigItem         configItem
	Content            content
	Credential         credential
	CredentialScope    credentialScope
	Event              event
	Group              group
	GroupAppBind       groupAppBind
	Hook               hook
	HookRelease        hookRelease
	IDGenerator        iDGenerator
	Release            release
	ReleasedConfigItem releasedConfigItem
	ReleasedGroup      releasedGroup
	ResourceLock       resourceLock
	Strategy           strategy
	Template           template
	TemplateRelease    templateRelease
	TemplateSet        templateSet
	TemplateSpace      templateSpace
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		App:                q.App.clone(db),
		ArchivedApp:        q.ArchivedApp.clone(db),
		Audit:              q.Audit.clone(db),
		Commit:             q.Commit.clone(db),
		ConfigHook:         q.ConfigHook.clone(db),
		ConfigItem:         q.ConfigItem.clone(db),
		Content:            q.Content.clone(db),
		Credential:         q.Credential.clone(db),
		CredentialScope:    q.CredentialScope.clone(db),
		Event:              q.Event.clone(db),
		Group:              q.Group.clone(db),
		GroupAppBind:       q.GroupAppBind.clone(db),
		Hook:               q.Hook.clone(db),
		HookRelease:        q.HookRelease.clone(db),
		IDGenerator:        q.IDGenerator.clone(db),
		Release:            q.Release.clone(db),
		ReleasedConfigItem: q.ReleasedConfigItem.clone(db),
		ReleasedGroup:      q.ReleasedGroup.clone(db),
		ResourceLock:       q.ResourceLock.clone(db),
		Strategy:           q.Strategy.clone(db),
		Template:           q.Template.clone(db),
		TemplateRelease:    q.TemplateRelease.clone(db),
		TemplateSet:        q.TemplateSet.clone(db),
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
		App:                q.App.replaceDB(db),
		ArchivedApp:        q.ArchivedApp.replaceDB(db),
		Audit:              q.Audit.replaceDB(db),
		Commit:             q.Commit.replaceDB(db),
		ConfigHook:         q.ConfigHook.replaceDB(db),
		ConfigItem:         q.ConfigItem.replaceDB(db),
		Content:            q.Content.replaceDB(db),
		Credential:         q.Credential.replaceDB(db),
		CredentialScope:    q.CredentialScope.replaceDB(db),
		Event:              q.Event.replaceDB(db),
		Group:              q.Group.replaceDB(db),
		GroupAppBind:       q.GroupAppBind.replaceDB(db),
		Hook:               q.Hook.replaceDB(db),
		HookRelease:        q.HookRelease.replaceDB(db),
		IDGenerator:        q.IDGenerator.replaceDB(db),
		Release:            q.Release.replaceDB(db),
		ReleasedConfigItem: q.ReleasedConfigItem.replaceDB(db),
		ReleasedGroup:      q.ReleasedGroup.replaceDB(db),
		ResourceLock:       q.ResourceLock.replaceDB(db),
		Strategy:           q.Strategy.replaceDB(db),
		Template:           q.Template.replaceDB(db),
		TemplateRelease:    q.TemplateRelease.replaceDB(db),
		TemplateSet:        q.TemplateSet.replaceDB(db),
		TemplateSpace:      q.TemplateSpace.replaceDB(db),
	}
}

type queryCtx struct {
	App                IAppDo
	ArchivedApp        IArchivedAppDo
	Audit              IAuditDo
	Commit             ICommitDo
	ConfigHook         IConfigHookDo
	ConfigItem         IConfigItemDo
	Content            IContentDo
	Credential         ICredentialDo
	CredentialScope    ICredentialScopeDo
	Event              IEventDo
	Group              IGroupDo
	GroupAppBind       IGroupAppBindDo
	Hook               IHookDo
	HookRelease        IHookReleaseDo
	IDGenerator        IIDGeneratorDo
	Release            IReleaseDo
	ReleasedConfigItem IReleasedConfigItemDo
	ReleasedGroup      IReleasedGroupDo
	ResourceLock       IResourceLockDo
	Strategy           IStrategyDo
	Template           ITemplateDo
	TemplateRelease    ITemplateReleaseDo
	TemplateSet        ITemplateSetDo
	TemplateSpace      ITemplateSpaceDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		App:                q.App.WithContext(ctx),
		ArchivedApp:        q.ArchivedApp.WithContext(ctx),
		Audit:              q.Audit.WithContext(ctx),
		Commit:             q.Commit.WithContext(ctx),
		ConfigHook:         q.ConfigHook.WithContext(ctx),
		ConfigItem:         q.ConfigItem.WithContext(ctx),
		Content:            q.Content.WithContext(ctx),
		Credential:         q.Credential.WithContext(ctx),
		CredentialScope:    q.CredentialScope.WithContext(ctx),
		Event:              q.Event.WithContext(ctx),
		Group:              q.Group.WithContext(ctx),
		GroupAppBind:       q.GroupAppBind.WithContext(ctx),
		Hook:               q.Hook.WithContext(ctx),
		HookRelease:        q.HookRelease.WithContext(ctx),
		IDGenerator:        q.IDGenerator.WithContext(ctx),
		Release:            q.Release.WithContext(ctx),
		ReleasedConfigItem: q.ReleasedConfigItem.WithContext(ctx),
		ReleasedGroup:      q.ReleasedGroup.WithContext(ctx),
		ResourceLock:       q.ResourceLock.WithContext(ctx),
		Strategy:           q.Strategy.WithContext(ctx),
		Template:           q.Template.WithContext(ctx),
		TemplateRelease:    q.TemplateRelease.WithContext(ctx),
		TemplateSet:        q.TemplateSet.WithContext(ctx),
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
