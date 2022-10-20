package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// HttpGet get请求资源，将reval对象传入进行序列化，然后返回
func HttpGet(client *http.Client, reqUrl string, reval any) error {
	//对url进行处理，如果传的参数为空，那么就不进行传参
	u, _ := url.ParseRequestURI(reqUrl)
	reqUrl = u.Scheme + "://" + u.Host + u.Path + "?"
	for k, v := range u.Query() {
		if len(v) > 0 && strings.Join(v, "") != "" {
			reqUrl += fmt.Sprintf("&%v=%v", k, url.QueryEscape(strings.Join(v, "+")))
		}
	}
	resp, err := client.Get(reqUrl)
	if err != nil {
		return errors.Wrap(err, "发送get请求失败")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "读取响应失败")
	}
	err = json.Unmarshal(body, reval)
	if err != nil {
		return errors.Wrap(err, "反序列化失败")
	}
	if gjson.Get(string(body), "errno").Raw != "" && gjson.Get(string(body), "errno").Raw != "0" {
		return errors.New(fmt.Sprintf("api返回错误，错误码：%v,错误信息：%v",
			gjson.Get(string(body), "errno").Raw,
			gjson.Get(string(body), "errmsg").Raw))
	}
	return nil
}

// HttpPostForm post请求资源，将reval对象传入进行序列化，然后返回
func HttpPostForm(client *http.Client, reqUrl string, reqValue url.Values, reval any) error {
	for k, v := range reqValue {
		if len(v) <= 0 || strings.Join(v, "") == "" {
			reqValue.Del(k)
		}
	}

	resp, err := client.PostForm(reqUrl, reqValue)
	if err != nil {
		return errors.Wrap(err, "发送post请求失败")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "读取响应失败")
	}
	err = json.Unmarshal(body, reval)
	if err != nil {
		return errors.Wrap(err, "反序列化失败")
	}
	if gjson.Get(string(body), "errno").Raw != "" && gjson.Get(string(body), "errno").Raw != "0" {
		return errors.New(fmt.Sprintf("api返回错误，错误码：%v,错误信息：%v",
			gjson.Get(string(body), "errno").Raw,
			gjson.Get(string(body), "errmsg").Raw))
	}
	return nil
}

// HttpPostFile Post上传文件
func HttpPostFile(client *http.Client, path string, reqUrl string, reval any) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = io.Copy(part, f)
	err = writer.Close()
	if err != nil {
		return errors.WithStack(err)
	}
	q, err := http.NewRequest("POST", reqUrl, payload)
	if err != nil {
		return errors.WithStack(err)
	}
	q.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(q)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "读取响应失败")
	}
	err = json.Unmarshal(body, reval)
	if err != nil {
		return errors.Wrap(err, "反序列化失败")
	}
	if gjson.Get(string(body), "errno").Raw != "" && gjson.Get(string(body), "errno").Raw != "0" {
		return errors.New(fmt.Sprintf("api返回错误，错误码：%v,错误信息：%v",
			gjson.Get(string(body), "errno").Raw,
			gjson.Get(string(body), "errmsg").Raw))
	}
	return err
}
