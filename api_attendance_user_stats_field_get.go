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

// GetAttendanceUserStatsField 查询考勤统计支持的日度统计或月度统计的统计表头。
//
// 调用统计开放接口api目前不返回休假加班新增字段类型
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_field/query
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_stats_data/query-2
func (r *AttendanceService) GetAttendanceUserStatsField(ctx context.Context, request *GetAttendanceUserStatsFieldReq, options ...MethodOptionFunc) (*GetAttendanceUserStatsFieldResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserStatsField != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserStatsField mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserStatsField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserStatsField",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_stats_fields/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserStatsFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserStatsField mock AttendanceGetAttendanceUserStatsField method
func (r *Mock) MockAttendanceGetAttendanceUserStatsField(f func(ctx context.Context, request *GetAttendanceUserStatsFieldReq, options ...MethodOptionFunc) (*GetAttendanceUserStatsFieldResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserStatsField = f
}

// UnMockAttendanceGetAttendanceUserStatsField un-mock AttendanceGetAttendanceUserStatsField method
func (r *Mock) UnMockAttendanceGetAttendanceUserStatsField() {
	r.mockAttendanceGetAttendanceUserStatsField = nil
}

// GetAttendanceUserStatsFieldReq ...
type GetAttendanceUserStatsFieldReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 响应体中的 user_id 的员工工号类型, 示例值: employee_id, 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	Locale       string       `json:"locale,omitempty"`        // 语言类型, 示例值: "zh", 可选值有: en: 英语, ja: 日语, zh: 中文
	StatsType    string       `json:"stats_type,omitempty"`    // 统计类型, 示例值: "daily", 可选值有: daily: 日度统计, month: 月度统计
	StartDate    int64        `json:"start_date,omitempty"`    // 开始时间, 示例值: 20210316
	EndDate      int64        `json:"end_date,omitempty"`      // 结束时间（时间间隔不超过 40 天）, 示例值: 20210323
}

// GetAttendanceUserStatsFieldResp ...
type GetAttendanceUserStatsFieldResp struct {
	UserStatsField *GetAttendanceUserStatsFieldRespUserStatsField `json:"user_stats_field,omitempty"` // 统计数据表头
}

// GetAttendanceUserStatsFieldRespUserStatsField ...
type GetAttendanceUserStatsFieldRespUserStatsField struct {
	StatsType string                                                `json:"stats_type,omitempty"` // 统计类型, 可选值有: daily: 日度统计, month: 月度统计
	UserID    string                                                `json:"user_id,omitempty"`    // 用户 ID
	Fields    []*GetAttendanceUserStatsFieldRespUserStatsFieldField `json:"fields,omitempty"`     // 字段列表
}

// GetAttendanceUserStatsFieldRespUserStatsFieldField ...
type GetAttendanceUserStatsFieldRespUserStatsFieldField struct {
	Code        string                                                          `json:"code,omitempty"`         // 字段编号
	Title       string                                                          `json:"title,omitempty"`        // 字段名称
	ChildFields []*GetAttendanceUserStatsFieldRespUserStatsFieldFieldChildField `json:"child_fields,omitempty"` // 子字段列表
}

// GetAttendanceUserStatsFieldRespUserStatsFieldFieldChildField ...
type GetAttendanceUserStatsFieldRespUserStatsFieldFieldChildField struct {
	Code     string `json:"code,omitempty"`      // 子字段编号
	Title    string `json:"title,omitempty"`     // 子字段名称
	TimeUnit string `json:"time_unit,omitempty"` // 时间单位(该字段已停止使用)
}

// getAttendanceUserStatsFieldResp ...
type getAttendanceUserStatsFieldResp struct {
	Code  int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                           `json:"msg,omitempty"`  // 错误描述
	Data  *GetAttendanceUserStatsFieldResp `json:"data,omitempty"`
	Error *ErrorDetail                     `json:"error,omitempty"`
}
