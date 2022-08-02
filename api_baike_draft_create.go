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

// CreateBaikeDraft 草稿并非百科词条, 而是指通过 API 发起创建新词条或更新现有词条的申请。百科管理员审核通过后, 草稿将变为新的词条或覆盖已有词条。
//
// 以用户身份创建草稿（即 Authorization 使用 user_access_token）, 对应用户将收到由企业百科 Bot 发送的审核结果；以应用身份创建草稿（即 Authorization 使用 tenant_access_toke）, 不会收到任何通知。
// · 创建新的百科词条时, 无需传入 entity_id 字段
// · 更新已有百科词条时, 请传入对应词条的 entity_id 或 outer_info
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/draft/create
func (r *BaikeService) CreateBaikeDraft(ctx context.Context, request *CreateBaikeDraftReq, options ...MethodOptionFunc) (*CreateBaikeDraftResp, *Response, error) {
	if r.cli.mock.mockBaikeCreateBaikeDraft != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Baike#CreateBaikeDraft mock enable")
		return r.cli.mock.mockBaikeCreateBaikeDraft(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "CreateBaikeDraft",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/drafts",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBaikeDraftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeCreateBaikeDraft mock BaikeCreateBaikeDraft method
func (r *Mock) MockBaikeCreateBaikeDraft(f func(ctx context.Context, request *CreateBaikeDraftReq, options ...MethodOptionFunc) (*CreateBaikeDraftResp, *Response, error)) {
	r.mockBaikeCreateBaikeDraft = f
}

// UnMockBaikeCreateBaikeDraft un-mock BaikeCreateBaikeDraft method
func (r *Mock) UnMockBaikeCreateBaikeDraft() {
	r.mockBaikeCreateBaikeDraft = nil
}

// CreateBaikeDraftReq ...
type CreateBaikeDraftReq struct {
	UserIDType  *IDType                         `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ID          *string                         `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）, 示例值: "enterprise_40217521"
	MainKeys    []*CreateBaikeDraftReqMainKey   `json:"main_keys,omitempty"`    // 词条名, 最大长度: `1`
	Aliases     []*CreateBaikeDraftReqAliase    `json:"aliases,omitempty"`      // 别名, 最大长度: `10`
	Description *string                         `json:"description,omitempty"`  // 词条释义（纯文本格式）, 示例值: "企业百科是飞书提供的一款知识管理工具, 通过企业百科可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通", 最大长度: `5000` 字符
	RelatedMeta *CreateBaikeDraftReqRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	OuterInfo   *CreateBaikeDraftReqOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    *string                         `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[企业百科指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分, 示例值: "<b>加粗</b><i>斜体</i><p><a href=\"https://feishu.cn\">链接</a></p><p><span>企业百科是飞书提供的一款知识管理工具, 通过企业百科可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通</span></p>", 最大长度: `5000` 字符
}

// CreateBaikeDraftReqAliase ...
type CreateBaikeDraftReqAliase struct {
	Key           string                                  `json:"key,omitempty"`            // 名称的值, 示例值: "企业百科"
	DisplayStatus *CreateBaikeDraftReqAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeDraftReqAliaseDisplayStatus ...
type CreateBaikeDraftReqAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// CreateBaikeDraftReqMainKey ...
type CreateBaikeDraftReqMainKey struct {
	Key           string                                   `json:"key,omitempty"`            // 名称的值, 示例值: "企业百科"
	DisplayStatus *CreateBaikeDraftReqMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeDraftReqMainKeyDisplayStatus ...
type CreateBaikeDraftReqMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// CreateBaikeDraftReqOuterInfo ...
type CreateBaikeDraftReqOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）, 示例值: "星云", 长度范围: `2` ～ `32` 字符
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）, 示例值: "client_6539i3498d", 长度范围: `1` ～ `64` 字符
}

// CreateBaikeDraftReqRelatedMeta ...
type CreateBaikeDraftReqRelatedMeta struct {
	Users           []*CreateBaikeDraftReqRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateBaikeDraftReqRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*CreateBaikeDraftReqRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*CreateBaikeDraftReqRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateBaikeDraftReqRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*CreateBaikeDraftReqRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateBaikeDraftReqRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
}

// CreateBaikeDraftReqRelatedMetaAbbreviation ...
type CreateBaikeDraftReqRelatedMetaAbbreviation struct {
	ID *string `json:"id,omitempty"` // 相关词条 ID, 示例值: "enterprise_51587960"
}

// CreateBaikeDraftReqRelatedMetaChat ...
type CreateBaikeDraftReqRelatedMetaChat struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// CreateBaikeDraftReqRelatedMetaClassification ...
type CreateBaikeDraftReqRelatedMetaClassification struct {
	ID       string  `json:"id,omitempty"`        // 二级分类 ID, 示例值: "7049606926702837761"
	FatherID *string `json:"father_id,omitempty"` // 对应一级分类 ID, 示例值: "7049606926702837777"
}

// CreateBaikeDraftReqRelatedMetaDoc ...
type CreateBaikeDraftReqRelatedMetaDoc struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "企业百科帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateBaikeDraftReqRelatedMetaLink ...
type CreateBaikeDraftReqRelatedMetaLink struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "企业百科帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateBaikeDraftReqRelatedMetaOncall ...
type CreateBaikeDraftReqRelatedMetaOncall struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// CreateBaikeDraftReqRelatedMetaUser ...
type CreateBaikeDraftReqRelatedMetaUser struct {
	ID    string  `json:"id,omitempty"`    // 对应相关信息 ID, 示例值: "格式请看请求体示例"
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "企业百科帮助中心"
}

// CreateBaikeDraftResp ...
type CreateBaikeDraftResp struct {
	Draft *CreateBaikeDraftRespDraft `json:"draft,omitempty"` // 草稿信息
}

// CreateBaikeDraftRespDraft ...
type CreateBaikeDraftRespDraft struct {
	DraftID string                           `json:"draft_id,omitempty"` // 草稿 ID
	Entity  *CreateBaikeDraftRespDraftEntity `json:"entity,omitempty"`   // 词条信息
}

// CreateBaikeDraftRespDraftEntity ...
type CreateBaikeDraftRespDraftEntity struct {
	ID          string                                      `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）
	MainKeys    []*CreateBaikeDraftRespDraftEntityMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*CreateBaikeDraftRespDraftEntityAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                      `json:"description,omitempty"`  // 词条释义（纯文本格式）
	CreateTime  string                                      `json:"create_time,omitempty"`  // 词条创建时间
	UpdateTime  string                                      `json:"update_time,omitempty"`  // 词条最近更新时间
	RelatedMeta *CreateBaikeDraftRespDraftEntityRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	Categories  []string                                    `json:"categories,omitempty"`   // 词条标签
	Statistics  *CreateBaikeDraftRespDraftEntityStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *CreateBaikeDraftRespDraftEntityOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    string                                      `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[企业百科指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
}

// CreateBaikeDraftRespDraftEntityAliase ...
type CreateBaikeDraftRespDraftEntityAliase struct {
	Key           string                                              `json:"key,omitempty"`            // 名称的值
	DisplayStatus *CreateBaikeDraftRespDraftEntityAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeDraftRespDraftEntityAliaseDisplayStatus ...
type CreateBaikeDraftRespDraftEntityAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// CreateBaikeDraftRespDraftEntityMainKey ...
type CreateBaikeDraftRespDraftEntityMainKey struct {
	Key           string                                               `json:"key,omitempty"`            // 名称的值
	DisplayStatus *CreateBaikeDraftRespDraftEntityMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeDraftRespDraftEntityMainKeyDisplayStatus ...
type CreateBaikeDraftRespDraftEntityMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// CreateBaikeDraftRespDraftEntityOuterInfo ...
type CreateBaikeDraftRespDraftEntityOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）
}

// CreateBaikeDraftRespDraftEntityRelatedMeta ...
type CreateBaikeDraftRespDraftEntityRelatedMeta struct {
	Users           []*CreateBaikeDraftRespDraftEntityRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateBaikeDraftRespDraftEntityRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*CreateBaikeDraftRespDraftEntityRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*CreateBaikeDraftRespDraftEntityRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateBaikeDraftRespDraftEntityRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*CreateBaikeDraftRespDraftEntityRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateBaikeDraftRespDraftEntityRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
}

// CreateBaikeDraftRespDraftEntityRelatedMetaAbbreviation ...
type CreateBaikeDraftRespDraftEntityRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关词条 ID
}

// CreateBaikeDraftRespDraftEntityRelatedMetaChat ...
type CreateBaikeDraftRespDraftEntityRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeDraftRespDraftEntityRelatedMetaClassification ...
type CreateBaikeDraftRespDraftEntityRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	Name     string `json:"name,omitempty"`      // 二级分类名称
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// CreateBaikeDraftRespDraftEntityRelatedMetaDoc ...
type CreateBaikeDraftRespDraftEntityRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeDraftRespDraftEntityRelatedMetaLink ...
type CreateBaikeDraftRespDraftEntityRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeDraftRespDraftEntityRelatedMetaOncall ...
type CreateBaikeDraftRespDraftEntityRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeDraftRespDraftEntityRelatedMetaUser ...
type CreateBaikeDraftRespDraftEntityRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeDraftRespDraftEntityStatistics ...
type CreateBaikeDraftRespDraftEntityStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 累计点赞
	DislikeCount int64 `json:"dislike_count,omitempty"` // 当前词条版本收到的负反馈数量
}

// createBaikeDraftResp ...
type createBaikeDraftResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *CreateBaikeDraftResp `json:"data,omitempty"`
}