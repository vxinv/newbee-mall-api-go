// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTzOrderSettlement = "tz_order_settlement"

// TzOrderSettlement mapped from table <tz_order_settlement>
type TzOrderSettlement struct {
	SettlementID uint64     `gorm:"column:settlement_id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:支付结算单据ID" json:"settlement_id"`    // 支付结算单据ID
	PayNo        *string    `gorm:"column:pay_no;type:varchar(36);comment:支付单号" json:"pay_no"`                                                            // 支付单号
	BizPayNo     *string    `gorm:"column:biz_pay_no;type:varchar(255);comment:外部订单流水号" json:"biz_pay_no"`                                                // 外部订单流水号
	OrderNumber  *string    `gorm:"column:order_number;type:varchar(36);uniqueIndex:primary_order_no,priority:1;comment:order表中的订单号" json:"order_number"` // order表中的订单号
	PayType      *int32     `gorm:"column:pay_type;type:int(1);comment:支付方式 1 微信支付 2 支付宝" json:"pay_type"`                                                // 支付方式 1 微信支付 2 支付宝
	PayTypeName  *string    `gorm:"column:pay_type_name;type:varchar(50);comment:支付方式名称" json:"pay_type_name"`                                            // 支付方式名称
	PayAmount    *float64   `gorm:"column:pay_amount;type:decimal(15,2);comment:支付金额" json:"pay_amount"`                                                  // 支付金额
	IsClearing   *int32     `gorm:"column:is_clearing;type:int(1);comment:是否清算 0:否 1:是" json:"is_clearing"`                                               // 是否清算 0:否 1:是
	UserID       *string    `gorm:"column:user_id;type:varchar(36);index:user_id,priority:1;comment:用户ID" json:"user_id"`                                 // 用户ID
	CreateTime   time.Time  `gorm:"column:create_time;type:datetime;not null;comment:创建时间" json:"create_time"`                                            // 创建时间
	ClearingTime *time.Time `gorm:"column:clearing_time;type:datetime;comment:清算时间" json:"clearing_time"`                                                 // 清算时间
	Version      *int32     `gorm:"column:version;type:int(11);comment:版本号" json:"version"`                                                               // 版本号
	PayStatus    *int32     `gorm:"column:pay_status;type:int(1);comment:支付状态" json:"pay_status"`                                                         // 支付状态
}

// TableName TzOrderSettlement's table name
func (*TzOrderSettlement) TableName() string {
	return TableNameTzOrderSettlement
}
