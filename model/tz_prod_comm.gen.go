// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTzProdComm = "tz_prod_comm"

// TzProdComm 商品评论
type TzProdComm struct {
	ProdCommID   uint64     `gorm:"column:prod_comm_id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"prod_comm_id"` // ID
	ProdID       uint64     `gorm:"column:prod_id;type:bigint(20) unsigned;not null;index:prod_id,priority:1;comment:商品ID" json:"prod_id"`     // 商品ID
	OrderItemID  *uint64    `gorm:"column:order_item_id;type:bigint(20) unsigned;comment:订单项ID" json:"order_item_id"`                          // 订单项ID
	UserID       *string    `gorm:"column:user_id;type:varchar(36);comment:评论用户ID" json:"user_id"`                                             // 评论用户ID
	Content      *string    `gorm:"column:content;type:varchar(500);comment:评论内容" json:"content"`                                              // 评论内容
	ReplyContent *string    `gorm:"column:reply_content;type:varchar(500);comment:掌柜回复" json:"reply_content"`                                  // 掌柜回复
	RecTime      *time.Time `gorm:"column:rec_time;type:datetime;comment:记录时间" json:"rec_time"`                                                // 记录时间
	ReplyTime    *time.Time `gorm:"column:reply_time;type:datetime;comment:回复时间" json:"reply_time"`                                            // 回复时间
	ReplySts     *int32     `gorm:"column:reply_sts;type:int(1);comment:是否回复 0:未回复  1:已回复" json:"reply_sts"`                                   // 是否回复 0:未回复  1:已回复
	Postip       *string    `gorm:"column:postip;type:varchar(16);comment:IP来源" json:"postip"`                                                 // IP来源
	Score        *int32     `gorm:"column:score;type:tinyint(2);comment:得分，0-5分" json:"score"`                                                 // 得分，0-5分
	UsefulCounts *int32     `gorm:"column:useful_counts;type:int(11);comment:有用的计数" json:"useful_counts"`                                      // 有用的计数
	Pics         *string    `gorm:"column:pics;type:varchar(1000);comment:晒图的json字符串" json:"pics"`                                             // 晒图的json字符串
	IsAnonymous  *int32     `gorm:"column:is_anonymous;type:int(1);comment:是否匿名(1:是  0:否)" json:"is_anonymous"`                                // 是否匿名(1:是  0:否)
	Status       *int32     `gorm:"column:status;type:int(1);comment:是否显示，1:为显示，0:待审核， -1：不通过审核，不显示。 如果需要审核评论，则是0,，否则1" json:"status"`         // 是否显示，1:为显示，0:待审核， -1：不通过审核，不显示。 如果需要审核评论，则是0,，否则1
	Evaluate     *int32     `gorm:"column:evaluate;type:tinyint(2);comment:评价(0好评 1中评 2差评)" json:"evaluate"`                                   // 评价(0好评 1中评 2差评)
}

// TableName TzProdComm's table name
func (*TzProdComm) TableName() string {
	return TableNameTzProdComm
}
