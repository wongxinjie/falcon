package loginaccout

import "time"

type Schema struct {
	tableName struct{}  `gorm:"login_account,alias:login_account"`
	Id        int64     `gorm:"column:id;primaryKey" db:"id"`
	Email     string    `gorm:"column:email;type:varchar(256)"`
	Phone     string    `gorm:"column:phone;type:varchar(16)"`
	Password  string    `gorm:"column:password;type:varchar(512)"`
	LastLogin time.Time `gorm:"column:last_login;autoCreateTime:milli;autoUpdateTime"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (d *Schema) TableName() string {
	return "login_account"
}
