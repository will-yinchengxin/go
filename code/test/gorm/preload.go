package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:13306)/will?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			// NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})
	if err != nil {
		return
	}
	DB = db
}

func PreLoad() {
	var u User
	// 一对多
	DB.Table("user").Preload("Company").Preload("Test").Where("id = ?", 1).Find(&u)
	fmt.Println(u)

	//DB.Table("user").Scopes(Paginate(&Page{1, 10})).Find(&u)
}


type Page struct {
	Page     int `form:"page" json:"page" validate:"omitempty,gt=0" label:"分页"`
	PageSize int `form:"pageSize" json:"pageSize" validate:"omitempty,gt=0" label:"每页条数"`
}

func Paginate(r *Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page == 0 {
			r.Page = 1
		}

		switch {
		case r.PageSize > 100:
			r.PageSize = 100
		case r.PageSize <= 0:
			r.PageSize = 10
		}

		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}

/*
// debug 打印 sql 语句
func (s *Activity) ActiveUserBuy(ctx context.Context, req dto.StatisticsDaoReq) (res []dto.StatisticsUserIdEntity, err error) {
	res = []dto.StatisticsUserIdEntity{}
	// debug 打印语句
	model := s.DB.WithContext(ctx).Debug().
		Model(entity.Order{}).
		Select(" FROM_UNIXTIME(create_time, '%Y-%m-%d' ) dates,  distributor_user_id").
		Where("delete_time", consts.UnDelete).
		Where("status", entity.OrderStatusSuccess)

	if req.StartTime > 0 {
		model.Where("create_time >= ?", req.StartTime)
	}
	if req.EndTime > 0 {
		model.Where("create_time <= ?", req.EndTime)
	}
	if len(req.ChannelId) > 0 {
		model.Where("channel_id = ?", req.ChannelId)
	}
	if req.AppId > 0 {
		model.Where("app_id = ?", req.AppId)
	}

	err = model.
		Group("dates, distributor_user_id").
		Find(&res).Error
	return
}
*/

/*
// 执行原生sql

func (a *Online) GetActivityList(req statistics.OnlineList, ctx context.Context) (result []entity.Activity) {
	// 进行sql替换
	startTime := strconv.Itoa(int(req.StartTime))
	endTime := strconv.Itoa(int(req.EndTime))
	sql := "WHERE sum_time BETWEEN " + startTime + " AND " + endTime
	// 通道id
	if len(req.ChannelId) > 0 {
		sql = sql + " AND channel_id = " + req.ChannelId
	}
	// 应用id
	if req.AppId > 0 {
		appId := strconv.Itoa(int(req.AppId))
		sql = sql + " AND app_id = " + appId
	}

	SQL := fmt.Sprintf(OnlineSQL, sql)
	a.DB.WithContext(ctx).Raw(SQL).Scan(&result)
	return result
}
*/