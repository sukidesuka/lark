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

// FinishUploadDriveMedia 调用[上传分片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_part)接口将分片全部上传完毕后, 你可调用本接口触发完成上传。了解完整的分片上传素材流程, 参考[分片上传素材概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/multipart-upload-media/introduction)。
//
// 该接口不支持较高并发, 且调用频率上限为 5 QPS。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_finish
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/media/multipart-upload-media/upload_finish
func (r *DriveService) FinishUploadDriveMedia(ctx context.Context, request *FinishUploadDriveMediaReq, options ...MethodOptionFunc) (*FinishUploadDriveMediaResp, *Response, error) {
	if r.cli.mock.mockDriveFinishUploadDriveMedia != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#FinishUploadDriveMedia mock enable")
		return r.cli.mock.mockDriveFinishUploadDriveMedia(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "FinishUploadDriveMedia",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/upload_finish",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(finishUploadDriveMediaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveFinishUploadDriveMedia mock DriveFinishUploadDriveMedia method
func (r *Mock) MockDriveFinishUploadDriveMedia(f func(ctx context.Context, request *FinishUploadDriveMediaReq, options ...MethodOptionFunc) (*FinishUploadDriveMediaResp, *Response, error)) {
	r.mockDriveFinishUploadDriveMedia = f
}

// UnMockDriveFinishUploadDriveMedia un-mock DriveFinishUploadDriveMedia method
func (r *Mock) UnMockDriveFinishUploadDriveMedia() {
	r.mockDriveFinishUploadDriveMedia = nil
}

// FinishUploadDriveMediaReq ...
type FinishUploadDriveMediaReq struct {
	UploadID string `json:"upload_id,omitempty"` // 分片上传事务 ID。通过调用[分片上传素材（预上传）](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_prepare)接口获取, 示例值: "7111211691345512356"
	BlockNum int64  `json:"block_num,omitempty"` // 分片数量。通过调用[分片上传素材（预上传）](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_prepare)接口获取, 示例值: 1
}

// FinishUploadDriveMediaResp ...
type FinishUploadDriveMediaResp struct {
	FileToken string `json:"file_token,omitempty"` // 新创建文件的 token
}

// finishUploadDriveMediaResp ...
type finishUploadDriveMediaResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *FinishUploadDriveMediaResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
