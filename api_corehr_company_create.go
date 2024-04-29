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

// CreateCoreHRCompany 创建公司。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/company/create
func (r *CoreHRService) CreateCoreHRCompany(ctx context.Context, request *CreateCoreHRCompanyReq, options ...MethodOptionFunc) (*CreateCoreHRCompanyResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRCompany != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRCompany mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRCompany(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRCompany",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/companies",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRCompanyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRCompany mock CoreHRCreateCoreHRCompany method
func (r *Mock) MockCoreHRCreateCoreHRCompany(f func(ctx context.Context, request *CreateCoreHRCompanyReq, options ...MethodOptionFunc) (*CreateCoreHRCompanyResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRCompany = f
}

// UnMockCoreHRCreateCoreHRCompany un-mock CoreHRCreateCoreHRCompany method
func (r *Mock) UnMockCoreHRCreateCoreHRCompany() {
	r.mockCoreHRCreateCoreHRCompany = nil
}

// CreateCoreHRCompanyReq ...
type CreateCoreHRCompanyReq struct {
	ClientToken                 *string                                            `query:"client_token" json:"-"`                   // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	HiberarchyCommon            *CreateCoreHRCompanyReqHiberarchyCommon            `json:"hiberarchy_common,omitempty"`              // 层级关系, 内层字段见实体
	Type                        *CreateCoreHRCompanyReqType                        `json:"type,omitempty"`                           // 性质, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)公司类型（company_type）枚举定义部分获得。该字段为通用字段, 若为公司维度则为必填。
	IndustryList                []*CreateCoreHRCompanyReqIndustry                  `json:"industry_list,omitempty"`                  // 行业, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)行业（industry）枚举定义部分获得
	LegalRepresentative         []*CreateCoreHRCompanyReqLegalRepresentative       `json:"legal_representative,omitempty"`           // 法定代表人, 仅注册地址中的 国家 / 地区为中国大陆时, 法人字段填入才有效, 若注册地址中的 国家 / 地区 不为中国大陆时, 则填入法人字段无效。
	PostCode                    *string                                            `json:"post_code,omitempty"`                      // 邮编, 示例值: "邮编"
	TaxPayerID                  *string                                            `json:"tax_payer_id,omitempty"`                   // 纳税人识别号, 示例值: "123456840"
	Confidential                *bool                                              `json:"confidential,omitempty"`                   // 是否保密, 示例值: true
	SubTypeList                 []*CreateCoreHRCompanyReqSubType                   `json:"sub_type_list,omitempty"`                  // 主体类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)主体类型（company_sub_type）枚举定义部分获得
	BranchCompany               *bool                                              `json:"branch_company,omitempty"`                 // 是否为分公司, 示例值: true
	PrimaryManager              []*CreateCoreHRCompanyReqPrimaryManager            `json:"primary_manager,omitempty"`                // 主要负责人
	CustomFields                []*CreateCoreHRCompanyReqCustomField               `json:"custom_fields,omitempty"`                  // 自定义字段
	Currency                    *CreateCoreHRCompanyReqCurrency                    `json:"currency,omitempty"`                       // 默认币种
	Phone                       *CreateCoreHRCompanyReqPhone                       `json:"phone,omitempty"`                          // 电话
	Fax                         *CreateCoreHRCompanyReqFax                         `json:"fax,omitempty"`                            // 传真
	RegisteredOfficeAddressInfo *CreateCoreHRCompanyReqRegisteredOfficeAddressInfo `json:"registered_office_address_info,omitempty"` // 注册地址详细信息。公共字段, 若请求对象为公司, 则该字段必填。
	OfficeAddressInfo           *CreateCoreHRCompanyReqOfficeAddressInfo           `json:"office_address_info,omitempty"`            // 办公地址详细信息。公共字段, 若请求对象为公司, 则该字段必填。
}

// CreateCoreHRCompanyReqCurrency ...
type CreateCoreHRCompanyReqCurrency struct {
	CurrencyName       []*CreateCoreHRCompanyReqCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        *int64                                        `json:"numeric_code,omitempty"`          // 对应币种的指代代码, 通过[查询货币信息v2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-currency/search)查询获取, 示例值: 12
	CurrencyAlpha3Code *string                                       `json:"currency_alpha_3_code,omitempty"` // 法定货币对应代码, 如CNY	、USD等, 通过[查询货币信息v2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-currency/search)查询获取, 示例值: "CNY"
}

// CreateCoreHRCompanyReqCurrencyCurrencyName ...
type CreateCoreHRCompanyReqCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 支持中文和英文。中文用zh-CN；英文用en-US, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 支持中文和英文, 示例值: "张三"
}

// CreateCoreHRCompanyReqCustomField ...
type CreateCoreHRCompanyReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// CreateCoreHRCompanyReqFax ...
type CreateCoreHRCompanyReqFax struct {
	AreaCode    *CreateCoreHRCompanyReqFaxAreaCode `json:"area_code,omitempty"`    // 区号对应的数字, 可通过, [请求接口](https://open.larkoffice.com/document/server-docs/corehr-v1/basic-infomation/custom_field/get_by_param)查询获取。请求参数: object_api_name=phone；custom_api_name=international_area_code, 示例值: 123123
	PhoneNumber string                             `json:"phone_number,omitempty"` // 号码, 示例值: "18812341234"
}

// CreateCoreHRCompanyReqFaxAreaCode ...
type CreateCoreHRCompanyReqFaxAreaCode struct {
	EnumName string `json:"enum_name,omitempty"` // 区号对应的名称, 示例值: "86_china"
}

// CreateCoreHRCompanyReqHiberarchyCommon ...
type CreateCoreHRCompanyReqHiberarchyCommon struct {
	ParentID       *string                                              `json:"parent_id,omitempty"`       // 上级组织 ID, 示例值: "4719168654814483759"
	Name           []*CreateCoreHRCompanyReqHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *CreateCoreHRCompanyReqHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得。该字段为通用字段, 若为公司维度则为必填。
	Active         bool                                                 `json:"active,omitempty"`          // 是否启用该公司, 示例值: true
	EffectiveTime  *string                                              `json:"effective_time,omitempty"`  // 生效时间, 示例值: "2020-05-01 00:00:00"
	ExpirationTime *string                                              `json:"expiration_time,omitempty"` // 失效时间, 示例值: "2020-05-02 00:00:00"
	Code           *string                                              `json:"code,omitempty"`            // 公司编码, 示例值: "12456"
	Description    []*CreateCoreHRCompanyReqHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	CustomFields   []*CreateCoreHRCompanyReqHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// CreateCoreHRCompanyReqHiberarchyCommonCustomField ...
type CreateCoreHRCompanyReqHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// CreateCoreHRCompanyReqHiberarchyCommonDescription ...
type CreateCoreHRCompanyReqHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文, 示例值: "xx有限科技公司"
}

// CreateCoreHRCompanyReqHiberarchyCommonName ...
type CreateCoreHRCompanyReqHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文, 示例值: "xx有限科技公司"
}

// CreateCoreHRCompanyReqHiberarchyCommonType ...
type CreateCoreHRCompanyReqHiberarchyCommonType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRCompanyReqIndustry ...
type CreateCoreHRCompanyReqIndustry struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRCompanyReqLegalRepresentative ...
type CreateCoreHRCompanyReqLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文, 示例值: "张三"
}

