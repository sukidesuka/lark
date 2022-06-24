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

// UpdateWikiSpaceSetting 根据space_id更新知识空间公共设置
//
// 知识库权限要求：
// - 为知识空间管理员
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-setting/update
func (r *DriveService) UpdateWikiSpaceSetting(ctx context.Context, request *UpdateWikiSpaceSettingReq, options ...MethodOptionFunc) (*UpdateWikiSpaceSettingResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateWikiSpaceSetting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateWikiSpaceSetting mock enable")
		return r.cli.mock.mockDriveUpdateWikiSpaceSetting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateWikiSpaceSetting",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/spaces/:space_id/setting",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateWikiSpaceSettingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateWikiSpaceSetting mock DriveUpdateWikiSpaceSetting method
func (r *Mock) MockDriveUpdateWikiSpaceSetting(f func(ctx context.Context, request *UpdateWikiSpaceSettingReq, options ...MethodOptionFunc) (*UpdateWikiSpaceSettingResp, *Response, error)) {
	r.mockDriveUpdateWikiSpaceSetting = f
}

// UnMockDriveUpdateWikiSpaceSetting un-mock DriveUpdateWikiSpaceSetting method
func (r *Mock) UnMockDriveUpdateWikiSpaceSetting() {
	r.mockDriveUpdateWikiSpaceSetting = nil
}

// UpdateWikiSpaceSettingReq ...
type UpdateWikiSpaceSettingReq struct {
	SpaceID         string  `path:"space_id" json:"-"`          // 知识空间id, 示例值："1565676577122621"
	CreateSetting   *string `json:"create_setting,omitempty"`   // 谁可以创建空间的一级页面： "admin_and_member" = 管理员和成员 "admin"  - 仅管理员, 示例值："admin/admin_and_member"
	SecuritySetting *string `json:"security_setting,omitempty"` // 可阅读用户可否创建副本/打印/导出/复制： "allow" - 允许 "not_allow" - 不允许, 示例值："allow/not_allow"
	CommentSetting  *string `json:"comment_setting,omitempty"`  // 可阅读用户可否评论： "allow" - 允许 "not_allow" - 不允许, 示例值："allow/not_allow"
}

// updateWikiSpaceSettingResp ...
type updateWikiSpaceSettingResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateWikiSpaceSettingResp `json:"data,omitempty"`
}

// UpdateWikiSpaceSettingResp ...
type UpdateWikiSpaceSettingResp struct {
	Setting *UpdateWikiSpaceSettingRespSetting `json:"setting,omitempty"` // 知识空间设置
}

// UpdateWikiSpaceSettingRespSetting ...
type UpdateWikiSpaceSettingRespSetting struct {
	CreateSetting   string `json:"create_setting,omitempty"`   // 谁可以创建空间的一级页面： "admin_and_member" = 管理员和成员 "admin"  - 仅管理员
	SecuritySetting string `json:"security_setting,omitempty"` // 可阅读用户可否创建副本/打印/导出/复制： "allow" - 允许 "not_allow" - 不允许
	CommentSetting  string `json:"comment_setting,omitempty"`  // 可阅读用户可否评论： "allow" - 允许 "not_allow" - 不允许
}
