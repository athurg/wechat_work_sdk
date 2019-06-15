package wechat

import (
	"strings"
)

//Message 定义了消息推送中的应用消息
type Message struct {
	ToTag   string `json:"totag"`   //标签ID列表，“|”分隔
	ToUser  string `json:"touser"`  //成员ID列表，“|”分隔
	ToParty string `json:"toparty"` //部门ID列表，“|”分隔

	Safe    int    `json:"safe"`    //是否保密消息，缺省0为否，1表示是
	AgentId int    `json:"agentid"` //企业应用ID
	MsgType string `json:"msgtype"` //消息类型

	Text              *TextMessage              `json:"text,omitempty"`
	Image             *MediaMessage             `json:"image,omitempty"`
	Voice             *MediaMessage             `json:"voice,omitempty"`
	File              *MediaMessage             `json:"file,omitempty"`
	Video             *VideoMessage             `json:"video,omitempty"`
	News              *NewsMessage              `json:"news,omitempty"`
	MpNews            *MpNewsMessage            `json:"mpnews,omitempty"`
	Markdown          *MarkdownMessage          `json:"markdown,omitempty"`
	TextCard          *TextCardMessage          `json:"textcard,omitempty"`
	TaskCard          *TaskCardMessage          `json:"taskcard,omitempty"`
	MiniprogramNotice *MiniprogramNoticeMessage `json:"miniprogramnotice,omitempty"`
}

//SetUser 设置消息的收件人
func (msg *Message) SetUser(users []string) {
	msg.ToUser = strings.Join(users, "|")
}

//TextMessage 定义了消息推送中的文本消息
type TextMessage struct {
	Content string `json:"content"`
}

//NewTextMessage 创建一条文本消息并设置内容，内容支持换行、超链接
func NewTextMessage(content string) *Message {
	return &Message{
		MsgType: "text",
		Text:    &TextMessage{Content: content},
	}
}

//MediaMessage 定义了消息推送中的多媒体消息
type MediaMessage struct {
	MediaId string `json:"media_id"`
}

//NewMediaMessage 创建一条素材消息（image、voice、file）
func NewMediaMessage(mediaType, mediaId string) *Message {
	message := Message{MsgType: mediaType}
	mediaMessage := &MediaMessage{MediaId: mediaId}
	switch mediaType {
	case "image":
		message.Image = mediaMessage
	case "voice":
		message.Voice = mediaMessage
	case "file":
		message.File = mediaMessage
	}

	return &message
}

//VideoMessage 定义了消息推送中的视频消息
type VideoMessage struct {
	MediaMessage
	Title       string `json:"title"`
	Description string `json:"description"`
}

//NewVideoMessage 创建一条视频消息（image、voice、file）
func NewVideoMessage(title, description, mediaId string) *Message {
	videoMessage := VideoMessage{
		Title:       title,
		Description: description,
	}
	videoMessage.MediaId = mediaId

	return &Message{
		MsgType: "video",
		Video:   &videoMessage,
	}
}

//TextCardMessage 定义了消息推送中的卡片消息
type TextCardMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

//NewsArticle 定义了消息推送中的图文消息
type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

//MaxNewsArticles 定义了图文消息中最大的图文消息条数
const MaxNewsArticles = 8

//NewsMessage 定义了图文消息中的图文内容
type NewsMessage struct {
	Articles []NewsArticle `json:"articles"`
}

//NewNewsMessage 用于创建一条图文消息
func NewNewsMessage() *NewsMessage {
	return &NewsMessage{
		Articles: []NewsArticle{},
	}
}

//Append 向图文消息添加一条新文章
func (nm *NewsMessage) Append(title, url, description, picUrl string) {
	article := NewsArticle{
		Title:       title,
		Description: description,
		Url:         url,
		PicUrl:      picUrl,
	}
	nm.Articles = append(nm.Articles, article)
}

//MpNewsArticle 定义了消息推送中的微信托管图文消息文章
type MpNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

//MaxMpNewsArticles 代表微信托管图文消息的最大文章数量
const MaxMpNewsArticles = 8

//MpNewsMessage 定义了消息推送中的微信托管图文消息
type MpNewsMessage struct {
	Articles []MpNewsArticle `json:"articles"`
}

//MarkdownMessage 定义了消息推送中的Markdown消息
type MarkdownMessage struct {
	Content string `json:"content"`
}

//NewMarkdownMessage 用于创建一条Markdown消息
func NewMarkdownMessage(content string) *Message {
	return &Message{
		MsgType:  "markdown",
		Markdown: &MarkdownMessage{Content: content},
	}
}

//MiniprogramNoticeContentItem 定义了消息推送中的小程序通知消息内容
type MiniprogramNoticeContentItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//MiniprogramNoticeMessage 定义了消息推送中的小程序通知消息
type MiniprogramNoticeMessage struct {
	Appid             string                         `json:"appid"`
	Page              string                         `json:"page"`
	Title             string                         `json:"title"`
	Description       string                         `json:"description"`
	EmphasisFirstItem bool                           `json:"emphasis_first_item"`
	ContentItem       []MiniprogramNoticeContentItem `json:"content_item"`
}

//TaskCardButton 定义了消息推送中的任务卡片消息按钮
type TaskCardButton struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	ReplaceName string `json:"replace_name"`
	Color       string `json:"color"`
	IsBold      bool   `json:"is_bold"`
}

//TaskCardMessage 定义了消息推送中的任务卡片消息
type TaskCardMessage struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Url         string           `json:"url"`
	TaskId      string           `json:"task_id"`
	Btn         []TaskCardButton `json:"btn"`
}
