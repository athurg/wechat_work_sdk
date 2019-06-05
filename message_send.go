package wechat

//给多个用户发送文本应用消息
func (cli *AgentClient) SendTextToUsers(content string, users ...string) (string, error) {
	message := NewTextMessage(content)
	message.SetUser(users)

	invalidUsers, _, _, err := cli.MessageSend(message)

	return invalidUsers, err
}

//消息推送-发送应用消息
func (cli *AgentClient) MessageSend(message *Message) (string, string, string, error) {
	message.AgentId = cli.AgentId

	var resp struct {
		InvalidUser  string
		InvalidParty string
		InvalidTag   string
	}

	err := cli.requestWithToken("POST", "/message/send", nil, message, &resp)
	if err != nil {
		return "", "", "", err
	}

	return resp.InvalidUser, resp.InvalidParty, resp.InvalidTag, nil
}
