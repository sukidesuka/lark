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

// UpdateDepartment 该接口用于更新当前部门所有信息。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。
//
// - 调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。
// - 没有填写的字段会被置为空值（order字段除外）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/update
func (r *ContactService) UpdateDepartment(ctx context.Context, request *UpdateDepartmentReq, options ...MethodOptionFunc) (*UpdateDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactUpdateDepartment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UpdateDepartment mock enable")
		return r.cli.mock.mockContactUpdateDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateDepartment",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/departments/:department_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateDepartment mock ContactUpdateDepartment method
func (r *Mock) MockContactUpdateDepartment(f func(ctx context.Context, request *UpdateDepartmentReq, options ...MethodOptionFunc) (*UpdateDepartmentResp, *Response, error)) {
	r.mockContactUpdateDepartment = f
}

// UnMockContactUpdateDepartment un-mock ContactUpdateDepartment method
func (r *Mock) UnMockContactUpdateDepartment() {
	r.mockContactUpdateDepartment = nil
}

// UpdateDepartmentReq ...
type UpdateDepartmentReq struct {
	UserIDType         *IDType                      `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType   *DepartmentIDType            `query:"department_id_type" json:"-"`   // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	DepartmentID       string                       `path:"department_id" json:"-"`         // 部门ID，需要与查询参数中传入的department_id_type类型保持一致。, 示例值："od-4e6ac4d14bcd5071a37a39de902c7141", 最大长度：`64` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`
	Name               string                       `json:"name,omitempty"`                 // 部门名称, 示例值："DemoName", 最小长度：`1` 字符
	I18nName           *UpdateDepartmentReqI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称
	ParentDepartmentID string                       `json:"parent_department_id,omitempty"` // 父部门的ID,* 创建根部门，该参数值为 “0”, 示例值："D067"
	LeaderUserID       *string                      `json:"leader_user_id,omitempty"`       // 部门主管用户ID, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	Order              *string                      `json:"order,omitempty"`                // 部门的排序，即部门在其同级部门的展示顺序, 示例值："100"
	UnitIDs            []string                     `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表，当前只支持一个, 示例值：custom_unit_id
	CreateGroupChat    *bool                        `json:"create_group_chat,omitempty"`    // 是否创建部门群，默认不创建, 示例值：false
}

// UpdateDepartmentReqI18nName ...
type UpdateDepartmentReqI18nName struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 部门的中文名, 示例值："Demo名称"
	JaJp *string `json:"ja_jp,omitempty"` // 部门的日文名, 示例值："デモ名"
	EnUs *string `json:"en_us,omitempty"` // 部门的英文名, 示例值："Demo Name"
}

// updateDepartmentResp ...
type updateDepartmentResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDepartmentResp `json:"data,omitempty"`
}

// UpdateDepartmentResp ...
type UpdateDepartmentResp struct {
	Department *UpdateDepartmentRespDepartment `json:"department,omitempty"` // 部门信息
}

// UpdateDepartmentRespDepartment ...
type UpdateDepartmentRespDepartment struct {
	Name               string                                  `json:"name,omitempty"`                 // 部门名称,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门基础信息,读取通讯录,以应用身份访问通讯录
	I18nName           *UpdateDepartmentRespDepartmentI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门基础信息,读取通讯录,以应用身份访问通讯录
	ParentDepartmentID string                                  `json:"parent_department_id,omitempty"` // 父部门的ID,* 创建根部门，该参数值为 “0”,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门组织架构信息,读取通讯录,以应用身份访问通讯录
	DepartmentID       string                                  `json:"department_id,omitempty"`        // 本部门的自定义部门ID,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门基础信息,读取通讯录,以应用身份访问通讯录
	OpenDepartmentID   string                                  `json:"open_department_id,omitempty"`   // 部门的open_id
	LeaderUserID       string                                  `json:"leader_user_id,omitempty"`       // 部门主管用户ID,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门组织架构信息,读取通讯录,以应用身份访问通讯录
	ChatID             string                                  `json:"chat_id,omitempty"`              // 部门群ID,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门基础信息,读取通讯录,以应用身份访问通讯录
	Order              string                                  `json:"order,omitempty"`                // 部门的排序，即部门在其同级部门的展示顺序,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门组织架构信息,读取通讯录,以应用身份访问通讯录
	UnitIDs            []string                                `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表，当前只支持一个,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门组织架构信息,读取通讯录,以应用身份访问通讯录
	MemberCount        int64                                   `json:"member_count,omitempty"`         // 部门下用户的个数,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门组织架构信息,读取通讯录,以应用身份访问通讯录
	Status             *UpdateDepartmentRespDepartmentStatus   `json:"status,omitempty"`               // 部门状态,**字段权限要求（满足任一）**：,以应用身份读取通讯录,获取部门基础信息,读取通讯录,以应用身份访问通讯录
}

// UpdateDepartmentRespDepartmentI18nName ...
type UpdateDepartmentRespDepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

// UpdateDepartmentRespDepartmentStatus ...
type UpdateDepartmentRespDepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}
