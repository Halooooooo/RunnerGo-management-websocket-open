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

	"RunnerGo-management/internal/pkg/dal/model"
)

func newStressPlanTimedTaskConf(db *gorm.DB, opts ...gen.DOOption) stressPlanTimedTaskConf {
	_stressPlanTimedTaskConf := stressPlanTimedTaskConf{}

	_stressPlanTimedTaskConf.stressPlanTimedTaskConfDo.UseDB(db, opts...)
	_stressPlanTimedTaskConf.stressPlanTimedTaskConfDo.UseModel(&model.StressPlanTimedTaskConf{})

	tableName := _stressPlanTimedTaskConf.stressPlanTimedTaskConfDo.TableName()
	_stressPlanTimedTaskConf.ALL = field.NewAsterisk(tableName)
	_stressPlanTimedTaskConf.ID = field.NewInt32(tableName, "id")
	_stressPlanTimedTaskConf.PlanID = field.NewString(tableName, "plan_id")
	_stressPlanTimedTaskConf.SceneID = field.NewString(tableName, "scene_id")
	_stressPlanTimedTaskConf.TeamID = field.NewString(tableName, "team_id")
	_stressPlanTimedTaskConf.UserID = field.NewString(tableName, "user_id")
	_stressPlanTimedTaskConf.Frequency = field.NewInt32(tableName, "frequency")
	_stressPlanTimedTaskConf.TaskExecTime = field.NewInt64(tableName, "task_exec_time")
	_stressPlanTimedTaskConf.TaskCloseTime = field.NewInt64(tableName, "task_close_time")
	_stressPlanTimedTaskConf.TaskType = field.NewInt32(tableName, "task_type")
	_stressPlanTimedTaskConf.TaskMode = field.NewInt32(tableName, "task_mode")
	_stressPlanTimedTaskConf.ControlMode = field.NewInt32(tableName, "control_mode")
	_stressPlanTimedTaskConf.DebugMode = field.NewString(tableName, "debug_mode")
	_stressPlanTimedTaskConf.ModeConf = field.NewString(tableName, "mode_conf")
	_stressPlanTimedTaskConf.RunUserID = field.NewString(tableName, "run_user_id")
	_stressPlanTimedTaskConf.Status = field.NewInt32(tableName, "status")
	_stressPlanTimedTaskConf.CreatedAt = field.NewTime(tableName, "created_at")
	_stressPlanTimedTaskConf.UpdatedAt = field.NewTime(tableName, "updated_at")
	_stressPlanTimedTaskConf.DeletedAt = field.NewField(tableName, "deleted_at")

	_stressPlanTimedTaskConf.fillFieldMap()

	return _stressPlanTimedTaskConf
}

type stressPlanTimedTaskConf struct {
	stressPlanTimedTaskConfDo stressPlanTimedTaskConfDo

	ALL           field.Asterisk
	ID            field.Int32  // 表id
	PlanID        field.String // 计划id
	SceneID       field.String // 场景id
	TeamID        field.String // 团队id
	UserID        field.String // 用户ID
	Frequency     field.Int32  // 任务执行频次: 0-一次，1-每天，2-每周，3-每月
	TaskExecTime  field.Int64  // 任务执行时间
	TaskCloseTime field.Int64  // 任务结束时间
	TaskType      field.Int32  // 任务类型：1-普通任务，2-定时任务
	TaskMode      field.Int32  // 压测模式：1-并发模式，2-阶梯模式，3-错误率模式，4-响应时间模式，5-每秒请求数模式，6 -每秒事务数模式
	ControlMode   field.Int32  // 控制模式：0-集中模式，1-单独模式
	DebugMode     field.String // debug模式：stop-关闭，all-开启全部日志，only_success-开启仅成功日志，only_error-开启仅错误日志
	ModeConf      field.String // 压测详细配置
	RunUserID     field.String // 运行人ID
	Status        field.Int32  // 任务状态：0-未启用，1-运行中，2-已过期
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s stressPlanTimedTaskConf) Table(newTableName string) *stressPlanTimedTaskConf {
	s.stressPlanTimedTaskConfDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stressPlanTimedTaskConf) As(alias string) *stressPlanTimedTaskConf {
	s.stressPlanTimedTaskConfDo.DO = *(s.stressPlanTimedTaskConfDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stressPlanTimedTaskConf) updateTableName(table string) *stressPlanTimedTaskConf {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.PlanID = field.NewString(table, "plan_id")
	s.SceneID = field.NewString(table, "scene_id")
	s.TeamID = field.NewString(table, "team_id")
	s.UserID = field.NewString(table, "user_id")
	s.Frequency = field.NewInt32(table, "frequency")
	s.TaskExecTime = field.NewInt64(table, "task_exec_time")
	s.TaskCloseTime = field.NewInt64(table, "task_close_time")
	s.TaskType = field.NewInt32(table, "task_type")
	s.TaskMode = field.NewInt32(table, "task_mode")
	s.ControlMode = field.NewInt32(table, "control_mode")
	s.DebugMode = field.NewString(table, "debug_mode")
	s.ModeConf = field.NewString(table, "mode_conf")
	s.RunUserID = field.NewString(table, "run_user_id")
	s.Status = field.NewInt32(table, "status")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *stressPlanTimedTaskConf) WithContext(ctx context.Context) *stressPlanTimedTaskConfDo {
	return s.stressPlanTimedTaskConfDo.WithContext(ctx)
}

func (s stressPlanTimedTaskConf) TableName() string { return s.stressPlanTimedTaskConfDo.TableName() }

func (s stressPlanTimedTaskConf) Alias() string { return s.stressPlanTimedTaskConfDo.Alias() }

func (s *stressPlanTimedTaskConf) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stressPlanTimedTaskConf) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 18)
	s.fieldMap["id"] = s.ID
	s.fieldMap["plan_id"] = s.PlanID
	s.fieldMap["scene_id"] = s.SceneID
	s.fieldMap["team_id"] = s.TeamID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["frequency"] = s.Frequency
	s.fieldMap["task_exec_time"] = s.TaskExecTime
	s.fieldMap["task_close_time"] = s.TaskCloseTime
	s.fieldMap["task_type"] = s.TaskType
	s.fieldMap["task_mode"] = s.TaskMode
	s.fieldMap["control_mode"] = s.ControlMode
	s.fieldMap["debug_mode"] = s.DebugMode
	s.fieldMap["mode_conf"] = s.ModeConf
	s.fieldMap["run_user_id"] = s.RunUserID
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s stressPlanTimedTaskConf) clone(db *gorm.DB) stressPlanTimedTaskConf {
	s.stressPlanTimedTaskConfDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stressPlanTimedTaskConf) replaceDB(db *gorm.DB) stressPlanTimedTaskConf {
	s.stressPlanTimedTaskConfDo.ReplaceDB(db)
	return s
}

