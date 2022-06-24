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

// UpdateChatModeration 更新群组的发言权限设置，可设置为全员可发言、仅管理员可发言  或 指定用户可发言。
//
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 若以用户授权调用接口，**当授权用户是群主**时，可更新群发言权限
// - 若以租户授权调用接口(即以机器人身份调用接口)，当**机器人是群主** 或者 **机器人是创建群组、具备[更新应用所创建群的群信息]权限且仍在群内**时，可更新群发言权限
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/update
func (r *ChatService) UpdateChatModeration(ctx context.Context, request *UpdateChatModerationReq, options ...MethodOptionFunc) (*UpdateChatModerationResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChatModeration != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChatModeration mock enable")
		return r.cli.mock.mockChatUpdateChatModeration(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChatModeration",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/moderation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatModerationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChatModeration mock ChatUpdateChatModeration method
func (r *Mock) MockChatUpdateChatModeration(f func(ctx context.Context, request *UpdateChatModerationReq, options ...MethodOptionFunc) (*UpdateChatModerationResp, *Response, error)) {
	r.mockChatUpdateChatModeration = f
}

// UnMockChatUpdateChatModeration un-mock ChatUpdateChatModeration method
func (r *Mock) UnMockChatUpdateChatModeration() {
	r.mockChatUpdateChatModeration = nil
}

// UpdateChatModerationReq ...
type UpdateChatModerationReq struct {
	UserIDType           *IDType  `query:"user_id_type" json:"-"`           // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`，字段权限要求: 获取用户 user ID
	ChatID               string   `path:"chat_id" json:"-"`                 // 群 ID，详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	ModerationSetting    *string  `json:"moderation_setting,omitempty"`     // 群发言模式（all_members/only_owner/moderator_list，其中 moderator_list 表示部分用户可发言的模式）, 示例值: "moderator_list"
	ModeratorAddedList   []string `json:"moderator_added_list,omitempty"`   // 选择部分用户可发言模式时，添加的可发言用户列表（自动过滤不在群内的用户）
	ModeratorRemovedList []string `json:"moderator_removed_list,omitempty"` // 选择部分用户可发言模式时，移除的可发言用户列表（自动过滤不在群内的用户）
}

// updateChatModerationResp ...
type updateChatModerationResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatModerationResp `json:"data,omitempty"`
}

// UpdateChatModerationResp ...
type UpdateChatModerationResp struct {
}
