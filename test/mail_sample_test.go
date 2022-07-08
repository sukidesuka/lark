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

func Test_Mail_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Mail

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailUser(ctx, &lark.GetMailUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Mail

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailUser(func(ctx context.Context, request *lark.GetMailUserReq, options ...lark.MethodOptionFunc) (*lark.GetMailUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailUser()

			_, _, err := moduleCli.GetMailUser(ctx, &lark.GetMailUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreateMailGroup(func(ctx context.Context, request *lark.CreateMailGroupReq, options ...lark.MethodOptionFunc) (*lark.CreateMailGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreateMailGroup()

			_, _, err := moduleCli.CreateMailGroup(ctx, &lark.CreateMailGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroup(func(ctx context.Context, request *lark.GetMailGroupReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroup()

			_, _, err := moduleCli.GetMailGroup(ctx, &lark.GetMailGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroupList(func(ctx context.Context, request *lark.GetMailGroupListReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroupList()

			_, _, err := moduleCli.GetMailGroupList(ctx, &lark.GetMailGroupListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailUpdateMailGroupPatch(func(ctx context.Context, request *lark.UpdateMailGroupPatchReq, options ...lark.MethodOptionFunc) (*lark.UpdateMailGroupPatchResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailUpdateMailGroupPatch()

			_, _, err := moduleCli.UpdateMailGroupPatch(ctx, &lark.UpdateMailGroupPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailUpdateMailGroup(func(ctx context.Context, request *lark.UpdateMailGroupReq, options ...lark.MethodOptionFunc) (*lark.UpdateMailGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailUpdateMailGroup()

			_, _, err := moduleCli.UpdateMailGroup(ctx, &lark.UpdateMailGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailGroup(func(ctx context.Context, request *lark.DeleteMailGroupReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailGroup()

			_, _, err := moduleCli.DeleteMailGroup(ctx, &lark.DeleteMailGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreateMailGroupMember(func(ctx context.Context, request *lark.CreateMailGroupMemberReq, options ...lark.MethodOptionFunc) (*lark.CreateMailGroupMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreateMailGroupMember()

			_, _, err := moduleCli.CreateMailGroupMember(ctx, &lark.CreateMailGroupMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroupMember(func(ctx context.Context, request *lark.GetMailGroupMemberReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroupMember()

			_, _, err := moduleCli.GetMailGroupMember(ctx, &lark.GetMailGroupMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroupMemberList(func(ctx context.Context, request *lark.GetMailGroupMemberListReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupMemberListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroupMemberList()

			_, _, err := moduleCli.GetMailGroupMemberList(ctx, &lark.GetMailGroupMemberListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailGroupMember(func(ctx context.Context, request *lark.DeleteMailGroupMemberReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailGroupMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailGroupMember()

			_, _, err := moduleCli.DeleteMailGroupMember(ctx, &lark.DeleteMailGroupMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreateMailGroupPermissionMember(func(ctx context.Context, request *lark.CreateMailGroupPermissionMemberReq, options ...lark.MethodOptionFunc) (*lark.CreateMailGroupPermissionMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreateMailGroupPermissionMember()

			_, _, err := moduleCli.CreateMailGroupPermissionMember(ctx, &lark.CreateMailGroupPermissionMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroupPermissionMember(func(ctx context.Context, request *lark.GetMailGroupPermissionMemberReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupPermissionMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroupPermissionMember()

			_, _, err := moduleCli.GetMailGroupPermissionMember(ctx, &lark.GetMailGroupPermissionMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroupPermissionMemberList(func(ctx context.Context, request *lark.GetMailGroupPermissionMemberListReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupPermissionMemberListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroupPermissionMemberList()

			_, _, err := moduleCli.GetMailGroupPermissionMemberList(ctx, &lark.GetMailGroupPermissionMemberListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailGroupPermissionMember(func(ctx context.Context, request *lark.DeleteMailGroupPermissionMemberReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailGroupPermissionMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailGroupPermissionMember()

			_, _, err := moduleCli.DeleteMailGroupPermissionMember(ctx, &lark.DeleteMailGroupPermissionMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreateMailGroupAlias(func(ctx context.Context, request *lark.CreateMailGroupAliasReq, options ...lark.MethodOptionFunc) (*lark.CreateMailGroupAliasResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreateMailGroupAlias()

			_, _, err := moduleCli.CreateMailGroupAlias(ctx, &lark.CreateMailGroupAliasReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailGroupAliasList(func(ctx context.Context, request *lark.GetMailGroupAliasListReq, options ...lark.MethodOptionFunc) (*lark.GetMailGroupAliasListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailGroupAliasList()

			_, _, err := moduleCli.GetMailGroupAliasList(ctx, &lark.GetMailGroupAliasListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailGroupAlias(func(ctx context.Context, request *lark.DeleteMailGroupAliasReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailGroupAliasResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailGroupAlias()

			_, _, err := moduleCli.DeleteMailGroupAlias(ctx, &lark.DeleteMailGroupAliasReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreatePublicMailbox(func(ctx context.Context, request *lark.CreatePublicMailboxReq, options ...lark.MethodOptionFunc) (*lark.CreatePublicMailboxResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreatePublicMailbox()

			_, _, err := moduleCli.CreatePublicMailbox(ctx, &lark.CreatePublicMailboxReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetPublicMailbox(func(ctx context.Context, request *lark.GetPublicMailboxReq, options ...lark.MethodOptionFunc) (*lark.GetPublicMailboxResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetPublicMailbox()

			_, _, err := moduleCli.GetPublicMailbox(ctx, &lark.GetPublicMailboxReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetPublicMailboxList(func(ctx context.Context, request *lark.GetPublicMailboxListReq, options ...lark.MethodOptionFunc) (*lark.GetPublicMailboxListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetPublicMailboxList()

			_, _, err := moduleCli.GetPublicMailboxList(ctx, &lark.GetPublicMailboxListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailUpdatePublicMailboxPatch(func(ctx context.Context, request *lark.UpdatePublicMailboxPatchReq, options ...lark.MethodOptionFunc) (*lark.UpdatePublicMailboxPatchResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailUpdatePublicMailboxPatch()

			_, _, err := moduleCli.UpdatePublicMailboxPatch(ctx, &lark.UpdatePublicMailboxPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailUpdatePublicMailbox(func(ctx context.Context, request *lark.UpdatePublicMailboxReq, options ...lark.MethodOptionFunc) (*lark.UpdatePublicMailboxResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailUpdatePublicMailbox()

			_, _, err := moduleCli.UpdatePublicMailbox(ctx, &lark.UpdatePublicMailboxReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeletePublicMailbox(func(ctx context.Context, request *lark.DeletePublicMailboxReq, options ...lark.MethodOptionFunc) (*lark.DeletePublicMailboxResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeletePublicMailbox()

			_, _, err := moduleCli.DeletePublicMailbox(ctx, &lark.DeletePublicMailboxReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreatePublicMailboxMember(func(ctx context.Context, request *lark.CreatePublicMailboxMemberReq, options ...lark.MethodOptionFunc) (*lark.CreatePublicMailboxMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreatePublicMailboxMember()

			_, _, err := moduleCli.CreatePublicMailboxMember(ctx, &lark.CreatePublicMailboxMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetPublicMailboxMember(func(ctx context.Context, request *lark.GetPublicMailboxMemberReq, options ...lark.MethodOptionFunc) (*lark.GetPublicMailboxMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetPublicMailboxMember()

			_, _, err := moduleCli.GetPublicMailboxMember(ctx, &lark.GetPublicMailboxMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetPublicMailboxMemberList(func(ctx context.Context, request *lark.GetPublicMailboxMemberListReq, options ...lark.MethodOptionFunc) (*lark.GetPublicMailboxMemberListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetPublicMailboxMemberList()

			_, _, err := moduleCli.GetPublicMailboxMemberList(ctx, &lark.GetPublicMailboxMemberListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeletePublicMailboxMember(func(ctx context.Context, request *lark.DeletePublicMailboxMemberReq, options ...lark.MethodOptionFunc) (*lark.DeletePublicMailboxMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeletePublicMailboxMember()

			_, _, err := moduleCli.DeletePublicMailboxMember(ctx, &lark.DeletePublicMailboxMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailClearPublicMailboxMember(func(ctx context.Context, request *lark.ClearPublicMailboxMemberReq, options ...lark.MethodOptionFunc) (*lark.ClearPublicMailboxMemberResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailClearPublicMailboxMember()

			_, _, err := moduleCli.ClearPublicMailboxMember(ctx, &lark.ClearPublicMailboxMemberReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreateMailPublicMailboxAlias(func(ctx context.Context, request *lark.CreateMailPublicMailboxAliasReq, options ...lark.MethodOptionFunc) (*lark.CreateMailPublicMailboxAliasResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreateMailPublicMailboxAlias()

			_, _, err := moduleCli.CreateMailPublicMailboxAlias(ctx, &lark.CreateMailPublicMailboxAliasReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailPublicMailboxAliasList(func(ctx context.Context, request *lark.GetMailPublicMailboxAliasListReq, options ...lark.MethodOptionFunc) (*lark.GetMailPublicMailboxAliasListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailPublicMailboxAliasList()

			_, _, err := moduleCli.GetMailPublicMailboxAliasList(ctx, &lark.GetMailPublicMailboxAliasListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailPublicMailboxAlias(func(ctx context.Context, request *lark.DeleteMailPublicMailboxAliasReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailPublicMailboxAliasResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailPublicMailboxAlias()

			_, _, err := moduleCli.DeleteMailPublicMailboxAlias(ctx, &lark.DeleteMailPublicMailboxAliasReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailCreateMailUserMailboxAlias(func(ctx context.Context, request *lark.CreateMailUserMailboxAliasReq, options ...lark.MethodOptionFunc) (*lark.CreateMailUserMailboxAliasResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailCreateMailUserMailboxAlias()

			_, _, err := moduleCli.CreateMailUserMailboxAlias(ctx, &lark.CreateMailUserMailboxAliasReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailUserMailboxAlias(func(ctx context.Context, request *lark.DeleteMailUserMailboxAliasReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailUserMailboxAliasResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailUserMailboxAlias()

			_, _, err := moduleCli.DeleteMailUserMailboxAlias(ctx, &lark.DeleteMailUserMailboxAliasReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailGetMailUserMailboxAliasList(func(ctx context.Context, request *lark.GetMailUserMailboxAliasListReq, options ...lark.MethodOptionFunc) (*lark.GetMailUserMailboxAliasListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailGetMailUserMailboxAliasList()

			_, _, err := moduleCli.GetMailUserMailboxAliasList(ctx, &lark.GetMailUserMailboxAliasListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockMailDeleteMailUserMailbox(func(ctx context.Context, request *lark.DeleteMailUserMailboxReq, options ...lark.MethodOptionFunc) (*lark.DeleteMailUserMailboxResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockMailDeleteMailUserMailbox()

			_, _, err := moduleCli.DeleteMailUserMailbox(ctx, &lark.DeleteMailUserMailboxReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Mail

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailUser(ctx, &lark.GetMailUserReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroup(ctx, &lark.CreateMailGroupReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroup(ctx, &lark.GetMailGroupReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupList(ctx, &lark.GetMailGroupListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateMailGroupPatch(ctx, &lark.UpdateMailGroupPatchReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateMailGroup(ctx, &lark.UpdateMailGroupReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroup(ctx, &lark.DeleteMailGroupReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroupMember(ctx, &lark.CreateMailGroupMemberReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupMember(ctx, &lark.GetMailGroupMemberReq{
				MailGroupID: "x",
				MemberID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupMemberList(ctx, &lark.GetMailGroupMemberListReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroupMember(ctx, &lark.DeleteMailGroupMemberReq{
				MailGroupID: "x",
				MemberID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroupPermissionMember(ctx, &lark.CreateMailGroupPermissionMemberReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupPermissionMember(ctx, &lark.GetMailGroupPermissionMemberReq{
				MailGroupID:        "x",
				PermissionMemberID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupPermissionMemberList(ctx, &lark.GetMailGroupPermissionMemberListReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroupPermissionMember(ctx, &lark.DeleteMailGroupPermissionMemberReq{
				MailGroupID:        "x",
				PermissionMemberID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroupAlias(ctx, &lark.CreateMailGroupAliasReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupAliasList(ctx, &lark.GetMailGroupAliasListReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroupAlias(ctx, &lark.DeleteMailGroupAliasReq{
				MailGroupID: "x",
				AliasID:     "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePublicMailbox(ctx, &lark.CreatePublicMailboxReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailbox(ctx, &lark.GetPublicMailboxReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailboxList(ctx, &lark.GetPublicMailboxListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdatePublicMailboxPatch(ctx, &lark.UpdatePublicMailboxPatchReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdatePublicMailbox(ctx, &lark.UpdatePublicMailboxReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeletePublicMailbox(ctx, &lark.DeletePublicMailboxReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePublicMailboxMember(ctx, &lark.CreatePublicMailboxMemberReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailboxMember(ctx, &lark.GetPublicMailboxMemberReq{
				PublicMailboxID: "x",
				MemberID:        "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailboxMemberList(ctx, &lark.GetPublicMailboxMemberListReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeletePublicMailboxMember(ctx, &lark.DeletePublicMailboxMemberReq{
				PublicMailboxID: "x",
				MemberID:        "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ClearPublicMailboxMember(ctx, &lark.ClearPublicMailboxMemberReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailPublicMailboxAlias(ctx, &lark.CreateMailPublicMailboxAliasReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailPublicMailboxAliasList(ctx, &lark.GetMailPublicMailboxAliasListReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailPublicMailboxAlias(ctx, &lark.DeleteMailPublicMailboxAliasReq{
				PublicMailboxID: "x",
				AliasID:         "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailUserMailboxAlias(ctx, &lark.CreateMailUserMailboxAliasReq{
				UserMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailUserMailboxAlias(ctx, &lark.DeleteMailUserMailboxAliasReq{
				UserMailboxID: "x",
				AliasID:       "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailUserMailboxAliasList(ctx, &lark.GetMailUserMailboxAliasListReq{
				UserMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailUserMailbox(ctx, &lark.DeleteMailUserMailboxReq{
				UserMailboxID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Mail
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailUser(ctx, &lark.GetMailUserReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroup(ctx, &lark.CreateMailGroupReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroup(ctx, &lark.GetMailGroupReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupList(ctx, &lark.GetMailGroupListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateMailGroupPatch(ctx, &lark.UpdateMailGroupPatchReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateMailGroup(ctx, &lark.UpdateMailGroupReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroup(ctx, &lark.DeleteMailGroupReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroupMember(ctx, &lark.CreateMailGroupMemberReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupMember(ctx, &lark.GetMailGroupMemberReq{
				MailGroupID: "x",
				MemberID:    "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupMemberList(ctx, &lark.GetMailGroupMemberListReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroupMember(ctx, &lark.DeleteMailGroupMemberReq{
				MailGroupID: "x",
				MemberID:    "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroupPermissionMember(ctx, &lark.CreateMailGroupPermissionMemberReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupPermissionMember(ctx, &lark.GetMailGroupPermissionMemberReq{
				MailGroupID:        "x",
				PermissionMemberID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupPermissionMemberList(ctx, &lark.GetMailGroupPermissionMemberListReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroupPermissionMember(ctx, &lark.DeleteMailGroupPermissionMemberReq{
				MailGroupID:        "x",
				PermissionMemberID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailGroupAlias(ctx, &lark.CreateMailGroupAliasReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailGroupAliasList(ctx, &lark.GetMailGroupAliasListReq{
				MailGroupID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailGroupAlias(ctx, &lark.DeleteMailGroupAliasReq{
				MailGroupID: "x",
				AliasID:     "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePublicMailbox(ctx, &lark.CreatePublicMailboxReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailbox(ctx, &lark.GetPublicMailboxReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailboxList(ctx, &lark.GetPublicMailboxListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdatePublicMailboxPatch(ctx, &lark.UpdatePublicMailboxPatchReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdatePublicMailbox(ctx, &lark.UpdatePublicMailboxReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeletePublicMailbox(ctx, &lark.DeletePublicMailboxReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreatePublicMailboxMember(ctx, &lark.CreatePublicMailboxMemberReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailboxMember(ctx, &lark.GetPublicMailboxMemberReq{
				PublicMailboxID: "x",
				MemberID:        "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetPublicMailboxMemberList(ctx, &lark.GetPublicMailboxMemberListReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeletePublicMailboxMember(ctx, &lark.DeletePublicMailboxMemberReq{
				PublicMailboxID: "x",
				MemberID:        "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ClearPublicMailboxMember(ctx, &lark.ClearPublicMailboxMemberReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailPublicMailboxAlias(ctx, &lark.CreateMailPublicMailboxAliasReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailPublicMailboxAliasList(ctx, &lark.GetMailPublicMailboxAliasListReq{
				PublicMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailPublicMailboxAlias(ctx, &lark.DeleteMailPublicMailboxAliasReq{
				PublicMailboxID: "x",
				AliasID:         "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateMailUserMailboxAlias(ctx, &lark.CreateMailUserMailboxAliasReq{
				UserMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailUserMailboxAlias(ctx, &lark.DeleteMailUserMailboxAliasReq{
				UserMailboxID: "x",
				AliasID:       "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetMailUserMailboxAliasList(ctx, &lark.GetMailUserMailboxAliasListReq{
				UserMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DeleteMailUserMailbox(ctx, &lark.DeleteMailUserMailboxReq{
				UserMailboxID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
