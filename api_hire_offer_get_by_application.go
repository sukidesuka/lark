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

// GetHireOfferByApplication 根据投递 ID 获取 Offer 信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/offer
func (r *HireService) GetHireOfferByApplication(ctx context.Context, request *GetHireOfferByApplicationReq, options ...MethodOptionFunc) (*GetHireOfferByApplicationResp, *Response, error) {
	if r.cli.mock.mockHireGetHireOfferByApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireOfferByApplication mock enable")
		return r.cli.mock.mockHireGetHireOfferByApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireOfferByApplication",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/applications/:application_id/offer",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireOfferByApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireOfferByApplication mock HireGetHireOfferByApplication method
func (r *Mock) MockHireGetHireOfferByApplication(f func(ctx context.Context, request *GetHireOfferByApplicationReq, options ...MethodOptionFunc) (*GetHireOfferByApplicationResp, *Response, error)) {
	r.mockHireGetHireOfferByApplication = f
}

// UnMockHireGetHireOfferByApplication un-mock HireGetHireOfferByApplication method
func (r *Mock) UnMockHireGetHireOfferByApplication() {
	r.mockHireGetHireOfferByApplication = nil
}

// GetHireOfferByApplicationReq ...
type GetHireOfferByApplicationReq struct {
	UserIDType    *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ApplicationID string  `path:"application_id" json:"-"` // 投递ID, 示例值："6949805467799537964"
}

// getHireOfferByApplicationResp ...
type getHireOfferByApplicationResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *GetHireOfferByApplicationResp `json:"data,omitempty"`
}

// GetHireOfferByApplicationResp ...
type GetHireOfferByApplicationResp struct {
	Offer *GetHireOfferByApplicationRespOffer `json:"offer,omitempty"` // Offer数据
}

// GetHireOfferByApplicationRespOffer ...
type GetHireOfferByApplicationRespOffer struct {
	ID            string                                        `json:"id,omitempty"`             // Offer id
	ApplicationID string                                        `json:"application_id,omitempty"` // 投递id
	BasicInfo     *GetHireOfferByApplicationRespOfferBasicInfo  `json:"basic_info,omitempty"`     // 基础信息
	SalaryPlan    *GetHireOfferByApplicationRespOfferSalaryPlan `json:"salary_plan,omitempty"`    // 薪酬计划
	SchemaID      string                                        `json:"schema_id,omitempty"`      // 当前 Offer 使用的 schema
	OfferStatus   int64                                         `json:"offer_status,omitempty"`   // Offer 状态, 可选值有: `0`：所有, `1`：未申请, `2`：审批中, `3`：审批已撤回, `4`：审批通过, `5`：审批不通过, `6`：Offer 已发出, `7`：候选人已接受, `8`：候选人已拒绝, `9`：Offer 已失效
	JobInfo       *GetHireOfferByApplicationRespOfferJobInfo    `json:"job_info,omitempty"`       // 职位信息
}

