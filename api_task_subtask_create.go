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

// CreateTaskSubtask 给一个任务创建一个子任务。
//
// 接口功能除了额外需要输入父任务的GUID之外, 和[创建任务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/create)接口功能完全一致。
// 创建子任务需要拥有父任务的编辑权限。详见[任务是如何鉴权的？](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/faq)
// 如果将新任务加入清单, 则需要清单的可编辑权限。详情见[任务功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)中的“任务是如何鉴权的？”章节。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task-subtask/create
func (r *TaskService) CreateTaskSubtask(ctx context.Context, request *CreateTaskSubtaskReq, options ...MethodOptionFunc) (*CreateTaskSubtaskResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskSubtask != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#CreateTaskSubtask mock enable")
		return r.cli.mock.mockTaskCreateTaskSubtask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskSubtask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/tasks/:task_guid/subtasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskSubtaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskSubtask mock TaskCreateTaskSubtask method
func (r *Mock) MockTaskCreateTaskSubtask(f func(ctx context.Context, request *CreateTaskSubtaskReq, options ...MethodOptionFunc) (*CreateTaskSubtaskResp, *Response, error)) {
	r.mockTaskCreateTaskSubtask = f
}

// UnMockTaskCreateTaskSubtask un-mock TaskCreateTaskSubtask method
func (r *Mock) UnMockTaskCreateTaskSubtask() {
	r.mockTaskCreateTaskSubtask = nil
}

// CreateTaskSubtaskReq ...
type CreateTaskSubtaskReq struct {
	TaskGuid       string                              `path:"task_guid" json:"-"`        // 父任务GUID, 示例值: "e297ddff-06ca-4166-b917-4ce57cd3a7a0", 最大长度: `100` 字符
	UserIDType     *IDType                             `query:"user_id_type" json:"-"`    // 用户 ID 类型, 示例值: open_id, 默认值: `open_id`
	Summary        string                              `json:"summary,omitempty"`         // 任务标题, 示例值: "针对全年销售进行一次复盘", 最大长度: `10000` 字符
	Description    *string                             `json:"description,omitempty"`     // 任务摘要, 示例值: "需要事先阅读复盘总结文档"
	Due            *CreateTaskSubtaskReqDue            `json:"due,omitempty"`             // 任务截止时间戳(ms), 截止时间戳和截止日期选择一个填写, 示例值: 1675742789470
	Origin         *CreateTaskSubtaskReqOrigin         `json:"origin,omitempty"`          // 任务关联的第三方平台来源信息。详见[如何使用Origin?](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/overview)
	Extra          *string                             `json:"extra,omitempty"`           // 调用者可以传入的任意附带到任务上的数据。在获取任务详情时会原样返回, 示例值: "dGVzdA[", 最大长度: `65536` 字符
	CompletedAt    *string                             `json:"completed_at,omitempty"`    // 任务的完成时刻时间戳(ms), 示例值: "1675742789470", 默认值: `0`, 最大长度: `20` 字符
	Members        []*CreateTaskSubtaskReqMember       `json:"members,omitempty"`         // 任务成员列表, 示例值: [ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f], 最大长度: `500`
	RepeatRule     *string                             `json:"repeat_rule,omitempty"`     // 如果设置, 则该任务为“重复任务”。该字段表示了重复任务的重复规则。详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“如何使用重复任务？”章节, 示例值: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO, TU, WE, TH, FR", 最大长度: `1000` 字符
	CustomComplete *CreateTaskSubtaskReqCustomComplete `json:"custom_complete,omitempty"` // 任务自定义完成规则。详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“如何使用自定义完成？”章节。
	Tasklists      []*CreateTaskSubtaskReqTasklist     `json:"tasklists,omitempty"`       // 任务所在清单的信息。如果设置, 则表示创建的任务要直接加入到指定清单。
	ClientToken    *string                             `json:"client_token,omitempty"`    // 幂等token。如果提供则触发后端实现幂等行为。详见[功能概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/task-v2/overview)中的“ 幂等调用 ”章节, 示例值: "daa2237f-8310-4707-a83b-52c8a81e0fb7", 长度范围: `10` ～ `100` 字符
	Start          *CreateTaskSubtaskReqStart          `json:"start,omitempty"`           // 任务的开始时间(ms)
	Reminders      []*CreateTaskSubtaskReqReminder     `json:"reminders,omitempty"`       // 任务提醒
}

// CreateTaskSubtaskReqCustomComplete ...
type CreateTaskSubtaskReqCustomComplete struct {
	Pc      *CreateTaskSubtaskReqCustomCompletePc      `json:"pc,omitempty"`      // pc客户端自定义完成配置（含mac和windows）
	Ios     *CreateTaskSubtaskReqCustomCompleteIos     `json:"ios,omitempty"`     // 飞书ios端的自定义完成配置
	Android *CreateTaskSubtaskReqCustomCompleteAndroid `json:"android,omitempty"` // 飞书android端的自定义完成配置
}

// CreateTaskSubtaskReqCustomCompleteAndroid ...
type CreateTaskSubtaskReqCustomCompleteAndroid struct {
	Href *string                                       `json:"href,omitempty"` // 自定义完成的跳转url, 示例值: "https://www.example.com"
	Tip  *CreateTaskSubtaskReqCustomCompleteAndroidTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// CreateTaskSubtaskReqCustomCompleteAndroidTip ...
type CreateTaskSubtaskReqCustomCompleteAndroidTip struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// CreateTaskSubtaskReqCustomCompleteIos ...
type CreateTaskSubtaskReqCustomCompleteIos struct {
	Href *string                                   `json:"href,omitempty"` // 自定义完成的跳转url, 示例值: "https://www.example.com"
	Tip  *CreateTaskSubtaskReqCustomCompleteIosTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// CreateTaskSubtaskReqCustomCompleteIosTip ...
type CreateTaskSubtaskReqCustomCompleteIosTip struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// CreateTaskSubtaskReqCustomCompletePc ...
type CreateTaskSubtaskReqCustomCompletePc struct {
	Href *string                                  `json:"href,omitempty"` // 自定义完成的跳转url, 示例值: "https://www.example.com"
	Tip  *CreateTaskSubtaskReqCustomCompletePcTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// CreateTaskSubtaskReqCustomCompletePcTip ...
type CreateTaskSubtaskReqCustomCompletePcTip struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// CreateTaskSubtaskReqDue ...
type CreateTaskSubtaskReqDue struct {
	Timestamp *string `json:"timestamp,omitempty"`  // 截止时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果截止时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true, 示例值: "1675454764000"
	IsAllDay  *bool   `json:"is_all_day,omitempty"` // 是否截止到一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储, 示例值: true
}

// CreateTaskSubtaskReqMember ...
type CreateTaskSubtaskReqMember struct {
	ID   *string `json:"id,omitempty"`   // 表示member的id, 示例值: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f", 最大长度: `100` 字符
	Type *string `json:"type,omitempty"` // 成员的类型, 示例值: "user", 默认值: `user`
	Role *string `json:"role,omitempty"` // 成员角色, 支持"assignee"或者"follower", 示例值: "assignee", 最大长度: `20` 字符
}

// CreateTaskSubtaskReqOrigin ...
type CreateTaskSubtaskReqOrigin struct {
	PlatformI18nName *CreateTaskSubtaskReqOriginPlatformI18nName `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。需提供多语言版本。
	Href             *CreateTaskSubtaskReqOriginHref             `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// CreateTaskSubtaskReqOriginHref ...
type CreateTaskSubtaskReqOriginHref struct {
	URL   *string `json:"url,omitempty"`   // 链接对应的地址, 示例值: "https://www.example.com", 长度范围: `0` ～ `1024` 字符
	Title *string `json:"title,omitempty"` // 链接对应的标题, 示例值: "反馈一个问题, 需要协助排查", 最大长度: `512` 字符
}

// CreateTaskSubtaskReqOriginPlatformI18nName ...
type CreateTaskSubtaskReqOriginPlatformI18nName struct {
	EnUs *string `json:"en_us,omitempty"` // 英文, 示例值: "workbench", 最大长度: `1000` 字符
	ZhCn *string `json:"zh_cn,omitempty"` // 中文, 示例值: "工作台", 最大长度: `1000` 字符
	ZhHk *string `json:"zh_hk,omitempty"` // 中文（香港地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	ZhTw *string `json:"zh_tw,omitempty"` // 中文（台湾地区）, 示例值: "工作臺", 最大长度: `1000` 字符
	JaJp *string `json:"ja_jp,omitempty"` // 日语, 示例值: "作業台", 最大长度: `1000` 字符
	FrFr *string `json:"fr_fr,omitempty"` // 法语, 示例值: "Table de travail"
	ItIt *string `json:"it_it,omitempty"` // 意大利语, 示例值: "banco di lavoro"
	DeDe *string `json:"de_de,omitempty"` // 德语, 示例值: "Werkbank"
	RuRu *string `json:"ru_ru,omitempty"` // 俄语, 示例值: "верстак"
	ThTh *string `json:"th_th,omitempty"` // 泰语, 示例值: "โต๊ะทำงาน"
	EsEs *string `json:"es_es,omitempty"` // 西班牙语, 示例值: "banco de trabajo"
	KoKr *string `json:"ko_kr,omitempty"` // 韩语, 示例值: "작업대"
}

// CreateTaskSubtaskReqReminder ...
type CreateTaskSubtaskReqReminder struct {
	RelativeFireMinute int64 `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间分钟数。例如30表示截止时间前30分钟提醒；0表示截止时提醒, 示例值: 30
}

// CreateTaskSubtaskReqStart ...
type CreateTaskSubtaskReqStart struct {
	Timestamp *string `json:"timestamp,omitempty"`  // 开始时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果开始时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true, 示例值: "1675454764000"
	IsAllDay  *bool   `json:"is_all_day,omitempty"` // 是否开始于一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储, 示例值: true
}

// CreateTaskSubtaskReqTasklist ...
type CreateTaskSubtaskReqTasklist struct {
	TasklistGuid *string `json:"tasklist_guid,omitempty"` // 任务要加入的清单的GUID, 示例值: "cc371766-6584-cf50-a222-c22cd9055004", 最大长度: `100` 字符
	SectionGuid  *string `json:"section_guid,omitempty"`  // 任务所在清单的自定义分组GUID。如果设置了清单GUID但没有设置自定义分组GUID, 则自动加入该清单的默认分组, 示例值: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
}

// CreateTaskSubtaskResp ...
type CreateTaskSubtaskResp struct {
	Subtask *CreateTaskSubtaskRespSubtask `json:"subtask,omitempty"` // 创建的任务
}

// CreateTaskSubtaskRespSubtask ...
type CreateTaskSubtaskRespSubtask struct {
	Guid           string                                      `json:"guid,omitempty"`             // 任务guid, 任务的唯一ID
	Summary        string                                      `json:"summary,omitempty"`          // 任务标题
	Description    string                                      `json:"description,omitempty"`      // 任务描述
	Due            *CreateTaskSubtaskRespSubtaskDue            `json:"due,omitempty"`              // 任务截止时间
	Reminders      []*CreateTaskSubtaskRespSubtaskReminder     `json:"reminders,omitempty"`        // 任务的提醒配置列表。目前每个任务最多有1个。
	Creator        *CreateTaskSubtaskRespSubtaskCreator        `json:"creator,omitempty"`          // 任务创建者
	Members        []*CreateTaskSubtaskRespSubtaskMember       `json:"members,omitempty"`          // 任务成员列表
	CompletedAt    string                                      `json:"completed_at,omitempty"`     // 任务完成的时间戳(ms)
	Attachments    []*CreateTaskSubtaskRespSubtaskAttachment   `json:"attachments,omitempty"`      // 任务的附件列表
	Origin         *CreateTaskSubtaskRespSubtaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息。创建是设置后就不可更改。
	Extra          string                                      `json:"extra,omitempty"`            // 任务附带的自定义数据。
	Tasklists      []*CreateTaskSubtaskRespSubtaskTasklist     `json:"tasklists,omitempty"`        // 任务所属清单的名字。调用者只能看到有权限访问的清单的列表。
	RepeatRule     string                                      `json:"repeat_rule,omitempty"`      // 如果任务为重复任务, 返回重复任务的配置
	ParentTaskGuid string                                      `json:"parent_task_guid,omitempty"` // 如果当前任务为某个任务的子任务, 返回父任务的guid
	Mode           int64                                       `json:"mode,omitempty"`             // 任务的模式。1 - 会签任务；2 - 或签任务
	Source         int64                                       `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知来源, 1: 任务中心, 2: 群组任务/消息转任务, 6: 通过开放平台以tenant_access_token授权创建的任务, 7: 通过开放平台以user_access_token授权创建的任务, 8: 文档任务
	CustomComplete *CreateTaskSubtaskRespSubtaskCustomComplete `json:"custom_complete,omitempty"`  // 任务的自定义完成配置
	TaskID         string                                      `json:"task_id,omitempty"`          // 任务界面上的代码
	CreatedAt      string                                      `json:"created_at,omitempty"`       // 任务创建时间戳(ms)
	UpdatedAt      string                                      `json:"updated_at,omitempty"`       // 任务最后一次更新的时间戳(ms)
	Status         string                                      `json:"status,omitempty"`           // 任务的状态, 支持"todo"和"done"两种状态
	URL            string                                      `json:"url,omitempty"`              // 任务的分享链接
	Start          *CreateTaskSubtaskRespSubtaskStart          `json:"start,omitempty"`            // 任务的开始时间, 如果同时设置任务的开始时间和截止时间, 开始时间必须<=截止时间, 并且开始/截止时间的is_all_day设置必须相同。
	SubtaskCount   int64                                       `json:"subtask_count,omitempty"`    // 该任务的子任务的个数。
}

// CreateTaskSubtaskRespSubtaskAttachment ...
type CreateTaskSubtaskRespSubtaskAttachment struct {
	Guid       string                                          `json:"guid,omitempty"`        // 附件guid
	FileToken  string                                          `json:"file_token,omitempty"`  // 附件在云文档系统中的token
	Name       string                                          `json:"name,omitempty"`        // 附件名
	Size       int64                                           `json:"size,omitempty"`        // 附件的字节大小
	Resource   *CreateTaskSubtaskRespSubtaskAttachmentResource `json:"resource,omitempty"`    // 附件归属的资源
	Uploader   *CreateTaskSubtaskRespSubtaskAttachmentUploader `json:"uploader,omitempty"`    // 附件上传者
	IsCover    bool                                            `json:"is_cover,omitempty"`    // 是否是封面图
	UploadedAt string                                          `json:"uploaded_at,omitempty"` // 上传时间戳(ms)
}

// CreateTaskSubtaskRespSubtaskAttachmentResource ...
type CreateTaskSubtaskRespSubtaskAttachmentResource struct {
	Type string `json:"type,omitempty"` // 资源类型
	ID   string `json:"id,omitempty"`   // 资源ID
}

// CreateTaskSubtaskRespSubtaskAttachmentUploader ...
type CreateTaskSubtaskRespSubtaskAttachmentUploader struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// CreateTaskSubtaskRespSubtaskCreator ...
type CreateTaskSubtaskRespSubtaskCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// CreateTaskSubtaskRespSubtaskCustomComplete ...
type CreateTaskSubtaskRespSubtaskCustomComplete struct {
	Pc      *CreateTaskSubtaskRespSubtaskCustomCompletePc      `json:"pc,omitempty"`      // pc客户端自定义完成配置（含mac和windows）
	Ios     *CreateTaskSubtaskRespSubtaskCustomCompleteIos     `json:"ios,omitempty"`     // ios端的自定义完成配置
	Android *CreateTaskSubtaskRespSubtaskCustomCompleteAndroid `json:"android,omitempty"` // android端的自定义完成配置
}

// CreateTaskSubtaskRespSubtaskCustomCompleteAndroid ...
type CreateTaskSubtaskRespSubtaskCustomCompleteAndroid struct {
	Href string                                                `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *CreateTaskSubtaskRespSubtaskCustomCompleteAndroidTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// CreateTaskSubtaskRespSubtaskCustomCompleteAndroidTip ...
type CreateTaskSubtaskRespSubtaskCustomCompleteAndroidTip struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// CreateTaskSubtaskRespSubtaskCustomCompleteIos ...
type CreateTaskSubtaskRespSubtaskCustomCompleteIos struct {
	Href string                                            `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *CreateTaskSubtaskRespSubtaskCustomCompleteIosTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// CreateTaskSubtaskRespSubtaskCustomCompleteIosTip ...
type CreateTaskSubtaskRespSubtaskCustomCompleteIosTip struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// CreateTaskSubtaskRespSubtaskCustomCompletePc ...
type CreateTaskSubtaskRespSubtaskCustomCompletePc struct {
	Href string                                           `json:"href,omitempty"` // 自定义完成的跳转url
	Tip  *CreateTaskSubtaskRespSubtaskCustomCompletePcTip `json:"tip,omitempty"`  // 自定义完成的弹出提示为
}

// CreateTaskSubtaskRespSubtaskCustomCompletePcTip ...
type CreateTaskSubtaskRespSubtaskCustomCompletePcTip struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// CreateTaskSubtaskRespSubtaskDue ...
type CreateTaskSubtaskRespSubtaskDue struct {
	Timestamp string `json:"timestamp,omitempty"`  // 截止时间/日期的时间戳, 距1970-01-01 00:00:00的毫秒数。如果截止时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否截止到一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// CreateTaskSubtaskRespSubtaskMember ...
type CreateTaskSubtaskRespSubtaskMember struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// CreateTaskSubtaskRespSubtaskOrigin ...
type CreateTaskSubtaskRespSubtaskOrigin struct {
	PlatformI18nName *CreateTaskSubtaskRespSubtaskOriginPlatformI18nName `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。需提供多语言版本。
	Href             *CreateTaskSubtaskRespSubtaskOriginHref             `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// CreateTaskSubtaskRespSubtaskOriginHref ...
type CreateTaskSubtaskRespSubtaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 链接对应的地址
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// CreateTaskSubtaskRespSubtaskOriginPlatformI18nName ...
type CreateTaskSubtaskRespSubtaskOriginPlatformI18nName struct {
	EnUs string `json:"en_us,omitempty"` // 英文
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	ZhHk string `json:"zh_hk,omitempty"` // 中文（香港地区）
	ZhTw string `json:"zh_tw,omitempty"` // 中文（台湾地区）
	JaJp string `json:"ja_jp,omitempty"` // 日语
	FrFr string `json:"fr_fr,omitempty"` // 法语
	ItIt string `json:"it_it,omitempty"` // 意大利语
	DeDe string `json:"de_de,omitempty"` // 德语
	RuRu string `json:"ru_ru,omitempty"` // 俄语
	ThTh string `json:"th_th,omitempty"` // 泰语
	EsEs string `json:"es_es,omitempty"` // 西班牙语
	KoKr string `json:"ko_kr,omitempty"` // 韩语
}

// CreateTaskSubtaskRespSubtaskReminder ...
type CreateTaskSubtaskRespSubtaskReminder struct {
	ID                 string `json:"id,omitempty"`                   // 提醒时间设置的 ID
	RelativeFireMinute int64  `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间分钟数。例如30表示截止时间前30分钟提醒；0表示截止时提醒。
}

// CreateTaskSubtaskRespSubtaskStart ...
type CreateTaskSubtaskRespSubtaskStart struct {
	Timestamp string `json:"timestamp,omitempty"`  // 开始时间/日期的时间戳, 距1970-01-01 00:00:00 UTC的毫秒数。如果开始时间是一个日期, 需要把日期转换成时间戳, 并设置 is_all_day=true。
	IsAllDay  bool   `json:"is_all_day,omitempty"` // 是否开始于一个日期。如果设为true, timestamp中只有日期的部分会被解析和存储。
}

// CreateTaskSubtaskRespSubtaskTasklist ...
type CreateTaskSubtaskRespSubtaskTasklist struct {
	TasklistGuid string `json:"tasklist_guid,omitempty"` // 任务所在清单的guid
	SectionGuid  string `json:"section_guid,omitempty"`  // 任务所在清单的自定义分组guid
}

// createTaskSubtaskResp ...
type createTaskSubtaskResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *CreateTaskSubtaskResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}