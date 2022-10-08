package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TranslateResult struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

var hash = md5.New()

type Resources struct {
	XMLName xml.Name `xml:"resources"`
	Strings []Values `xml:"string"`
}

type Values struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Text    string   `xml:",innerxml"`
}

// 翻译
func translate(query, toLang string) (string, error) {
	response, err := http.PostForm(apiUrl, url.Values{
		"q":     {query},
		"from":  {fromLang},
		"to":    {toLang},
		"appid": {appid},
		"salt":  {stringToMD5(query)},
		"sign":  {buildSign(query)},
	})
	response.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	translate := TranslateResult{}
	err = json.Unmarshal(body, &translate)
	if err != nil {
		return "", err
	}

	if len(translate.TransResult) > 0 {
		dst := ""
		for index, text := range translate.TransResult {
			dst += text.Dst
			if index < len(translate.TransResult)-1 {
				dst += "\n"
			}
		}
		return dst, nil
	}

	return "", errors.New(string(body))
}

// 构建签名
// 格式：appid+q+salt+密钥的MD5值
func buildSign(query string) string {
	return stringToMD5(fmt.Sprintf("%s%s%s%s", appid, query, stringToMD5(query), key))
}

// stringToMD5
func stringToMD5(text string) string {
	hash.Reset()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
