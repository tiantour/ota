package order

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (
	// MFWCommentList comment list
	MFWCommentList struct {
		Total int32             `json:"total"` // 点评数量
		List  []*MFWCommentItem // 点评信息列表
	}
	// MFWCommentItem comment item
	MFWCommentItem struct {
		ID        int32           `json:"id "`        // 点评ID
		OrderID   string          `json:"order_id"`   // 旅行商城订单号
		UserName  string          `json:"username"`   // 用户昵称
		UID       int32           `json:"uid"`        // 点评用户的UID
		UserPhone string          `json:"user_phone"` // 点评用户的手机号
		Ctime     string          `json:"ctime"`      // 点评首次提交时间
		Mtime     string          `json:"mtime"`      // 点评最后一次更新时间
		Star      int32           `json:"star"`       // 1，2 表示差评 3表示中评 4，5表示好评
		SubStar   []*MFWSubStar   `json:"sub_star"`   // 维度评分信息
		StarTags  []*MFWStarTags  `json:"star_tags"`  // 点评关联的标签
		Content   string          `json:"content"`    // 点评内容
		Reply     string          `json:"reply"`      // 商家关于点评的回复
		SalesID   int32           `json:"sales_id"`   // 马蜂窝产品id，产品唯一标识
		SalesName string          `json:"sales_name"` // 产品名称
		ImgCount  int32           `json:"img_count"`  // 用户在点评中上传的照片数量
		ImageInfo []*MFWImageInfo `json:"image_info"` // 	用户在点评中上传的图片信息
		Status    int32           `json:"status"`     // 点评状态 0是待审核 1是审核通过
	}

	// MFWSubStar sub star
	MFWSubStar struct {
		Alias string `json:"alias"` // 点评维度名称
		Star  int32  `json:"star"`  // 维度评价得分 1，2 表示差评 3表示中评 4，5表示好评
	}

	// MFWStarTags star tag
	MFWStarTags struct {
		Name string `json:"name"` // 点评标签名称
	}

	// MFWImageInfo image info
	MFWImageInfo struct {
		ImgSrc string `json:"img_src"` // 点评照片链接
	}
)

// Comment comment
type Comment struct {
	StartTime       string  `json:"start_time,omitempty"`        // 点评开始时间
	EndTime         string  `json:"end_time,omitempty"`          // 点评结束时间
	SalesID         int32   `json:"sales_id,omitempty"`          // 马蜂窝产品id，产品唯一标识
	OrderID         string  `json:"order_id,omitempty"`          // 旅行商城业务订单号
	ModifyStartTime string  `json:"modify_start_time,omitempty"` // 点评更新开始时间
	ModifyEndTime   string  `json:"modify_end_time,omitempty"`   // 点评更新结束时间
	PageNo          int32   `json:"page_no,omitempty"`           // 当前页数
	PageSize        int32   `json:"page_size,omitempty"`         // 单页条数（最大值20） 超过20默认选择20
	Star            []int32 `json:"star,omitempty"`              // 点评得分
	CommentID       int32   `json:"comment_id,omitempty"`        // 点评ID
	Content         string  `json:"content,omitempty"`           // 回复内容
}

// NewComment new comment
func NewComment() *Comment {
	return &Comment{}
}

// List get comment list
func (c *Comment) List(args *Comment) (*MFWCommentList, error) {
	action := "sales.order.comment.list.get"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWCommentList{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Reply set comment reply
func (c *Comment) Reply(args *Comment) (*Comment, error) {
	action := "sales.order.comment.reply"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Comment{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
