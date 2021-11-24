package main

import "github.com/axgle/mahonia"

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result
}

func GBKToUTF8(i string) string {
	return ConvertToString(i, "gbk", "utf8")
}

func UTF8ToGBK(i string) string {
	return ConvertToString(i, "utf8", "gbk")
}
