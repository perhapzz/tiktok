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

	"github.com/41197-yhkt/tiktok/video/gen/dal/model"
)

func newVideo(db *gorm.DB, opts ...gen.DOOption) video {
	_video := video{}

	_video.videoDo.UseDB(db, opts...)
	_video.videoDo.UseModel(&model.Video{})

	tableName := _video.videoDo.TableName()
	_video.ALL = field.NewAsterisk(tableName)
	_video.ID = field.NewUint(tableName, "id")
	_video.CreatedAt = field.NewTime(tableName, "created_at")
	_video.UpdatedAt = field.NewTime(tableName, "updated_at")
	_video.DeletedAt = field.NewField(tableName, "deleted_at")
	_video.AuthorId = field.NewInt64(tableName, "author")
	_video.PlayUrl = field.NewString(tableName, "play_url")
	_video.CoverUrl = field.NewString(tableName, "cover_url")
	_video.FavoriteCount = field.NewString(tableName, "favorite_count")
	_video.CommentCount = field.NewString(tableName, "comment_count")
	_video.Title = field.NewString(tableName, "title")

	_video.fillFieldMap()

	return _video
}

type video struct {
	videoDo videoDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	AuthorId      field.Int64
	PlayUrl       field.String
	CoverUrl      field.String
	FavoriteCount field.String
	CommentCount  field.String
	Title         field.String

	fieldMap map[string]field.Expr
}

func (v video) Table(newTableName string) *video {
	v.videoDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v video) As(alias string) *video {
	v.videoDo.DO = *(v.videoDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *video) updateTableName(table string) *video {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewUint(table, "id")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")
	v.AuthorId = field.NewInt64(table, "author")
	v.PlayUrl = field.NewString(table, "play_url")
	v.CoverUrl = field.NewString(table, "cover_url")
	v.FavoriteCount = field.NewString(table, "favorite_count")
	v.CommentCount = field.NewString(table, "comment_count")
	v.Title = field.NewString(table, "title")

	v.fillFieldMap()

	return v
}

func (v *video) WithContext(ctx context.Context) *videoDo { return v.videoDo.WithContext(ctx) }

func (v video) TableName() string { return v.videoDo.TableName() }

func (v video) Alias() string { return v.videoDo.Alias() }

func (v *video) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *video) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 10)
	v.fieldMap["id"] = v.ID
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
	v.fieldMap["author"] = v.AuthorId
	v.fieldMap["play_url"] = v.PlayUrl
	v.fieldMap["cover_url"] = v.CoverUrl
	v.fieldMap["favorite_count"] = v.FavoriteCount
	v.fieldMap["comment_count"] = v.CommentCount
	v.fieldMap["title"] = v.Title
}

func (v video) clone(db *gorm.DB) video {
	v.videoDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v video) replaceDB(db *gorm.DB) video {
	v.videoDo.ReplaceDB(db)
	return v
}

type videoDo struct{ gen.DO }

// where(id=@id)
func (v videoDo) FindByID(id int64) (result model.Video, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select * from @@table where AuthorId = @Authorid)
func (v videoDo) FindByAuthorId(Authorid int) (result []*model.Video, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, Authorid)
	generateSQL.WriteString("select * from videos where AuthorId = ? ")

	var executeSQL *gorm.DB
	executeSQL = v.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (v videoDo) Debug() *videoDo {
	return v.withDO(v.DO.Debug())
}

func (v videoDo) WithContext(ctx context.Context) *videoDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoDo) ReadDB() *videoDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoDo) WriteDB() *videoDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoDo) Session(config *gorm.Session) *videoDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoDo) Clauses(conds ...clause.Expression) *videoDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoDo) Returning(value interface{}, columns ...string) *videoDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoDo) Not(conds ...gen.Condition) *videoDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoDo) Or(conds ...gen.Condition) *videoDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoDo) Select(conds ...field.Expr) *videoDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoDo) Where(conds ...gen.Condition) *videoDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *videoDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v videoDo) Order(conds ...field.Expr) *videoDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoDo) Distinct(cols ...field.Expr) *videoDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoDo) Omit(cols ...field.Expr) *videoDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoDo) Join(table schema.Tabler, on ...field.Expr) *videoDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *videoDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoDo) RightJoin(table schema.Tabler, on ...field.Expr) *videoDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoDo) Group(cols ...field.Expr) *videoDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoDo) Having(conds ...gen.Condition) *videoDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoDo) Limit(limit int) *videoDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoDo) Offset(offset int) *videoDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *videoDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoDo) Unscoped() *videoDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoDo) Create(values ...*model.Video) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoDo) CreateInBatches(values []*model.Video, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoDo) Save(values ...*model.Video) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoDo) First() (*model.Video, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) Take() (*model.Video, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) Last() (*model.Video, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) Find() ([]*model.Video, error) {
	result, err := v.DO.Find()
	return result.([]*model.Video), err
}

func (v videoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Video, err error) {
	buf := make([]*model.Video, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoDo) FindInBatches(result *[]*model.Video, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoDo) Attrs(attrs ...field.AssignExpr) *videoDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoDo) Assign(attrs ...field.AssignExpr) *videoDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoDo) Joins(fields ...field.RelationField) *videoDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoDo) Preload(fields ...field.RelationField) *videoDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoDo) FirstOrInit() (*model.Video, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) FirstOrCreate() (*model.Video, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Video), nil
	}
}

func (v videoDo) FindByPage(offset int, limit int) (result []*model.Video, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v videoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoDo) Delete(models ...*model.Video) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoDo) withDO(do gen.Dao) *videoDo {
	v.DO = *do.(*gen.DO)
	return v
}