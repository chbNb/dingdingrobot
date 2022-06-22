package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"time"
	"unsafe"
)

func main() {
	req := &Request{}
	req.Msgtype = "text"
	req.Text.Content = "本周值班同学："
	req.At.isAtAll = false
	// 当前值班组
	idx := 5
	c := cron.New()
	//_ = c.AddFunc(cron_, func() {
	fmt.Println(time.Now().String())
	ids := timeMap[idx]
	idOne, idTwo := ids[0], ids[1]
	req.Text.Content += member[idOne] + "，" + member[idTwo]
	// 必须添加电话才能@
	req.At.AtUserIds = append(req.At.AtUserIds, userIds[idOne])
	req.At.AtUserIds = append(req.At.AtUserIds, userIds[idTwo])
	req.At.AtMobiles = append(req.At.AtMobiles, phone[idOne])
	req.At.AtMobiles = append(req.At.AtMobiles, phone[idTwo])
	bytesData, err := json.Marshal(req)
	if err != nil {
		fmt.Errorf("AddFunc error : %v", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	// 清空
	idx = (idx + 1) % len(userIds)
	req.Text.Content = "本周值班同学："
	req.At.AtUserIds = req.At.AtUserIds[:0]
	req.At.AtMobiles = req.At.AtMobiles[:0]
	//})
	c.Start()
	defer c.Stop()
	select {}
}
