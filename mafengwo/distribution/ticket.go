package distribution

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWTicketData mafengwo ticket data
	MFWTicketData struct {
		OrderID        string         `json:"order_id"`         // 马蜂窝编号
		PartnerOrderID string         `json:"partner_order_id"` // 商家编号
		OrderInfo      MFWTicketOrder `json:"order_info"`       // 订单详情
	}
	// MFWTicketOrder mafengwo ticket order
	MFWTicketOrder struct {
		SalesID    string  `json:"sales_id"`    // 产品ID
		PayTime    string  `json:"pay_time"`    // 支付时间
		Date       string  `json:"date"`        // 入园时间
		TotalPrice float64 `json:"total_price"` // 订单总金额
		PaymentFee float64 `json:"payment_fee"` // 支付金额
		Items      []int32 `json:"items"`       // 购买项
	}
	// MFWTicketList mafengwo ticket list
	MFWTicketList struct {
		Num      int32  `json:"num"`       // 购买个数
		UserType int32  `json:"user_type"` // 适用人群
		Name     string `json:"name"`      // 游客姓名
	}
)

// Ticket ticket
type Ticket struct {
	PartnerOrderID string        `json:"partner_order_id"` // 商家编号
	SalesID        string        `json:"sales_id"`         // 产品ID
	Date           string        `json:"date"`             // 入园时间
	Mobile         string        `json:"mobile"`           // 预定人手机号
	TicketList     MFWTicketList `json:"ticket_list"`      // 票务预定信息列表
}

// NewTicket new ticket
func NewTicket() *Ticket {
	return &Ticket{}
}

// Order ticket order create
func (t *Ticket) Order(args *Ticket) (*MFWTicketData, error) {
	action := "action：sales.distribution.ticket.order.create"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWTicketData{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
