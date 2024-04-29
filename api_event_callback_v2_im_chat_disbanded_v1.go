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

// EventV2IMChatDisbandedV1 群组被解散后触发此事件。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 需要订阅 [消息与群组] 分类下的 [解散群] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
// - 解散群并勾选[为群成员保留历史记录]不会触发该事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/disbanded
// new doc: https://open.feishu.cn/document/server-docs/group/chat/events/disbanded
func (r *EventCallbackService) HandlerEventV2IMChatDisbandedV1(f EventV2IMChatDisbandedV1Handler) {
	r.cli.eventHandler.eventV2IMChatDisbandedV1Handler = f
}

// EventV2IMChatDisbandedV1Handler event EventV2IMChatDisbandedV1 handler
type EventV2IMChatDisbandedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatDisbandedV1) (string, error)

// EventV2IMChatDisbandedV1 ...
type EventV2IMChatDisbandedV1 struct {
	ChatID            string                              `json:"chat_id,omitempty"`             // 群组 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	OperatorID        *EventV2IMChatDisbandedV1OperatorID `json:"operator_id,omitempty"`         // 操作者的ID
	External          bool                                `json:"external,omitempty"`            // 被解散的群是否是外部群
	OperatorTenantKey string                              `json:"operator_tenant_key,omitempty"` // 操作者的租户 Key, 为租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用中的唯一标识
	Name              string                              `json:"name,omitempty"`                // 群名称
	I18nNames         *I18nNames                          `json:"i18n_names,omitempty"`          // 群国际化名称
}

// EventV2IMChatDisbandedV1OperatorID ...
type EventV2IMChatDisbandedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