// GetHireOfferByApplicationRespOfferBasicInfo ...
type GetHireOfferByApplicationRespOfferBasicInfo struct {
	OfferType         int64                                                       `json:"offer_type,omitempty"`          // Offer 类型, 可选值有: `1`：Social, `2`：Campus, `3`：Intern, `4`：InternTransfer
	Remark            string                                                      `json:"remark,omitempty"`              // 备注
	ExpireTime        int64                                                       `json:"expire_time,omitempty"`         // Offer 过期时间
	OwnerUserID       string                                                      `json:"owner_user_id,omitempty"`       // Offer 负责人 ID
	CreatorUserID     string                                                      `json:"creator_user_id,omitempty"`     // Offer 创建人 ID
	EmployeeType      *GetHireOfferByApplicationRespOfferBasicInfoEmployeeType    `json:"employee_type,omitempty"`       // Offer 人员类型
	CreateTime        string                                                      `json:"create_time,omitempty"`         // 创建时间
	LeaderUserID      string                                                      `json:"leader_user_id,omitempty"`      // 直属上级 ID
	OnboardDate       string                                                      `json:"onboard_date,omitempty"`        // 入职日期
	DepartmentID      string                                                      `json:"department_id,omitempty"`       // 入职部门
	ProbationMonth    int64                                                       `json:"probation_month,omitempty"`     // 试用期, 比如试用期6个月
	ContractYear      int64                                                       `json:"contract_year,omitempty"`       // 合同期, 比如3年
	RecruitmentType   *GetHireOfferByApplicationRespOfferBasicInfoRecruitmentType `json:"recruitment_type,omitempty"`    // 雇员类型
	Sequence          *GetHireOfferByApplicationRespOfferBasicInfoSequence        `json:"sequence,omitempty"`            // 序列
	Level             *GetHireOfferByApplicationRespOfferBasicInfoLevel           `json:"level,omitempty"`               // 级别
	OnboardAddress    *GetHireOfferByApplicationRespOfferBasicInfoOnboardAddress  `json:"onboard_address,omitempty"`     // 入职地点
	WorkAddress       *GetHireOfferByApplicationRespOfferBasicInfoWorkAddress     `json:"work_address,omitempty"`        // 工作地点
	CustomizeInfoList []*GetHireOfferByApplicationRespOfferBasicInfoCustomizeInfo `json:"customize_info_list,omitempty"` // 自定义字段信息
}

// GetHireOfferByApplicationRespOfferBasicInfoEmployeeType ...
type GetHireOfferByApplicationRespOfferBasicInfoEmployeeType struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
}

// GetHireOfferByApplicationRespOfferBasicInfoRecruitmentType ...
type GetHireOfferByApplicationRespOfferBasicInfoRecruitmentType struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
}

// GetHireOfferByApplicationRespOfferBasicInfoSequence ...
type GetHireOfferByApplicationRespOfferBasicInfoSequence struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
}

// GetHireOfferByApplicationRespOfferBasicInfoLevel ...
type GetHireOfferByApplicationRespOfferBasicInfoLevel struct {
	ID     string `json:"id,omitempty"`      // ID
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
}

// GetHireOfferByApplicationRespOfferBasicInfoOnboardAddress ...
type GetHireOfferByApplicationRespOfferBasicInfoOnboardAddress struct {
	ID       string                                                             `json:"id,omitempty"`       // ID
	ZhName   string                                                             `json:"zh_name,omitempty"`  // 中文名称
	EnName   string                                                             `json:"en_name,omitempty"`  // 英文名称
	District *GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressDistrict `json:"district,omitempty"` // 区域信息
	City     *GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressCity     `json:"city,omitempty"`     // 城市信息
	State    *GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressState    `json:"state,omitempty"`    // 省信息
	Country  *GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressCountry  `json:"country,omitempty"`  // 国家信息
}

// GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressDistrict ...
type GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressDistrict struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型
}

// GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressCity ...
type GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressCity struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型, 可选值有: `1`：COUNTRY, `2`：STATE, `3`：CITY, `4`：DISTRICT, `5`：ADDRESS
}

// GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressState ...
type GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressState struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型, 可选值有: `1`：COUNTRY, `2`：STATE, `3`：CITY, `4`：DISTRICT, `5`：ADDRESS
}

// GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressCountry ...
type GetHireOfferByApplicationRespOfferBasicInfoOnboardAddressCountry struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型, 可选值有: `1`：COUNTRY, `2`：STATE, `3`：CITY, `4`：DISTRICT, `5`：ADDRESS
}

// GetHireOfferByApplicationRespOfferBasicInfoWorkAddress ...
type GetHireOfferByApplicationRespOfferBasicInfoWorkAddress struct {
	ID       string                                                          `json:"id,omitempty"`       // ID
	ZhName   string                                                          `json:"zh_name,omitempty"`  // 中文名称
	EnName   string                                                          `json:"en_name,omitempty"`  // 英文名称
	District *GetHireOfferByApplicationRespOfferBasicInfoWorkAddressDistrict `json:"district,omitempty"` // 区域信息
	City     *GetHireOfferByApplicationRespOfferBasicInfoWorkAddressCity     `json:"city,omitempty"`     // 城市信息
	State    *GetHireOfferByApplicationRespOfferBasicInfoWorkAddressState    `json:"state,omitempty"`    // 省信息
	Country  *GetHireOfferByApplicationRespOfferBasicInfoWorkAddressCountry  `json:"country,omitempty"`  // 国家信息
}

