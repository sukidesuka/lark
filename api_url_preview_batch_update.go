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

// BatchUpdateURLPreview 主动更新 URL 预览, 调用后会重新触发一次客户端拉取, 需要回调服务返回更新后的数据。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/url_preview/batch_update
func (r *MessageService) BatchUpdateURLPreview(ctx context.Context, request *BatchUpdateURLPreviewReq, options ...MethodOptionFunc) (*BatchUpdateURLPreviewResp, *Response, error) {
	if r.cli.mock.mockMessageBatchUpdateURLPreview != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#BatchUpdateURLPreview mock enable")
		return r.cli.mock.mockMessageBatchUpdateURLPreview(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "BatchUpdateURLPreview",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v2/url_previews/batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchUpdateURLPreviewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageBatchUpdateURLPreview mock MessageBatchUpdateURLPreview method
func (r *Mock) MockMessageBatchUpdateURLPreview(f func(ctx context.Context, request *BatchUpdateURLPreviewReq, options ...MethodOptionFunc) (*BatchUpdateURLPreviewResp, *Response, error)) {
	r.mockMessageBatchUpdateURLPreview = f
}

// UnMockMessageBatchUpdateURLPreview un-mock MessageBatchUpdateURLPreview method
func (r *Mock) UnMockMessageBatchUpdateURLPreview() {
	r.mockMessageBatchUpdateURLPreview = nil
}

// BatchUpdateURLPreviewReq ...
type BatchUpdateURLPreviewReq struct {
	PreviewTokens []string `json:"preview_tokens,omitempty"` // URL 预览的 token 列表。单个 token 限制更新频率为 1次/5秒, 示例值: ["952te0c8-9ccf-463d-ad73-593f8f768a5c"], 长度范围: `1` ～ `10`
	OpenIDs       *string  `json:"open_ids,omitempty"`       // 需要更新 URL 预览的用户 open_id。若不传, 则默认更新 URL 预览所在会话的所有成员；若用户不在 URL 所在会话, 则无法触发更新该用户对应的 URL 预览结果, 示例值: ["ou_xxx"], 长度范围: `0` ～ `100`
}

// BatchUpdateURLPreviewResp ...
type BatchUpdateURLPreviewResp struct {
}

// batchUpdateURLPreviewResp ...
type batchUpdateURLPreviewResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *BatchUpdateURLPreviewResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}