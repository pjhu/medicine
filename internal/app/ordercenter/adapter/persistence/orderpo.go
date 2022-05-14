package persistence

import "time"

type UserOrderPO struct {
	Id               uint      `gorm:"primaryKey;comment:主键"`
	OrderAmountTotal int       `gorm:"not null;default:0;comment:订单金额"`
	PayChannel       string    `gorm:"type:varchar(200);comment:付款渠道"`
	OrderStatus      string    `gorm:"type:varchar(200);comment:订单状态"`
	CreatedAt        time.Time `gorm:"type:datetime;not null;default:current_timestam;comment:创建时间"`
	CreatedBy        uint      `gorm:"not null;comment:创建人"`
	LastModifiedAt   time.Time `gorm:"type:datetime;not null;default:current_timestam;comment:创建时间"`
}
