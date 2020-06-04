package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 发送短信验证码
func SendSMSLuosinao(mobile string, msg string) (err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST",
		"http://sms-api.luosimao.com/v1/send.json",
		strings.NewReader("mobile="+mobile+"&message="+msg+"【叮当工匠】"))
	if err != nil {
		return errors.New(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", "key-8038fe9d1e3f6f5710c6df1205c6e5aa")
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println(string(body))
	return nil
}
