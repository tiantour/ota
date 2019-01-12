package refund

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWRefundList refund list
	MFWRefundList struct {
		List  []*MFWRefundItem `json:"list"`  // 退款订单列表
		Total int32            `json:"total"` // 退款单总数量
	}

	// MFWRefundItem refund item
	MFWRefundItem struct {
		RefundID           int32                `json:"refund_id"`           // 马蜂窝订单关联的退款单号
		OrderID            string               `json:"order_id"`            // 旅行商城业务订单号
		ApplyCtime         string               `json:"apply_ctime"`         // 退款申请时间
		UpdateTime         string               `json:"update_time"`         // 最后更新时间
		ApplyReason        int32                `json:"apply_reason"`        // "20:其他原因 21:行程有变 22:供应商无货了 23:数量排错了"
		RefundRemark       []*MFWRefundRemark   `json:"refund_remark"`       // 退款备注
		RefundFlag         int32                `json:"refund_flag"`         // 1:已完成退款 2:已申请退款 3:拒绝退款 4:已确认退款"
		TotalPrice         string               `json:"total_price"`         // 订单原始金额
		PaymentFee         string               `json:"payment_fee"`         // 商家原始在此订单上的结算金额
		RefundedItems      []*MFWRefundData     `json:"refunded_items"`      // 已退款项
		RefundFee          string               `json:"refund_fee"`          // 已退款金额，单位：元
		RefundingItems     []*MFWRefundData     `json:"refunding_items"`     // 正在退款项
		Items              []*MFWOrderItem      `json:"items"`               // 购买项
		Travelers          []*MFWTraveler       `json:"travelers"`           // 出行人信息
		RefundedTravelers  []*MFWRefundTraveler `json:"refunded_travelers"`  // 已经退款的出行人信息
		RefundingTravelers []*MFWRefundTraveler `json:"refunding_travelers"` // 退款中的出行人信息
	}

	// MFWRefundData refund itemx
	MFWRefundData struct {
		ID         int    `json:"id "`         // 退款项ID
		Num        int    `json:"num"`         // 本项可退个数
		Name       string `json:"name"`        // 退款项描述
		RefundSold string `json:"refund_sold"` // 实付金额
		RefundFee  string `json:"refund_fee"`  // 退款金额
	}

	// MFWOrderItem order item
	MFWOrderItem struct {
		ID               int32   `json:"id"`                 // 退款项ID
		SkuID            int32   `json:"skuId"`              // 马蜂窝SKU ID，SKU唯一标识
		Num              int32   `json:"num"`                // 本项可退个数
		Price            float64 `json:"price"`              // 本项单价可退金额
		PriceType        int32   `json:"price_type"`         // 费用项 具体说明请见 费用类型说明表
		Name             string  `json:"name"`               // 退款项描述
		PaymentFee       float64 `json:"payment_fee"`        // 本项支付金额
		TotalPrice       float64 `json:"total_price"`        // 本项可退总金额
		RemainNum        int32   `json:"remain_num"`         // 剩余可退数量
		RemainPaymentFee float64 `json:"remain_payment_fee"` // 剩余可退金额
	}

	// MFWRefundRemark refund remark
	MFWRefundRemark struct {
		UID    int32  `json:"uid"`    // 退款备注添加人UID
		Remark string `json:"remark"` // 退款备注
		Ctime  string `json:"ctime"`  // 添加备注时间
	}

	// MFWTraveler traveler
	MFWTraveler struct {
		TravelerID      int32  `json:"traveler_id"`       // 出行人ID
		Name            string `json:"name"`              // 出行人姓名
		Cellphone       string `json:"cellphone"`         // 联系电话
		IDCard          string `json:"id_card"`           // 身份证号码
		Passport        string `json:"passport"`          // 护照证件号
		LaissezPasserTW string `json:"laissez_passer_tw"` // 台湾通行证
		LaissezPasser   string `json:"laissez_passer"`    // 港澳通行证
	}

	// MFWRefundTraveler refund tradveler
	MFWRefundTraveler struct {
		TravelerID int32 `json:"traveler_id"` // 出行人ID
	}
)

// Refund refund
type Refund struct {
	OrderID      string          `json:"order_id,omitempty"`      // 旅行商城业务订单号
	RefundStatus int32           `json:"refund_status,omitempty"` // 退款状态 0:全部 1:已完成退款 2:已申请退款 3:拒绝退款 4:已确认退款
	PageNo       int32           `json:"page_no,omitempty"`       // 页码
	PageSize     int32           `json:"page_size,omitempty"`     // 单页条数（最大值20）
	Items        []*MFWOrderItem `json:"items,omitempty"`         // 购买项
	Reason       int32           `json:"reason,omitempty"`        // 退款原因。 商家提交的退款信息客人可见。 20:其他原因 21:行程有变 22:供应商无货了 23:数量排错了
	Remark       string          `json:"remark,omitempty"`        // 退款备注
	Travelers    []*MFWTraveler  `json:"travelers,omitempty"`     // 出行人信息
	RefundID     int32           `json:"refund_id,omitempty"`     // 马蜂窝订单关联的退款单号
	CalcType     int32           `json:"calc_type,omitempty"`     // 退款金额计算类型 具体说明请见 退款确认接口调用说明
	IsTotal      int32           `json:"is_total,omitempty"`      // 订单退款完成标识： 0-退款后订单继续进行。 1-退款后服务结束，订单关闭。

}

// NewRefund new refund
func NewRefund() *Refund {
	return &Refund{}
}

// List get refund list
func (r *Refund) List(args *Refund) (*MFWRefundList, error) {
	action := "sales.refund.list.get"
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

// Item get refund item
func (r *Refund) Item(orderID string) (*MFWRefundItem, error) {
	action := "sales.refund.detail.get"
	data, err := json.Marshal(&Refund{
		OrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWRefundItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Apply apply refund
func (r *Refund) Apply(args *Refund) (*Refund, error) {
	action := "sales.refund.apply"
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

// Confirm confirm refund
func (r *Refund) Confirm(args *Refund) (*Refund, error) {
	action := "sales.refund.confirm"
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

// Refuse refuse refund
func (r *Refund) Refuse(args *Refund) (*Refund, error) {
	action := "sales.refund.refuse"
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

// Repeal repeal refund
func (r *Refund) Repeal(args *Refund) (*Refund, error) {
	action := "sales.refund.repeal"
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
