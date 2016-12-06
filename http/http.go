package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func ApiPost(apiUrl string, data url.Values) (res map[string]interface{}, err error) {
	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		return
	}
	defer func() {
		resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var cont interface{}
	err = json.Unmarshal(body, &cont)
	if err != nil {
		return
	}
	if cont == nil {
		err = errors.New("api response empty")
		return
	}

	res = cont.(map[string]interface{})
	return
}

func ApiGet(url string, data url.Values) (res map[string]interface{}, err error) {
	url += "?" + data.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var cont interface{}
	err = json.Unmarshal(body, &cont)
	if err != nil {
		return
	}
	if cont == nil {
		err = errors.New("api response empty.")
		return
	}

	res = cont.(map[string]interface{})
	return
}

func ReadMultipart(req *http.Request) (formData map[string][]string, err error) {
	mediaType, ps, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return
	}
	if !strings.HasPrefix(mediaType, "multipart/") {
		err = errors.New(fmt.Sprintf("invalid media type: %v", mediaType))
		return
	}

	mr := multipart.NewReader(req.Body, ps["boundary"])
	form, err := mr.ReadForm(40960)
	if err != nil {
		return
	}
	formData = form.Value
	return
}
