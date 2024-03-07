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

// GetCalendarEventInstanceViewList 该接口用于以用户身份查询某日历下的日程视图（重复性日程展开）。
//
// 身份由 Header Authorization 的 Token 类型决定。
// - 当前身份必须对日历有reader、writer或owner权限才会返回日程详细信息（调用获取日历接口, role字段可查看权限）。
// - 根据日历ID获取该日历下的日程视图。仅支持primary、shared类型的日历, 暂不支持google、exchange类型的日历。
// - 本接口与查询日程列表的区别在于对重复性日程的处理。本接口会将重复性日程根据重复性规则展开成日程实例, 根据查询的时间区间返回相应日程的实例
// - 查询日程视图的时间跨度需要小于40天
// - 查询日程视图的日程实例数量需要小于1000
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/instance_view
func (r *CalendarService) GetCalendarEventInstanceViewList(ctx context.Context, request *GetCalendarEventInstanceViewListReq, options ...MethodOptionFunc) (*GetCalendarEventInstanceViewListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEventInstanceViewList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEventInstanceViewList mock enable")
		return r.cli.mock.mockCalendarGetCalendarEventInstanceViewList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEventInstanceViewList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/instance_view",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventInstanceViewListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarEventInstanceViewList mock CalendarGetCalendarEventInstanceViewList method
func (r *Mock) MockCalendarGetCalendarEventInstanceViewList(f func(ctx context.Context, request *GetCalendarEventInstanceViewListReq, options ...MethodOptionFunc) (*GetCalendarEventInstanceViewListResp, *Response, error)) {
	r.mockCalendarGetCalendarEventInstanceViewList = f
}

// UnMockCalendarGetCalendarEventInstanceViewList un-mock CalendarGetCalendarEventInstanceViewList method
func (r *Mock) UnMockCalendarGetCalendarEventInstanceViewList() {
	r.mockCalendarGetCalendarEventInstanceViewList = nil
}

// GetCalendarEventInstanceViewListReq ...
type GetCalendarEventInstanceViewListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"`   // 日历ID, 示例值: "feishu.cn_HF9U2MbibE8PPpjro6xjqa@group.calendar.feishu.cn"
	StartTime  string  `query:"start_time" json:"-"`   // 日程开始Unix时间戳, 单位为秒, 起止时间跨度小于40天, 示例值: 1631777271
	EndTime    string  `query:"end_time" json:"-"`     // 日程结束Unix时间戳, 单位为秒, 起止时间跨度小于40天, 示例值: 1631777271
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetCalendarEventInstanceViewListResp ...
type GetCalendarEventInstanceViewListResp struct {
	Items []*GetCalendarEventInstanceViewListRespItem `json:"items,omitempty"` // 日程instance列表
}

// GetCalendarEventInstanceViewListRespItem ...
type GetCalendarEventInstanceViewListRespItem struct {
	EventID             string                                                  `json:"event_id,omitempty"`              // 日程实例ID
	Summary             string                                                  `json:"summary,omitempty"`               // 日程主题
	Description         string                                                  `json:"description,omitempty"`           // 日程描述
	StartTime           *GetCalendarEventInstanceViewListRespItemStartTime      `json:"start_time,omitempty"`            // 开始时间
	EndTime             *GetCalendarEventInstanceViewListRespItemEndTime        `json:"end_time,omitempty"`              // 结束时间
	Status              string                                                  `json:"status,omitempty"`                // 日程状态, 可选值有: tentative: 未回应, confirmed: 已确认, cancelled: 日程已取消
	IsException         bool                                                    `json:"is_exception,omitempty"`          // 是否是例外日程实例
	AppLink             string                                                  `json:"app_link,omitempty"`              // 日程的app_link, 跳转到具体的某个日程
	OrganizerCalendarID string                                                  `json:"organizer_calendar_id,omitempty"` // 日程组织者日历ID
	Vchat               *GetCalendarEventInstanceViewListRespItemVchat          `json:"vchat,omitempty"`                 // 视频会议信息, 仅当日程至少有一位attendee时生效
	Visibility          string                                                  `json:"visibility,omitempty"`            // 日程公开范围, 新建日程默认为Default；仅新建日程时对所有参与人生效, 之后修改该属性仅对当前身份生效, 可选值有: default: 默认权限, 仅向他人显示是否“忙碌”, public: 公开, 显示日程详情, private: 私密, 仅自己可见
	AttendeeAbility     string                                                  `json:"attendee_ability,omitempty"`      // 参与人权限, 可选值有: none: 无法编辑日程、无法邀请其它参与人、无法查看参与人列表, can_see_others: 无法编辑日程、无法邀请其它参与人、可以查看参与人列表, can_invite_others: 无法编辑日程、可以邀请其它参与人、可以查看参与人列表, can_modify_event: 可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus      string                                                  `json:"free_busy_status,omitempty"`      // 日程占用的忙闲状态, 新建日程默认为Busy；仅新建日程时对所有参与人生效, 之后修改该属性仅对当前身份生效, 可选值有: busy: 忙碌, free: 空闲
	Location            *GetCalendarEventInstanceViewListRespItemLocation       `json:"location,omitempty"`              // 日程地点
	Color               int64                                                   `json:"color,omitempty"`                 // 日程颜色, 颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	RecurringEventID    string                                                  `json:"recurring_event_id,omitempty"`    // 例外日程的原重复日程的event_id
	EventOrganizer      *GetCalendarEventInstanceViewListRespItemEventOrganizer `json:"event_organizer,omitempty"`       // 日程组织者信息
	Attendees           []*GetCalendarEventInstanceViewListRespItemAttendee     `json:"attendees,omitempty"`             // 日程参与人信息, 当前只返回会议室, 需要其他类型参与人信息请使用「获取日程参与人列表」
}

// GetCalendarEventInstanceViewListRespItemAttendee ...
type GetCalendarEventInstanceViewListRespItemAttendee struct {
	Type                  CalendarEventAttendeeType                                                `json:"type,omitempty"`                   // 参与人类型, 仅当新建参与人时可设置类型, type为User时, 值为open_id/user_id/union_id, type为Chat时, 值为open_chat_id, type为Resource时, 值为open_room_id, type为ThirdParty时, 值为third_party_email；不支持通过API新建该类型参与人, 可选值有: user: 用户, chat: 群组, resource: 会议室, third_party: 邮箱
	AttendeeID            string                                                                   `json:"attendee_id,omitempty"`            // 参与人ID
	RsvpStatus            string                                                                   `json:"rsvp_status,omitempty"`            // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional            bool                                                                     `json:"is_optional,omitempty"`            // 参与人是否为「可选参加」, 无法编辑群参与人的此字段
	IsOrganizer           bool                                                                     `json:"is_organizer,omitempty"`           // 参与人是否为日程组织者
	IsExternal            bool                                                                     `json:"is_external,omitempty"`            // 参与人是否为外部参与人；外部参与人不支持编辑
	DisplayName           string                                                                   `json:"display_name,omitempty"`           // 参与人名称
	ChatMembers           []*GetCalendarEventInstanceViewListRespItemAttendeeChatMember            `json:"chat_members,omitempty"`           // 群中的群成员, 当type为Chat时有效；群成员不支持编辑
	UserID                string                                                                   `json:"user_id,omitempty"`                // 参与人的用户id, 依赖于user_id_type返回对应的取值, 当is_external为true时, 此字段只会返回open_id或者union_id
	ChatID                string                                                                   `json:"chat_id,omitempty"`                // chat类型参与人的群组chat_id
	RoomID                string                                                                   `json:"room_id,omitempty"`                // resource类型参与人的会议室room_id
	ThirdPartyEmail       string                                                                   `json:"third_party_email,omitempty"`      // third_party类型参与人的邮箱
	OperateID             string                                                                   `json:"operate_id,omitempty"`             // bot身份操作时, 为预定的会议室指定实际预定人
	ResourceCustomization []*GetCalendarEventInstanceViewListRespItemAttendeeResourceCustomization `json:"resource_customization,omitempty"` // 会议室的个性化配置
	ApprovalReason        string                                                                   `json:"approval_reason,omitempty"`        // 会议室审批原因
}

// GetCalendarEventInstanceViewListRespItemAttendeeChatMember ...
type GetCalendarEventInstanceViewListRespItemAttendeeChatMember struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态, 可选值有: needs_action: 参与人尚未回复状态, 或表示会议室预约中, accept: 参与人回复接受, 或表示会议室预约成功, tentative: 参与人回复待定, decline: 参与人回复拒绝, 或表示会议室预约失败, removed: 参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}

