package model

type OrderStatus int8

// 订单
const (
	OrderStatusCREATE  OrderStatus = 0 //新订单
	OrderStatusPAYMENT OrderStatus = 7 //已付款未发货
	OrderStatusCANCEL  OrderStatus = 8 //取消
	OrderStatusSUCCESS OrderStatus = 9 //发货成功
)

func init() {
	Register(&Order{})
}

type Order struct {
	Model    `bson:"inline"`
	Trade    string      `json:"trade" bson:"trade" index:"name:"` //平台订单号，用于对账
	Create   int64       `json:"create" bson:"create" `            //订单创建时间
	Expire   int64       `json:"expire" bson:"expire" `            //过期时间,未完成订单过期后无法继续
	Status   OrderStatus `json:"status" bson:"status"`             //状态
	Amount   float32     `json:"amount" bson:"amount"`             //订单金额(元，默认货币)
	Receive  float32     `json:"receive" bson:"receive"`           //实际到账金额(元),平台折扣，代金券之类会抵消部分金额
	Platform int32       `json:"pfm" bson:"pfm"`                   //支付平台
	Currency string      `json:"cyt" bson:"cyt"`                   //平台对账默认币种
}

// TableName 数据库表名
func (this *Order) TableName() string {
	return "orders"
}
