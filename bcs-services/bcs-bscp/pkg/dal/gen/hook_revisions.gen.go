// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bscp.io/pkg/dal/table"
)

func newHookRevision(db *gorm.DB, opts ...gen.DOOption) hookRevision {
	_hookRevision := hookRevision{}

	_hookRevision.hookRevisionDo.UseDB(db, opts...)
	_hookRevision.hookRevisionDo.UseModel(&table.HookRevision{})

	tableName := _hookRevision.hookRevisionDo.TableName()
	_hookRevision.ALL = field.NewAsterisk(tableName)
	_hookRevision.ID = field.NewUint32(tableName, "id")
	_hookRevision.Name = field.NewString(tableName, "name")
	_hookRevision.State = field.NewString(tableName, "state")
	_hookRevision.Content = field.NewString(tableName, "content")
	_hookRevision.Memo = field.NewString(tableName, "memo")
	_hookRevision.BizID = field.NewUint32(tableName, "biz_id")
	_hookRevision.HookID = field.NewUint32(tableName, "hook_id")
	_hookRevision.Creator = field.NewString(tableName, "creator")
	_hookRevision.Reviser = field.NewString(tableName, "reviser")
	_hookRevision.CreatedAt = field.NewTime(tableName, "created_at")
	_hookRevision.UpdatedAt = field.NewTime(tableName, "updated_at")

	_hookRevision.fillFieldMap()

	return _hookRevision
}

type hookRevision struct {
	hookRevisionDo hookRevisionDo

	ALL       field.Asterisk
	ID        field.Uint32
	Name      field.String
	State     field.String
	Content   field.String
	Memo      field.String
	BizID     field.Uint32
	HookID    field.Uint32
	Creator   field.String
	Reviser   field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (h hookRevision) Table(newTableName string) *hookRevision {
	h.hookRevisionDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hookRevision) As(alias string) *hookRevision {
	h.hookRevisionDo.DO = *(h.hookRevisionDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hookRevision) updateTableName(table string) *hookRevision {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewUint32(table, "id")
	h.Name = field.NewString(table, "name")
	h.State = field.NewString(table, "state")
	h.Content = field.NewString(table, "content")
	h.Memo = field.NewString(table, "memo")
	h.BizID = field.NewUint32(table, "biz_id")
	h.HookID = field.NewUint32(table, "hook_id")
	h.Creator = field.NewString(table, "creator")
	h.Reviser = field.NewString(table, "reviser")
	h.CreatedAt = field.NewTime(table, "created_at")
	h.UpdatedAt = field.NewTime(table, "updated_at")

	h.fillFieldMap()

	return h
}

func (h *hookRevision) WithContext(ctx context.Context) IHookRevisionDo {
	return h.hookRevisionDo.WithContext(ctx)
}

func (h hookRevision) TableName() string { return h.hookRevisionDo.TableName() }

func (h hookRevision) Alias() string { return h.hookRevisionDo.Alias() }

func (h hookRevision) Columns(cols ...field.Expr) gen.Columns {
	return h.hookRevisionDo.Columns(cols...)
}

func (h *hookRevision) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hookRevision) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 11)
	h.fieldMap["id"] = h.ID
	h.fieldMap["name"] = h.Name
	h.fieldMap["state"] = h.State
	h.fieldMap["content"] = h.Content
	h.fieldMap["memo"] = h.Memo
	h.fieldMap["biz_id"] = h.BizID
	h.fieldMap["hook_id"] = h.HookID
	h.fieldMap["creator"] = h.Creator
	h.fieldMap["reviser"] = h.Reviser
	h.fieldMap["created_at"] = h.CreatedAt
	h.fieldMap["updated_at"] = h.UpdatedAt
}

func (h hookRevision) clone(db *gorm.DB) hookRevision {
	h.hookRevisionDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hookRevision) replaceDB(db *gorm.DB) hookRevision {
	h.hookRevisionDo.ReplaceDB(db)
	return h
}

type hookRevisionDo struct{ gen.DO }