// GetCalendarEventInstanceViewListRespItemAttendeeResourceCustomization ...
type GetCalendarEventInstanceViewListRespItemAttendeeResourceCustomization struct {
	IndexKey     string                                                                         `json:"index_key,omitempty"`     // 每个配置的唯一ID
	InputContent string                                                                         `json:"input_content,omitempty"` // 当type类型为填空时, 该参数有返回值
	Options      []*GetCalendarEventInstanceViewListRespItemAttendeeResourceCustomizationOption `json:"options,omitempty"`       // 每个配置的选项
}

// GetCalendarEventInstanceViewListRespItemAttendeeResourceCustomizationOption ...
type GetCalendarEventInstanceViewListRespItemAttendeeResourceCustomizationOption struct {
	OptionKey     string `json:"option_key,omitempty"`     // 每个选项的唯一ID
	OthersContent string `json:"others_content,omitempty"` // 当type类型为其它选项时, 该参数有返回值
}

// GetCalendarEventInstanceViewListRespItemEndTime ...
type GetCalendarEventInstanceViewListRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// GetCalendarEventInstanceViewListRespItemEventOrganizer ...
type GetCalendarEventInstanceViewListRespItemEventOrganizer struct {
	UserID      string `json:"user_id,omitempty"`      // 日程组织者user ID
	DisplayName string `json:"display_name,omitempty"` // 日程组织者姓名
}

// GetCalendarEventInstanceViewListRespItemLocation ...
type GetCalendarEventInstanceViewListRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称
	Address   string  `json:"address,omitempty"`   // 地点地址
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
}

// GetCalendarEventInstanceViewListRespItemStartTime ...
type GetCalendarEventInstanceViewListRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// GetCalendarEventInstanceViewListRespItemVchat ...
type GetCalendarEventInstanceViewListRespItemVchat struct {
	VCType      string `json:"vc_type,omitempty"`     // 视频会议类型, 可选值有: vc: 飞书视频会议, third_party: 第三方链接视频会议, no_meeting: 无视频会议, lark_live: Lark直播, unknown: 未知类型
	IconType    string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型, 可选值有: vc: 飞书视频会议icon, live: 直播视频会议icon, default: 默认icon
	Description string `json:"description,omitempty"` // 第三方视频会议文案, 可以为空, 为空展示默认文案
	MeetingURL  string `json:"meeting_url,omitempty"` // 视频会议URL
}

// getCalendarEventInstanceViewListResp ...
type getCalendarEventInstanceViewListResp struct {
	Code int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventInstanceViewListResp `json:"data,omitempty"`
}