// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetVCTopUserReport 获取一段时间内组织内会议使用的 Top 用户列表。
//
// 支持最近90天内的数据查询；默认返回前10位, 最多可查询前100位
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_top_user
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/report/get_top_user
func (r *VCService) GetVCTopUserReport(ctx context.Context, request *GetVCTopUserReportReq, options ...MethodOptionFunc) (*GetVCTopUserReportResp, *Response, error) {
	if r.cli.mock.mockVCGetVCTopUserReport != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#GetVCTopUserReport mock enable")
		return r.cli.mock.mockVCGetVCTopUserReport(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCTopUserReport",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reports/get_top_user",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getVCTopUserReportResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCTopUserReport mock VCGetVCTopUserReport method
func (r *Mock) MockVCGetVCTopUserReport(f func(ctx context.Context, request *GetVCTopUserReportReq, options ...MethodOptionFunc) (*GetVCTopUserReportResp, *Response, error)) {
	r.mockVCGetVCTopUserReport = f
}

// UnMockVCGetVCTopUserReport un-mock VCGetVCTopUserReport method
func (r *Mock) UnMockVCGetVCTopUserReport() {
	r.mockVCGetVCTopUserReport = nil
}

// GetVCTopUserReportReq ...
type GetVCTopUserReportReq struct {
	StartTime  string  `query:"start_time" json:"-"`   // 开始时间（unix时间, 单位sec）, 示例值: 1608888867
	EndTime    string  `query:"end_time" json:"-"`     // 结束时间（unix时间, 单位sec）, 示例值: 1608889966
	Limit      int64   `query:"limit" json:"-"`        // 取前多少位, 示例值: 10
	OrderBy    int64   `query:"order_by" json:"-"`     // 排序依据（降序）, 示例值: 1, 可选值有: 1: 会议数量, 2: 会议时长
	Unit       *int64  `query:"unit" json:"-"`         // 数据驻留地（传参前提是租户存在多个驻留地数据且开通了该查询功能）, 示例值: 0, 可选值有: 0: 中国大陆, 1: 美国, 2: 新加坡, 3: 日本
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCTopUserReportResp ...
type GetVCTopUserReportResp struct {
	TopUserReport []*GetVCTopUserReportRespTopUserReport `json:"top_user_report,omitempty"` // top用户列表
}

// GetVCTopUserReportRespTopUserReport ...
type GetVCTopUserReportRespTopUserReport struct {
	ID              string `json:"id,omitempty"`               // 用户ID
	Name            string `json:"name,omitempty"`             // 用户名
	UserType        int64  `json:"user_type,omitempty"`        // 用户类型, 可选值有: 1: 飞书用户, 2: rooms用户, 3: 文档用户, 4: neo单品用户, 5: neo单品游客用户, 6: pstn用户, 7: sip用户
	MeetingCount    string `json:"meeting_count,omitempty"`    // 会议数量
	MeetingDuration string `json:"meeting_duration,omitempty"` // 会议时长（单位min）
}

// getVCTopUserReportResp ...
type getVCTopUserReportResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *GetVCTopUserReportResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
