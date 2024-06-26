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

// GetDrivePublicPermissionV2 该接口用于根据 filetoken 获取云文档的权限设置。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIzNzUjLyczM14iM3MTN/drive-v2/permission-public/get
// new doc: https://open.feishu.cn/document/server-docs/docs/permission/permission-public/get-2
func (r *DriveService) GetDrivePublicPermissionV2(ctx context.Context, request *GetDrivePublicPermissionV2Req, options ...MethodOptionFunc) (*GetDrivePublicPermissionV2Resp, *Response, error) {
	if r.cli.mock.mockDriveGetDrivePublicPermissionV2 != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetDrivePublicPermissionV2 mock enable")
		return r.cli.mock.mockDriveGetDrivePublicPermissionV2(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDrivePublicPermissionV2",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v2/permissions/:token/public",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDrivePublicPermissionV2Resp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDrivePublicPermissionV2 mock DriveGetDrivePublicPermissionV2 method
func (r *Mock) MockDriveGetDrivePublicPermissionV2(f func(ctx context.Context, request *GetDrivePublicPermissionV2Req, options ...MethodOptionFunc) (*GetDrivePublicPermissionV2Resp, *Response, error)) {
	r.mockDriveGetDrivePublicPermissionV2 = f
}

// UnMockDriveGetDrivePublicPermissionV2 un-mock DriveGetDrivePublicPermissionV2 method
func (r *Mock) UnMockDriveGetDrivePublicPermissionV2() {
	r.mockDriveGetDrivePublicPermissionV2 = nil
}

// GetDrivePublicPermissionV2Req ...
type GetDrivePublicPermissionV2Req struct {
	Token string `path:"token" json:"-"` // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	Type  string `query:"type" json:"-"` // 文件类型, 需要与文件的 token 相匹配, 示例值: doc, 可选值有: doc: 旧版文档, sheet: 电子表格, file: 云空间文件, wiki: 知识库节点, bitable: 多维表格, docx: 新版文档, mindnote: 思维笔记, minutes: 妙记, slides: 幻灯片
}

// GetDrivePublicPermissionV2Resp ...
type GetDrivePublicPermissionV2Resp struct {
	PermissionPublic *GetDrivePublicPermissionV2RespPermissionPublic `json:"permission_public,omitempty"` // 返回的文档公共设置
}

// GetDrivePublicPermissionV2RespPermissionPublic ...
type GetDrivePublicPermissionV2RespPermissionPublic struct {
	ExternalAccessEntity     string `json:"external_access_entity,omitempty"`     // 允许内容被分享到组织外, 可选值有: open: 打开, closed: 关闭, allow_share_partner_tenant: 允许分享给关联组织
	SecurityEntity           string `json:"security_entity,omitempty"`            // 谁可以创建副本、打印、下载, 可选值有: anyone_can_view: 拥有可阅读权限的用户, anyone_can_edit: 拥有可编辑权限的用户, only_full_access: 拥有可管理权限（包括我）的用户
	CommentEntity            string `json:"comment_entity,omitempty"`             // 谁可以评论, 可选值有: anyone_can_view: 拥有可阅读权限的用户, anyone_can_edit: 拥有可编辑权限的用户
	ShareEntity              string `json:"share_entity,omitempty"`               // 谁可以添加和管理协作者-组织维度, 可选值有: anyone: 所有可阅读或编辑此文档的用户, same_tenant: 组织内所有可阅读或编辑此文档的用户
	ManageCollaboratorEntity string `json:"manage_collaborator_entity,omitempty"` // 谁可以添加和管理协作者-协作者维度, 可选值有: collaborator_can_view: 拥有可阅读权限的协作者, collaborator_can_edit: 拥有可编辑权限的协作者, collaborator_full_access: 拥有可管理权限（包括我）的协作者
	LinkShareEntity          string `json:"link_share_entity,omitempty"`          // 链接分享设置, 可选值有: tenant_readable: 组织内获得链接的人可阅读, tenant_editable: 组织内获得链接的人可编辑, partner_tenant_readable: 关联组织的人可阅读, partner_tenant_editable: 关联组织的人可编辑, anyone_readable: 互联网上获得链接的任何人可阅读（仅external_access=“open”时有效）, anyone_editable: 互联网上获得链接的任何人可编辑（仅external_access=“open”时有效）, closed: 关闭链接分享
	CopyEntity               string `json:"copy_entity,omitempty"`                // 谁可以复制内容, 可选值有: anyone_can_view: 拥有可阅读权限的用户, anyone_can_edit: 拥有可编辑权限的用户, only_full_access: 拥有可管理权限（包括我）的协作者
	LockSwitch               bool   `json:"lock_switch,omitempty"`                // 节点是否已加锁, 加锁之后不再继承父级页面的权限
}

// getDrivePublicPermissionV2Resp ...
type getDrivePublicPermissionV2Resp struct {
	Code  int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                          `json:"msg,omitempty"`  // 错误描述
	Data  *GetDrivePublicPermissionV2Resp `json:"data,omitempty"`
	Error *ErrorDetail                    `json:"error,omitempty"`
}
