package task1_error

import (
	"time"
)

type Item struct {
	ID     uint      `gorm:"primaryKey;column:runoob_id"`
	Title  string    `gorm:"column:runoob_title;type:varchar(100);"`
	Author string    `gorm:"column:runoob_author;type:varchar(40);"`
	Date   time.Time `gorm:"column:submission_date;type:timestamp;"`
}

func (Item) TableName() string {
	return "runoob_tbl"
}
