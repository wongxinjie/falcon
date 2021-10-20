package dynasty

import "time"

type Schema struct {
	tableName struct{}  `gorm:"dynasty,alias:dna"`
	Id        int64     `gorm:"column:id;primaryKey" db:"id"`
	Name      string    `gorm:"column:name;type:varchar(16);index" db:"name"`
	Count     int64     `gorm:"column:count;count"`
	CreatedAt time.Time `gorm:"column:created_at;created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;updated_at"`
}

func (d *Schema) TableName() string {
	return "dynasty"
}
