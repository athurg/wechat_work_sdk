package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-http/wechat_work"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func main() {
	client := wechat.NewAgentClientFromEnv()

	client.SendTextToUsers("中华英豪", "fengjianbo")

	message := wechat.NewNewsMessage()
	message.Append("标题标题3", "http://b22aiodu.com", "描sfa描述", "")
	message.Append("标题标题44", "http://ba22iodu.com", "描述描述s", "")

	log.Println(client.SendNewsMessageToUsers(message, "fengjianbo"))

	if _, ok := os.LookupEnv("TENCENTCLOUD_RUNENV"); ok {
		cloudfunction.Start(hello)
		return
	}

	hello()
}

func hello() error {
	client := wechat.NewAgentClientFromEnv()
	message := wechat.NewNewsMessage()
	message.Append("标题标题", "http://baiodu.com", "描述描述", "")
	message.Append("标题标题2", "http://baiodu.com", "描述描述s", "")

	client.SendNewsMessageToUsers(message, "fengjianbo")

	token, err := client.GetAccessTokenFromCache()
	if err != nil {
		return err
	}
	log.Println("AccessToken is ", token)

	resp, err := http.Get("https://img.xjh.me/random_img.php?return=302")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	mediaId, expiredAt, err := client.ImageMediaUpload(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("MediaUpload:%#v,%#v", mediaId, expiredAt)

	invalidUsers, err := client.SendImageToUsers(mediaId, "fengjianbo")
	if err != nil {
		return err
	}
	log.Println("MessageSend", invalidUsers)

	return nil
}
