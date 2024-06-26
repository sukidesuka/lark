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

// GetDriveExportTask 根据[创建导出任务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/create)返回的导出任务 ID（ticket）轮询导出任务结果, 并返回导出文件的 token。你可使用该 token 继续调用[下载导出文件](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/download)接口将导出的产物下载到本地。了解完整的导出文件步骤, 参考[导出飞书云文档概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/export-user-guide)。
//
// ::: warning
// 调用该接口的用户或应用需与调用创建导出任务接口的用户或应用保持一致。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/get
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/export_task/get
func (r *DriveService) GetDriveExportTask(ctx context.Context, request *GetDriveExportTaskReq, options ...MethodOptionFunc) (*GetDriveExportTaskResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveExportTask != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetDriveExportTask mock enable")
		return r.cli.mock.mockDriveGetDriveExportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveExportTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/export_tasks/:ticket",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveExportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveExportTask mock DriveGetDriveExportTask method
func (r *Mock) MockDriveGetDriveExportTask(f func(ctx context.Context, request *GetDriveExportTaskReq, options ...MethodOptionFunc) (*GetDriveExportTaskResp, *Response, error)) {
	r.mockDriveGetDriveExportTask = f
}

// UnMockDriveGetDriveExportTask un-mock DriveGetDriveExportTask method
func (r *Mock) UnMockDriveGetDriveExportTask() {
	r.mockDriveGetDriveExportTask = nil
}

// GetDriveExportTaskReq ...
type GetDriveExportTaskReq struct {
	Ticket string `path:"ticket" json:"-"` // 导出任务 ID。调用[创建导出任务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/create) 获取, 示例值: "6933093124755412345"
	Token  string `query:"token" json:"-"` // 要导出的云文档的 token。获取方式参考[如何获取云文档相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)。你可参考以下请求示例了解如何使用查询参数, 示例值: docbcZVGtv1papC6jAVGiyabcef, 最大长度: `27` 字符
}

// GetDriveExportTaskResp ...
type GetDriveExportTaskResp struct {
	Result *GetDriveExportTaskRespResult `json:"result,omitempty"` // 导出任务结果
}

// GetDriveExportTaskRespResult ...
type GetDriveExportTaskRespResult struct {
	FileExtension string `json:"file_extension,omitempty"` // 导出的文件的扩展名, 可选值有: docx: Microsoft Word 格式, pdf: PDF 格式, xlsx: Microsoft Excel (XLSX) 格式, csv: CSV 格式
	Type          string `json:"type,omitempty"`           // 要导出的云文档的类型。可通过云文档的链接判断, 可选值有: doc: 旧版飞书文档。支持导出扩展名为 docx 和 pdf 的文件。, sheet: 飞书电子表格。支持导出扩展名为 xlsx 和 csv 的文件。, bitable: 飞书多维表格。支持导出扩展名为 xlsx 和 csv 格式的文件。, docx: 新版飞书文档。支持导出扩展名为 docx 和 pdf 格式的文件。
	FileName      string `json:"file_name,omitempty"`      // 导出的文件名称
	FileToken     string `json:"file_token,omitempty"`     // 导出的文件的 token。可用于调用[下载导出文件](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/download)接口将导出的产物下载到本地。
	FileSize      int64  `json:"file_size,omitempty"`      // 导出文件的大小, 单位字节。
	JobErrorMsg   string `json:"job_error_msg,omitempty"`  // 导出任务失败的原因
	JobStatus     int64  `json:"job_status,omitempty"`     // 导出任务状态, 可选值有: 0: 成功, 1: 初始化, 2: 处理中, 3: 内部错误, 107: 导出文档过大, 108: 处理超时, 109: 导出内容块无权限, 110: 无权限, 111: 导出文档已删除, 122: 创建副本中禁止导出, 123: 导出文档不存在, 6000: 导出文档图片过多
}

// getDriveExportTaskResp ...
type getDriveExportTaskResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *GetDriveExportTaskResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
