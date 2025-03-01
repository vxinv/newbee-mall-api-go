// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTzAttachFile = "tz_attach_file"

// TzAttachFile mapped from table <tz_attach_file>
type TzAttachFile struct {
	FileID       int64      `gorm:"column:file_id;type:bigint(20);primaryKey;autoIncrement:true" json:"file_id"`
	FilePath     *string    `gorm:"column:file_path;type:varchar(255);comment:文件路径" json:"file_path"`                                // 文件路径
	FileType     *string    `gorm:"column:file_type;type:varchar(20);comment:文件类型" json:"file_type"`                                 // 文件类型
	FileSize     *int32     `gorm:"column:file_size;type:int(10);comment:文件大小" json:"file_size"`                                     // 文件大小
	UploadTime   *time.Time `gorm:"column:upload_time;type:datetime;comment:上传时间" json:"upload_time"`                                // 上传时间
	FileJoinID   *int64     `gorm:"column:file_join_id;type:bigint(20);comment:文件关联的表主键id" json:"file_join_id"`                      // 文件关联的表主键id
	FileJoinType *int32     `gorm:"column:file_join_type;type:tinyint(2);comment:文件关联表类型：1 商品表  FileJoinType" json:"file_join_type"` // 文件关联表类型：1 商品表  FileJoinType
}

// TableName TzAttachFile's table name
func (*TzAttachFile) TableName() string {
	return TableNameTzAttachFile
}
