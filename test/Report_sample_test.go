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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Report_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Report

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.QueryReportRule(ctx, &lark.QueryReportRuleReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Report

		t.Run("", func(t *testing.T) {

			cli.Mock().MockReportQueryReportRule(func(ctx context.Context, request *lark.QueryReportRuleReq, options ...lark.MethodOptionFunc) (*lark.QueryReportRuleResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockReportQueryReportRule()

			_, _, err := moduleCli.QueryReportRule(ctx, &lark.QueryReportRuleReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockReportRemoveReportView(func(ctx context.Context, request *lark.RemoveReportViewReq, options ...lark.MethodOptionFunc) (*lark.RemoveReportViewResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockReportRemoveReportView()

			_, _, err := moduleCli.RemoveReportView(ctx, &lark.RemoveReportViewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockReportQueryReportTask(func(ctx context.Context, request *lark.QueryReportTaskReq, options ...lark.MethodOptionFunc) (*lark.QueryReportTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockReportQueryReportTask()

			_, _, err := moduleCli.QueryReportTask(ctx, &lark.QueryReportTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Report

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.QueryReportRule(ctx, &lark.QueryReportRuleReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RemoveReportView(ctx, &lark.RemoveReportViewReq{
				RuleID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.QueryReportTask(ctx, &lark.QueryReportTaskReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Report
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.QueryReportRule(ctx, &lark.QueryReportRuleReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.RemoveReportView(ctx, &lark.RemoveReportViewReq{
				RuleID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.QueryReportTask(ctx, &lark.QueryReportTaskReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