// CreateCoreHRCompanyReqOfficeAddressInfo ...
type CreateCoreHRCompanyReqOfficeAddressInfo struct {
	CountryRegionID   string  `json:"country_region_id,omitempty"`   // 国家 / 地区id, 若选择中国大陆、中国香港、中国澳门, 则需要指定主要行政区（中国大陆为省份）、城市、区/县（中国香港为地区） ；, 若选择中国台湾, 则需要指定主要行政区（中国台湾为省/地区）、城市、区/县（中国台湾为区）。可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)查询获取, 示例值: "6862995757234914824"
	RegionID          *string `json:"region_id,omitempty"`           // 主要行政区id。可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region_subdivision/search)查询获取, 示例值: "6863326815667095047"
	CityID            *string `json:"city_id,omitempty"`             // 城市id, 可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-city/search)查询获取, 示例值: "6863333254578046471"
	DistinctID        *string `json:"distinct_id,omitempty"`         // 区/县id, 可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-district/search)查询获取, 示例值: "6863333516579440141"
	LocalAddressLine1 *string `json:"local_address_line1,omitempty"` // 地址行 1（非拉丁语系的本地文字）, 示例值: "丹佛测试地址-纽埃时区"
	LocalAddressLine2 *string `json:"local_address_line2,omitempty"` // 地址行 2（非拉丁语系的本地文字）, 示例值: "PoewH"
	LocalAddressLine3 *string `json:"local_address_line3,omitempty"` // 地址行 3（非拉丁语系的本地文字）, 示例值: "PoewH"
	LocalAddressLine4 *string `json:"local_address_line4,omitempty"` // 地址行 4（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine5 *string `json:"local_address_line5,omitempty"` // 地址行 5（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine6 *string `json:"local_address_line6,omitempty"` // 地址行 6（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine7 *string `json:"local_address_line7,omitempty"` // 地址行 7（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine8 *string `json:"local_address_line8,omitempty"` // 地址行 8（非拉丁语系的本地文字）, 示例值: "rafSu"
	LocalAddressLine9 *string `json:"local_address_line9,omitempty"` // 地址行 9（非拉丁语系的本地文字）, 示例值: "McPRG"
	PostalCode        *string `json:"postal_code,omitempty"`         // 邮政编码, 示例值: "611530"
}

