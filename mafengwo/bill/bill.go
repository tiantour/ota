package bill

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWBillList bill list
	MFWBillList struct {
		TotalResults int32          `json:"total_results"` //
		List         []*MFWBillItem `json:"list"`          //
	}
	// MFWBillItem bill item
	MFWBillItem struct {
		DetailID           int32   `json:"detail_id"`            // 交易号
		OrderID            string  `json:"order_id"`             // 	旅行商城业务订单号
		SalesAmount        float64 `json:"sales_amount"`         // 销售金额
		CommisionRate      float64 `json:"commision_rate"`       // 佣金比例
		CommisionAmount    float64 `json:"commision_amount"`     // 金/手续费
		PayAmount          float64 `json:"pay_amount"`           // 支付金额
		BonusAmount        float64 `json:"bonus_amount"`         // 马蜂窝补贴，单位元
		OtaBonusAmount     float64 `json:"ota_bonus_amount"`     // 商家补贴
		ExpectSettleAmount float64 `json:"expect_settle_amount"` // 应结算金额
		SettleAmount       float64 `json:"settle_amount"`        // 结算金额(原币种)
		Currency           string  `json:"currency"`             //	币种
		ExchangeRate       float64 `json:"exchange_rate"`        // 汇率
		PayTime            string  `json:"pay_time"`             // 支付时间
		SettleTime         string  `json:"settle_time"`          // 预计结算时间
		ISReplenishOrder   int32   `json:"is_replenish_order"`   // 是否是补款单 1是 0否
		ConfirmTime        float64 `json:"confirm_time"`         // 确认结算时间
	}
)

// Bill bill
type Bill struct {
	SettleStatus int32 `json:"settle_status,omitempty"` // 0:未开放 1:可结算
	PageNo       int32 `json:"page_no,omitempty"`       // 页码
	PageSize     int32 `json:"page_size,omitempty"`     // 单页条数（最大值1000）
}

// NewBill new bill
func NewBill() *Bill {
	return &Bill{}
}

// Transaction transaction
func (b *Bill) Transaction(args *Bill) (*MFWBillList, error) {
	action := "sales.bills.transaction.get"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWBillList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
