# 翻译安卓strings.xml文件

该项目使用了百度翻译的api使用之前请申请appid和key。

+ 支持暂停后再翻译。
+ 支持翻译多种语言。

## 使用

修改`config.go`文件的内容。

```go
const (
filename = "需要翻译的文件"
fromLang = "源语言"
appid    = "百度翻译appid"
key = "百度翻译key"
apiUrl = "https://fanyi-api.baidu.com/api/trans/vip/translate"
)

// 需要翻译的目标语言。
var toLang = [...]string{"en", "cht", "yue", "wyw"}
```

开始翻译

```bash
go run main.go translate.go config.go 
开始翻译。
正在翻译第[1\1]条数据。
翻译完成output/strings_en.xml。
正在翻译第[1\1]条数据。
翻译完成output/strings_cht.xml。
正在翻译第[1\1]条数据。
翻译完成output/strings_yue.xml。
正在翻译第[1\1]条数据。
翻译完成output/strings_wyw.xml。
```

使用`cat`命令查看结果

```bash
cat output/strings_*
# strings_cht.xml
<?xml version="1.0" encoding="utf-8"?>
<resources>
        <string name="test">測試</string>
</resources>
# strings_en.xml
<?xml version="1.0" encoding="utf-8"?>
<resources>
        <string name="test">test</string>
</resources>
# strings_wyw.xml
<?xml version="1.0" encoding="utf-8"?>
<resources>
        <string name="test">试检试</string>
</resources>
# strings_yue.xml
<?xml version="1.0" encoding="utf-8"?>
<resources>
        <string name="test">测试</string>
</resources>
```

## 暂停后再翻译

```bash
 ./translate_strings 
开始翻译。
正在翻译第[1\1]条数据。
翻译完成output/strings_en.xml。
正在翻译第[1\1]条数据。^C
 ./translate_strings 
开始翻译。
翻译完成output/strings_en.xml。
翻译完成output/strings_cht.xml。
正在翻译第[1\1]条数据。
翻译完成output/strings_yue.xml。
正在翻译第[1\1]条数据。
翻译完成output/strings_wyw.xml。
```

## 常见语种列表

| 名称    | 代码  | 名称     | 代码  | 名称    | 代码  |
|-------|-----|--------|-----|-------|-----|
| 中文    | zh  | 繁体中文   | cht | 希腊语   | el  |
| 粤语    | yue | 文言文    | wyw | 意大利语  | it  |
| 英语    | en  | 日语     | jp  | 德语    | de  |
| 韩语    | kor | 法语     | fra | 葡萄牙语  | pt  |
| 西班牙语  | spa | 阿拉伯语   | ara | 俄语    | ru  |
| 荷兰语   | nl  | 保加利亚语  | bul | 爱沙尼亚语 | est |
| 丹麦语   | dan | 芬兰语    | fin | 捷克语   | cs  |
| 罗马尼亚语 | rom | 斯洛文尼亚语 | slo | 瑞典语   | swe |
| 匈牙利语  | hu  | 越南语    | vie |

## 完整语种列表

+ [百度翻译开放平台](http://api.fanyi.baidu.com/product/113)

## 错误码列表

未添加appid和key。

```bash
2022/10/08 14:58:40 翻译失败：{"error_code":"52003","error_msg":"UNAUTHORIZED USER"}
```

| 错误码   | 含义         | 解决方案                               |
|-------|------------|------------------------------------|
| 52000 | 成功         |                                    |
| 52001 | 请求超时       | 请重试                                |
| 52002 | 系统错误       | 请重试                                |
| 52003 | 未授权用户      | 请检查appid是否正确或者服务是否开通               |
| 54000 | 必填参数为空     | 请检查是否少传参数                          |
| 54001 | 签名错误       | 请检查您的签名生成方法                        |
| 54003 | 访问频率受限     | 请降低您的调用频率，或进行身份认证后切换为高级版/尊享版       |
| 54004 | 账户余额不足     | 请前往管理控制台为账户充值                      |
| 54005 | 长query请求频繁 | 请降低长query的发送频率，3s后再试               |
| 58000 | 客户端IP非法    | 检查个人资料里填写的IP地址是否正确，可前往开发者信息-基本信息修改 |
| 58001 | 译文语言方向不支持  | 检查译文语言是否在语言列表里                     |
| 58002 | 服务当前已关闭    | 请前往管理控制台开启服务                       |
| 90107 | 认证未通过或未生效  | 请前往我的认证查看认证进度                      |
