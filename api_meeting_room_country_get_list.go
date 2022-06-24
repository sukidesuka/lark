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

// GetMeetingRoomCountryList 新建建筑时需要标明所处国家/地区，该接口用于获得系统预先提供的可供选择的国家 /地区列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQTNwYjL0UDM24CN1AjN
func (r *MeetingRoomService) GetMeetingRoomCountryList(ctx context.Context, request *GetMeetingRoomCountryListReq, options ...MethodOptionFunc) (*GetMeetingRoomCountryListResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomGetMeetingRoomCountryList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#GetMeetingRoomCountryList mock enable")
		return r.cli.mock.mockMeetingRoomGetMeetingRoomCountryList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "GetMeetingRoomCountryList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/meeting_room/country/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMeetingRoomCountryListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMeetingRoomGetMeetingRoomCountryList mock MeetingRoomGetMeetingRoomCountryList method
func (r *Mock) MockMeetingRoomGetMeetingRoomCountryList(f func(ctx context.Context, request *GetMeetingRoomCountryListReq, options ...MethodOptionFunc) (*GetMeetingRoomCountryListResp, *Response, error)) {
	r.mockMeetingRoomGetMeetingRoomCountryList = f
}

// UnMockMeetingRoomGetMeetingRoomCountryList un-mock MeetingRoomGetMeetingRoomCountryList method
func (r *Mock) UnMockMeetingRoomGetMeetingRoomCountryList() {
	r.mockMeetingRoomGetMeetingRoomCountryList = nil
}

// GetMeetingRoomCountryListReq ...
type GetMeetingRoomCountryListReq struct {
}

// getMeetingRoomCountryListResp ...
type getMeetingRoomCountryListResp struct {
	Code int64                          `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *GetMeetingRoomCountryListResp `json:"data,omitempty"` // 返回业务信息
}

// GetMeetingRoomCountryListResp ...
type GetMeetingRoomCountryListResp struct {
	Countries *GetMeetingRoomCountryListRespCountries `json:"countries,omitempty"` // 国家地区列表
}

// GetMeetingRoomCountryListRespCountries ...
type GetMeetingRoomCountryListRespCountries struct {
	CountryID string `json:"country_id,omitempty"` // 国家地区ID
	Name      string `json:"name,omitempty"`       // 国家地区名称
}
