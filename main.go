package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func crawler() {
	//构造请求
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://tool.chinaz.com/tools/use", nil)
	if err != nil {
		fmt.Println("请求构建失败", err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败", err)
	}
	defer resp.Body.Close()
	//解析网页
	reader, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析失败", err)
	}
	//获取节点信息
	//body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto

	//body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto > p:nth-child(9)  描述
	//body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto > div:nth-child(10) 颜色

	//body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto > p:nth-child(12) 描述
	//body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto > div:nth-child(13) 颜色

	//body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto > div:nth-child(15)

	text := reader.Find("body > div.Tool-MainWrap.wrapper.mb10 > div.CorContent.auto > div:nth-child(7)").Text()
	var arr []string
	arr = append(arr, text)
	for index, value := range arr {
		fmt.Printf("%v:%v\n", index, value)
	}
}

func main() {
	crawler()
	fmt.Println("complete")
}
