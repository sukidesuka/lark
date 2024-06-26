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

// UpdateCoreHRNationalIDType 更新国家证件类型。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/national_id_type/patch
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/national_id_type/patch
func (r *CoreHRService) UpdateCoreHRNationalIDType(ctx context.Context, request *UpdateCoreHRNationalIDTypeReq, options ...MethodOptionFunc) (*UpdateCoreHRNationalIDTypeResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHRNationalIDType != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHRNationalIDType mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHRNationalIDType(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHRNationalIDType",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/national_id_types/:national_id_type_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHRNationalIDTypeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHRNationalIDType mock CoreHRUpdateCoreHRNationalIDType method
func (r *Mock) MockCoreHRUpdateCoreHRNationalIDType(f func(ctx context.Context, request *UpdateCoreHRNationalIDTypeReq, options ...MethodOptionFunc) (*UpdateCoreHRNationalIDTypeResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHRNationalIDType = f
}

// UnMockCoreHRUpdateCoreHRNationalIDType un-mock CoreHRUpdateCoreHRNationalIDType method
func (r *Mock) UnMockCoreHRUpdateCoreHRNationalIDType() {
	r.mockCoreHRUpdateCoreHRNationalIDType = nil
}

// UpdateCoreHRNationalIDTypeReq ...
type UpdateCoreHRNationalIDTypeReq struct {
	NationalIDTypeID          string                                                    `path:"national_id_type_id" json:"-"`          // 证件类型ID, 示例值: "1616161616"
	ClientToken               *string                                                   `query:"client_token" json:"-"`                // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	CountryRegionID           *string                                                   `json:"country_region_id,omitempty"`           // 国家 / 地区, 示例值: "6862995747139225096"
	Name                      []*UpdateCoreHRNationalIDTypeReqName                      `json:"name,omitempty"`                        // 名称
	Active                    *bool                                                     `json:"active,omitempty"`                      // 是否启用, 示例值: true
	ValidationRule            *string                                                   `json:"validation_rule,omitempty"`             // 校验规则, 示例值: "^\d{9}$"
	ValidationRuleDescription []*UpdateCoreHRNationalIDTypeReqValidationRuleDescription `json:"validation_rule_description,omitempty"` // 校验规则描述
	Code                      *string                                                   `json:"code,omitempty"`                        // 编码, 示例值: "AUS-TFN"
	IdentificationType        *UpdateCoreHRNationalIDTypeReqIdentificationType          `json:"identification_type,omitempty"`         // 证件类型
	CustomFields              []*UpdateCoreHRNationalIDTypeReqCustomField               `json:"custom_fields,omitempty"`               // 自定义字段
}

// UpdateCoreHRNationalIDTypeReqCustomField ...
type UpdateCoreHRNationalIDTypeReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// UpdateCoreHRNationalIDTypeReqIdentificationType ...
type UpdateCoreHRNationalIDTypeReqIdentificationType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHRNationalIDTypeReqName ...
type UpdateCoreHRNationalIDTypeReqName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHRNationalIDTypeReqValidationRuleDescription ...
type UpdateCoreHRNationalIDTypeReqValidationRuleDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHRNationalIDTypeResp ...
type UpdateCoreHRNationalIDTypeResp struct {
	NationalIDType *UpdateCoreHRNationalIDTypeRespNationalIDType `json:"national_id_type,omitempty"` // 国家证件类型
}

// UpdateCoreHRNationalIDTypeRespNationalIDType ...
type UpdateCoreHRNationalIDTypeRespNationalIDType struct {
	ID                        string                                                                   `json:"id,omitempty"`                          // 证件类型 ID
	CountryRegionID           string                                                                   `json:"country_region_id,omitempty"`           // 国家 / 地区
	Name                      []*UpdateCoreHRNationalIDTypeRespNationalIDTypeName                      `json:"name,omitempty"`                        // 名称
	Active                    bool                                                                     `json:"active,omitempty"`                      // 是否启用
	ValidationRule            string                                                                   `json:"validation_rule,omitempty"`             // 校验规则
	ValidationRuleDescription []*UpdateCoreHRNationalIDTypeRespNationalIDTypeValidationRuleDescription `json:"validation_rule_description,omitempty"` // 校验规则描述
	Code                      string                                                                   `json:"code,omitempty"`                        // 编码
	IdentificationType        *UpdateCoreHRNationalIDTypeRespNationalIDTypeIdentificationType          `json:"identification_type,omitempty"`         // 证件类型
	CustomFields              []*UpdateCoreHRNationalIDTypeRespNationalIDTypeCustomField               `json:"custom_fields,omitempty"`               // 自定义字段
}

// UpdateCoreHRNationalIDTypeRespNationalIDTypeCustomField ...
type UpdateCoreHRNationalIDTypeRespNationalIDTypeCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHRNationalIDTypeRespNationalIDTypeIdentificationType ...
type UpdateCoreHRNationalIDTypeRespNationalIDTypeIdentificationType struct {
	EnumName string                                                                   `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRNationalIDTypeRespNationalIDTypeIdentificationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRNationalIDTypeRespNationalIDTypeIdentificationTypeDisplay ...
type UpdateCoreHRNationalIDTypeRespNationalIDTypeIdentificationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHRNationalIDTypeRespNationalIDTypeName ...
type UpdateCoreHRNationalIDTypeRespNationalIDTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHRNationalIDTypeRespNationalIDTypeValidationRuleDescription ...
type UpdateCoreHRNationalIDTypeRespNationalIDTypeValidationRuleDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHRNationalIDTypeResp ...
type updateCoreHRNationalIDTypeResp struct {
	Code  int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                          `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateCoreHRNationalIDTypeResp `json:"data,omitempty"`
	Error *ErrorDetail                    `json:"error,omitempty"`
}
