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

// DeleteMailGroupPermissionMember 从自定义成员中删除单个成员，删除后该成员无法发送邮件到该邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/delete
func (r *MailService) DeleteMailGroupPermissionMember(ctx context.Context, request *DeleteMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupPermissionMemberResp, *Response, error) {
	if r.cli.mock.mockMailDeleteMailGroupPermissionMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailGroupPermissionMember mock enable")
		return r.cli.mock.mockMailDeleteMailGroupPermissionMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "DeleteMailGroupPermissionMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteMailGroupPermissionMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailDeleteMailGroupPermissionMember mock MailDeleteMailGroupPermissionMember method
func (r *Mock) MockMailDeleteMailGroupPermissionMember(f func(ctx context.Context, request *DeleteMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupPermissionMemberResp, *Response, error)) {
	r.mockMailDeleteMailGroupPermissionMember = f
}

// UnMockMailDeleteMailGroupPermissionMember un-mock MailDeleteMailGroupPermissionMember method
func (r *Mock) UnMockMailDeleteMailGroupPermissionMember() {
	r.mockMailDeleteMailGroupPermissionMember = nil
}

// DeleteMailGroupPermissionMemberReq ...
type DeleteMailGroupPermissionMemberReq struct {
	MailGroupID        string `path:"mailgroup_id" json:"-"`         // The unique ID or email address of a mail group, 示例值："xxxxxxxxxxxxxxx or test_mail_group@xxx.xx"
	PermissionMemberID string `path:"permission_member_id" json:"-"` // The unique ID of a member in this permission group, 示例值："xxxxxxxxxxxxxxx"
}

// deleteMailGroupPermissionMemberResp ...
type deleteMailGroupPermissionMemberResp struct {
	Code int64                                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailGroupPermissionMemberResp `json:"data,omitempty"`
}

// DeleteMailGroupPermissionMemberResp ...
type DeleteMailGroupPermissionMemberResp struct {
}
