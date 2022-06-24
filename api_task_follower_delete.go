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

// DeleteTaskFollower 该接口用于删除任务关注者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/delete
func (r *TaskService) DeleteTaskFollower(ctx context.Context, request *DeleteTaskFollowerReq, options ...MethodOptionFunc) (*DeleteTaskFollowerResp, *Response, error) {
	if r.cli.mock.mockTaskDeleteTaskFollower != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#DeleteTaskFollower mock enable")
		return r.cli.mock.mockTaskDeleteTaskFollower(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "DeleteTaskFollower",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/followers/:follower_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteTaskFollowerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskDeleteTaskFollower mock TaskDeleteTaskFollower method
func (r *Mock) MockTaskDeleteTaskFollower(f func(ctx context.Context, request *DeleteTaskFollowerReq, options ...MethodOptionFunc) (*DeleteTaskFollowerResp, *Response, error)) {
	r.mockTaskDeleteTaskFollower = f
}

// UnMockTaskDeleteTaskFollower un-mock TaskDeleteTaskFollower method
func (r *Mock) UnMockTaskDeleteTaskFollower() {
	r.mockTaskDeleteTaskFollower = nil
}

// DeleteTaskFollowerReq ...
type DeleteTaskFollowerReq struct {
	TaskID     string `path:"task_id" json:"-"`     // 任务 ID, 示例值："83912691-2e43-47fc-94a4-d512e03984fa"
	FollowerID string `path:"follower_id" json:"-"` // 任务关注者 ID（Open ID）, 示例值："ou_87e1a581b36ecc4862cbfbce473f346a"
}

// deleteTaskFollowerResp ...
type deleteTaskFollowerResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteTaskFollowerResp `json:"data,omitempty"`
}

// DeleteTaskFollowerResp ...
type DeleteTaskFollowerResp struct {
}