type stressPlanTimedTaskConfDo struct{ gen.DO }

func (s stressPlanTimedTaskConfDo) Debug() *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Debug())
}

func (s stressPlanTimedTaskConfDo) WithContext(ctx context.Context) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stressPlanTimedTaskConfDo) ReadDB() *stressPlanTimedTaskConfDo {
	return s.Clauses(dbresolver.Read)
}

func (s stressPlanTimedTaskConfDo) WriteDB() *stressPlanTimedTaskConfDo {
	return s.Clauses(dbresolver.Write)
}

func (s stressPlanTimedTaskConfDo) Session(config *gorm.Session) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Session(config))
}

func (s stressPlanTimedTaskConfDo) Clauses(conds ...clause.Expression) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stressPlanTimedTaskConfDo) Returning(value interface{}, columns ...string) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stressPlanTimedTaskConfDo) Not(conds ...gen.Condition) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stressPlanTimedTaskConfDo) Or(conds ...gen.Condition) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stressPlanTimedTaskConfDo) Select(conds ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stressPlanTimedTaskConfDo) Where(conds ...gen.Condition) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stressPlanTimedTaskConfDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *stressPlanTimedTaskConfDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s stressPlanTimedTaskConfDo) Order(conds ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stressPlanTimedTaskConfDo) Distinct(cols ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stressPlanTimedTaskConfDo) Omit(cols ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stressPlanTimedTaskConfDo) Join(table schema.Tabler, on ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stressPlanTimedTaskConfDo) LeftJoin(table schema.Tabler, on ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stressPlanTimedTaskConfDo) RightJoin(table schema.Tabler, on ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stressPlanTimedTaskConfDo) Group(cols ...field.Expr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stressPlanTimedTaskConfDo) Having(conds ...gen.Condition) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stressPlanTimedTaskConfDo) Limit(limit int) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stressPlanTimedTaskConfDo) Offset(offset int) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stressPlanTimedTaskConfDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stressPlanTimedTaskConfDo) Unscoped() *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stressPlanTimedTaskConfDo) Create(values ...*model.StressPlanTimedTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stressPlanTimedTaskConfDo) CreateInBatches(values []*model.StressPlanTimedTaskConf, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stressPlanTimedTaskConfDo) Save(values ...*model.StressPlanTimedTaskConf) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stressPlanTimedTaskConfDo) First() (*model.StressPlanTimedTaskConf, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTimedTaskConf), nil
	}
}

func (s stressPlanTimedTaskConfDo) Take() (*model.StressPlanTimedTaskConf, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTimedTaskConf), nil
	}
}

func (s stressPlanTimedTaskConfDo) Last() (*model.StressPlanTimedTaskConf, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTimedTaskConf), nil
	}
}

func (s stressPlanTimedTaskConfDo) Find() ([]*model.StressPlanTimedTaskConf, error) {
	result, err := s.DO.Find()
	return result.([]*model.StressPlanTimedTaskConf), err
}

func (s stressPlanTimedTaskConfDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StressPlanTimedTaskConf, err error) {
	buf := make([]*model.StressPlanTimedTaskConf, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stressPlanTimedTaskConfDo) FindInBatches(result *[]*model.StressPlanTimedTaskConf, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stressPlanTimedTaskConfDo) Attrs(attrs ...field.AssignExpr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stressPlanTimedTaskConfDo) Assign(attrs ...field.AssignExpr) *stressPlanTimedTaskConfDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stressPlanTimedTaskConfDo) Joins(fields ...field.RelationField) *stressPlanTimedTaskConfDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stressPlanTimedTaskConfDo) Preload(fields ...field.RelationField) *stressPlanTimedTaskConfDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stressPlanTimedTaskConfDo) FirstOrInit() (*model.StressPlanTimedTaskConf, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTimedTaskConf), nil
	}
}

func (s stressPlanTimedTaskConfDo) FirstOrCreate() (*model.StressPlanTimedTaskConf, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanTimedTaskConf), nil
	}
}

func (s stressPlanTimedTaskConfDo) FindByPage(offset int, limit int) (result []*model.StressPlanTimedTaskConf, count int64, err error) {
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

func (s stressPlanTimedTaskConfDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stressPlanTimedTaskConfDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stressPlanTimedTaskConfDo) Delete(models ...*model.StressPlanTimedTaskConf) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stressPlanTimedTaskConfDo) withDO(do gen.Dao) *stressPlanTimedTaskConfDo {
	s.DO = *do.(*gen.DO)
	return s
}