// CreateCoreHRCompanyReqPhone ...
type CreateCoreHRCompanyReqPhone struct {
	AreaCode    *CreateCoreHRCompanyReqPhoneAreaCode `json:"area_code,omitempty"`    // 区号对应的数字, 可通过, [请求接口](https://open.larkoffice.com/document/server-docs/corehr-v1/basic-infomation/custom_field/get_by_param)查询获取。请求参数: object_api_name=phone；custom_api_name=international_area_code, 示例值: 123123
	PhoneNumber string                               `json:"phone_number,omitempty"` // 号码, 示例值: "18812341234"
}

// CreateCoreHRCompanyReqPhoneAreaCode ...
type CreateCoreHRCompanyReqPhoneAreaCode struct {
	EnumName string `json:"enum_name,omitempty"` // 区号对应名称, 示例值: "86_china"
}

// CreateCoreHRCompanyReqPrimaryManager ...
type CreateCoreHRCompanyReqPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文, 示例值: "张三"
}

// CreateCoreHRCompanyReqRegisteredOfficeAddressInfo ...
type CreateCoreHRCompanyReqRegisteredOfficeAddressInfo struct {
	CountryRegionID   string  `json:"country_region_id,omitempty"`   // 国家 / 地区id, 若选择中国大陆、中国香港、中国澳门, 则需要指定主要行政区（中国大陆为省份）、城市、区/县（中国香港为地区） ；, 若选择中国台湾, 则需要指定主要行政区（中国台湾为省/地区）、城市、区/县（中国台湾为区）, 可通过, [请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)查询获取, 示例值: "6862995757234914824"
	RegionID          *string `json:"region_id,omitempty"`           // 主要行政区id, 可通过, [请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region_subdivision/search)查询获取, 示例值: "6863326815667095047"
	CityID            *string `json:"city_id,omitempty"`             // 城市id, 可通过, [请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-city/search)查询获取, 示例值: "6863333254578046471"
	DistinctID        *string `json:"distinct_id,omitempty"`         // 区/县id, 可通过, [请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-district/search)查询获取, 示例值: "6863333516579440141"
	LocalAddressLine1 *string `json:"local_address_line1,omitempty"` // 地址行 1（非拉丁语系的本地文字）, 示例值: "丹佛测试地址-纽埃时区"
	LocalAddressLine2 *string `json:"local_address_line2,omitempty"` // 地址行 2（非拉丁语系的本地文字）, 示例值: "PoewH"
	LocalAddressLine3 *string `json:"local_address_line3,omitempty"` // 地址行 3（非拉丁语系的本地文字）, 示例值: "PoewH"
	LocalAddressLine4 *string `json:"local_address_line4,omitempty"` // 地址行 4（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine5 *string `json:"local_address_line5,omitempty"` // 地址行 5（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine6 *string `json:"local_address_line6,omitempty"` // 地址行 6（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine7 *string `json:"local_address_line7,omitempty"` // 地址行 7（非拉丁语系的本地文字）, 示例值: "jmwJc"
	LocalAddressLine8 *string `json:"local_address_line8,omitempty"` // 地址行 8（非拉丁语系的本地文字）, 示例值: "rafSu"
	LocalAddressLine9 *string `json:"local_address_line9,omitempty"` // 地址行 9（非拉丁语系的本地文字）, 示例值: "McPRG"
	PostalCode        *string `json:"postal_code,omitempty"`         // 邮政编码, 示例值: "611530"
}