type IHookRevisionDo interface {
	gen.SubQuery
	Debug() IHookRevisionDo
	WithContext(ctx context.Context) IHookRevisionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHookRevisionDo
	WriteDB() IHookRevisionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHookRevisionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHookRevisionDo
	Not(conds ...gen.Condition) IHookRevisionDo
	Or(conds ...gen.Condition) IHookRevisionDo
	Select(conds ...field.Expr) IHookRevisionDo
	Where(conds ...gen.Condition) IHookRevisionDo
	Order(conds ...field.Expr) IHookRevisionDo
	Distinct(cols ...field.Expr) IHookRevisionDo
	Omit(cols ...field.Expr) IHookRevisionDo
	Join(table schema.Tabler, on ...field.Expr) IHookRevisionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHookRevisionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHookRevisionDo
	Group(cols ...field.Expr) IHookRevisionDo
	Having(conds ...gen.Condition) IHookRevisionDo
	Limit(limit int) IHookRevisionDo
	Offset(offset int) IHookRevisionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHookRevisionDo
	Unscoped() IHookRevisionDo
	Create(values ...*table.HookRevision) error
	CreateInBatches(values []*table.HookRevision, batchSize int) error
	Save(values ...*table.HookRevision) error
	First() (*table.HookRevision, error)
	Take() (*table.HookRevision, error)
	Last() (*table.HookRevision, error)
	Find() ([]*table.HookRevision, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.HookRevision, err error)
	FindInBatches(result *[]*table.HookRevision, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.HookRevision) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHookRevisionDo
	Assign(attrs ...field.AssignExpr) IHookRevisionDo
	Joins(fields ...field.RelationField) IHookRevisionDo
	Preload(fields ...field.RelationField) IHookRevisionDo
	FirstOrInit() (*table.HookRevision, error)
	FirstOrCreate() (*table.HookRevision, error)
	FindByPage(offset int, limit int) (result []*table.HookRevision, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHookRevisionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hookRevisionDo) Debug() IHookRevisionDo {
	return h.withDO(h.DO.Debug())
}

func (h hookRevisionDo) WithContext(ctx context.Context) IHookRevisionDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hookRevisionDo) ReadDB() IHookRevisionDo {
	return h.Clauses(dbresolver.Read)
}

func (h hookRevisionDo) WriteDB() IHookRevisionDo {
	return h.Clauses(dbresolver.Write)
}

func (h hookRevisionDo) Session(config *gorm.Session) IHookRevisionDo {
	return h.withDO(h.DO.Session(config))
}

func (h hookRevisionDo) Clauses(conds ...clause.Expression) IHookRevisionDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hookRevisionDo) Returning(value interface{}, columns ...string) IHookRevisionDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hookRevisionDo) Not(conds ...gen.Condition) IHookRevisionDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hookRevisionDo) Or(conds ...gen.Condition) IHookRevisionDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hookRevisionDo) Select(conds ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hookRevisionDo) Where(conds ...gen.Condition) IHookRevisionDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hookRevisionDo) Order(conds ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hookRevisionDo) Distinct(cols ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hookRevisionDo) Omit(cols ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hookRevisionDo) Join(table schema.Tabler, on ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hookRevisionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hookRevisionDo) RightJoin(table schema.Tabler, on ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hookRevisionDo) Group(cols ...field.Expr) IHookRevisionDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hookRevisionDo) Having(conds ...gen.Condition) IHookRevisionDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hookRevisionDo) Limit(limit int) IHookRevisionDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hookRevisionDo) Offset(offset int) IHookRevisionDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hookRevisionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHookRevisionDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hookRevisionDo) Unscoped() IHookRevisionDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hookRevisionDo) Create(values ...*table.HookRevision) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hookRevisionDo) CreateInBatches(values []*table.HookRevision, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hookRevisionDo) Save(values ...*table.HookRevision) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hookRevisionDo) First() (*table.HookRevision, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.HookRevision), nil
	}
}

func (h hookRevisionDo) Take() (*table.HookRevision, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.HookRevision), nil
	}
}

func (h hookRevisionDo) Last() (*table.HookRevision, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.HookRevision), nil
	}
}

func (h hookRevisionDo) Find() ([]*table.HookRevision, error) {
	result, err := h.DO.Find()
	return result.([]*table.HookRevision), err
}

func (h hookRevisionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.HookRevision, err error) {
	buf := make([]*table.HookRevision, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hookRevisionDo) FindInBatches(result *[]*table.HookRevision, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hookRevisionDo) Attrs(attrs ...field.AssignExpr) IHookRevisionDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hookRevisionDo) Assign(attrs ...field.AssignExpr) IHookRevisionDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hookRevisionDo) Joins(fields ...field.RelationField) IHookRevisionDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hookRevisionDo) Preload(fields ...field.RelationField) IHookRevisionDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hookRevisionDo) FirstOrInit() (*table.HookRevision, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.HookRevision), nil
	}
}

func (h hookRevisionDo) FirstOrCreate() (*table.HookRevision, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.HookRevision), nil
	}
}

func (h hookRevisionDo) FindByPage(offset int, limit int) (result []*table.HookRevision, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hookRevisionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hookRevisionDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hookRevisionDo) Delete(models ...*table.HookRevision) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hookRevisionDo) withDO(do gen.Dao) *hookRevisionDo {
	h.DO = *do.(*gen.DO)
	return h
}
