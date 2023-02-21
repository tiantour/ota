package order

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type Confirm struct {
	OrderID string `json:"order_id,omitempty"` // 订单号
	Status  string `json:"status,omitempty"`   // 订单状态，0待确认，12，已确认，13，已出单
}

func NewConfirm() *Confirm {
	return &Confirm{}
}

// Item get confirm item
func (c *Confirm) Item(orderID string) (*Confirm, error) {
	data, err := json.Marshal(&Confirm{
		OrderID: orderID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.order.confirm.status"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Confirm{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
