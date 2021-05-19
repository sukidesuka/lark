// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteDriveComment 删除云文档中的某条回复。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/delete
func (r *DriveService) DeleteDriveComment(ctx context.Context, request *DeleteDriveCommentReq, options ...MethodOptionFunc) (*DeleteDriveCommentResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteDriveComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteDriveComment mock enable")
		return r.cli.mock.mockDriveDeleteDriveComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteDriveComment",
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(deleteDriveCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveDeleteDriveComment(f func(ctx context.Context, request *DeleteDriveCommentReq, options ...MethodOptionFunc) (*DeleteDriveCommentResp, *Response, error)) {
	r.mockDriveDeleteDriveComment = f
}

func (r *Mock) UnMockDriveDeleteDriveComment() {
	r.mockDriveDeleteDriveComment = nil
}

type DeleteDriveCommentReq struct {
	FileType  FileType `query:"file_type" json:"-"` // 文档类型, 示例值："doc", 可选值有: `doc`：文档, `sheet`：表格, `file`：文件
	FileToken string   `path:"file_token" json:"-"` // 文档token, 示例值："doccnHh7U87HOFpii5u5G*****"
	CommentID string   `path:"comment_id" json:"-"` // 评论ID, 示例值："6916106822734578184"
	ReplyID   string   `path:"reply_id" json:"-"`   // 回复ID, 示例值："6916106822734594568"
}

type deleteDriveCommentResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteDriveCommentResp `json:"data,omitempty"`
}

type DeleteDriveCommentResp struct{}