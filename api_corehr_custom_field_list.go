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

// GetCoreHRCustomFieldList 获取「飞书人事」具体对象下的自定义字段列表。注: 在「人员档案信息配置」-「个人信息」功能中添加的分组, 实际上是一个自定义对象, 可以通过该接口查询自定义对象的所有字段。使用方式可参考[操作手册]如何通过 OpenAPI 维护自定义字段](https://feishu.feishu.cn/docx/QlUudBfCtosWMbxx3vxcOFDknn7)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/query
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/custom_field/query
func (r *CoreHRService) GetCoreHRCustomFieldList(ctx context.Context, request *GetCoreHRCustomFieldListReq, options ...MethodOptionFunc) (*GetCoreHRCustomFieldListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCustomFieldList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCustomFieldList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCustomFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCustomFieldList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/custom_fields/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCustomFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCustomFieldList mock CoreHRGetCoreHRCustomFieldList method
func (r *Mock) MockCoreHRGetCoreHRCustomFieldList(f func(ctx context.Context, request *GetCoreHRCustomFieldListReq, options ...MethodOptionFunc) (*GetCoreHRCustomFieldListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCustomFieldList = f
}

// UnMockCoreHRGetCoreHRCustomFieldList un-mock CoreHRGetCoreHRCustomFieldList method
func (r *Mock) UnMockCoreHRGetCoreHRCustomFieldList() {
	r.mockCoreHRGetCoreHRCustomFieldList = nil
}

// GetCoreHRCustomFieldListReq ...
type GetCoreHRCustomFieldListReq struct {
	ObjectApiNameList []string `query:"object_api_name_list" json:"-"` // 所属对象 apiname, 支持一个或多个, 当前数量限制为 20 个, 示例值: person
}

// GetCoreHRCustomFieldListResp ...
type GetCoreHRCustomFieldListResp struct {
	Items []*GetCoreHRCustomFieldListRespItem `json:"items,omitempty"` // 自定义字段列表
}

// GetCoreHRCustomFieldListRespItem ...
type GetCoreHRCustomFieldListRespItem struct {
	CustomApiName      string                                              `json:"custom_api_name,omitempty"`      // 自定义字段 apiname, 即自定义字段的唯一标识
	Name               *GetCoreHRCustomFieldListRespItemName               `json:"name,omitempty"`                 // 自定义字段名称
	Description        *GetCoreHRCustomFieldListRespItemDescription        `json:"description,omitempty"`          // 描述
	IsOpen             bool                                                `json:"is_open,omitempty"`              // 是否启用
	IsRequired         bool                                                `json:"is_required,omitempty"`          // 是否必填
	IsUnique           bool                                                `json:"is_unique,omitempty"`            // 是否唯一
	ObjectApiName      string                                              `json:"object_api_name,omitempty"`      // 所属对象 apiname
	Type               int64                                               `json:"type,omitempty"`                 // 自定义字段类型, 可选值有: 1: 文本 Text, “文本”和“超链接”属于该类型, 2: 布尔 Boolean, 3: 数字 Number, 4: 枚举 Option, “单选”和“多选”为该类型, 5: 查找 Lookup, “人员（单选）”、“人员（多选）”和个人信息中的自定义分组为该类型, 6: 自动编码 Autonumber, 7: 日期时间 Datetime, 8: 附件 Attachment, “附件单选”和“附件多选”为该类型, 9: 图片 Image, 10: 计算字段 Calculated, 11: 反向查找 Backlookup
	CommonSchemaConfig *GetCoreHRCustomFieldListRespItemCommonSchemaConfig `json:"common_schema_config,omitempty"` // 配置信息, 当前仅字段类型为「文本」、「布尔」、「数字」、「选项」、「日期时间」、「查找」、「附件」、「图片」时返回具体的配置信息, 其余类型的自定义字段暂不返回
	CreateTime         string                                              `json:"create_time,omitempty"`          // 创建时间, 秒级时间戳
	UpdateTime         string                                              `json:"update_time,omitempty"`          // 更新时间, 秒级时间戳
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfig ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfig struct {
	TextFieldSetting       *GetCoreHRCustomFieldListRespItemCommonSchemaConfigTextFieldSetting       `json:"text_field_setting,omitempty"`       // 文本配置信息
	NumberFieldSetting     *GetCoreHRCustomFieldListRespItemCommonSchemaConfigNumberFieldSetting     `json:"number_field_setting,omitempty"`     // 数字配置信息
	EnumFieldSetting       *GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSetting       `json:"enum_field_setting,omitempty"`       // 枚举配置信息
	LookupFieldSetting     *GetCoreHRCustomFieldListRespItemCommonSchemaConfigLookupFieldSetting     `json:"lookup_field_setting,omitempty"`     // 查找字段配置信息
	DateTimeFieldSetting   *GetCoreHRCustomFieldListRespItemCommonSchemaConfigDateTimeFieldSetting   `json:"date_time_field_setting,omitempty"`  // 日期时间配置信息
	AttachmentFieldSetting *GetCoreHRCustomFieldListRespItemCommonSchemaConfigAttachmentFieldSetting `json:"attachment_field_setting,omitempty"` // 附件配置信息
	ImageFieldSetting      *GetCoreHRCustomFieldListRespItemCommonSchemaConfigImageFieldSetting      `json:"image_field_setting,omitempty"`      // 图片配置信息
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigAttachmentFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigAttachmentFieldSetting struct {
	IsMultiple bool     `json:"is_multiple,omitempty"` // 是否支持多个文件
	FileType   FileType `json:"file_type,omitempty"`   // 文件类型枚举, 可选值有: `1`: jpeg, `2`: png, `3`: gif, `4`: pdf, `5`: docx, `6`: doc, `7`: csv, `8`: xls, `9`: txt, `10`: xlsx, `11`: mp4, `12`: pptx, `13`: ppt, `14`: json, `15`: zip, `16`: rar
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigDateTimeFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigDateTimeFieldSetting struct {
	DateTimeType int64 `json:"date_time_type,omitempty"` // 时间类型枚举, 可选值有: `1`: Date 日期, 如 2020-01-01, `2`: Time  时间, 如 11:52:00, `3`: Datetime 日期时间, 如 2020-01-01 11:52:00, `4`: CusDatetime // timestamp 时间戳
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSetting struct {
	EnumFieldOptionList []*GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOption `json:"enum_field_option_list,omitempty"` // 枚举选项信息
	IsMultiple          bool                                                                                 `json:"is_multiple,omitempty"`            // 是否为多选
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOption ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOption struct {
	ApiName     string                                                                                        `json:"api_name,omitempty"`    // 选项 api_name
	Name        *GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOptionName        `json:"name,omitempty"`        // 选项名称
	Description *GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOptionDescription `json:"description,omitempty"` // 选项描述
	IsOpen      bool                                                                                          `json:"is_open,omitempty"`     // 是否启用
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOptionDescription ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOptionDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOptionName ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigEnumFieldSettingEnumFieldOptionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigImageFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigImageFieldSetting struct {
	ImageType    int64 `json:"image_type,omitempty"`    // 图片类型枚举, 可选值有: `1`: Avatar 头像, `2`: BadgePhoto 工卡照片, `3`: Logo 标志
	DisplayStyle int64 `json:"display_style,omitempty"` // 显示样式, 可选值有: `1`: SquareImage 方形, `2`: RoundImage  圆形
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigLookupFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigLookupFieldSetting struct {
	LookupObjApiName string `json:"lookup_obj_api_name,omitempty"` // 查找字段对应的对象 apiname, 可通过[获取自定义字段列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/query)接口获取这个对象中定义的自定义字段
	IsMultiple       bool   `json:"is_multiple,omitempty"`         // 是否为多值
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigNumberFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigNumberFieldSetting struct {
	NumberFieldType    int64 `json:"number_field_type,omitempty"`    // 数字类型, 可选值有: `1`: Percent 百分比, `2`: Integer 整数, `3`: Value 数值（浮点数）, `4`: Money 金额（浮点数）
	DecimalPlaces      int64 `json:"decimal_places,omitempty"`       // 小数部分位数（浮点数整数部分和小数部分分别最大30位）
	RoundType          int64 `json:"round_type,omitempty"`           // 四舍五入规则, 可选值有: `0`: Round 四舍五入, `1`: Ceil 向上舍入, `2`: Floor 向下舍入
	DecimalTotalPlaces int64 `json:"decimal_total_places,omitempty"` // 整数+小数总位数
}

// GetCoreHRCustomFieldListRespItemCommonSchemaConfigTextFieldSetting ...
type GetCoreHRCustomFieldListRespItemCommonSchemaConfigTextFieldSetting struct {
	IsMultilingual bool  `json:"is_multilingual,omitempty"` // 是否多语言
	IsMultiline    bool  `json:"is_multiline,omitempty"`    // 是否多行
	MaxLength      int64 `json:"max_length,omitempty"`      // 最大长度
	IsURLType      bool  `json:"is_url_type,omitempty"`     // 是否是URL类型
}

// GetCoreHRCustomFieldListRespItemDescription ...
type GetCoreHRCustomFieldListRespItemDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHRCustomFieldListRespItemName ...
type GetCoreHRCustomFieldListRespItemName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getCoreHRCustomFieldListResp ...
type getCoreHRCustomFieldListResp struct {
	Code  int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                        `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRCustomFieldListResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}
