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

// BatchGetCoreHRCompany 通过公司 ID 批量获取公司信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/company/batch_get
func (r *CoreHRService) BatchGetCoreHRCompany(ctx context.Context, request *BatchGetCoreHRCompanyReq, options ...MethodOptionFunc) (*BatchGetCoreHRCompanyResp, *Response, error) {
	if r.cli.mock.mockCoreHRBatchGetCoreHRCompany != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#BatchGetCoreHRCompany mock enable")
		return r.cli.mock.mockCoreHRBatchGetCoreHRCompany(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "BatchGetCoreHRCompany",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/companies/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetCoreHRCompanyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRBatchGetCoreHRCompany mock CoreHRBatchGetCoreHRCompany method
func (r *Mock) MockCoreHRBatchGetCoreHRCompany(f func(ctx context.Context, request *BatchGetCoreHRCompanyReq, options ...MethodOptionFunc) (*BatchGetCoreHRCompanyResp, *Response, error)) {
	r.mockCoreHRBatchGetCoreHRCompany = f
}

// UnMockCoreHRBatchGetCoreHRCompany un-mock CoreHRBatchGetCoreHRCompany method
func (r *Mock) UnMockCoreHRBatchGetCoreHRCompany() {
	r.mockCoreHRBatchGetCoreHRCompany = nil
}

// BatchGetCoreHRCompanyReq ...
type BatchGetCoreHRCompanyReq struct {
	CompanyIDs []string `json:"company_ids,omitempty"` // 公司 ID 列表, 示例值: ["151515"], 长度范围: `1` ～ `100`
}

// BatchGetCoreHRCompanyResp ...
type BatchGetCoreHRCompanyResp struct {
	Items []*BatchGetCoreHRCompanyRespItem `json:"items,omitempty"` // 查询的公司信息
}

// BatchGetCoreHRCompanyRespItem ...
type BatchGetCoreHRCompanyRespItem struct {
	CompanyID               string                                                 `json:"company_id,omitempty"`                // 公司 ID
	HiberarchyCommon        *BatchGetCoreHRCompanyRespItemHiberarchyCommon         `json:"hiberarchy_common,omitempty"`         // 公司基本信息
	Type                    *BatchGetCoreHRCompanyRespItemType                     `json:"type,omitempty"`                      // 性质
	IndustryList            []*BatchGetCoreHRCompanyRespItemIndustry               `json:"industry_list,omitempty"`             // 行业
	LegalRepresentative     []*BatchGetCoreHRCompanyRespItemLegalRepresentative    `json:"legal_representative,omitempty"`      // 法定代表人
	PostCode                string                                                 `json:"post_code,omitempty"`                 // 邮编
	TaxPayerID              string                                                 `json:"tax_payer_id,omitempty"`              // 纳税人识别号
	Confidential            bool                                                   `json:"confidential,omitempty"`              // 是否保密
	SubTypeList             []*BatchGetCoreHRCompanyRespItemSubType                `json:"sub_type_list,omitempty"`             // 主体类型
	BranchCompany           bool                                                   `json:"branch_company,omitempty"`            // 是否为分公司
	PrimaryManager          []*BatchGetCoreHRCompanyRespItemPrimaryManager         `json:"primary_manager,omitempty"`           // 主要负责人
	Currency                *BatchGetCoreHRCompanyRespItemCurrency                 `json:"currency,omitempty"`                  // 默认币种
	Phone                   *BatchGetCoreHRCompanyRespItemPhone                    `json:"phone,omitempty"`                     // 电话
	Fax                     *BatchGetCoreHRCompanyRespItemFax                      `json:"fax,omitempty"`                       // 传真
	RegisteredOfficeAddress []*BatchGetCoreHRCompanyRespItemRegisteredOfficeAddres `json:"registered_office_address,omitempty"` // 注册地址
	OfficeAddress           []*BatchGetCoreHRCompanyRespItemOfficeAddres           `json:"office_address,omitempty"`            // 办公地址
	CustomFields            []*BatchGetCoreHRCompanyRespItemCustomField            `json:"custom_fields,omitempty"`             // 自定义字段
}

// BatchGetCoreHRCompanyRespItemCurrency ...
type BatchGetCoreHRCompanyRespItemCurrency struct {
	CurrencyID         string                                               `json:"currency_id,omitempty"`           // 货币 ID
	CountryRegionID    string                                               `json:"country_region_id,omitempty"`     // 货币所属国家/地区 ID, 详细信息可通过[查询国家/地区信息](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/list)接口查询获得
	CurrencyName       []*BatchGetCoreHRCompanyRespItemCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                                `json:"numeric_code,omitempty"`          // 数字代码
	CurrencyAlpha3Code string                                               `json:"currency_alpha_3_code,omitempty"` // 三位字母代码
	Status             int64                                                `json:"status,omitempty"`                // 状态, 可选值有: 1: 生效, 0: 失效
}

// BatchGetCoreHRCompanyRespItemCurrencyCurrencyName ...
type BatchGetCoreHRCompanyRespItemCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemCustomField ...
type BatchGetCoreHRCompanyRespItemCustomField struct {
	CustomApiName string                                        `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *BatchGetCoreHRCompanyRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                         `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                        `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// BatchGetCoreHRCompanyRespItemCustomFieldName ...
type BatchGetCoreHRCompanyRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// BatchGetCoreHRCompanyRespItemFax ...
type BatchGetCoreHRCompanyRespItemFax struct {
	AreaCode    *BatchGetCoreHRCompanyRespItemFaxAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                    `json:"phone_number,omitempty"` // 号码
}

// BatchGetCoreHRCompanyRespItemFaxAreaCode ...
type BatchGetCoreHRCompanyRespItemFaxAreaCode struct {
	EnumName string                                             `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRCompanyRespItemFaxAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRCompanyRespItemFaxAreaCodeDisplay ...
type BatchGetCoreHRCompanyRespItemFaxAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemHiberarchyCommon ...
type BatchGetCoreHRCompanyRespItemHiberarchyCommon struct {
	ParentID       string                                                      `json:"parent_id,omitempty"`       // 上级
	Name           []*BatchGetCoreHRCompanyRespItemHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *BatchGetCoreHRCompanyRespItemHiberarchyCommonType          `json:"type,omitempty"`            // 类型
	Active         bool                                                        `json:"active,omitempty"`          // 启用
	EffectiveTime  string                                                      `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                      `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                      `json:"code,omitempty"`            // 编码
	Description    []*BatchGetCoreHRCompanyRespItemHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                      `json:"tree_order,omitempty"`      // 树形排序
	ListOrder      string                                                      `json:"list_order,omitempty"`      // 列表排序
	CustomFields   []*BatchGetCoreHRCompanyRespItemHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// BatchGetCoreHRCompanyRespItemHiberarchyCommonCustomField ...
type BatchGetCoreHRCompanyRespItemHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// BatchGetCoreHRCompanyRespItemHiberarchyCommonDescription ...
type BatchGetCoreHRCompanyRespItemHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemHiberarchyCommonName ...
type BatchGetCoreHRCompanyRespItemHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemHiberarchyCommonType ...
type BatchGetCoreHRCompanyRespItemHiberarchyCommonType struct {
	EnumName string                                                      `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRCompanyRespItemHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRCompanyRespItemHiberarchyCommonTypeDisplay ...
type BatchGetCoreHRCompanyRespItemHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemIndustry ...
type BatchGetCoreHRCompanyRespItemIndustry struct {
	EnumName string                                          `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRCompanyRespItemIndustryDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRCompanyRespItemIndustryDisplay ...
type BatchGetCoreHRCompanyRespItemIndustryDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemLegalRepresentative ...
type BatchGetCoreHRCompanyRespItemLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemOfficeAddres ...
type BatchGetCoreHRCompanyRespItemOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemPhone ...
type BatchGetCoreHRCompanyRespItemPhone struct {
	AreaCode    *BatchGetCoreHRCompanyRespItemPhoneAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                      `json:"phone_number,omitempty"` // 号码
}

// BatchGetCoreHRCompanyRespItemPhoneAreaCode ...
type BatchGetCoreHRCompanyRespItemPhoneAreaCode struct {
	EnumName string                                               `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRCompanyRespItemPhoneAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRCompanyRespItemPhoneAreaCodeDisplay ...
type BatchGetCoreHRCompanyRespItemPhoneAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemPrimaryManager ...
type BatchGetCoreHRCompanyRespItemPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemRegisteredOfficeAddres ...
type BatchGetCoreHRCompanyRespItemRegisteredOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemSubType ...
type BatchGetCoreHRCompanyRespItemSubType struct {
	EnumName string                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRCompanyRespItemSubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRCompanyRespItemSubTypeDisplay ...
type BatchGetCoreHRCompanyRespItemSubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRCompanyRespItemType ...
type BatchGetCoreHRCompanyRespItemType struct {
	EnumName string                                      `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRCompanyRespItemTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRCompanyRespItemTypeDisplay ...
type BatchGetCoreHRCompanyRespItemTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// batchGetCoreHRCompanyResp ...
type batchGetCoreHRCompanyResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *BatchGetCoreHRCompanyResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
