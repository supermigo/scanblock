package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

func DataZipString(data interface{}) string {

	return Encode_Go_Base64(Struct2JsonString(data))
}
func Struct2JsonString(data interface{}) string {
	bytestr, _ := json.Marshal(data)
	return string(bytestr)
}

// 编码
func Encode_Go_Base64(input string) string {

	//进行数据解码
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	//去掉传输干扰
	encoded = Replace_Data(encoded)

	return encoded
}

// 数据替换
func Replace_Data(str string) string {

	//替换特殊字符 去掉斜杠干扰
	str = strings.Replace(str, "/", "@", -1)

	//去掉 +号
	str = strings.Replace(str, "+", "_", -1)

	str = strings.Replace(str, "=", "$", -1)

	return str
}

// 获取系统现在有的json的合约。
