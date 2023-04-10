package sku

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// SkuItem sku item
	SkuItem struct {
		SkuID          int32  `json:"sku_id"`           // 马蜂窝sku_id
		OtaSkuID       string `json:"ota_sku_id"`       // 商家设置的SKU外部编码
		SalesID        int32  `json:"sales_id"`         // 马蜂窝产品id，产品唯一标识
		SalesName      string `json:"sales_name"`       // 产品名称
		SkuName        string `json:"sku_name"`         // sku名称
		SkuCtime       string `json:"sku_ctime"`        // sku创建时间
		SkuMtime       string `json:"sku_mtime"`        // sku修改时间
		SkuOnline      int32  `json:"sku_online"`       // 上下线状态 0：下线 1：上线
		Del            int32  `json:"del"`              // 删除状态 0：未删除 1：已删除
		RangeCleanDay  int32  `json:"range_clean_day"`  // 提前预订天数
		RangeCleanTime string `json:"range_clean_time"` // 提前预订时间
	}
)

// Sku sku
type Sku struct {
	SkuID          int32  `json:"sku_id,omitempty"`           // 马蜂窝sku_id，马蜂窝sku_id与ota_sku_id商家编码二选一
	OtaSkuID       string `json:"ota_sku_id,omitempty"`       // 商家设置的SKU外部编码，sku_id与ota_sku_id二选一
	RangeCleanDay  int32  `json:"range_clean_day,omitempty"`  // 提前预订天数
	RangeCleanTime string `json:"range_clean_time,omitempty"` // 提前预订时间
}

// NewSku new sku
func NewSku() *Sku {
	return &Sku{}
}

// Item get sku item
func (s *Sku) Item(skuID int32) (*SkuItem, error) {
	action := "sales.sku.info.get"
	data, err := json.Marshal(&Sku{
		SkuID: skuID,
	})
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := SkuItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Time set sku time
func (s *Sku) Time(args *Sku) (*SkuItem, error) {
	action := "sales.sku.bookingtime.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := SkuItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// OnLine set sku online
func (s *Sku) OnLine(skuID int32) (*SkuItem, error) {
	action := "sales.sku.online"
	data, err := json.Marshal(&Sku{
		SkuID: skuID,
	})
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := SkuItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// OffLine set sku offline
func (s *Sku) OffLine(skuID int32) (*SkuItem, error) {
	action := "sales.sku.offline"
	data, err := json.Marshal(&Sku{
		SkuID: skuID,
	})
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := SkuItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Delete set sku delete
func (s *Sku) Delete(skuID int32) (*SkuItem, error) {
	action := "sales.sku.delete"
	data, err := json.Marshal(&Sku{
		SkuID: skuID,
	})
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := SkuItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
