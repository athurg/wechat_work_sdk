package main

import (
	"log"
	"net/http"
	"wechat"

    "github.com/tencentyun/scf-go-lib/cloudfunction"
)

func main() {
    cloudfunction.Start(hello)
}

func hello() error {
	client := wechat.NewAgentClientFromEnv()

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
