package api

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	ErrorCode   = 400
	SuccessCode = 0

	MaxTimeDev int64 = 60 * 1
)

func CheckParams(req *http.Request, required []string, secretKey string) (params url.Values, err error) {
	if req.Method != "POST" {
		err = fmt.Errorf("unvalid request method: %v", req.Method)
		return
	}

	err = req.ParseForm()
	if err != nil {
		return
	}
	params = req.Form

	err = CheckRequiredParams(params, required)
	if err != nil {
		return
	}

	sign := CreateSign(params, secretKey)
	if sign != params.Get("sign") {
		err = errors.New("sign mismatch")
		return
	}

	err = checkTimeDev(params)
	if err != nil {
		return
	}

	return
}

func CheckRequiredParams(params url.Values, required []string) (err error) {
	for _, name := range required {
		if _, ok := params[name]; !ok {
			err = fmt.Errorf("lack of required param: %v", name)
			return
		}
	}
	return
}

func CreateSign(params url.Values, secretKey string) (str string) {
	keys := SortStrKeys(params)

	for _, key := range keys {
		if key == "sign" {
			continue
		}
		str += key + "=" + params.Get(key)
	}

	str += secretKey
	str = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return
}

func checkTimeDev(params map[string][]string) (err error) {
	strs, ok := params["timestamp"]
	if !ok {
		return
	}

	ts, err := strconv.ParseInt(strs[0], 10, 64)
	if err != nil {
		return
	}

	dev := time.Now().Unix() - ts
	if dev > MaxTimeDev || dev < -MaxTimeDev {
		return errors.New("timestamp exceeds max deviation")
	}
	return
}

func SortStrKeys(m map[string][]string) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func ResponseFail(w http.ResponseWriter, code, msg interface{}) {
	data, _ := json.Marshal(map[string]interface{}{
		"code": code,
		"msg":  msg,
	})

	w.Write(data)
}

func ResponseDone(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(map[string]interface{}{
		"code": SuccessCode,
		"data": data,
	})
	if err != nil {
		panic(err)
	}
	w.Write(b)
}