// CreateCoreHRCompanyReqSubType ...
type CreateCoreHRCompanyReqSubType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRCompanyReqType ...
type CreateCoreHRCompanyReqType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRCompanyResp ...
type CreateCoreHRCompanyResp struct {
	Company *CreateCoreHRCompanyRespCompany `json:"company,omitempty"` // 创建成功的公司信息
}

// CreateCoreHRCompanyRespCompany ...
type CreateCoreHRCompanyRespCompany struct {
	ID                          string                                                     `json:"id,omitempty"`                             // 公司 ID
	HiberarchyCommon            *CreateCoreHRCompanyRespCompanyHiberarchyCommon            `json:"hiberarchy_common,omitempty"`              // 层级关系, 内层字段见实体
	Type                        *CreateCoreHRCompanyRespCompanyType                        `json:"type,omitempty"`                           // 性质, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)公司类型（company_type）枚举定义部分获得。该字段为通用字段, 若为公司维度则为必填。
	IndustryList                []*CreateCoreHRCompanyRespCompanyIndustry                  `json:"industry_list,omitempty"`                  // 行业, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)行业（industry）枚举定义部分获得
	LegalRepresentative         []*CreateCoreHRCompanyRespCompanyLegalRepresentative       `json:"legal_representative,omitempty"`           // 法定代表人, 仅注册地址中的 国家 / 地区为中国大陆时, 法人字段填入才有效, 若注册地址中的 国家 / 地区 不为中国大陆时, 则填入法人字段无效。
	PostCode                    string                                                     `json:"post_code,omitempty"`                      // 邮编
	TaxPayerID                  string                                                     `json:"tax_payer_id,omitempty"`                   // 纳税人识别号
	Confidential                bool                                                       `json:"confidential,omitempty"`                   // 是否保密
	SubTypeList                 []*CreateCoreHRCompanyRespCompanySubType                   `json:"sub_type_list,omitempty"`                  // 主体类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)主体类型（company_sub_type）枚举定义部分获得
	BranchCompany               bool                                                       `json:"branch_company,omitempty"`                 // 是否为分公司
	PrimaryManager              []*CreateCoreHRCompanyRespCompanyPrimaryManager            `json:"primary_manager,omitempty"`                // 主要负责人
	CustomFields                []*CreateCoreHRCompanyRespCompanyCustomField               `json:"custom_fields,omitempty"`                  // 自定义字段
	Currency                    *CreateCoreHRCompanyRespCompanyCurrency                    `json:"currency,omitempty"`                       // 默认币种
	Phone                       *CreateCoreHRCompanyRespCompanyPhone                       `json:"phone,omitempty"`                          // 电话
	Fax                         *CreateCoreHRCompanyRespCompanyFax                         `json:"fax,omitempty"`                            // 传真
	RegisteredOfficeAddress     []*CreateCoreHRCompanyRespCompanyRegisteredOfficeAddres    `json:"registered_office_address,omitempty"`      // 完整注册地址
	OfficeAddress               []*CreateCoreHRCompanyRespCompanyOfficeAddres              `json:"office_address,omitempty"`                 // 完整办公地址
	RegisteredOfficeAddressInfo *CreateCoreHRCompanyRespCompanyRegisteredOfficeAddressInfo `json:"registered_office_address_info,omitempty"` // 注册地址详细信息。公共字段, 若请求对象为公司, 则该字段必填。
	OfficeAddressInfo           *CreateCoreHRCompanyRespCompanyOfficeAddressInfo           `json:"office_address_info,omitempty"`            // 办公地址详细信息。公共字段, 若请求对象为公司, 则该字段必填。
}

// CreateCoreHRCompanyRespCompanyCurrency ...
type CreateCoreHRCompanyRespCompanyCurrency struct {
	ID                 string                                                `json:"id,omitempty"`                    // 货币id
	CountryRegionID    string                                                `json:"country_region_id,omitempty"`     // 货币所属国家/地区id, 详细信息可通过[查询国家/地区信息](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/list)接口查询获得
	CurrencyName       []*CreateCoreHRCompanyRespCompanyCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                                 `json:"numeric_code,omitempty"`          // 对应币种的指代代码, 通过[查询货币信息v2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-currency/search)查询获取。
	CurrencyAlpha3Code string                                                `json:"currency_alpha_3_code,omitempty"` // 法定货币对应代码, 如CNY	、USD等, 通过[查询货币信息v2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-currency/search)查询获取。
}

