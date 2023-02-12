// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/41197-yhkt/tiktok/composite/gen/dal/model"
)

func newUserFavorite(db *gorm.DB, opts ...gen.DOOption) userFavorite {
	_userFavorite := userFavorite{}

	_userFavorite.userFavoriteDo.UseDB(db, opts...)
	_userFavorite.userFavoriteDo.UseModel(&model.UserFavorite{})

	tableName := _userFavorite.userFavoriteDo.TableName()
	_userFavorite.ALL = field.NewAsterisk(tableName)
	_userFavorite.ID = field.NewUint(tableName, "id")
	_userFavorite.CreatedAt = field.NewTime(tableName, "created_at")
	_userFavorite.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userFavorite.DeletedAt = field.NewField(tableName, "deleted_at")
	_userFavorite.UserId = field.NewInt64(tableName, "user_id")
	_userFavorite.VideoId = field.NewInt64(tableName, "video_id")

	_userFavorite.fillFieldMap()

	return _userFavorite
}

type userFavorite struct {
	userFavoriteDo userFavoriteDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserId    field.Int64
	VideoId   field.Int64

	fieldMap map[string]field.Expr
}

func (u userFavorite) Table(newTableName string) *userFavorite {
	u.userFavoriteDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userFavorite) As(alias string) *userFavorite {
	u.userFavoriteDo.DO = *(u.userFavoriteDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userFavorite) updateTableName(table string) *userFavorite {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UserId = field.NewInt64(table, "user_id")
	u.VideoId = field.NewInt64(table, "video_id")

	u.fillFieldMap()

	return u
}

func (u *userFavorite) WithContext(ctx context.Context) *userFavoriteDo {
	return u.userFavoriteDo.WithContext(ctx)
}

func (u userFavorite) TableName() string { return u.userFavoriteDo.TableName() }

func (u userFavorite) Alias() string { return u.userFavoriteDo.Alias() }

func (u *userFavorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userFavorite) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["user_id"] = u.UserId
	u.fieldMap["video_id"] = u.VideoId
}

func (u userFavorite) clone(db *gorm.DB) userFavorite {
	u.userFavoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userFavorite) replaceDB(db *gorm.DB) userFavorite {
	u.userFavoriteDo.ReplaceDB(db)
	return u
}

type userFavoriteDo struct{ gen.DO }

// sql(select video_id from @@table where user_id = @userId)
func (u userFavoriteDo) FindByUserid(userId int64) (result []*model.UserFavorite, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userId)
	generateSQL.WriteString("select video_id from user_favorites where user_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select user_id from @@table where video_id = @videoId)
func (u userFavoriteDo) FindByVideoid(videoId int64) (result []*model.UserFavorite, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, videoId)
	generateSQL.WriteString("select user_id from user_favorites where video_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select * from @@table where video_id = @videoId and user_id = @userId)
func (u userFavoriteDo) FindByUseridAndVideoid(userId int64, videoId int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, videoId)
	params = append(params, userId)
	generateSQL.WriteString("select * from user_favorites where video_id = ? and user_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select count(*) from @@table where video_id = @videoId)
func (u userFavoriteDo) CountByVideoid(videoId int64) (result int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, videoId)
	generateSQL.WriteString("select count(*) from user_favorites where video_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(delete from @@table where user_id = @userId and video_id = @videoId)
func (u userFavoriteDo) DeleteByUseridAndVideoid(userId int64, videoId int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, userId)
	params = append(params, videoId)
	generateSQL.WriteString("delete from user_favorites where user_id = ? and video_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (u userFavoriteDo) Debug() *userFavoriteDo {
	return u.withDO(u.DO.Debug())
}

func (u userFavoriteDo) WithContext(ctx context.Context) *userFavoriteDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userFavoriteDo) ReadDB() *userFavoriteDo {
	return u.Clauses(dbresolver.Read)
}

func (u userFavoriteDo) WriteDB() *userFavoriteDo {
	return u.Clauses(dbresolver.Write)
}

func (u userFavoriteDo) Session(config *gorm.Session) *userFavoriteDo {
	return u.withDO(u.DO.Session(config))
}

func (u userFavoriteDo) Clauses(conds ...clause.Expression) *userFavoriteDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userFavoriteDo) Returning(value interface{}, columns ...string) *userFavoriteDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userFavoriteDo) Not(conds ...gen.Condition) *userFavoriteDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userFavoriteDo) Or(conds ...gen.Condition) *userFavoriteDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userFavoriteDo) Select(conds ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userFavoriteDo) Where(conds ...gen.Condition) *userFavoriteDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userFavoriteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userFavoriteDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userFavoriteDo) Order(conds ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userFavoriteDo) Distinct(cols ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userFavoriteDo) Omit(cols ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userFavoriteDo) Join(table schema.Tabler, on ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userFavoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userFavoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userFavoriteDo) Group(cols ...field.Expr) *userFavoriteDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userFavoriteDo) Having(conds ...gen.Condition) *userFavoriteDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userFavoriteDo) Limit(limit int) *userFavoriteDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userFavoriteDo) Offset(offset int) *userFavoriteDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userFavoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userFavoriteDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userFavoriteDo) Unscoped() *userFavoriteDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userFavoriteDo) Create(values ...*model.UserFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userFavoriteDo) CreateInBatches(values []*model.UserFavorite, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userFavoriteDo) Save(values ...*model.UserFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userFavoriteDo) First() (*model.UserFavorite, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) Take() (*model.UserFavorite, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) Last() (*model.UserFavorite, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) Find() ([]*model.UserFavorite, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserFavorite), err
}

func (u userFavoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFavorite, err error) {
	buf := make([]*model.UserFavorite, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userFavoriteDo) FindInBatches(result *[]*model.UserFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userFavoriteDo) Attrs(attrs ...field.AssignExpr) *userFavoriteDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userFavoriteDo) Assign(attrs ...field.AssignExpr) *userFavoriteDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userFavoriteDo) Joins(fields ...field.RelationField) *userFavoriteDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userFavoriteDo) Preload(fields ...field.RelationField) *userFavoriteDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userFavoriteDo) FirstOrInit() (*model.UserFavorite, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) FirstOrCreate() (*model.UserFavorite, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorite), nil
	}
}

func (u userFavoriteDo) FindByPage(offset int, limit int) (result []*model.UserFavorite, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userFavoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userFavoriteDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userFavoriteDo) Delete(models ...*model.UserFavorite) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userFavoriteDo) withDO(do gen.Dao) *userFavoriteDo {
	u.DO = *do.(*gen.DO)
	return u
}
