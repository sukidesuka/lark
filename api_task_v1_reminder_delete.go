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

// DeleteTaskV1Reminder 删除提醒时间, 返回结果状态。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/delete
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task-reminder/delete
//
// Deprecated
func (r *TaskV1Service) DeleteTaskV1Reminder(ctx context.Context, request *DeleteTaskV1ReminderReq, options ...MethodOptionFunc) (*DeleteTaskV1ReminderResp, *Response, error) {
	if r.cli.mock.mockTaskV1DeleteTaskV1Reminder != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] TaskV1#DeleteTaskV1Reminder mock enable")
		return r.cli.mock.mockTaskV1DeleteTaskV1Reminder(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "TaskV1",
		API:                   "DeleteTaskV1Reminder",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/reminders/:reminder_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteTaskV1ReminderResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskV1DeleteTaskV1Reminder mock TaskV1DeleteTaskV1Reminder method
func (r *Mock) MockTaskV1DeleteTaskV1Reminder(f func(ctx context.Context, request *DeleteTaskV1ReminderReq, options ...MethodOptionFunc) (*DeleteTaskV1ReminderResp, *Response, error)) {
	r.mockTaskV1DeleteTaskV1Reminder = f
}

// UnMockTaskV1DeleteTaskV1Reminder un-mock TaskV1DeleteTaskV1Reminder method
func (r *Mock) UnMockTaskV1DeleteTaskV1Reminder() {
	r.mockTaskV1DeleteTaskV1Reminder = nil
}

// DeleteTaskV1ReminderReq ...
type DeleteTaskV1ReminderReq struct {
	TaskID     string `path:"task_id" json:"-"`     // 任务 ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	ReminderID string `path:"reminder_id" json:"-"` // 任务提醒时间设置的 ID（即 reminder.id）, 示例值: "1"
}

// DeleteTaskV1ReminderResp ...
type DeleteTaskV1ReminderResp struct {
}

// deleteTaskV1ReminderResp ...
type deleteTaskV1ReminderResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteTaskV1ReminderResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
