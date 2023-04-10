package shop

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// ShopList get shop list
	ShopList struct {
		Total int32       `json:"total"`
		List  []*ShopItem `json:"list"`
	}

	// ShopItem get shop item
	ShopItem struct {
		SalesID       string  `json:"sales_id,omitempty"`        // 产品id
		SalesName     string  `json:"sales_name,omitempty"`      // 产品名称
		SalesType     string  `json:"sales_type,omitempty"`      // 品类ID
		SalesTypeName string  `json:"sales_type_name,omitempty"` // 品类名称
		MddName       string  `json:"mdd_name,omitempty"`        // 目的地名称：产品关联目的地名称
		StartDate     string  `json:"start_date,omitempty"`      // 开始日期
		EndDate       string  `json:"end_date,omitempty"`        // 开始日期	结束日期
		OnlinePrdNum  int32   `json:"online_prd_num,omitempty"`  // 统计周期内，在线产品数量
		PrdUv         int32   `json:"prd_uv"`                    // 统计周期内，所有产品详情页的访客数去重
		PrdPv         int32   `json:"prd_pv"`                    // 统计周期内，所有产品详情页的浏览流累计
		PaidUIDNum    int32   `json:"paid_uid_num"`              // 统计周期内，所有支付订单的用户去重后的支付人数
		PaidOrdNum    int32   `json:"paid_ord_num"`              // 支付订单数：统计周期内，支付订单总数
		PaidStockNum  int32   `json:"paid_stock_num"`            // 支付件数：统计周期内，所有支付订单的支付祝库存件数累计
		PaidGmv       float64 `json:"paid_gmv"`                  // 支付金额：统计周期内，所有支付订单的订单金额累计（包含订单优惠金额及退款金额）
		AvgUIDGmv     float64 `json:"avg_uid_gmv"`               // 客单价：统计周期内，支付订单金额/支付人数
		AvgOrdGmv     float64 `json:"avg_ord_gmv"`               // 均单价：统计周期内，支付订单金额/支付订单数
		AvgStockGmv   float64 `json:"avg_stock_gmv"`             // 库存单价：统计周期内，支付订单金额/支付件数
		Reduce        float64 `json:"reduce_"`                   // 马蜂窝优惠金额：统计周期内，支付订单的马蜂窝优惠金额累计（马蜂窝补贴）
		ReduceSum     float64 `json:"reduce_sum"`                // 优惠金额：统计周期内，支付订单的优惠金额累计
		ReduceOta     float64 `json:"reduce_ota"`                // 商家优惠金额：统计周期内，支付订单的商家优惠金额累计（商家自补贴）
		RefundSum     float64 `json:"refund_sum"`                // 退款总额：统计周期内，所有产生退款订单的退款金额累计
	}
)

// Shop shop
type Shop struct {
	StartTime string  `json:"start_time,omitempty"` // 查询开始日期
	EndTime   string  `json:"end_time,omitempty"`   // 查询结束时间
	Type      int32   `json:"type,omitempty"`       // 汇总周期类型：“3-自然天”／“1-自然周”／“2-自然月”
	PageNo    int32   `json:"page_no,omitempty"`    // 当前页数
	PageSize  int32   `json:"page_size,omitempty"`  // 单页条数（最大值20）
	SalesID   []int32 `json:"sales_id,omitempty"`   // 产品ID
}

// NewShop new shop
func NewShop() *Shop {
	return &Shop{}
}

// Product get shop product data
func (s *Shop) Product(start, end string, types, page, size int32) (*ShopList, error) {
	action := "sales.shop.product.trade.get"
	data, err := json.Marshal(&Shop{
		StartTime: start,
		EndTime:   end,
		Type:      types,
		PageNo:    page,
		PageSize:  size,
	})
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := ShopList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Trade get shop trade data
func (s *Shop) Trade(start, end string, types, page, size int32) (*ShopList, error) {
	action := "sales.shop.trade.get"
	data, err := json.Marshal(&Shop{
		StartTime: start,
		EndTime:   end,
		Type:      types,
		PageNo:    page,
		PageSize:  size,
	})
	if err != nil {
		return nil, err
	}

	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := ShopList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
