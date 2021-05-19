// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBitableFieldList 根据 app_token 和 table_id，获取数据表的所有字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list
func (r *BitableService) GetBitableFieldList(ctx context.Context, request *GetBitableFieldListReq, options ...MethodOptionFunc) (*GetBitableFieldListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableFieldList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableFieldList mock enable")
		return r.cli.mock.mockBitableGetBitableFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:        "Bitable",
		API:          "GetBitableFieldList",
		Method:       "GET",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(getBitableFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableGetBitableFieldList(f func(ctx context.Context, request *GetBitableFieldListReq, options ...MethodOptionFunc) (*GetBitableFieldListResp, *Response, error)) {
	r.mockBitableGetBitableFieldList = f
}

func (r *Mock) UnMockBitableGetBitableFieldList() {
	r.mockBitableGetBitableFieldList = nil
}

type GetBitableFieldListReq struct {
	ViewID    *string `query:"view_id" json:"-"`    // 视图 ID, 示例值："vewOVMEXPF"
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："fldwJ4YrtB"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`100`
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID   string  `path:"table_id" json:"-"`    // table id, 示例值："tblsRc9GRRXKqhvW"
}

type getBitableFieldListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableFieldListResp `json:"data,omitempty"` //
}

type GetBitableFieldListResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetBitableFieldListRespItem `json:"items,omitempty"`      // 字段信息
}

type GetBitableFieldListRespItem struct {
	FieldID   string      `json:"field_id,omitempty"`   // 多维表格字段 id
	FieldName string      `json:"field_name,omitempty"` // 多维表格字段名
	Type      int64       `json:"type,omitempty"`       // 多维表格字段类型
	Property  interface{} `json:"property,omitempty"`   // 字段属性
}