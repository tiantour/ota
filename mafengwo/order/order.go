package order

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// OrderList order list
	OrderList struct {
		List  []*OrderItem `json:"list"`
		Total int32        `json:"total"`
	}

	// OrderItem order item
	OrderItem struct {
		OrderID         string           `json:"orderId"`         // 订单号
		Status          *OrderStatus     `json:"status"`          // 订单状态
		GoDate          string           `json:"goDate"`          // 旅行出行日期
		EndDate         string           `json:"endDate"`         // 旅行结束日期
		Paytime         string           `json:"payTime"`         // 订单支付时间
		Ctime           string           `json:"ctime"`           // 订单创建时间
		BookingPeople   *BookingPeople   `json:"bookingPeople"`   // 预订人信息
		SalesID         int32            `json:"salesId"`         // 马蜂窝产品id，产品唯一标识
		SalesName       string           `json:"salesName"`       // 产品名称
		OtaSalesName    string           `json:"otaSalesName"`    // 商家设置的产品外部编码
		SalesType       int32            `json:"salesType"`       // 具体说明请见 马蜂窝品类说明表
		Mdd             string           `json:"mdd"`             // 目的地
		From            string           `json:"from"`            // 订单关联产品出发地
		SkuID           int32            `json:"skuId"`           // 马蜂窝SKU ID，SKU唯一标识
		OtaSkuID        string           `json:"otaSkuId"`        // 商家设置的SKU外部编码
		SkuName         string           `json:"skuName"`         // 门票 成人票	SKU名称
		TotalPrice      string           `json:"totalPrice"`      // 订单原始金额
		PaymentFee      string           `json:"paymentFee"`      // 用户实际支付金额
		Items           []*BookingItem   `json:"items"`           // 订单购买项详细信息
		PromotionDetail *PromotionDetail `json:"promotionDetail"` // 订单优惠信息
		TravelPeople    *TravelPeople    `json:"travel_people"`   // 出行人信息
		Skus            []*Sku           `json:"skus"`            // 库存信息
		ConfirmInfo     *ConfirmInfo     `json:"confirmInfo"`     // 确认方式信息
		IssueInfo       *IssueInfo       `json:"issueInfo"`       // 出单方式信息
	}

	// OrderStatus order static
	OrderStatus struct {
		OrderStatus   int32 `json:"orderStatus"`   // 订单状态：1-待支付，2-待出单，4-已出单，5-已完成，6-已关闭
		AllRefundFlag int32 `json:"allRefundFlag"` // 全退标识：0 无退款 1 未全部退款 2 已全部退款
		RefundStatus  int32 `json:"refundStatus"`  // 退款状态 1-已完成退款 2-已申请退款 3-拒绝退款 4-已确认退款 0-可发起退款
	}

	// BookingPeople booking people
	BookingPeople struct {
		UID       int32  `json:"uid"`        // 预订人马蜂窝的UID
		Name      string `json:"name"`       // 预订人姓名
		Email     string `json:"email"`      // 预订人邮箱
		Phone     string `json:"phone"`      // 预订人手机号
		PhoneArea string `json:"phone_area"` // 预订人手机区号
		Wechat    string `json:"wechat"`     // 预订人微信
		Remark    string `json:"remark"`     // 用户下单时添加的订单备注
	}

	// BookingItem booking item
	BookingItem struct {
		RemainPaymentFee float64 `json:"remain_payment_fee"` // 剩余可退金额
		RemainNum        int32   `json:"remain_num"`         // 剩余可退数量
		SkuID            int32   `json:"skuId"`              // 马蜂窝SKU ID，SKU唯一标识
		PriceType        int32   `json:"price_type"`         // 费用项 具体说明请见 费用类型说明表
		PaymentFee       float64 `json:"payment_price"`      // 本项支付金额
		TotalPrice       float64 `json:"total_price"`        // 本项总金额
		Name             string  `json:"name"`               //	购买项描述
		Price            float64 `json:"price"`              // 本项单价金额
		Num              int32   `json:"num"`                // 本项购买个数
		ID               int32   `json:"id"`                 // 购买项ID
	}

	// PromotionDetail promotion detail
	PromotionDetail struct {
		Reduce    float64 `json:"reduce_"`    // 马蜂窝补贴金额
		ReduceOta float64 `json:"reduce_ota"` // 商家补贴金额
	}

	// Sku sku
	Sku struct {
		StockName string `json:"stockName"` // 不限量套餐	库存名称
		OtaSkuID  string `json:"otaSkuId"`  // 商家设置的SKU外部编码
		SkuID     int32  `json:"skuId"`     // 马蜂窝SKU ID，SKU唯一标识
	}

	// ConfirmInfo confirmInfo
	ConfirmInfo struct {
		Type     int32  `json:"type"`      // 确认方式：0-普通确认，1-即时确认，2-二次确认
		OverTime string `json:"over_time"` // 类型为二次确认时，需进行确认的时限
	}

	// IssueInfo issueInfo
	IssueInfo struct {
		Type int32 `json:"type"` // 出单方式：1-商家出单，2-自动出单
		Days int32 `json:"days"` // 出单期限天数，0表示不限制
	}

	// OrderParam order param
	OrderParam struct {
		TimeFrom    string `json:"time_from,omitempty"`    // 用户下单起始时间
		TimeTo      string `json:"time_to,omitempty"`      // 用户下单结束时间
		SalesType   int32  `json:"sales_type,omitempty"`   // 订单关联产品品类 具体说明请见 马蜂窝品类说明表
		Color       int32  `json:"color,omitempty"`        // 订单标记的颜色 0-白色,1-红色,2-橙色,3-黄色,4-绿色,5-蓝色,6-紫色,7-灰色
		OrderStatus int32  `json:"order_status,omitempty"` // 订单状态：1-待支付，2-待出单，4-已出单，5-已完成，6-已关闭
	}
)

// Order order
type Order struct {
	OrderID   string      `json:"order_id,omitemtpy"`   // 订单号
	Status    int32       `json:"status,omitemtpy"`     // 订单状态值（已联系用户并确认库存(12)，已发确认单(13)）
	AddStatus int32       `json:"add_status,omitemtpy"` // 签证订单的子状态值(默认(0),已快递(1),已自取(2),已发邮件(3),已完成预约(4),已拒签(5))
	Memo      string      `json:"memo,omitemtpy"`       // 订单状态的备注信息。该备注对用户可见，请谨慎填写，并有300个字符的限制
	PageNo    int32       `json:"page_no,omitempty"`    // 当前页数
	PageSize  int32       `json:"page_size,omitempty"`  // 单页条数（最大值20）
	Params    *OrderParam `json:"params,omitempty"`     // 拓展参数
}

// NewOrder new order
func NewOrder() *Order {
	return &Order{}
}

// List get order lst
func (o *Order) List(schedule, page, size int32) (*OrderList, error) {
	data, err := json.Marshal(&Order{
		PageNo:   page,
		PageSize: size,
		Params: &OrderParam{
			OrderStatus: schedule,
		},
	})
	if err != nil {
		return nil, err
	}

	action := "sales.order.list.get"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := OrderList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Item get order item
func (o *Order) Item(orderID string) (*OrderItem, error) {
	data, err := json.Marshal(&Order{
		OrderID: orderID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.order.detail.get"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := OrderItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Update update status
func (o *Order) Update(args *Order) (*Order, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	action := "sales.order.status.update"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Order{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
