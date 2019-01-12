package sku

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWPriceItem price item
	MFWPriceItem struct {
		Calendar map[string]map[string]*MFWPriceInfo `json:"calendar"` // 日历库存
		Stable   map[string]*MFWPriceInfo            `json:"stable"`   // 非日历库存
	}

	// MFWPriceInfo price info
	MFWPriceInfo struct {
		PriceSettle float64 `json:"price_settle"` // 结算价
		Remain      int32   `json:"remain"`       // 日历库存余量
		MinNum      int32   `json:"min_num"`      // 最小购买数
		MaxNum      int32   `json:"max_num"`      // 最大购买数
	}
)

// Price price
type Price struct {
	SkuID     int32   `json:"sku_id,omitempty"`     // 马蜂窝sku_id，马蜂窝sku_id与ota_sku_id商家编码二选一
	OtaSkuID  string  `json:"ota_sku_id,omitempty"` // 商家设置的SKU外部编码，sku_id与ota_sku_id二选一
	PriceType int32   `json:"price_type,omitempty"` // 费用类型 具体说明请见 费用类型说明表
	Price     float64 `json:"price,omitempty"`      // 售卖价格
	Remain    int32   `json:"remain,omitempty"`     // 日历库存余量
	MinNum    int32   `json:"min_num,omitempty"`    // 最小购买数
	MaxNum    int32   `json:"max_num,omitempty"`    // 最大购买数
}

// NewPrice new price
func NewPrice() *Price {
	return &Price{}
}

// Item get sku price item
func (p *Price) Item(skuID int32) (*MFWPriceItem, error) {
	action := "sales.sku.price.get"
	data, err := json.Marshal(&Price{
		SkuID: skuID,
	})
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWPriceItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Update update sku price
func (p *Price) Update(args *Price) (*MFWPriceItem, error) {
	action := "sales.sku.price.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWPriceItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
