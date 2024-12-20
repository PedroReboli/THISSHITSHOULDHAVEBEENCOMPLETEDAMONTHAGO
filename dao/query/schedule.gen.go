// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"server/dao/model"
)

func newSchedule(db *gorm.DB, opts ...gen.DOOption) schedule {
	_schedule := schedule{}

	_schedule.scheduleDo.UseDB(db, opts...)
	_schedule.scheduleDo.UseModel(&model.Schedule{})

	tableName := _schedule.scheduleDo.TableName()
	_schedule.ALL = field.NewAsterisk(tableName)
	_schedule.ID = field.NewInt32(tableName, "id")
	_schedule.Eventid = field.NewInt32(tableName, "eventid")
	_schedule.Managerid = field.NewInt32(tableName, "managerid")
	_schedule.Volunteerid = field.NewInt32(tableName, "volunteerid")
	_schedule.Startdate = field.NewTime(tableName, "startdate")
	_schedule.Enddate = field.NewTime(tableName, "enddate")
	_schedule.Weakly = field.NewBool(tableName, "weakly")

	_schedule.fillFieldMap()

	return _schedule
}

type schedule struct {
	scheduleDo scheduleDo

	ALL         field.Asterisk
	ID          field.Int32
	Eventid     field.Int32
	Managerid   field.Int32
	Volunteerid field.Int32
	Startdate   field.Time
	Enddate     field.Time
	Weakly      field.Bool

	fieldMap map[string]field.Expr
}

func (s schedule) Table(newTableName string) *schedule {
	s.scheduleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s schedule) As(alias string) *schedule {
	s.scheduleDo.DO = *(s.scheduleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *schedule) updateTableName(table string) *schedule {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Eventid = field.NewInt32(table, "eventid")
	s.Managerid = field.NewInt32(table, "managerid")
	s.Volunteerid = field.NewInt32(table, "volunteerid")
	s.Startdate = field.NewTime(table, "startdate")
	s.Enddate = field.NewTime(table, "enddate")
	s.Weakly = field.NewBool(table, "weakly")

	s.fillFieldMap()

	return s
}

func (s *schedule) WithContext(ctx context.Context) *scheduleDo { return s.scheduleDo.WithContext(ctx) }

func (s schedule) TableName() string { return s.scheduleDo.TableName() }

func (s schedule) Alias() string { return s.scheduleDo.Alias() }

func (s schedule) Columns(cols ...field.Expr) gen.Columns { return s.scheduleDo.Columns(cols...) }

func (s *schedule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *schedule) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["eventid"] = s.Eventid
	s.fieldMap["managerid"] = s.Managerid
	s.fieldMap["volunteerid"] = s.Volunteerid
	s.fieldMap["startdate"] = s.Startdate
	s.fieldMap["enddate"] = s.Enddate
	s.fieldMap["weakly"] = s.Weakly
}

func (s schedule) clone(db *gorm.DB) schedule {
	s.scheduleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s schedule) replaceDB(db *gorm.DB) schedule {
	s.scheduleDo.ReplaceDB(db)
	return s
}

type scheduleDo struct{ gen.DO }

func (s scheduleDo) Debug() *scheduleDo {
	return s.withDO(s.DO.Debug())
}

func (s scheduleDo) WithContext(ctx context.Context) *scheduleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s scheduleDo) ReadDB() *scheduleDo {
	return s.Clauses(dbresolver.Read)
}

func (s scheduleDo) WriteDB() *scheduleDo {
	return s.Clauses(dbresolver.Write)
}

func (s scheduleDo) Session(config *gorm.Session) *scheduleDo {
	return s.withDO(s.DO.Session(config))
}

func (s scheduleDo) Clauses(conds ...clause.Expression) *scheduleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s scheduleDo) Returning(value interface{}, columns ...string) *scheduleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s scheduleDo) Not(conds ...gen.Condition) *scheduleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s scheduleDo) Or(conds ...gen.Condition) *scheduleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s scheduleDo) Select(conds ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s scheduleDo) Where(conds ...gen.Condition) *scheduleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s scheduleDo) Order(conds ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s scheduleDo) Distinct(cols ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s scheduleDo) Omit(cols ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s scheduleDo) Join(table schema.Tabler, on ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s scheduleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s scheduleDo) RightJoin(table schema.Tabler, on ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s scheduleDo) Group(cols ...field.Expr) *scheduleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s scheduleDo) Having(conds ...gen.Condition) *scheduleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s scheduleDo) Limit(limit int) *scheduleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s scheduleDo) Offset(offset int) *scheduleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s scheduleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *scheduleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s scheduleDo) Unscoped() *scheduleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s scheduleDo) Create(values ...*model.Schedule) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s scheduleDo) CreateInBatches(values []*model.Schedule, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s scheduleDo) Save(values ...*model.Schedule) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s scheduleDo) First() (*model.Schedule, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Schedule), nil
	}
}

func (s scheduleDo) Take() (*model.Schedule, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Schedule), nil
	}
}

func (s scheduleDo) Last() (*model.Schedule, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Schedule), nil
	}
}

func (s scheduleDo) Find() ([]*model.Schedule, error) {
	result, err := s.DO.Find()
	return result.([]*model.Schedule), err
}

func (s scheduleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Schedule, err error) {
	buf := make([]*model.Schedule, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s scheduleDo) FindInBatches(result *[]*model.Schedule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s scheduleDo) Attrs(attrs ...field.AssignExpr) *scheduleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s scheduleDo) Assign(attrs ...field.AssignExpr) *scheduleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s scheduleDo) Joins(fields ...field.RelationField) *scheduleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s scheduleDo) Preload(fields ...field.RelationField) *scheduleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s scheduleDo) FirstOrInit() (*model.Schedule, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Schedule), nil
	}
}

func (s scheduleDo) FirstOrCreate() (*model.Schedule, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Schedule), nil
	}
}

func (s scheduleDo) FindByPage(offset int, limit int) (result []*model.Schedule, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s scheduleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s scheduleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s scheduleDo) Delete(models ...*model.Schedule) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *scheduleDo) withDO(do gen.Dao) *scheduleDo {
	s.DO = *do.(*gen.DO)
	return s
}
