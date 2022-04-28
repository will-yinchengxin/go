package gorm

type Users struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Name       string `gorm:"column:name" json:"name" form:"name"`
	ComId      int64  `gorm:"column:com_id" json:"com_id" form:"com_id"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time" form:"delete_time"`
}

type User struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Name       string `gorm:"column:name" json:"name" form:"name"`
	ComId      int64  `gorm:"column:com_id" json:"com_id" form:"com_id"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time" form:"delete_time"`
	// 一对多
	Company []Company `gorm:"foreignKey:UserId;references:ID"`
	Test    Test      `gorm:"foreignKey:Status;references:ID"`
}

type Company struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Industry int64  `gorm:"column:industry" json:"industry" form:"industry"`
	Name     string `gorm:"column:name" json:"name" form:"name"`
	Job      string `gorm:"column:job" json:"job" form:"job"`
	UserId   int64  `gorm:"column:user_id" json:"user_id" form:"user_id"`
}

type Test struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Title      string `gorm:"column:title" json:"title" form:"title"`
	Content    string `gorm:"column:content" json:"content" form:"content"`
	Status     int64  `gorm:"column:status" json:"status" form:"status"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time" form:"delete_time"`
}
