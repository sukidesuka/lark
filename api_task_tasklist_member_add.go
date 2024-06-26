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

// AddTaskTasklistMember 向一个清单添加1个或多个协作成员。成员信息通过设置`members`字段实现。关于member的格式, 详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 如何表示任务和清单的成员？”章节。
//
// 一个清单协作成员可以是一个用户, 应用或者群组。每个成员可以设置“可编辑”或者“可阅读”的角色。群组作为协作成员表示该群里所有群成员都自动拥有群组协作成员的角色。
// 如果要添加的成员已经是清单成员, 且角色和请求中设置是一样的, 则会被自动忽略, 接口返回成功。
// 如果要添加的成员已经是清单成员, 且角色和请求中设置是不一样的（比如原来的角色是可阅读, 请求中设为可编辑）, 则相当于更新其角色。
// 如果要添加的成员已经是清单的所有者, 则会被自动忽略。接口返回成功。其所有者的角色不会改变。
// 本接口不能用来设置清单所有者, 如要设置, 可以使用[更新清单](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/patch)接口。
// 需要清单编辑权限。详情见[清单功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/overview)中的“清单是如何鉴权的？“章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/tasklist/add_members
func (r *TaskService) AddTaskTasklistMember(ctx context.Context, request *AddTaskTasklistMemberReq, options ...MethodOptionFunc) (*AddTaskTasklistMemberResp, *Response, error) {
	if r.cli.mock.mockTaskAddTaskTasklistMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#AddTaskTasklistMember mock enable")
		return r.cli.mock.mockTaskAddTaskTasklistMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "AddTaskTasklistMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasklists/:tasklist_guid/add_members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(addTaskTasklistMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskAddTaskTasklistMember mock TaskAddTaskTasklistMember method
func (r *Mock) MockTaskAddTaskTasklistMember(f func(ctx context.Context, request *AddTaskTasklistMemberReq, options ...MethodOptionFunc) (*AddTaskTasklistMemberResp, *Response, error)) {
	r.mockTaskAddTaskTasklistMember = f
}

// UnMockTaskAddTaskTasklistMember un-mock TaskAddTaskTasklistMember method
func (r *Mock) UnMockTaskAddTaskTasklistMember() {
	r.mockTaskAddTaskTasklistMember = nil
}

// AddTaskTasklistMemberReq ...
type AddTaskTasklistMemberReq struct {
	TasklistGuid string                            `path:"tasklist_guid" json:"-"` // 要添加成员的清单的全局唯一ID, 示例值: "d300a75f-c56a-4be9-80d1-e47653028ceb"
	UserIDType   *IDType                           `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
	Members      []*AddTaskTasklistMemberReqMember `json:"members,omitempty"`      // 要添加的成员列表。关于member的格式, 详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 如何表示任务和清单的成员？”章节, 长度范围: `1` ～ `500`
}

// AddTaskTasklistMemberReqMember ...
type AddTaskTasklistMemberReqMember struct {
	ID   *string `json:"id,omitempty"`   // 表示member的id, 示例值: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f", 最大长度: `100` 字符
	Type *string `json:"type,omitempty"` // 成员的类型, 支持: user: 普通用户, 此时member的id是一个表示用户的ID, 比如open_id。具体格式取决于user_id_type参数, chat: 群组, 此时member的id是一个Open Chat ID, app: 应用, 此时member的id是一个应用的ID, 示例值: "user", 默认值: `user`
	Role *string `json:"role,omitempty"` // 成员角色。支持: editor: 可编辑, viewer: 可阅读, 默认为"viewer", 不能通过该字段设置清单所有者角色, 示例值: "editor", 最大长度: `20` 字符
}

// AddTaskTasklistMemberResp ...
type AddTaskTasklistMemberResp struct {
	Tasklist *AddTaskTasklistMemberRespTasklist `json:"tasklist,omitempty"` // 完成更新后的清单实体
}

// AddTaskTasklistMemberRespTasklist ...
type AddTaskTasklistMemberRespTasklist struct {
	Guid      string                                     `json:"guid,omitempty"`       // 清单的全局唯一ID
	Name      string                                     `json:"name,omitempty"`       // 清单名
	Creator   *AddTaskTasklistMemberRespTasklistCreator  `json:"creator,omitempty"`    // 清单创建者
	Owner     *AddTaskTasklistMemberRespTasklistOwner    `json:"owner,omitempty"`      // 清单所有者
	Members   []*AddTaskTasklistMemberRespTasklistMember `json:"members,omitempty"`    // 清单协作成员
	URL       string                                     `json:"url,omitempty"`        // 该清单分享的applink
	CreatedAt string                                     `json:"created_at,omitempty"` // 清单创建时间戳(ms)
	UpdatedAt string                                     `json:"updated_at,omitempty"` // 清单最后一次更新时间戳（ms)
}

// AddTaskTasklistMemberRespTasklistCreator ...
type AddTaskTasklistMemberRespTasklistCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员类型
}

// AddTaskTasklistMemberRespTasklistMember ...
type AddTaskTasklistMemberRespTasklistMember struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// AddTaskTasklistMemberRespTasklistOwner ...
type AddTaskTasklistMemberRespTasklistOwner struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 清单角色
}

// addTaskTasklistMemberResp ...
type addTaskTasklistMemberResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *AddTaskTasklistMemberResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
