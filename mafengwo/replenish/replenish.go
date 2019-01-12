package replenish

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWReplenishList replenish list
	MFWReplenishList struct {
		Total int32               `json:"total"` // 数量
		List  []*MFWReplenishItem `json:"list"`  //
	}

	// MFWReplenishItem replenish item
	MFWReplenishItem struct {
		OrderID     string  `json:"order_id"`     // 旅行商城业务订单号
		ReplenishID string  `json:"replenish_id"` // 旅行商城业务订单号关联补款单号
		Status      int32   `json:"status"`       // 补款单状态0-待支付；1-已支付；2-申请退款中；3-部分退款成功；4-全部退款成功；5-已关闭
		Ctime       string  `json:"ctime"`        // 补款单创建时间
		Reason      int32   `json:"reason"`       // 创建补款单原因/类型：0-酒店升级/变更；1-房型升级/变更；2-航班升级/变更；3-产品服务升级/变更；4-套餐变更；5-增订项目；76-其他；
		Fee         float64 `json:"fee"`          // 补款单具体金额，精确到小数点后两位
		Remark      string  `json:"remark"`       // 补款单备注
	}
)

// Replenish replenish
type Replenish struct {
	OrderID     string  `json:"order_id,omitempty"`     // 旅行商城业务订单号
	BeginDate   string  `json:"begin_date,omitempty"`   // [条件-起始时间]补款单创建时间
	EndDate     string  `json:"end_date,omitempty"`     // [条件-结束时间]补款单创建时间
	ReplenishID string  `json:"replenish_id,omitempty"` // 旅行商城业务订单号关联补款单号
	Fee         float64 `json:"fee,omitempty"`          // 补款金额
	Reason      int32   `json:"reason,omitempty"`       // 创建补款单原因/类型：0-酒店升级/变更；1-房型升级/变更；2-航班升级/变更；3-产品服务升级/变更；4-套餐变更；5-增订项目；6-其他
	Remark      string  `json:"remark,omitempty"`       // 补款单备注
}

// NewReplenish new replenish
func NewReplenish() *Replenish {
	return &Replenish{}
}

// List get replenish list
func (r *Replenish) List(args *Replenish) (*MFWReplenishList, error) {
	action := "sales.replenish.list.get"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWReplenishList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Create create replenish
func (r *Replenish) Create(args *Replenish) (*Replenish, error) {
	action := "sales.replenish.create"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Replenish{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Cancel create cancel
func (r *Replenish) Cancel(args *Replenish) (*Replenish, error) {
	action := "sales.replenish.cancel"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Replenish{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
