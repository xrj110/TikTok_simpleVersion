package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GenerateCover(videoPath string, coverPath string, name string) error {
	url := "http://localhost:5000/process_video" // 替换成实际的目标URL
	// 构建请求体
	payload := strings.NewReader("videoPath=" + videoPath + "&coverPath=" + coverPath + "&name=" + name) // 替换成实际的path和coverPath值

	// 发送POST请求
	response, err := http.Post(url, "application/x-www-form-urlencoded", payload)
	if err != nil {
		fmt.Println("发送请求时出现错误:", err)
		return err
	}
	defer response.Body.Close()

	// 读取响应的内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应时出现错误:", err)
		return err
	}

	// 打印响应的内容
	fmt.Println("响应内容:", string(body))
	return nil
}

func AbsPath() string {
	workingDir, _ := os.Getwd()

	// 将当前工作目录转换为绝对路径
	absolutePath, _ := filepath.Abs(workingDir)
	return absolutePath
}
