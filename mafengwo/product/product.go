package product

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWProductList product list
	MFWProductList struct {
		Total int32             `json:"total"`
		List  []*MFWProductItem `json:"list"`
	}

	// MFWProductItem product item
	MFWProductItem struct {
		SalesID       int32     `json:"sales_id"`        // 产品id
		SalesName     string    `json:"sales_name"`      // 一日游	产品名称
		SalesType     int32     `json:"sales_type"`      // 品类id（二级品类）
		SalesTypeName string    `json:"sales_type_name"` // 景区/场馆	品类名称（二级品类）
		SalesStatus   int32     `json:"sales_status"`    // 产品状态，主要包括：0-未上线，1-已上线，2-人工下线，3-系统下线，4-审核下线；
		SkuList       []*MFWSku `json:"sku_list"`        // 产品关联的sku信息
	}
	// MFWSku sku
	MFWSku struct {
		SkuID          int32  `json:"sku_id"`           // 马蜂窝SKU ID，SKU唯一标识
		OtaSkuID       string `json:"ota_sku_id"`       // 商家设置的SKU外部编码
		SkuName        string `json:"sku_name"`         // 门票 成人票	SKU名称
		RangeCleanDay  int32  `json:"range_clean_day"`  // 提前预订天数，与提前预订时间配合使用；
		RangeCleanTime string `json:"range_clean_time"` // 提前预订时间，如range_clean_day＝0，range_clean_time＝10:00:00，表示当天10点前SKU可订
	}
)

// Product product
type Product struct {
	PageNo      int32  `json:"page_no,omitempty"`      // 当前页数
	PageSize    int32  `json:"page_size,omitempty"`    // 单页条数（最大值20）
	SalesStatus int32  `json:"sales_status,omitempty"` // 产品状态，主要包括：0-未上线，1-已上线，2-人工下线，3-系统下线，4-审核下线；不传此参数默认返回全部
	OtaSalesID  string `json:"ota_sales_id,omitempty"` // 商家设置的外部商家产品编码
	SalesID     int32  `json:"sales_id,omitempty"`     // 产品id
}

// NewProduct new product
func NewProduct() *Product {
	return &Product{}
}

// List get product list
func (p *Product) List(page, size int32) (*MFWProductList, error) {
	action := "sales.product.list.get"
	data, err := json.Marshal(&Product{
		PageNo:   page,
		PageSize: size,
	})
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWProductList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Item get product item
func (p *Product) Item(salesID int32) (*MFWProductItem, error) {
	action := "sales.product.detail.get"
	data, err := json.Marshal(&Product{
		SalesID: salesID,
	})
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWProductItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
