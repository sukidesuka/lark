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

// SendUrgentAppMessage 对指定消息进行应用内加急。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 只能加急机器人自己发送的消息
// - 加急时机器人仍需要在会话内
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_app
func (r *MessageService) SendUrgentAppMessage(ctx context.Context, request *SendUrgentAppMessageReq, options ...MethodOptionFunc) (*SendUrgentAppMessageResp, *Response, error) {
	if r.cli.mock.mockMessageSendUrgentAppMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#SendUrgentAppMessage mock enable")
		return r.cli.mock.mockMessageSendUrgentAppMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "SendUrgentAppMessage",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/urgent_app",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendUrgentAppMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageSendUrgentAppMessage mock MessageSendUrgentAppMessage method
func (r *Mock) MockMessageSendUrgentAppMessage(f func(ctx context.Context, request *SendUrgentAppMessageReq, options ...MethodOptionFunc) (*SendUrgentAppMessageResp, *Response, error)) {
	r.mockMessageSendUrgentAppMessage = f
}

// UnMockMessageSendUrgentAppMessage un-mock MessageSendUrgentAppMessage method
func (r *Mock) UnMockMessageSendUrgentAppMessage() {
	r.mockMessageSendUrgentAppMessage = nil
}

// SendUrgentAppMessageReq ...
type SendUrgentAppMessageReq struct {
	UserIDType IDType   `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MessageID  string   `path:"message_id" json:"-"`    // 待加急的消息ID。注意不支持批量消息ID（bm_xxx）, 示例值："om_dc13264520392913993dd051dba21dcf"
	UserIDList []string `json:"user_id_list,omitempty"` // 目标用户的ID。列表不可为空。, 示例值：["ou_6yf8af6bgb9100449565764t3382b168"]
}

// sendUrgentAppMessageResp ...
type sendUrgentAppMessageResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *SendUrgentAppMessageResp `json:"data,omitempty"`
}

// SendUrgentAppMessageResp ...
type SendUrgentAppMessageResp struct {
	InvalidUserIDList []string `json:"invalid_user_id_list,omitempty"` // 无效的用户ID
}
