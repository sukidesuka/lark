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

// GetApplication 根据app_id获取应用的基础信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/get
// new doc: https://open.feishu.cn/document/server-docs/application-v6/application/get
func (r *ApplicationService) GetApplication(ctx context.Context, request *GetApplicationReq, options ...MethodOptionFunc) (*GetApplicationResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplication mock enable")
		return r.cli.mock.mockApplicationGetApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplication",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplication mock ApplicationGetApplication method
func (r *Mock) MockApplicationGetApplication(f func(ctx context.Context, request *GetApplicationReq, options ...MethodOptionFunc) (*GetApplicationResp, *Response, error)) {
	r.mockApplicationGetApplication = f
}

// UnMockApplicationGetApplication un-mock ApplicationGetApplication method
func (r *Mock) UnMockApplicationGetApplication() {
	r.mockApplicationGetApplication = nil
}

// GetApplicationReq ...
type GetApplicationReq struct {
	AppID      string  `path:"app_id" json:"-"`        // 应用的 app_id, 需要查询其他应用信息时, 必须申请[获取应用信息](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)权限, 仅查询本应用信息时, 可填入 "me" 或者应用自身 app_id, 示例值: "cli_9b445f5258795107"
	Lang       string  `query:"lang" json:"-"`         // 指定获取应用在该语言下的信息, 示例值: zh_cn, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文, 最小长度: `1` 字符
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetApplicationResp ...
type GetApplicationResp struct {
	App *GetApplicationRespApp `json:"app,omitempty"` // 应用数据
}

// GetApplicationRespApp ...
type GetApplicationRespApp struct {
	AppID            string                        `json:"app_id,omitempty"`             // 应用的 app_id
	CreatorID        string                        `json:"creator_id,omitempty"`         // 应用创建者（所有者）
	Status           int64                         `json:"status,omitempty"`             // 应用状态, 可选值有: 0: 停用状态, 1: 启用状态, 2: 未启用状态, 3: 未知状态
	SceneType        int64                         `json:"scene_type,omitempty"`         // 应用类型, 可选值有: 0: 自建应用, 1: 应用商店应用, 2: 个人应用商店应用, 3: 未知应用类型
	PaymentType      int64                         `json:"payment_type,omitempty"`       // 付费类型, 可选值有: 0: 免费, 1: 付费
	RedirectURLs     []string                      `json:"redirect_urls,omitempty"`      // 安全设置中的重定向 URL
	OnlineVersionID  string                        `json:"online_version_id,omitempty"`  // 发布在线上的应用版本 ID, 若没有则为空
	UnauditVersionID string                        `json:"unaudit_version_id,omitempty"` // 在审核中的版本 ID, 若没有则为空
	AppName          string                        `json:"app_name,omitempty"`           // 应用名称
	AvatarURL        string                        `json:"avatar_url,omitempty"`         // 应用图标 url
	Description      string                        `json:"description,omitempty"`        // 应用默认描述
	Scopes           []*GetApplicationRespAppScope `json:"scopes,omitempty"`             // 应用权限列表
	BackHomeURL      string                        `json:"back_home_url,omitempty"`      // 后台主页地址
	I18n             []*GetApplicationRespAppI18n  `json:"i18n,omitempty"`               // 应用的国际化信息列表
	PrimaryLanguage  string                        `json:"primary_language,omitempty"`   // 应用主语言, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	CommonCategories []string                      `json:"common_categories,omitempty"`  // 应用分类的国际化描述
	Owner            *GetApplicationRespAppOwner   `json:"owner,omitempty"`              // 应用的所有者信息
}

// GetApplicationRespAppI18n ...
type GetApplicationRespAppI18n struct {
	I18nKey     string `json:"i18n_key,omitempty"`    // 国际化语言的 key, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	Name        string `json:"name,omitempty"`        // 应用国际化名称
	Description string `json:"description,omitempty"` // 应用国际化描述（副标题）
	HelpUse     string `json:"help_use,omitempty"`    // 国际化帮助文档链接
}

// GetApplicationRespAppOwner ...
type GetApplicationRespAppOwner struct {
	Type     int64  `json:"type,omitempty"`      // 应用所有者类型, 可选值有: 0: 飞书科技, 1: 飞书合作伙伴, 2: 企业内成员
	OwnerID  string `json:"owner_id,omitempty"`  // 应用所有者ID
	Name     string `json:"name,omitempty"`      // 应用开发商名称(仅商店应用返回)
	HelpDesk string `json:"help_desk,omitempty"` // 应用开发商服务台链接(仅商店应用返回)
	Email    string `json:"email,omitempty"`     // 应用开发商的邮箱(仅商店应用返回)
	Phone    string `json:"phone,omitempty"`     // 应用开发商的手机号(仅商店应用返回)
}

// GetApplicationRespAppScope ...
type GetApplicationRespAppScope struct {
	Scope       string `json:"scope,omitempty"`       // 应用权限
	Description string `json:"description,omitempty"` // 应用权限的国际化描述
	Level       int64  `json:"level,omitempty"`       // 权限等级描述, 可选值有: 1: 普通权限, 2: 高级权限, 3: 超敏感权限, 0: 未知等级
}

// getApplicationResp ...
type getApplicationResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetApplicationResp `json:"data,omitempty"`
}
