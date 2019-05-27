package ticket

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWQuantity quantity
	MFWQuantity struct {
		NotUsedQuantity  int32 `json:"not_used_quantity"` // 未使用数量
		UsedQuantity     int32 `json:"used_quantity"`     // 已使用数量
		RefundedQuantity int32 `json:"refunded_quantity"` // 已退款数量
	}
	// MFWVoucherList voucher list
	MFWVoucherList struct {
		SKUID    int32      `json:"sku_id,omitempty"`     // 马蜂窝sku_id，马蜂窝sku_id与ota_sku_id商家编码二选一
		OTASKUID string     `json:"ota_sku_id,omitempty"` // 商家设置的SKU外部编码，sku_id与ota_sku_id二选一
		Type     int32      `json:"type,omitempty"`       // 1－一码一验，一个库存对应一个凭证码； 2－一码多验，多个库存对应一个凭证码；
		Vouchers MFWVoucher `json:"vouchers,omitempty"`   // 凭证码列表
		Quantity int32      `json:"quantity,omitempty"`   // 购买数量
	}
	// MFWVoucher voucher
	MFWVoucher struct {
		Voucher    string  `json:"voucher,omitempty"`     // 凭证码
		VoucherPic string  `json:"voucher_pic,omitempty"` // 图片码凭证，传图片链接地址, 有多个时，顺序必须与vouchers的顺序一致，马蜂窝将以此顺序与voucher顺序进行一一对应；
		Travelers  []int32 `json:"travelers,omitempty"`   // 出行人信息
		Status     int32   `json:"status,omitempty"`      // 凭证状态：1-未使用，2-已使用，3-已退款，4-已废弃（对应的门票还未消费，但是此凭证码作废了）
	}
)

// Update update
type Update struct {
	OrderID            string           `json:"order_id,omitempty"`             // 旅行商城业务订单号
	SKUID              int32            `json:"sku_id,omitempty"`               // 马蜂窝sku_id，马蜂窝sku_id与ota_sku_id商家编码二选一
	OTASKUID           string           `json:"ota_sku_id,omitempty"`           // 商家设置的SKU外部编码，sku_id与ota_sku_id二选一
	QuantityStatusInfo MFWQuantity      `json:"quantity_status_info,omitempty"` // 一码多验 门票数量状态信息
	Memo               string           `json:"memo,omitempty"`                 // 订单状态的备注信息。该备注对用户可见，请谨慎填写，并有300个字符的限制
	TicketVouchers     []MFWVoucherList `json:"ticket_vouchers,omitempty"`      //
}

// NewUpdate new update
func NewUpdate() *Update {
	return &Update{}
}

// Quantity quantity
func (u *Update) Quantity(args *Update) (*Update, error) {
	action := "sales.ticket.quantity.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Update{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Order order
func (u *Update) Order(args *Update) (*Update, error) {
	action := "sales.ticket.order.status.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Update{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
