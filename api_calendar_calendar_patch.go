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

// UpdateCalendar
//
// 该接口用于以当前身份（应用 / 用户）修改日历信息。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份对日历有 owner 权限时，可修改全局字段：summary, description, permission。
// 当前身份对日历不具有 owner 权限时，仅可修改对自己生效的字段：color, summary_alias。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/patch
func (r *CalendarService) UpdateCalendar(ctx context.Context, request *UpdateCalendarReq, options ...MethodOptionFunc) (*UpdateCalendarResp, *Response, error) {
	if r.cli.mock.mockCalendarUpdateCalendar != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UpdateCalendar mock enable")
		return r.cli.mock.mockCalendarUpdateCalendar(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "UpdateCalendar",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateCalendarResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarUpdateCalendar mock CalendarUpdateCalendar method
func (r *Mock) MockCalendarUpdateCalendar(f func(ctx context.Context, request *UpdateCalendarReq, options ...MethodOptionFunc) (*UpdateCalendarResp, *Response, error)) {
	r.mockCalendarUpdateCalendar = f
}

// UnMockCalendarUpdateCalendar un-mock CalendarUpdateCalendar method
func (r *Mock) UnMockCalendarUpdateCalendar() {
	r.mockCalendarUpdateCalendar = nil
}

// UpdateCalendarReq ...
type UpdateCalendarReq struct {
	CalendarID   string              `path:"calendar_id" json:"-"`    // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	Summary      *string             `json:"summary,omitempty"`       // 日历标题, 示例值："测试日历", 长度范围：`1` ～ `255` 字符
	Description  *string             `json:"description,omitempty"`   // 日历描述, 示例值："使用开放接口创建日历", 最大长度：`255` 字符
	Permissions  *CalendarPermission `json:"permissions,omitempty"`   // 日历公开范围, 示例值："private", 可选值有: `private`：私密, `show_only_free_busy`：仅展示忙闲信息, `public`：他人可查看日程详情
	Color        *int64              `json:"color,omitempty"`         // 日历颜色，颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效, 示例值：-1
	SummaryAlias *string             `json:"summary_alias,omitempty"` // 日历备注名，修改或添加后仅对当前身份生效, 示例值："日历备注名", 最大长度：`255` 字符
}

// updateCalendarResp ...
type updateCalendarResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCalendarResp `json:"data,omitempty"`
}

// UpdateCalendarResp ...
type UpdateCalendarResp struct {
	Calendar *UpdateCalendarRespCalendar `json:"calendar,omitempty"` // 更新后的日历实体
}

// UpdateCalendarRespCalendar ...
type UpdateCalendarRespCalendar struct {
	CalendarID   string             `json:"calendar_id,omitempty"`    // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction)
	Summary      string             `json:"summary,omitempty"`        // 日历标题
	Description  string             `json:"description,omitempty"`    // 日历描述
	Permissions  CalendarPermission `json:"permissions,omitempty"`    // 日历公开范围, 可选值有: `private`：私密, `show_only_free_busy`：仅展示忙闲信息, `public`：他人可查看日程详情
	Color        int64              `json:"color,omitempty"`          // 日历颜色，颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效
	Type         CalendarType       `json:"type,omitempty"`           // 日历类型, 可选值有: `unknown`：未知类型, `primary`：用户或应用的主日历, `shared`：由用户或应用创建的共享日历, `google`：用户绑定的谷歌日历, `resource`：会议室日历, `exchange`：用户绑定的Exchange日历
	SummaryAlias string             `json:"summary_alias,omitempty"`  // 日历备注名，修改或添加后仅对当前身份生效
	IsDeleted    bool               `json:"is_deleted,omitempty"`     // 对于当前身份，日历是否已经被标记为删除
	IsThirdParty bool               `json:"is_third_party,omitempty"` // 当前日历是否是第三方数据；三方日历及日程只支持读，不支持写入
	Role         CalendarRole       `json:"role,omitempty"`           // 当前身份对于该日历的访问权限, 可选值有: `unknown`：未知权限, `free_busy_reader`：游客，只能看到忙碌/空闲信息, `reader`：订阅者，查看所有日程详情, `writer`：编辑者，创建及修改日程, `owner`：管理员，管理日历及共享设置
}
