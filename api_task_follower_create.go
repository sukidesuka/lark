// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateTaskFollower 该接口用于创建任务关注者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/create
func (r *TaskService) CreateTaskFollower(ctx context.Context, request *CreateTaskFollowerReq, options ...MethodOptionFunc) (*CreateTaskFollowerResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskFollower != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CreateTaskFollower mock enable")
		return r.cli.mock.mockTaskCreateTaskFollower(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskFollower",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/followers",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createTaskFollowerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockTaskCreateTaskFollower(f func(ctx context.Context, request *CreateTaskFollowerReq, options ...MethodOptionFunc) (*CreateTaskFollowerResp, *Response, error)) {
	r.mockTaskCreateTaskFollower = f
}

func (r *Mock) UnMockTaskCreateTaskFollower() {
	r.mockTaskCreateTaskFollower = nil
}

type CreateTaskFollowerReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`,, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	TaskID     string  `path:"task_id" json:"-"`       // 任务 ID, 示例值："83912691-2e43-47fc-94a4-d512e03984fa"
	ID         string  `json:"id,omitempty"`           // 任务关注者 ID, 示例值："ou_99e1a581b36ecc4862cbfbce473f3123"
}

type createTaskFollowerResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskFollowerResp `json:"data,omitempty"`
}

type CreateTaskFollowerResp struct {
	Follower *CreateTaskFollowerRespFollower `json:"follower,omitempty"` // 创建后的任务关注者
}

type CreateTaskFollowerRespFollower struct {
	ID string `json:"id,omitempty"` // 任务关注者 ID
}