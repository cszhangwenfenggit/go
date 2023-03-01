package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const XkcdUrl = "http://xkcd.com"

const Omdbapi = "http://www.omdbapi.com"

type Info struct {
	Title      string
	Link       string
	Year       string
	Transcript string
}

func Xkcd(targets []string) {
	for _, v := range targets {
		url := XkcdUrl + "/" + v + "/info.0.json"
		data, err := GetFromUrl(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(url)
		fmt.Printf("Title:%s Year:%s Link:%s\n", data.Title, data.Year, data.Link)
	}
}

func GetFromUrl(url string) (*Info, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("get url %s failed:%s", url, resp.Status)
	}

	var result Info
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, err
}

func DownloadImg(imgUrl string) (int64, error) {
	path := strings.Split(imgUrl, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}

	//生成一个文件，用于报错图片
	out, err := os.Create(name)
	defer out.Close()
	//todo 请求图片地址,获得内容
	resp, err := http.Get(imgUrl)
	defer resp.Body.Close()
	//todo 写入新创建的文件
	pix, err := ioutil.ReadAll(resp.Body)
	n, err := io.Copy(out, bytes.NewReader(pix))

	return n, err
}
