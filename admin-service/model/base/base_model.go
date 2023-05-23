package base

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type GcsModel interface {
	TableName() string
	ModelName() string
}

type GcsModelPrefix struct {
	ID uint `json:"id" gorm:"column:id;primaryKey;comment:主键" AliasName:"编号"`
}

type GcsModelSuffix struct {
	CreatedAt   Date                   `json:"createdAt" gorm:"column:created_at;comment:创建时间" AliasName:"创建时间"`
	UpdatedAt   Date                   `json:"updatedAt" gorm:"column:updated_at;comment:更新时间" AliasName:"修改时间"`
	DeletedAt   gorm.DeletedAt         `json:"deletedAt" gorm:"index;comment:删除时间" `
	ExpandProps map[string]interface{} `json:"expandProps" gorm:"-"`
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Trim(s, `"`)
	date, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		d.Time, _ = time.Parse("2006-01-02 15:04:05", "1000-01-01")
		return nil
	}
	*d = Date{date}
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	formattedDate := d.Time.Format("2006-01-02 15:04:05")
	return []byte(`"` + formattedDate + `"`), nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan Date")
	}
	d.Time = t
	return nil
}
