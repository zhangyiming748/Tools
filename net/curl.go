package net

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

func HttpProxyFileUpload(file *multipart.FileHeader, fileKey string, addFields map[string]string, addHeaders map[string]string, urlPath string) (body []byte, err error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile(fileKey, file.Filename)
	if err != nil {
		log.Println(err)
		return
	}

	// 从文件读取数据，写入表单
	srcFile, err := file.Open()
	if err != nil {
		log.Println(err)
		return
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Println(err)
		return
	}
	for fieldKey, fieldVal := range addFields {
		if err = writer.WriteField(fieldKey, fieldVal); err != nil {
			log.Println(err)
			return
		}
	}
	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	req, err := http.NewRequest("POST", urlPath, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}

func HttpPostJson(addHeaders map[string]string, data interface{}, urlPath string) (body []byte, err error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpGet(addHeaders map[string]string, data map[string]string, urlPath string) (body []byte, err error) {
	params := url.Values{}
	urlInfo, err := url.Parse(urlPath)
	if err != nil {
		log.Println(err)
	}
	for dataKey, dataVal := range data {
		params.Set(dataKey, dataVal)
	}
	urlInfo.RawQuery = params.Encode()
	fullUrl := urlInfo.String()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
