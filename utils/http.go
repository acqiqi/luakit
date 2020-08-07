package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Http Get Json 请求
func HttpGetJson(url string, cb interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	JsonDecode(string(b), &cb)
	return nil
}

// Http Post Json 请求
func HttpPostJson(url string, body interface{}, cb interface{}) error {

	requestBody := JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	log.Println(string(b))
	JsonDecode(string(b), &cb)
	return nil
}

// Http Post Json 请求 不带回调
func HttpPostJsonNotCallback(url string, body interface{}, platform_key string) error {
	requestBody := JsonEncode(body)
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PlatformKey", platform_key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(b))
	}
	log.Println(string(b))
	return nil
}
