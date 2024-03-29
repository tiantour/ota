package order

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MemoItem memo item
	MemoItem struct {
		ID       int32  `json:"id"`        // 备注id
		OrderID  string `json:"order_id"`  // 旅行商城业务订单号
		AdminUID int32  `json:"admin_uid"` // 备注添加人UID
		Content  string `json:"content"`   // 备注内容
		Ctime    string `json:"ctime"`     // 备注时间
	}
)

// Memo memo
type Memo struct {
	OrderID string `json:"order_id,omitempty"` // 订单号
	Memo    string `json:"memo,omitempty"`     // 备注信息
}

// NewMemo new memo
func NewMemo() *Memo {
	return &Memo{}
}

// Item get memo item
func (m *Memo) Item(orderID string) (*MemoItem, error) {
	data, err := json.Marshal(&Memo{
		OrderID: orderID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.order.memo.get"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := MemoItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Add add memo
func (m *Memo) Add(args *Memo) (*Memo, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	action := "sales.order.memo.add"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Memo{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
