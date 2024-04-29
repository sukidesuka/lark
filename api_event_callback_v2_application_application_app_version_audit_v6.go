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

// EventV2ApplicationApplicationAppVersionAuditV6 通过订阅该事件, 可接收应用审核（通过 / 拒绝）事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/events/audit
// new doc: https://open.feishu.cn/document/server-docs/application-v6/event/audit
func (r *EventCallbackService) HandlerEventV2ApplicationApplicationAppVersionAuditV6(f EventV2ApplicationApplicationAppVersionAuditV6Handler) {
	r.cli.eventHandler.eventV2ApplicationApplicationAppVersionAuditV6Handler = f
}

// EventV2ApplicationApplicationAppVersionAuditV6Handler event EventV2ApplicationApplicationAppVersionAuditV6 handler
type EventV2ApplicationApplicationAppVersionAuditV6Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ApplicationApplicationAppVersionAuditV6) (string, error)

// EventV2ApplicationApplicationAppVersionAuditV6 ...
type EventV2ApplicationApplicationAppVersionAuditV6 struct {
	OperatorID  *EventV2ApplicationApplicationAppVersionAuditV6OperatorID `json:"operator_id,omitempty"`  // 通过 / 拒绝应用审核的管理员 id
	VersionID   string                                                    `json:"version_id,omitempty"`   // 被审核的应用版本 id
	CreatorID   *EventV2ApplicationApplicationAppVersionAuditV6CreatorID  `json:"creator_id,omitempty"`   // 应用创建者的 id
	AppID       string                                                    `json:"app_id,omitempty"`       // 撤回应用的 id
	Operation   string                                                    `json:"operation,omitempty"`    // 审核通过 / 拒绝, 可选值有: audited: 审核通过, reject: 审核拒绝
	Remark      string                                                    `json:"remark,omitempty"`       // 审核信息, 当审核拒绝时, 管理员填写的拒绝理由
	AuditSource string                                                    `json:"audit_source,omitempty"` // 应用审核的方式, 可选值有: administrator: 管理员审核, auto: 自动免审
}

// EventV2ApplicationApplicationAppVersionAuditV6CreatorID ...
type EventV2ApplicationApplicationAppVersionAuditV6CreatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2ApplicationApplicationAppVersionAuditV6OperatorID ...
type EventV2ApplicationApplicationAppVersionAuditV6OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
