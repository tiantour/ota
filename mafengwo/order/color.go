package order

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

// Color color
type Color struct {
	OrderID string `json:"order_id,omitempty"` // 订单号
	Color   string `json:"color,omitempty"`    // 订单标记的颜色 0-白色,1-红色,2-橙色,3-黄色,4-绿色,5-蓝色,6-紫色,7-灰色
}

// NewColor new color
func NewColor() *Color {
	return &Color{}
}

// Item get color item
func (c *Color) Item(orderID string) (*Color, error) {
	data, err := json.Marshal(&Color{
		OrderID: orderID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.order.color.get"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Color{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Add add color
func (c *Color) Add(args *Color) (*Color, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	action := "sales.order.color.add"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Color{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
