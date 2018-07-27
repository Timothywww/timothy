package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"runtime"
)

var (
	wg sync.WaitGroup
	p  = 0
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Print("\n启动壁纸下载程序\n\n确认开始下载(y/n): ")
	key := ""
	fmt.Scan(&key)
	if key == "y" {
		fmt.Println("\n\t这可能需要一些时间...\n")
		index := 300
		for i := 1; i <= index; i++ {
			getImages("插图壁纸", "http://papers.co/android/page/"+strconv.Itoa(i)+"/")
		}
	}
}

// 图片匹配正则
var imageExp = regexp.MustCompile(`http://papers.co/android/wp-content/uploads/.+\.jpg`)

// 下载方法
func getImages(dirName, url string) {
	body, err := GetGoal(url)
	if err != nil {
		panic(err)
	}
	// 匹配图片链接
	fmt.Print(string(body))
	imgs := imageExp.FindAllStringSubmatch(string(body), 1000)


	// 链接处理
	imgUrl := make([]string, 0)
	for _, v := range imgs {
		u := strings.Split(v[0], "-250x400")
		fmt.Print(u[0]+u[1])
		imgUrl = append(imgUrl, u[0]+u[1])
	}
	//for _, v := range imgUrl {
	//	fmt.Println(v)
	//}

	// 创建保存目录
	if err := os.MkdirAll(dirName, 0777); err != nil {
		panic(err)
	}

	// 下载图片(并行)
	wg.Add(len(imgUrl))
	for i, v := range imgUrl {
		go func(i int, v string) {

			defer wg.Done()
			p++
			filename := dirName + "/wallpaper-" + strconv.Itoa(p) + ".jpg"
			data, err := GetGoal(v)
			if err != nil {
				// 请求超时
				fmt.Print("!")
			} else {
				// 保存本地
				if err := ioutil.WriteFile(filename, data, 0777); err != nil {
					fmt.Print(" ERR ")
				} else {

					fmt.Print("#")
				}
			}
		}(i, v)
	}
	wg.Wait()
	fmt.Println()
}

// 获取目标资源函数
func GetGoal(url string) (data []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// 资源序列化
	data, err = ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	return
}
