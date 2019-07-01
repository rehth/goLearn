package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// Go 同时支持标准的和 URL 兼容的 base64 格式
	date := "abc123!?$*&()'-=@~"

	// 标准 base64 编码，解码
	sEnc := b64.StdEncoding.EncodeToString([]byte(date))
	fmt.Println("Std:", sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Std:", string(sDec))

	// URL 兼容的 base64 编码，解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(date))
	fmt.Println("URL:", uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println("URL:", string(uDec))

}
