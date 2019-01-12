package replenish

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWRefundList refund list
	MFWRefundList struct {
		List  []*MFWRefundItem `json:"list"`  // 马蜂窝自由行补款单退款信息
		Total int32            `json:"total"` // 补款退款单总个数
	}

	// MFWRefundItem refund item
	MFWRefundItem struct {
		ReplenishID  string  `json:"replenish_id"`  // 旅行商城业务订单号关联补款单号
		RefundFlag   int32   `json:"refund_flag"`   // 退款单状态1-退款中2-退款完成3-退款驳回4-未申请
		TotalPrice   float64 `json:"total_price"`   // 补款单原始金额
		RefundFee    float64 `json:"refund_fee"`    // 可退款金额
		RefundingFee float64 `json:"refunding_fee"` // 补款单正在退款金额
		RefundedFee  float64 `json:"refunded_fee"`  // 补款单已退金额
	}
)

// Refund refund
type Refund struct {
	OrderID     string  `json:"order_id,omitempty"`     // 旅行商城业务订单号
	ReplenishID string  `json:"replenish_id,omitempty"` // 旅行商城业务订单号关联补款单号
	RefundFee   float64 `json:"refund_fee,omitempty"`   // 退款金额
	Remark      string  `json:"remark,omitempty"`       // 补款单退款备注
}

// NewRefund new refund
func NewRefund() *Refund {
	return &Refund{}
}

// List get refund list
func (r *Refund) List(args *Refund) (*MFWRefundList, error) {
	action := "sales.replenish.refund.get"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWRefundList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Create create refund
func (r *Refund) Create(args *Refund) (*Refund, error) {
	action := "sales.replenish.refund.create"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Refund{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Cancel cancel refund
func (r *Refund) Cancel(args *Refund) (*Refund, error) {
	action := "sales.replenish.refund.cancel"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Refund{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
