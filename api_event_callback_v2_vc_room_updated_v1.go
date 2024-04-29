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

// EventV2VCRoomUpdatedV1 当更新会议室时, 会触发该事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room/events/updated
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/room/events/updated
func (r *EventCallbackService) HandlerEventV2VCRoomUpdatedV1(f EventV2VCRoomUpdatedV1Handler) {
	r.cli.eventHandler.eventV2VCRoomUpdatedV1Handler = f
}

// EventV2VCRoomUpdatedV1Handler event EventV2VCRoomUpdatedV1 handler
type EventV2VCRoomUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCRoomUpdatedV1) (string, error)

// EventV2VCRoomUpdatedV1 ...
type EventV2VCRoomUpdatedV1 struct {
	Room *EventV2VCRoomUpdatedV1Room `json:"room,omitempty"` // 会议室详情
}

// EventV2VCRoomUpdatedV1Room ...
type EventV2VCRoomUpdatedV1Room struct {
	RoomID       string                                `json:"room_id,omitempty"`        // 会议室ID
	Name         string                                `json:"name,omitempty"`           // 会议室名称
	Capacity     int64                                 `json:"capacity,omitempty"`       // 会议室能容纳的人数
	Description  string                                `json:"description,omitempty"`    // 会议室的相关描述
	DisplayID    string                                `json:"display_id,omitempty"`     // 会议室的展示ID
	CustomRoomID string                                `json:"custom_room_id,omitempty"` // 自定义的会议室ID
	RoomLevelID  string                                `json:"room_level_id,omitempty"`  // 层级ID
	Path         []string                              `json:"path,omitempty"`           // 层级路径
	RoomStatus   *EventV2VCRoomUpdatedV1RoomRoomStatus `json:"room_status,omitempty"`    // 会议室状态
	Device       []*EventV2VCRoomUpdatedV1RoomDevice   `json:"device,omitempty"`         // 设施信息列表
}

// EventV2VCRoomUpdatedV1RoomDevice ...
type EventV2VCRoomUpdatedV1RoomDevice struct {
	Name string `json:"name,omitempty"` // 设施名称
}

// EventV2VCRoomUpdatedV1RoomRoomStatus ...
type EventV2VCRoomUpdatedV1RoomRoomStatus struct {
	Status           bool     `json:"status,omitempty"`             // 是否启用会议室
	ScheduleStatus   bool     `json:"schedule_status,omitempty"`    // 会议室未来状态为启用或禁用
	DisableStartTime string   `json:"disable_start_time,omitempty"` // 禁用开始时间（unix时间, 单位sec）
	DisableEndTime   string   `json:"disable_end_time,omitempty"`   // 禁用结束时间（unix时间, 单位sec, 数值0表示永久禁用）
	DisableReason    string   `json:"disable_reason,omitempty"`     // 禁用原因
	ContactIDs       []string `json:"contact_ids,omitempty"`        // 联系人列表
	DisableNotice    bool     `json:"disable_notice,omitempty"`     // 是否在禁用时发送通知给预定了该会议室的员工
	ResumeNotice     bool     `json:"resume_notice,omitempty"`      // 是否在恢复启用时发送通知给预定了该会议室的员工
}
