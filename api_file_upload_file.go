package lark

import (
	"context"
	"io"
)

// UploadFile 上传文件，可以上传视频，音频和常见的文件类型
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create
func (r *FileAPI) UploadFile(ctx context.Context, request *UploadFileReq) (*UploadFileResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/files",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                true,
	}
	resp := new(uploadFileResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("File", "UploadFile", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UploadFileReq struct {
	FileType FileType  `json:"file_type,omitempty"` // 文件类型,**示例值**："mp4",**可选值有**：,- `opus`：上传opus音频文件；,其他格式的音频文件，请转为opus格式后上传，转换方式可参考：ffmpeg -i  SourceFile.mp3 -acodec libopus -ac 1 -ar 16000 TargetFile.opus,- `mp4`：上传mp4视频文件,- `pdf`：上传pdf格式文件,- `doc`：上传doc格式文件,- `xls`：上传xls格式文件,- `ppt`：上传ppt格式文件,- `stream`：上传stream格式文件
	FileName string    `json:"file_name,omitempty"` // 带后缀的文件名,**示例值**："测试视频.mp4"
	Duration *int      `json:"duration,omitempty"`  // 文件的时长(视频，音频),单位:毫秒,**示例值**：3000
	File     io.Reader `json:"file,omitempty"`      // 文件内容,**示例值**：二进制文件
}

type uploadFileResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *UploadFileResp `json:"data,omitempty"` // \-
}

type UploadFileResp struct {
	FileKey string `json:"file_key,omitempty"` // 文件的key
}