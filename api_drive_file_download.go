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
	"io"
)

// DownloadDriveFile 下载云空间下的文件, 不含飞书文档、电子表格以及多维表格等在线文档, 支持指定文件 Range 进行下载。
//
// 本接口提供文件下载能力, 如要下载云文档素材, 需调用[下载素材](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/download)接口。素材表示云文档中的资源文件, 比如新版文档中的图片及附件等, 素材不会呈现在云空间, 只会显示在对应云文档内。
// 该接口支持调用频率上限为 5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/download
func (r *DriveService) DownloadDriveFile(ctx context.Context, request *DownloadDriveFileReq, options ...MethodOptionFunc) (*DownloadDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveDownloadDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DownloadDriveFile mock enable")
		return r.cli.mock.mockDriveDownloadDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DownloadDriveFile",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/download",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(downloadDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDownloadDriveFile mock DriveDownloadDriveFile method
func (r *Mock) MockDriveDownloadDriveFile(f func(ctx context.Context, request *DownloadDriveFileReq, options ...MethodOptionFunc) (*DownloadDriveFileResp, *Response, error)) {
	r.mockDriveDownloadDriveFile = f
}

// UnMockDriveDownloadDriveFile un-mock DriveDownloadDriveFile method
func (r *Mock) UnMockDriveDownloadDriveFile() {
	r.mockDriveDownloadDriveFile = nil
}

// DownloadDriveFileReq ...
type DownloadDriveFileReq struct {
	FileToken string   `path:"file_token" json:"-"` // 文件的 token, 获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction), 示例值: "boxcnabCdefg12345"
	Range     [2]int64 `header:"range" json:"-"`    // 指定文件下载部分, 示例值: "bytes=0-1024"
}

// downloadDriveFileResp ...
type downloadDriveFileResp struct {
	Code int64                  `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
	Data *DownloadDriveFileResp `json:"data,omitempty"`
}

func (r *downloadDriveFileResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadDriveFileResp{}
	}
	r.Data.File = file
}

func (r *downloadDriveFileResp) SetFilename(filename string) {
	if r.Data == nil {
		r.Data = &DownloadDriveFileResp{}
	}
	r.Data.Filename = filename
}

// DownloadDriveFileResp ...
type DownloadDriveFileResp struct {
	File     io.Reader `json:"file,omitempty"`
	Filename string    `json:"filename,omitempty"`
}
