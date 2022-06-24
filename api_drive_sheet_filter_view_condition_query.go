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

// QuerySheetFilterViewCondition 查询一个筛选视图的所有筛选条件，返回筛选视图的筛选范围内的筛选条件。
//
// 筛选条件含义可参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/query
func (r *DriveService) QuerySheetFilterViewCondition(ctx context.Context, request *QuerySheetFilterViewConditionReq, options ...MethodOptionFunc) (*QuerySheetFilterViewConditionResp, *Response, error) {
	if r.cli.mock.mockDriveQuerySheetFilterViewCondition != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#QuerySheetFilterViewCondition mock enable")
		return r.cli.mock.mockDriveQuerySheetFilterViewCondition(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "QuerySheetFilterViewCondition",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(querySheetFilterViewConditionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveQuerySheetFilterViewCondition mock DriveQuerySheetFilterViewCondition method
func (r *Mock) MockDriveQuerySheetFilterViewCondition(f func(ctx context.Context, request *QuerySheetFilterViewConditionReq, options ...MethodOptionFunc) (*QuerySheetFilterViewConditionResp, *Response, error)) {
	r.mockDriveQuerySheetFilterViewCondition = f
}

// UnMockDriveQuerySheetFilterViewCondition un-mock DriveQuerySheetFilterViewCondition method
func (r *Mock) UnMockDriveQuerySheetFilterViewCondition() {
	r.mockDriveQuerySheetFilterViewCondition = nil
}

// QuerySheetFilterViewConditionReq ...
type QuerySheetFilterViewConditionReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值："pH9hbVcCXA"
}

// querySheetFilterViewConditionResp ...
type querySheetFilterViewConditionResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *QuerySheetFilterViewConditionResp `json:"data,omitempty"`
}

// QuerySheetFilterViewConditionResp ...
type QuerySheetFilterViewConditionResp struct {
	Items []*QuerySheetFilterViewConditionRespItem `json:"items,omitempty"` // 筛选视图设置的所有筛选条件
}

// QuerySheetFilterViewConditionRespItem ...
type QuerySheetFilterViewConditionRespItem struct {
	ConditionID string   `json:"condition_id,omitempty"` // 设置筛选条件的列，使用字母号
	FilterType  string   `json:"filter_type,omitempty"`  // 筛选类型
	CompareType string   `json:"compare_type,omitempty"` // 比较类型
	Expected    []string `json:"expected,omitempty"`     // 筛选参数
}
