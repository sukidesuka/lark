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

// CancelSendHelpdeskNotification 取消推送接口，审核通过后待调度可以调用，发送过程中可以调用（会撤回已发送的消息），发送完成后可以需要推送（会撤回所有已发送的消息）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/cancel_send
func (r *HelpdeskService) CancelSendHelpdeskNotification(ctx context.Context, request *CancelSendHelpdeskNotificationReq, options ...MethodOptionFunc) (*CancelSendHelpdeskNotificationResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCancelSendHelpdeskNotification != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#CancelSendHelpdeskNotification mock enable")
		return r.cli.mock.mockHelpdeskCancelSendHelpdeskNotification(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "CancelSendHelpdeskNotification",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/notifications/:notification_id/cancel_send",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(cancelSendHelpdeskNotificationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskCancelSendHelpdeskNotification mock HelpdeskCancelSendHelpdeskNotification method
func (r *Mock) MockHelpdeskCancelSendHelpdeskNotification(f func(ctx context.Context, request *CancelSendHelpdeskNotificationReq, options ...MethodOptionFunc) (*CancelSendHelpdeskNotificationResp, *Response, error)) {
	r.mockHelpdeskCancelSendHelpdeskNotification = f
}

// UnMockHelpdeskCancelSendHelpdeskNotification un-mock HelpdeskCancelSendHelpdeskNotification method
func (r *Mock) UnMockHelpdeskCancelSendHelpdeskNotification() {
	r.mockHelpdeskCancelSendHelpdeskNotification = nil
}

// CancelSendHelpdeskNotificationReq ...
type CancelSendHelpdeskNotificationReq struct {
	NotificationID string `path:"notification_id" json:"-"` // 唯一ID, 示例值："6981801914270744596"
	IsRecall       bool   `json:"is_recall,omitempty"`      // 是否召回已发送的消息,新人入职消息同样适用, 示例值：true
}

// cancelSendHelpdeskNotificationResp ...
type cancelSendHelpdeskNotificationResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *CancelSendHelpdeskNotificationResp `json:"data,omitempty"`
}

// CancelSendHelpdeskNotificationResp ...
type CancelSendHelpdeskNotificationResp struct {
}
