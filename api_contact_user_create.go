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

// CreateUser 使用该接口向通讯录创建一个用户, 可以理解为员工入职。创建用户后只返回有数据权限的数据。具体的数据权限的与字段的对应关系请参照[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。
//
// - 当在[混合license模式](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant-product_assign_info/query)下, 席位字段为必填。
// - 新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户, 如果想要在根部门下新增用户, 必须要有全员权限。
// - 应用商店应用无权限调用此接口。
// - 创建用户后, 会给用户发送邀请短信/邮件, 用户在操作同意后才可访问团队。
// - 返回数据中不返回手机号, 如果需要请重新查询用户信息获取手机号。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/user/create
func (r *ContactService) CreateUser(ctx context.Context, request *CreateUserReq, options ...MethodOptionFunc) (*CreateUserResp, *Response, error) {
	if r.cli.mock.mockContactCreateUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#CreateUser mock enable")
		return r.cli.mock.mockContactCreateUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "CreateUser",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/users",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactCreateUser mock ContactCreateUser method
func (r *Mock) MockContactCreateUser(f func(ctx context.Context, request *CreateUserReq, options ...MethodOptionFunc) (*CreateUserResp, *Response, error)) {
	r.mockContactCreateUser = f
}

// UnMockContactCreateUser un-mock ContactCreateUser method
func (r *Mock) UnMockContactCreateUser() {
	r.mockContactCreateUser = nil
}

// CreateUserReq ...
type CreateUserReq struct {
	UserIDType              *IDType                    `query:"user_id_type" json:"-"`                // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType        *DepartmentIDType          `query:"department_id_type" json:"-"`          // 此次调用中使用的部门ID的类型, 不同 ID 的说明参见[部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0), 示例值: open_department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	ClientToken             *string                    `query:"client_token" json:"-"`                // 用于幂等判断是否为同一请求, 避免重复创建。字符串类型, 自行生成, 示例值: xxxx-xxxxx-xxx
	UserID                  *string                    `json:"user_id,omitempty"`                     // 用户的user_id, 租户内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "3e3cf96b"
	Name                    string                     `json:"name,omitempty"`                        // 用户名, 示例值: "张三", 最小长度: `1` 字符
	EnName                  *string                    `json:"en_name,omitempty"`                     // 英文名, 示例值: "San Zhang"
	Nickname                *string                    `json:"nickname,omitempty"`                    // 别名, 示例值: "Alex Zhang"
	Email                   *string                    `json:"email,omitempty"`                       // 邮箱, 注意: 1. 非中国大陆手机号成员必须同时添加邮箱, 2. 邮箱不可重复, 示例值: "zhangsan@gmail.com"
	Mobile                  string                     `json:"mobile,omitempty"`                      // 手机号, 注意: 1. 在本企业内不可重复, 2. 未认证企业仅支持添加中国大陆手机号, 通过飞书认证的企业允许添加海外手机号, 3. 国际电话区号前缀中必须包含加号 +, 4. 该 mobile 字段在海外版飞书非必填, 示例值: "13011111111 (其他例子, 中国大陆手机号: 13011111111 或 +8613011111111, 非中国大陆手机号: +41446681800)"
	MobileVisible           *bool                      `json:"mobile_visible,omitempty"`              // 手机号码可见性, true 为可见, false 为不可见, 目前默认为 true。不可见时, 组织员工将无法查看该员工的手机号码, 示例值: false
	Gender                  *int64                     `json:"gender,omitempty"`                      // 性别, 示例值: 1, 可选值有: 0: 保密, 1: 男, 2: 女
	AvatarKey               *string                    `json:"avatar_key,omitempty"`                  // 头像的文件Key, 可通过“消息与群组/消息/图片信息”中的“上传图片”接口上传并获取头像文件 Key, “上传图片”功能参见[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create), 示例值: "2500c7a9-5fff-4d9a-a2de-3d59614ae28g"
	DepartmentIDs           []string                   `json:"department_ids,omitempty"`              // 用户所属部门的ID列表, 一个用户可属于多个部门, ID值的类型与查询参数中的department_id_type 对应, 不同 ID 的说明与department_id的获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0), 示例值: ["od-4e6ac4d14bcd5071a37a39de902c714111111"]
	LeaderUserID            *string                    `json:"leader_user_id,omitempty"`              // 用户的直接主管的用户ID, ID值与查询参数中的user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 获取方式参见[如何获取user_id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	City                    *string                    `json:"city,omitempty"`                        // 工作城市, 示例值: "杭州"
	Country                 *string                    `json:"country,omitempty"`                     // 国家或地区Code缩写, 具体写入格式请参考 [国家/地区码表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/country-code-description), 示例值: "CN"
	WorkStation             *string                    `json:"work_station,omitempty"`                // 工位, 示例值: "北楼-H34"
	JoinTime                *int64                     `json:"join_time,omitempty"`                   // 入职时间, 时间戳格式, 表示从1970年1月1日开始所经过的秒数。创建用户时不指定入职时间则默认填充当前时间, 示例值: 2147483647
	EmployeeNo              *string                    `json:"employee_no,omitempty"`                 // 工号, 示例值: "1"
	EmployeeType            int64                      `json:"employee_type,omitempty"`               // 员工类型, 可选值有: `1`: 正式员工, `2`: 实习生, `3`: 外包, `4`: 劳务, `5`: 顾问, 同时可读取到自定义员工类型的 int 值, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 示例值: 1
	Orders                  []*CreateUserReqOrder      `json:"orders,omitempty"`                      // 用户排序信息, 用于标记通讯录下组织架构的人员顺序, 人员可能存在多个部门中, 且有不同的排序。
	CustomAttrs             []*CreateUserReqCustomAttr `json:"custom_attrs,omitempty"`                // 自定义字段, 请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“, 否则该字段不会生效/返回, 更多详情参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	EnterpriseEmail         *string                    `json:"enterprise_email,omitempty"`            // 企业邮箱, 请先确保已在管理后台启用飞书邮箱服务, 创建用户时, 企业邮箱的使用方式参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 示例值: "demo@mail.com"
	JobTitle                *string                    `json:"job_title,omitempty"`                   // 职务, 示例值: "xxxxx"
	Geo                     *string                    `json:"geo,omitempty"`                         // 数据驻留地, 示例值: "cn"
	JobLevelID              *string                    `json:"job_level_id,omitempty"`                // 职级ID, 示例值: "mga5oa8ayjlp9rb"
	JobFamilyID             *string                    `json:"job_family_id,omitempty"`               // 序列ID, 示例值: "mga5oa8ayjlp9rb"
	SubscriptionIDs         []string                   `json:"subscription_ids,omitempty"`            // 分配给用户的席位ID列表, 需开通“分配用户席位”权限。可通过下方接口获取到该租户的可用席位ID, 参见[获取企业席位信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant-product_assign_info/query)。当在混合license模式下, 此字段为必填, 示例值: ["23213213213123123"]
	DottedLineLeaderUserIDs []string                   `json:"dotted_line_leader_user_ids,omitempty"` // 虚线上级ID, 示例值: ["ou_7dab8a3d3cdcc9da365777c7ad535d62"]
}

// CreateUserReqCustomAttr ...
type CreateUserReqCustomAttr struct {
	Type  *string                       `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, 具体说明参见常见问题的[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 示例值: "TEXT"
	ID    *string                       `json:"id,omitempty"`    // 自定义字段ID, 示例值: "DemoId"
	Value *CreateUserReqCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// CreateUserReqCustomAttrValue ...
type CreateUserReqCustomAttrValue struct {
	Text        *string                                  `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填, 示例值: "DemoText"
	URL         *string                                  `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页, 示例值: "http://www.fs.cn"
	PcURL       *string                                  `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL, 示例值: "http://www.fs.cn"
	OptionID    *string                                  `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时, 该参数定义选项值, 示例值: "edcvfrtg"
	GenericUser *CreateUserReqCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// CreateUserReqCustomAttrValueGenericUser ...
type CreateUserReqCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 具体参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "9b2fabg5"
	Type int64  `json:"type,omitempty"` // 用户类型: 1: 用户, 目前固定为1, 表示用户类型, 示例值: 1
}

// CreateUserReqOrder ...
type CreateUserReqOrder struct {
	DepartmentID    *string `json:"department_id,omitempty"`    // 排序信息对应的部门ID, ID值与查询参数中的department_id_type 对应, 表示用户所在的、且需要排序的部门, 不同 ID 的说明参见及获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview), 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	UserOrder       *int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序, 数值越大, 排序越靠前, 示例值: 100
	DepartmentOrder *int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序, 数值越大, 排序越靠前, 示例值: 100
	IsPrimaryDept   *bool   `json:"is_primary_dept,omitempty"`  // 标识用户的唯一主部门, 主部门为用户所属部门中排序第一的部门(department_order最大), 示例值: true
}

// CreateUserResp ...
type CreateUserResp struct {
	User *CreateUserRespUser `json:"user,omitempty"` // 用户信息
}

// CreateUserRespUser ...
type CreateUserRespUser struct {
	UnionID                 string                          `json:"union_id,omitempty"`                    // 用户的union_id, 应用开发商发布的不同应用中同一用户的标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	UserID                  string                          `json:"user_id,omitempty"`                     // 用户的user_id, 租户内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 字段权限要求: 获取用户 user ID
	OpenID                  string                          `json:"open_id,omitempty"`                     // 用户的open_id, 应用内用户的唯一标识, 不同ID的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Name                    string                          `json:"name,omitempty"`                        // 用户名, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	EnName                  string                          `json:"en_name,omitempty"`                     // 英文名, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	Nickname                string                          `json:"nickname,omitempty"`                    // 别名, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	Email                   string                          `json:"email,omitempty"`                       // 邮箱, 注意: 1. 非中国大陆手机号成员必须同时添加邮箱, 2. 邮箱不可重复, 字段权限要求: 获取用户邮箱信息
	Mobile                  string                          `json:"mobile,omitempty"`                      // 手机号, 注意: 1. 在本企业内不可重复, 2. 未认证企业仅支持添加中国大陆手机号, 通过飞书认证的企业允许添加海外手机号, 3. 国际电话区号前缀中必须包含加号 +, 4. 该 mobile 字段在海外版飞书非必填, 字段权限要求: 获取用户手机号
	MobileVisible           bool                            `json:"mobile_visible,omitempty"`              // 手机号码可见性, true 为可见, false 为不可见, 目前默认为 true。不可见时, 组织员工将无法查看该员工的手机号码
	Gender                  int64                           `json:"gender,omitempty"`                      // 性别, 可选值有: 0: 保密, 1: 男, 2: 女, 字段权限要求（满足任一）: 获取用户性别, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	AvatarKey               string                          `json:"avatar_key,omitempty"`                  // 头像的文件Key, 可通过“消息与群组/消息/图片信息”中的“上传图片”接口上传并获取头像文件 Key, “上传图片”功能参见[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)
	Avatar                  *CreateUserRespUserAvatar       `json:"avatar,omitempty"`                      // 用户头像信息, 字段权限要求（满足任一）: 获取用户基本信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	Status                  *CreateUserRespUserStatus       `json:"status,omitempty"`                      // 用户状态, 枚举类型, 包括is_frozen、is_resigned、is_activated、is_exited, 用户状态转移参见: [用户状态图](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/field-overview#4302b5a1), 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	DepartmentIDs           []string                        `json:"department_ids,omitempty"`              // 用户所属部门的ID列表, 一个用户可属于多个部门, ID值的类型与查询参数中的department_id_type 对应, 不同 ID 的说明与department_id的获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview#23857fe0), 字段权限要求（满足任一）: 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	LeaderUserID            string                          `json:"leader_user_id,omitempty"`              // 用户的直接主管的用户ID, ID值与查询参数中的user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 获取方式参见[如何获取user_id](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 字段权限要求（满足任一）: 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	City                    string                          `json:"city,omitempty"`                        // 工作城市, 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	Country                 string                          `json:"country,omitempty"`                     // 国家或地区Code缩写, 具体写入格式请参考 [国家/地区码表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/country-code-description), 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	WorkStation             string                          `json:"work_station,omitempty"`                // 工位, 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	JoinTime                int64                           `json:"join_time,omitempty"`                   // 入职时间, 时间戳格式, 表示从1970年1月1日开始所经过的秒数, 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	IsTenantManager         bool                            `json:"is_tenant_manager,omitempty"`           // 是否是租户超级管理员, 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	EmployeeNo              string                          `json:"employee_no,omitempty"`                 // 工号, 字段权限要求（满足任一）: 查看成员工号, 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	EmployeeType            int64                           `json:"employee_type,omitempty"`               // 员工类型, 可选值有: `1`: 正式员工, `2`: 实习生, `3`: 外包, `4`: 劳务, `5`: 顾问, 同时可读取到自定义员工类型的 int 值, 可通过下方接口获取到该租户的自定义员工类型的名称, 参见[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	Orders                  []*CreateUserRespUserOrder      `json:"orders,omitempty"`                      // 用户排序信息, 用于标记通讯录下组织架构的人员顺序, 人员可能存在多个部门中, 且有不同的排序, 字段权限要求（满足任一）: 获取用户组织架构信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	CustomAttrs             []*CreateUserRespUserCustomAttr `json:"custom_attrs,omitempty"`                // 自定义字段, 请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“, 否则该字段不会生效/返回, 更多详情参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	EnterpriseEmail         string                          `json:"enterprise_email,omitempty"`            // 企业邮箱, 请先确保已在管理后台启用飞书邮箱服务, 创建用户时, 企业邮箱的使用方式参见[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525), 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	JobTitle                string                          `json:"job_title,omitempty"`                   // 职务, 字段权限要求（满足任一）: 获取用户受雇信息, 以应用身份访问通讯录, 读取通讯录, 以应用身份读取通讯录
	IsFrozen                bool                            `json:"is_frozen,omitempty"`                   // 是否暂停用户
	Geo                     string                          `json:"geo,omitempty"`                         // 数据驻留地, 字段权限要求: 查看成员数据驻留地
	JobLevelID              string                          `json:"job_level_id,omitempty"`                // 职级ID, 字段权限要求: 查询用户职级
	JobFamilyID             string                          `json:"job_family_id,omitempty"`               // 序列ID, 字段权限要求: 查询用户所属的工作序列
	DottedLineLeaderUserIDs []string                        `json:"dotted_line_leader_user_ids,omitempty"` // 虚线上级ID, 字段权限要求: 查看成员的虚线上级 ID
}

// CreateUserRespUserAvatar ...
type CreateUserRespUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

// CreateUserRespUserCustomAttr ...
type CreateUserRespUserCustomAttr struct {
	Type  string                             `json:"type,omitempty"`  // 自定义字段类型, `TEXT`: 文本, `HREF`: 网页, `ENUMERATION`: 枚举, `PICTURE_ENUM`: 图片, `GENERIC_USER`: 用户, 具体说明参见常见问题的[用户接口相关问题](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN#77061525)
	ID    string                             `json:"id,omitempty"`    // 自定义字段ID
	Value *CreateUserRespUserCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

// CreateUserRespUserCustomAttrValue ...
type CreateUserRespUserCustomAttrValue struct {
	Text        string                                        `json:"text,omitempty"`         // 字段类型为`TEXT`时该参数定义字段值, 必填；字段类型为`HREF`时该参数定义网页标题, 必填
	URL         string                                        `json:"url,omitempty"`          // 字段类型为 HREF 时, 该参数定义默认 URL, 例如手机端跳转小程序, PC端跳转网页
	PcURL       string                                        `json:"pc_url,omitempty"`       // 字段类型为 HREF 时, 该参数定义PC端 URL
	OptionID    string                                        `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时, 该参数定义选项值
	OptionValue string                                        `json:"option_value,omitempty"` // 选项类型的值, 表示 成员详情/自定义字段 中选项选中的值
	Name        string                                        `json:"name,omitempty"`         // 选项类型为图片时, 表示图片的名称
	PictureURL  string                                        `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *CreateUserRespUserCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时, 该参数定义引用人员
}

// CreateUserRespUserCustomAttrValueGenericUser ...
type CreateUserRespUserCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 具体参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Type int64  `json:"type,omitempty"` // 用户类型: 1: 用户, 目前固定为1, 表示用户类型
}

// CreateUserRespUserOrder ...
type CreateUserRespUserOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID, ID值与查询参数中的department_id_type 对应, 表示用户所在的、且需要排序的部门, 不同 ID 的说明参见及获取方式参见 [部门ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview)
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序, 数值越大, 排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序, 数值越大, 排序越靠前
	IsPrimaryDept   bool   `json:"is_primary_dept,omitempty"`  // 标识用户的唯一主部门, 主部门为用户所属部门中排序第一的部门(department_order最大)
}

// CreateUserRespUserStatus ...
type CreateUserRespUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
	IsExited    bool `json:"is_exited,omitempty"`    // 是否主动退出, 主动退出一段时间后用户会自动转为已离职
	IsUnjoin    bool `json:"is_unjoin,omitempty"`    // 是否未加入, 需要用户自主确认才能加入团队
}

// createUserResp ...
type createUserResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateUserResp `json:"data,omitempty"`
}
