package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	// 打开需要翻译文件
	oldStringsFile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer oldStringsFile.Close()

	resource := Resources{}
	// 读取需要翻译文件中的内容。
	data, err := io.ReadAll(oldStringsFile)
	if err != nil {
		log.Fatalln("读取失败：" + err.Error())
		return
	}
	// 将文件中的内容转换成Resources结构体
	err = xml.Unmarshal(data, &resource)
	if err != nil {
		log.Fatalln("解析失败：" + err.Error())
		return
	}
	fmt.Println("开始翻译。")
	// 需要翻译的语言
	for _, lang := range toLang {
		newFileName := fmt.Sprintf("output/strings_%s.xml", lang)
		// 创建新文件
		newStringsFile, err := os.OpenFile(newFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Println("打开新文件失败：" + err.Error())
			newStringsFile.Close()
			break
		}
		lineCount := 0
		// 读取新文件的内容
		newStringFileText, err := io.ReadAll(newStringsFile)
		if err != nil {
			log.Println("读取新文件所有内容失败：" + err.Error())
			newStringsFile.Close()
			break
		}
		// 判断新文件是否存在翻译过的数据。
		compile, err := regexp.Compile("<string name=\".+\">.+")
		if err != nil {
			log.Println("创建正则表达式失败：" + err.Error())
			newStringsFile.Close()
			break
		}
		lineCount = len(compile.FindAll(newStringFileText, -1))
		// 如果没有翻译过的数据就添加xml文件头
		if lineCount <= 0 {
			newStringsFile.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<resources>")
		}
		// 判断文件是否翻译完成
		if lineCount >= len(resource.Strings) {
			fmt.Printf("翻译完成%s。\n", newFileName)
			continue
		}

		// 开始翻译
		for index, value := range resource.Strings[lineCount:] {
			data, err := translate(value.Text, lang)
			if err != nil {
				log.Println("翻译失败：" + err.Error())
				newStringsFile.Close()
				return
			}
			if data == "" {
				continue
			}
			if data != "" {
				value.Text = fixErrorFormat(data)
			}

			fmt.Printf("\r正在翻译第\033[1;00;31m[%d\033[0m\\\033[1;00;32m%d]\033[0m条数据。", lineCount+index+1, len(resource.Strings))
			result, err := xml.MarshalIndent(value, "", "")
			if err != nil {
				log.Println("创建XML失败：" + err.Error())
				newStringsFile.Close()
				break
			}
			newStringsFile.WriteString("\n\t")
			newStringsFile.Write(result)
			time.Sleep(time.Second)
		}
		newStringsFile.WriteString("\n</resources>")
		fmt.Printf("\n翻译完成%s。\n", newFileName)
	}
}
