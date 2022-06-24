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

// DeleteEphemeralMessage
//
// 在群会话中删除指定用户可见的临时消息卡片<br>
// 临时卡片消息可以通过该接口进行显式删除，临时卡片消息删除后将不会在该设备上留下任何痕迹。
// **权限说明** ：需要启用机器人能力；需要机器人在会话群里
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN
func (r *MessageService) DeleteEphemeralMessage(ctx context.Context, request *DeleteEphemeralMessageReq, options ...MethodOptionFunc) (*DeleteEphemeralMessageResp, *Response, error) {
	if r.cli.mock.mockMessageDeleteEphemeralMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#DeleteEphemeralMessage mock enable")
		return r.cli.mock.mockMessageDeleteEphemeralMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "DeleteEphemeralMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/ephemeral/v1/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteEphemeralMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageDeleteEphemeralMessage mock MessageDeleteEphemeralMessage method
func (r *Mock) MockMessageDeleteEphemeralMessage(f func(ctx context.Context, request *DeleteEphemeralMessageReq, options ...MethodOptionFunc) (*DeleteEphemeralMessageResp, *Response, error)) {
	r.mockMessageDeleteEphemeralMessage = f
}

// UnMockMessageDeleteEphemeralMessage un-mock MessageDeleteEphemeralMessage method
func (r *Mock) UnMockMessageDeleteEphemeralMessage() {
	r.mockMessageDeleteEphemeralMessage = nil
}

// DeleteEphemeralMessageReq ...
type DeleteEphemeralMessageReq struct {
	MessageID string `json:"message_id,omitempty"` // 临时消息ID
}

// deleteEphemeralMessageResp ...
type deleteEphemeralMessageResp struct {
	Code int64                       `json:"code,omitempty"`
	Msg  string                      `json:"msg,omitempty"`
	Data *DeleteEphemeralMessageResp `json:"data,omitempty"`
}

// DeleteEphemeralMessageResp ...
type DeleteEphemeralMessageResp struct {
}
