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

// EventV2CorehrPreHireUpdatedV1 待入职人员信息更新后, 触发此事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/pre_hire/events/updated
func (r *EventCallbackService) HandlerEventV2CorehrPreHireUpdatedV1(f EventV2CorehrPreHireUpdatedV1Handler) {
	r.cli.eventHandler.eventV2CorehrPreHireUpdatedV1Handler = f
}

// EventV2CorehrPreHireUpdatedV1Handler event EventV2CorehrPreHireUpdatedV1 handler
type EventV2CorehrPreHireUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrPreHireUpdatedV1) (string, error)

// EventV2CorehrPreHireUpdatedV1 ...
type EventV2CorehrPreHireUpdatedV1 struct {
	PreHireID    string   `json:"pre_hire_id,omitempty"`   // 待入职 ID
	FieldChanges []string `json:"field_changes,omitempty"` // 变更的字段
}
