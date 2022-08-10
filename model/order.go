package model

import "time"

type Order struct {
	OrderID     string `gorm:"primarykey"`
	OrderType   int    `json:"order_type"` // 订单类型  打赏/会员/商品交易/酒店/出行/吃住
	OrderStatus string
	FinishTime  int64 // 完成取消过期最后时间
	Price       int64 // 分

	UserID    string
	AddressID uint
	PaymentID string // 支付信息
	StoreId   string // 商家
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderDispatch struct {
	OrderInfo   Order           `json:"order_info"`
	ProductInfo OrderProduct    `json:"product_info"`
	PaymentInfo Payment         `json:"payment_info"`
	AddressInfo DeliveryAddress `json:"address_info"`
	UserInfo    User            `json:"user_info"`
}

type OrderVipMemberRequest struct {
	UserId      string `json:"user_id"`
	OrderId     string `json:"order_id"`
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Count       int64  `json:"count"`
}

type OrderRequest struct {
	UserId       string         `json:"user_id"` // 付款方
	OrderId      string         `json:"order_id"`
	StoreId      string         `json:"store_id"`       // 销售方
	AcceptUserId string         `json:"accept_user_id"` // 收货方  代买情况
	ProductList  []StoreProduct `json:"product_list"`
}

type OrderView struct {
	Order
	OrderProduct
}
