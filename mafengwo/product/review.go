package product

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

// Review review
type Review struct {
	SalesID int32 `json:"sales_id,omitempty"` // 产品id
	Status  int32 `json:"status,omitempty"`   // 产品状态，1-商品已在线，2-未上线，未提交上线申请，3-未上线，审核中，4-未上线，提交申请被驳回，5-已上线，提交申请被通过
}

// NewReview new review
func NewReview() *Review {
	return &Review{}
}

// Item get review item
func (r *Review) Item(salesID int32) (*Review, error) {
	data, err := json.Marshal(&Review{
		SalesID: salesID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.product.review.status.get"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Review{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Apply apply review
func (r *Review) Apply(salesID int32) (*Review, error) {
	data, err := json.Marshal(&Review{
		SalesID: salesID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.product.online.review.apply"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Review{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Cancel review cancel
func (r *Review) Cancel(salesID int32) (*Review, error) {
	data, err := json.Marshal(&Review{
		SalesID: salesID,
	})
	if err != nil {
		return nil, err
	}

	action := "sales.product.online.review.cancel"
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}

	result := Review{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