// CreateCoreHRCompanyRespCompanyCurrencyCurrencyName ...
type CreateCoreHRCompanyRespCompanyCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 货币语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 货币内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyCustomField ...
type CreateCoreHRCompanyRespCompanyCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHRCompanyRespCompanyFax ...
type CreateCoreHRCompanyRespCompanyFax struct {
	AreaCode    *CreateCoreHRCompanyRespCompanyFaxAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                     `json:"phone_number,omitempty"` // 号码
}

// CreateCoreHRCompanyRespCompanyFaxAreaCode ...
type CreateCoreHRCompanyRespCompanyFaxAreaCode struct {
	EnumName string                                              `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRCompanyRespCompanyFaxAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRCompanyRespCompanyFaxAreaCodeDisplay ...
type CreateCoreHRCompanyRespCompanyFaxAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyHiberarchyCommon ...
type CreateCoreHRCompanyRespCompanyHiberarchyCommon struct {
	ParentID       string                                                       `json:"parent_id,omitempty"`       // 上级组织 ID
	Name           []*CreateCoreHRCompanyRespCompanyHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *CreateCoreHRCompanyRespCompanyHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得。该字段为通用字段, 若为公司维度则为必填。
	Active         bool                                                         `json:"active,omitempty"`          // 是否启用
	EffectiveTime  string                                                       `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                       `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                       `json:"code,omitempty"`            // 编码
	Description    []*CreateCoreHRCompanyRespCompanyHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                       `json:"tree_order,omitempty"`      // 树形排序, 代表同层级的部门排序序号
	ListOrder      string                                                       `json:"list_order,omitempty"`      // 列表排序, 代表所有部门的混排序号
	CustomFields   []*CreateCoreHRCompanyRespCompanyHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// CreateCoreHRCompanyRespCompanyHiberarchyCommonCustomField ...
type CreateCoreHRCompanyRespCompanyHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHRCompanyRespCompanyHiberarchyCommonDescription ...
type CreateCoreHRCompanyRespCompanyHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyHiberarchyCommonName ...
type CreateCoreHRCompanyRespCompanyHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文。
}

// CreateCoreHRCompanyRespCompanyHiberarchyCommonType ...
type CreateCoreHRCompanyRespCompanyHiberarchyCommonType struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay ...
type CreateCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyIndustry ...
type CreateCoreHRCompanyRespCompanyIndustry struct {
	EnumName string                                           `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRCompanyRespCompanyIndustryDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRCompanyRespCompanyIndustryDisplay ...
type CreateCoreHRCompanyRespCompanyIndustryDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyLegalRepresentative ...
type CreateCoreHRCompanyRespCompanyLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyOfficeAddres ...
type CreateCoreHRCompanyRespCompanyOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyOfficeAddressInfo ...
type CreateCoreHRCompanyRespCompanyOfficeAddressInfo struct {
	CountryRegionID   string `json:"country_region_id,omitempty"`   // 国家 / 地区id。若选择中国大陆、中国香港、中国澳门, 则需要指定主要行政区（中国大陆为省份）、城市、区/县（中国香港为地区） ；, 若选择中国台湾, 则需要指定主要行政区（中国台湾为省/地区）、城市、区/县（中国台湾为区）。可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)查询获取。
	RegionID          string `json:"region_id,omitempty"`           // 主要行政区id, 可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region_subdivision/search)查询获取。
	CityID            string `json:"city_id,omitempty"`             // 城市id, 可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-city/search)查询获取。
	DistinctID        string `json:"distinct_id,omitempty"`         // 区/县id, 可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-district/search)查询获取。
	LocalAddressLine1 string `json:"local_address_line1,omitempty"` // 地址行 1（非拉丁语系的本地文字）
	LocalAddressLine2 string `json:"local_address_line2,omitempty"` // 地址行 2（非拉丁语系的本地文字）
	LocalAddressLine3 string `json:"local_address_line3,omitempty"` // 地址行 3（非拉丁语系的本地文字）
	LocalAddressLine4 string `json:"local_address_line4,omitempty"` // 地址行 4（非拉丁语系的本地文字）
	LocalAddressLine5 string `json:"local_address_line5,omitempty"` // 地址行 5（非拉丁语系的本地文字）
	LocalAddressLine6 string `json:"local_address_line6,omitempty"` // 地址行 6（非拉丁语系的本地文字）
	LocalAddressLine7 string `json:"local_address_line7,omitempty"` // 地址行 7（非拉丁语系的本地文字）
	LocalAddressLine8 string `json:"local_address_line8,omitempty"` // 地址行 8（非拉丁语系的本地文字）
	LocalAddressLine9 string `json:"local_address_line9,omitempty"` // 地址行 9（非拉丁语系的本地文字）
	PostalCode        string `json:"postal_code,omitempty"`         // 邮政编码
}

// CreateCoreHRCompanyRespCompanyPhone ...
type CreateCoreHRCompanyRespCompanyPhone struct {
	AreaCode    *CreateCoreHRCompanyRespCompanyPhoneAreaCode `json:"area_code,omitempty"`    // 区号对应的数字, 可通过, [请求接口](https://open.larkoffice.com/document/server-docs/corehr-v1/basic-infomation/custom_field/get_by_param)查询获取。请求参数: object_api_name=phone；custom_api_name=international_area_code
	PhoneNumber string                                       `json:"phone_number,omitempty"` // 号码
}

// CreateCoreHRCompanyRespCompanyPhoneAreaCode ...
type CreateCoreHRCompanyRespCompanyPhoneAreaCode struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 区号对应的名称
	Display  []*CreateCoreHRCompanyRespCompanyPhoneAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRCompanyRespCompanyPhoneAreaCodeDisplay ...
type CreateCoreHRCompanyRespCompanyPhoneAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyPrimaryManager ...
type CreateCoreHRCompanyRespCompanyPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyRegisteredOfficeAddres ...
type CreateCoreHRCompanyRespCompanyRegisteredOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyRegisteredOfficeAddressInfo ...
type CreateCoreHRCompanyRespCompanyRegisteredOfficeAddressInfo struct {
	CountryRegionID   string `json:"country_region_id,omitempty"`   // 国家 / 地区id。若选择中国大陆、中国香港、中国澳门, 则需要指定主要行政区（中国大陆为省份）、城市、区/县（中国香港为地区）, 若选择中国台湾, 则需要指定主要行政区（中国台湾为省/地区）、城市、区/县（中国台湾为区）。可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)查询获取。
	RegionID          string `json:"region_id,omitempty"`           // 主要行政区id。可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region_subdivision/search)查询获取。
	CityID            string `json:"city_id,omitempty"`             // 城市id。可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-city/search)查询获取。
	DistinctID        string `json:"distinct_id,omitempty"`         // 区/县id, 可通过[请求接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-district/search)查询获取。
	LocalAddressLine1 string `json:"local_address_line1,omitempty"` // 地址行 1（非拉丁语系的本地文字）
	LocalAddressLine2 string `json:"local_address_line2,omitempty"` // 地址行 2（非拉丁语系的本地文字）
	LocalAddressLine3 string `json:"local_address_line3,omitempty"` // 地址行 3（非拉丁语系的本地文字）
	LocalAddressLine4 string `json:"local_address_line4,omitempty"` // 地址行 4（非拉丁语系的本地文字）
	LocalAddressLine5 string `json:"local_address_line5,omitempty"` // 地址行 5（非拉丁语系的本地文字）
	LocalAddressLine6 string `json:"local_address_line6,omitempty"` // 地址行 6（非拉丁语系的本地文字）
	LocalAddressLine7 string `json:"local_address_line7,omitempty"` // 地址行 7（非拉丁语系的本地文字）
	LocalAddressLine8 string `json:"local_address_line8,omitempty"` // 地址行 8（非拉丁语系的本地文字）
	LocalAddressLine9 string `json:"local_address_line9,omitempty"` // 地址行 9（非拉丁语系的本地文字）
	PostalCode        string `json:"postal_code,omitempty"`         // 邮政编码
}

// CreateCoreHRCompanyRespCompanySubType ...
type CreateCoreHRCompanyRespCompanySubType struct {
	EnumName string                                          `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRCompanyRespCompanySubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRCompanyRespCompanySubTypeDisplay ...
type CreateCoreHRCompanyRespCompanySubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// CreateCoreHRCompanyRespCompanyType ...
type CreateCoreHRCompanyRespCompanyType struct {
	EnumName string                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRCompanyRespCompanyTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRCompanyRespCompanyTypeDisplay ...
type CreateCoreHRCompanyRespCompanyTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 支持中文和英文。中文用zh-CN；英文用en-US。
	Value string `json:"value,omitempty"` // 名称信息的内容, 支持中文和英文
}

// createCoreHRCompanyResp ...
type createCoreHRCompanyResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCoreHRCompanyResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