// GetHireOfferByApplicationRespOfferBasicInfoWorkAddressDistrict ...
type GetHireOfferByApplicationRespOfferBasicInfoWorkAddressDistrict struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型
}

// GetHireOfferByApplicationRespOfferBasicInfoWorkAddressCity ...
type GetHireOfferByApplicationRespOfferBasicInfoWorkAddressCity struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型, 可选值有: `1`：COUNTRY, `2`：STATE, `3`：CITY, `4`：DISTRICT, `5`：ADDRESS
}

// GetHireOfferByApplicationRespOfferBasicInfoWorkAddressState ...
type GetHireOfferByApplicationRespOfferBasicInfoWorkAddressState struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型, 可选值有: `1`：COUNTRY, `2`：STATE, `3`：CITY, `4`：DISTRICT, `5`：ADDRESS
}

// GetHireOfferByApplicationRespOfferBasicInfoWorkAddressCountry ...
type GetHireOfferByApplicationRespOfferBasicInfoWorkAddressCountry struct {
	ZhName       string `json:"zh_name,omitempty"`       // 中文名称
	EnName       string `json:"en_name,omitempty"`       // 英文名称
	Code         string `json:"code,omitempty"`          // 编码
	LocationType int64  `json:"location_type,omitempty"` // 地址类型, 可选值有: `1`：COUNTRY, `2`：STATE, `3`：CITY, `4`：DISTRICT, `5`：ADDRESS
}

// GetHireOfferByApplicationRespOfferBasicInfoCustomizeInfo ...
type GetHireOfferByApplicationRespOfferBasicInfoCustomizeInfo struct {
	ObjectID       string `json:"object_id,omitempty"`       // 自定义字段 ID
	CustomizeValue string `json:"customize_value,omitempty"` // 自定义字段 value
}

// GetHireOfferByApplicationRespOfferSalaryPlan ...
type GetHireOfferByApplicationRespOfferSalaryPlan struct {
	Currency                  string                                                       `json:"currency,omitempty"`                    // 币种
	BasicSalary               string                                                       `json:"basic_salary,omitempty"`                // 基本薪资，为JSON 格式，amount 代表基本薪资的金额，peroid 代表基本薪资的周期单位，如："{"amount":"10000","period":2}"
	ProbationSalaryPercentage string                                                       `json:"probation_salary_percentage,omitempty"` // 试用期百分比
	AwardSalaryMultiple       string                                                       `json:"award_salary_multiple,omitempty"`       // 年终奖月数
	OptionShares              string                                                       `json:"option_shares,omitempty"`               // 期权股数
	QuarterlyBonus            string                                                       `json:"quarterly_bonus,omitempty"`             // 季度奖金额
	HalfYearBonus             string                                                       `json:"half_year_bonus,omitempty"`             // 半年奖金额
	TotalAnnualCash           string                                                       `json:"total_annual_cash,omitempty"`           // 年度现金总额(数量，非公式)
	CustomizeInfoList         []*GetHireOfferByApplicationRespOfferSalaryPlanCustomizeInfo `json:"customize_info_list,omitempty"`         // 自定义字段的 value 信息
}

// GetHireOfferByApplicationRespOfferSalaryPlanCustomizeInfo ...
type GetHireOfferByApplicationRespOfferSalaryPlanCustomizeInfo struct {
	ObjectID       string `json:"object_id,omitempty"`       // 自定义字段 ID
	CustomizeValue string `json:"customize_value,omitempty"` // 自定义字段 value
}

// GetHireOfferByApplicationRespOfferJobInfo ...
type GetHireOfferByApplicationRespOfferJobInfo struct {
	JobID   string `json:"job_id,omitempty"`   // Offer 职位 ID
	JobName string `json:"job_name,omitempty"` // Offer 职位名称
}
