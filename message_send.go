package wechat

//SendImageToUsers 用于给多个用户发送文本应用消息
func (cli *AgentClient) SendImageToUsers(mediaId string, users ...string) (string, error) {
	message := NewMediaMessage("image", mediaId)
	message.SetUser(users)

	invalidUsers, _, _, err := cli.MessageSend(message)

	return invalidUsers, err
}

//SendTextToUsers 用于给多个用户发送文本应用消息
func (cli *AgentClient) SendTextToUsers(content string, users ...string) (string, error) {
	message := NewTextMessage(content)
	message.SetUser(users)

	invalidUsers, _, _, err := cli.MessageSend(message)

	return invalidUsers, err
}

//SendTextCardToUsers 用于给多个用户发送文本应用消息
func (cli *AgentClient) SendTextCardToUsers(title, description, url string, users ...string) (string, error) {
	message := NewTextCardMessage(title, description, url)
	message.SetUser(users)

	invalidUsers, _, _, err := cli.MessageSend(message)

	return invalidUsers, err
}

//SendNewsMessageToUsers 用于给多个用户发送图文消息
func (cli *AgentClient) SendNewsMessageToUsers(newsMessage *NewsMessage, users ...string) (string, error) {
	message := &Message{
		MsgType: "news",
		News:    newsMessage,
	}
	message.SetUser(users)

	invalidUsers, _, _, err := cli.MessageSend(message)

	return invalidUsers, err
}

//SendBatchNewsMessageToUsers 给多个用户发送批量图文消息，当图文信息超过单条上限时自动分隔成多条,并且每一条消息都用原来的第一条作为封面消息
func (cli *AgentClient) SendBatchNewsMessageToUsers(newsMessage *NewsMessage, users ...string) (string, error) {
	size := len(newsMessage.Articles)
	if size <= 1 {
		return cli.SendNewsMessageToUsers(newsMessage, users...)
	}

	coverArticle := newsMessage.Articles[0]

	normalArticles := newsMessage.Articles[1:size]
	normalSize := size - 1
	normalStep := MaxNewsArticles - 1

	var allInvalidUsers string
	for from := 0; from < normalSize; from += normalStep {
		to := from + normalStep
		if to > normalSize {
			to = normalSize
		}

		nm := &NewsMessage{
			Articles: append([]NewsArticle{coverArticle}, normalArticles[from:to]...),
		}

		invalidUsers, err := cli.SendNewsMessageToUsers(nm, users...)
		if err != nil {
			return "", err
		}

		if invalidUsers != "" {
			allInvalidUsers += "|" + invalidUsers
		}

	}

	if allInvalidUsers != "" {
		allInvalidUsers = allInvalidUsers[1:]
	}

	return allInvalidUsers, nil
}

//MessageSend 用于消息推送-发送应用消息
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
