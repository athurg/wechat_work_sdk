package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//GetAccessTokenFromCache 用于获取AccessToken，如果有缓存过期则自动刷新
func (cli *AgentClient) GetAccessTokenFromCache() (string, error) {
	var err error
	if time.Now().After(cli.AccessTokenExpiresAt) {
		err = cli.RefreshAccessToken()
	}

	return cli.AccessToken, err
}

//RefreshAccessToken 用于刷新AccessToken
func (cli *AgentClient) RefreshAccessToken() error {
	p := url.Values{
		"corpid":     {cli.CorpId},
		"corpsecret": {cli.Secret},
	}

	resp, err := http.Get(WechatApiUrl + "/gettoken?" + p.Encode())
	if err != nil {
		return fmt.Errorf("请求错误: %s", err)
	}

	defer resp.Body.Close()

	var respInfo struct {
		CommonResponse
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return fmt.Errorf("读取错误: %s", err)
	}

	if respInfo.ErrCode != 0 {
		return fmt.Errorf("API错误: [%d]%s", respInfo.ErrCode, respInfo.ErrMsg)
	}

	cli.AccessToken = respInfo.AccessToken
	cli.AccessTokenExpiresAt = time.Now().Add(time.Duration(respInfo.ExpiresIn-2) * time.Second)

	return nil
}
