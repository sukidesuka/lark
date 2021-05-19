// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateBitableRecord 该接口用于更新数据表中的一条记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/update
func (r *BitableService) UpdateBitableRecord(ctx context.Context, request *UpdateBitableRecordReq, options ...MethodOptionFunc) (*UpdateBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableUpdateBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#UpdateBitableRecord mock enable")
		return r.cli.mock.mockBitableUpdateBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:        "Bitable",
		API:          "UpdateBitableRecord",
		Method:       "PUT",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(updateBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableUpdateBitableRecord(f func(ctx context.Context, request *UpdateBitableRecordReq, options ...MethodOptionFunc) (*UpdateBitableRecordResp, *Response, error)) {
	r.mockBitableUpdateBitableRecord = f
}

func (r *Mock) UnMockBitableUpdateBitableRecord() {
	r.mockBitableUpdateBitableRecord = nil
}

type UpdateBitableRecordReq struct {
	UserIDType *IDType                `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	AppToken   string                 `path:"app_token" json:"-"`     // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID    string                 `path:"table_id" json:"-"`      // table id, 示例值："tblsRc9GRRXKqhvW"
	RecordID   string                 `path:"record_id" json:"-"`     // 单条记录的 id, 示例值："recqwIwhc6"
	Fields     map[string]interface{} `json:"fields,omitempty"`       // 记录字段
}

type updateBitableRecordResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateBitableRecordResp `json:"data,omitempty"` //
}

type UpdateBitableRecordResp struct {
	Record *UpdateBitableRecordRespRecord `json:"record,omitempty"` // {,    "fields": {,        "人力评估": 2,,        "任务执行人": [,            {,                "id": "ou_debc524b2d8cb187704df652b43d29de",            },        ],,        "任务描述": "多渠道收集用户反馈",,        "对应 OKR": [,            "recqwIwhc6",,            "recOuEJMvN",        ],,        "截止日期": 1609516800000,,        "是否完成": true,,        "状态": "已结束",,        "相关部门": [,            "销售",,            "客服",        ],    },}
}

type UpdateBitableRecordRespRecord struct {
	RecordID string                 `json:"record_id,omitempty"` // 记录 id
	Fields   map[string]interface{} `json:"fields,omitempty"`    // 记录字段
}