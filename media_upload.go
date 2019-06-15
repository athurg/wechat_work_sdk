package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

//构造素材上传表单
func buildMediaUploadForm(fileReader io.Reader) (io.Reader, string, error) {
	buffer := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(buffer)

	fileWriter, err := multipartWriter.CreateFormFile("media", "filename")
	if err != nil {
		return nil, "", err
	}

	_, err = io.Copy(fileWriter, fileReader)
	if err != nil {
		return nil, "", err
	}

	//必须手动关闭以写入boundary结束标记
	err = multipartWriter.Close()
	if err != nil {
		return nil, "", err
	}

	return buffer, multipartWriter.FormDataContentType(), nil
}

//MediaUpload 用于上传临时素材，返回素材ID和上传时间戳（三天后素材自动过期）
func (cli *AgentClient) MediaUpload(mediaType string, mediaReader io.Reader) (string, int, error) {
	token, err := cli.GetAccessTokenFromCache()
	if err != nil {
		return "", 0, fmt.Errorf("获取Token错误: %s", err)
	}

	query := url.Values{
		"access_token": {token},
		"type":         {mediaType},
	}

	reqBody, contentType, err := buildMediaUploadForm(mediaReader)
	if err != nil {
		return "", 0, fmt.Errorf("构造表单请求错误: %s", err)
	}

	resp, err := http.Post(WechatApiUrl+"/media/upload?"+query.Encode(), contentType, reqBody)
	if err != nil {
		return "", 0, fmt.Errorf("执行请求错误: %s", err)
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("读取响应错误: %s", err)
	}

	var info CommonResponse
	err = json.Unmarshal(buf, &info)
	if err != nil {
		return "", 0, fmt.Errorf("解析响应错误: %s", err)
	}

	if info.ErrCode != 0 {
		return "", 0, fmt.Errorf("API错误: [%d]%s", info.ErrCode, info.ErrMsg)
	}

	var respInfo struct {
		MediaId   string `json:"media_id"`
		CreatedAt int    `json:"created_at,string"`
	}
	err = json.Unmarshal(buf, &respInfo)
	if err != nil {
		return "", 0, fmt.Errorf("解析响应错误: %s", err)
	}

	return respInfo.MediaId, respInfo.CreatedAt, nil
}

//ImageMediaUpload 用于上传临时图片素材，返回素材ID和上传时间戳（三天后素材自动过期）
func (cli *AgentClient) ImageMediaUpload(mediaReader io.Reader) (string, int, error) {
	return cli.MediaUpload("image", mediaReader)
}

//VoiceMediaUpload 用于上传临时语音素材，返回素材ID和上传时间戳（三天后素材自动过期）
func (cli *AgentClient) VoiceMediaUpload(mediaReader io.Reader) (string, int, error) {
	return cli.MediaUpload("voice", mediaReader)
}

//VideoMediaUpload 用于上传临时视频素材，返回素材ID和上传时间戳（三天后素材自动过期）
func (cli *AgentClient) VideoMediaUpload(mediaReader io.Reader) (string, int, error) {
	return cli.MediaUpload("video", mediaReader)
}

//FileMediaUpload 用于上传临时文件素材，返回素材ID和上传时间戳（三天后素材自动过期）
func (cli *AgentClient) FileMediaUpload(mediaReader io.Reader) (string, int, error) {
	return cli.MediaUpload("file", mediaReader)
}
