package sku

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWCalendarList calendar list
	MFWCalendarList struct {
		SkuID         int32              `json:"sku_id,omitempty"`         // 马蜂窝sku_id，马蜂窝sku_id与ota_sku_id商家编码二选一
		OtaSkuID      string             `json:"ota_sku_id,omitempty"`     // 商家设置的SKU外部编码，sku_id与ota_sku_id二选一
		CalendarItems []*MFWCalendarItem `json:"calendar_items,omitempty"` // 日历库存每天数据，天数最多180天，单次最多540条数据（目前单天数据对应最多的费用项为3个：成人、儿童、单房差）
	}
	// MFWCalendarItem calendar item
	MFWCalendarItem struct {
		Date      string  `json:"date,omitempty"`       // 需要更新的日期
		PriceType int32   `json:"price_type,omitempty"` // 费用类型 具体说明请见 费用类型说明表
		Price     float64 `json:"price,omitempty"`      // 售卖价格
		Remain    int32   `json:"remain,omitempty"`     // 日历库存余量
		MinNum    int32   `json:"min_num,omitempty"`    // 最小购买数
		MaxNum    int32   `json:"max_num,omitempty"`    // 最大购买数
	}
)

// Calendar calendar
type Calendar struct {
	SkuID     int32    `json:"sku_id,omitempty"`     // 马蜂窝sku_id，马蜂窝sku_id与ota_sku_id商家编码二选一
	OtaSkuID  string   `json:"ota_sku_id,omitempty"` // 商家设置的SKU外部编码，sku_id与ota_sku_id二选一
	Date      []string `json:"date,omitempty"`       // 需要更新的日期
	PriceType int32    `json:"price_type,omitempty"` // 费用类型 具体说明请见 费用类型说明表
	Price     float64  `json:"price,omitempty"`      // 售卖价格
	Remain    int32    `json:"remain,omitempty"`     // 日历库存余量
	MinNum    int32    `json:"min_num,omitempty"`    // 最小购买数
	MaxNum    int32    `json:"max_num,omitempty"`    // 最大购买数
}

// NewCalendar new calendar
func NewCalendar() *Calendar {
	return &Calendar{}
}

// Update update calendar once
func (c *Calendar) Update(args *Calendar) (*Calendar, error) {
	action := "sales.sku.calendar.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Calendar{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Batch update calendar batch
func (c *Calendar) Batch(args *MFWCalendarList) (*Calendar, error) {
	action := "sales.sku.calendar.batch.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Calendar{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Override update calendar override
func (c *Calendar) Override(args []*MFWCalendarList) (*Calendar, error) {
	action := "sales.sku.calendar.override"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Calendar{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
