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

// CreateHireExternalInterviewAssessment 导入来自其他系统的面评信息, 创建为外部面评。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/external_interview_assessment/create
func (r *HireService) CreateHireExternalInterviewAssessment(ctx context.Context, request *CreateHireExternalInterviewAssessmentReq, options ...MethodOptionFunc) (*CreateHireExternalInterviewAssessmentResp, *Response, error) {
	if r.cli.mock.mockHireCreateHireExternalInterviewAssessment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#CreateHireExternalInterviewAssessment mock enable")
		return r.cli.mock.mockHireCreateHireExternalInterviewAssessment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CreateHireExternalInterviewAssessment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/external_interview_assessments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createHireExternalInterviewAssessmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCreateHireExternalInterviewAssessment mock HireCreateHireExternalInterviewAssessment method
func (r *Mock) MockHireCreateHireExternalInterviewAssessment(f func(ctx context.Context, request *CreateHireExternalInterviewAssessmentReq, options ...MethodOptionFunc) (*CreateHireExternalInterviewAssessmentResp, *Response, error)) {
	r.mockHireCreateHireExternalInterviewAssessment = f
}

// UnMockHireCreateHireExternalInterviewAssessment un-mock HireCreateHireExternalInterviewAssessment method
func (r *Mock) UnMockHireCreateHireExternalInterviewAssessment() {
	r.mockHireCreateHireExternalInterviewAssessment = nil
}

// CreateHireExternalInterviewAssessmentReq ...
type CreateHireExternalInterviewAssessmentReq struct {
	ExternalID              *string                                                        `json:"external_id,omitempty"`               // 外部系统面评主键（仅用于幂等）, 示例值: "123"
	Username                *string                                                        `json:"username,omitempty"`                  // 面试官姓名, 示例值: "shaojiale"
	Conclusion              *int64                                                         `json:"conclusion,omitempty"`                // 面试结果, 示例值: 1, 可选值有: 1: 不通过, 2: 通过, 3: 待定
	AssessmentDimensionList []*CreateHireExternalInterviewAssessmentReqAssessmentDimension `json:"assessment_dimension_list,omitempty"` // 评价维度列表
	Content                 *string                                                        `json:"content,omitempty"`                   // 综合记录, 示例值: "hello world"
	ExternalInterviewID     string                                                         `json:"external_interview_id,omitempty"`     // 外部面试 ID, 示例值: "6986199832494934316"
}

// CreateHireExternalInterviewAssessmentReqAssessmentDimension ...
type CreateHireExternalInterviewAssessmentReqAssessmentDimension struct {
	Score          *int64   `json:"score,omitempty"`           // 打分题分数（当题目类型为「打分题」时使用）, 示例值: 99
	Option         *string  `json:"option,omitempty"`          // 单选选项（当题目类型为「单选题」时使用）, 示例值: "opt"
	Options        []string `json:"options,omitempty"`         // 多选选项（当题目类型为「多选题」时使用）, 示例值: ["6989181065243969836"]
	Content        *string  `json:"content,omitempty"`         // 描述内容（当题目类型为「描述题」时使用）, 示例值: "content"
	AssessmentType *int64   `json:"assessment_type,omitempty"` // 题目类型, 示例值: 1, 可选值有: 1: 打分题, 2: 单选题, 3: 描述题, 4: 多选题
	Title          *string  `json:"title,omitempty"`           // 题目标题, 示例值: "title"
	Description    *string  `json:"description,omitempty"`     // 题目描述, 示例值: "desc"
}

// CreateHireExternalInterviewAssessmentResp ...
type CreateHireExternalInterviewAssessmentResp struct {
	ExternalInterviewAssessment *CreateHireExternalInterviewAssessmentRespExternalInterviewAssessment `json:"external_interview_assessment,omitempty"` // 外部面评信息
}

// CreateHireExternalInterviewAssessmentRespExternalInterviewAssessment ...
type CreateHireExternalInterviewAssessmentRespExternalInterviewAssessment struct {
	ID                      string                                                                                     `json:"id,omitempty"`                        // 外部面评 ID
	ExternalID              string                                                                                     `json:"external_id,omitempty"`               // 外部系统面评主键（仅用于幂等）
	Username                string                                                                                     `json:"username,omitempty"`                  // 面试官姓名
	Conclusion              int64                                                                                      `json:"conclusion,omitempty"`                // 面试结果, 可选值有: 1: 不通过, 2: 通过, 3: 待定
	AssessmentDimensionList []*CreateHireExternalInterviewAssessmentRespExternalInterviewAssessmentAssessmentDimension `json:"assessment_dimension_list,omitempty"` // 评价维度列表
	Content                 string                                                                                     `json:"content,omitempty"`                   // 综合记录
	ExternalInterviewID     string                                                                                     `json:"external_interview_id,omitempty"`     // 外部面试 ID
}

// CreateHireExternalInterviewAssessmentRespExternalInterviewAssessmentAssessmentDimension ...
type CreateHireExternalInterviewAssessmentRespExternalInterviewAssessmentAssessmentDimension struct {
	Score          int64    `json:"score,omitempty"`           // 打分题分数（当题目类型为「打分题」时使用）
	Option         string   `json:"option,omitempty"`          // 单选选项（当题目类型为「单选题」时使用）
	Options        []string `json:"options,omitempty"`         // 多选选项（当题目类型为「多选题」时使用）
	Content        string   `json:"content,omitempty"`         // 描述内容（当题目类型为「描述题」时使用）
	AssessmentType int64    `json:"assessment_type,omitempty"` // 题目类型, 可选值有: 1: 打分题, 2: 单选题, 3: 描述题, 4: 多选题
	Title          string   `json:"title,omitempty"`           // 题目标题
	Description    string   `json:"description,omitempty"`     // 题目描述
}

// createHireExternalInterviewAssessmentResp ...
type createHireExternalInterviewAssessmentResp struct {
	Code int64                                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateHireExternalInterviewAssessmentResp `json:"data,omitempty"`
}
