package wechat

import (
	"strings"
)

//消息推送中的应用消息定义
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

//设置消息的收件人
func (msg *Message) SetUser(users []string) {
	msg.ToUser = strings.Join(users, "|")
}

type TextMessage struct {
	Content string `json:"content"`
}

//创建一条文本消息并设置内容，内容支持换行、超链接
func NewTextMessage(content string) *Message {
	return &Message{
		MsgType: "text",
		Text:    &TextMessage{Content: content},
	}
}

type MediaMessage struct {
	MediaId string `json:"media_id"`
}

//创建一条素材消息（image、voice、file）
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

type VideoMessage struct {
	MediaMessage
	Title       string `json:"title"`
	Description string `json:"description"`
}

//创建一条视频消息（image、voice、file）
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

type TextCardMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type NewsMessage struct {
	Articles []NewsArticle `json:"articles"`
}

func NewNewsMessage() *NewsMessage {
	return &NewsMessage{
		Articles: []NewsArticle{},
	}
}

func (nm *NewsMessage) Append(title, url, description, picUrl string) {
	article := NewsArticle{
		Title:       title,
		Description: url,
		Url:         description,
		PicUrl:      picUrl,
	}
	nm.Articles = append(nm.Articles, article)
}

type MpNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

type MpNewsMessage struct {
	Articles []MpNewsArticle `json:"articles"`
}

type MarkdownMessage struct {
	Content string `json:"content"`
}

func NewMarkdownMessage(content string) *Message {
	return &Message{
		MsgType:  "markdown",
		Markdown: &MarkdownMessage{Content: content},
	}
}

type MiniprogramNoticeContentItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MiniprogramNoticeMessage struct {
	Appid             string                         `json:"appid"`
	Page              string                         `json:"page"`
	Title             string                         `json:"title"`
	Description       string                         `json:"description"`
	EmphasisFirstItem bool                           `json:"description"`
	ContentItem       []MiniprogramNoticeContentItem `json:"content_item"`
}

type TaskCardButton struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	ReplaceName string `json:"replace_name"`
	Color       string `json:"color"`
	IsBold      bool   `json:"is_bold"`
}

type TaskCardMessage struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Url         string           `json:"url"`
	TaskId      string           `json:"task_id"`
	Btn         []TaskCardButton `json:"btn"`
}
