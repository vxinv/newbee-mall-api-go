// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTzArea = "tz_area"

// TzArea mapped from table <tz_area>
type TzArea struct {
	AreaID   int64   `gorm:"column:area_id;type:bigint(20);primaryKey;autoIncrement:true" json:"area_id"`
	AreaName *string `gorm:"column:area_name;type:varchar(50)" json:"area_name"`
	ParentID *int64  `gorm:"column:parent_id;type:bigint(20);index:parent_id,priority:1" json:"parent_id"`
	Level    *int32  `gorm:"column:level;type:int(1)" json:"level"`
}

// TableName TzArea's table name
func (*TzArea) TableName() string {
	return TableNameTzArea
}
