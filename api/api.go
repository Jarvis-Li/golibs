package api

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"

	"ggs/log"

	"ybztserv/conf"
)

const (
	ErrorCode   = 400
	SuccessCode = 0
)

func CheckParams(req *http.Request, required []string) (params url.Values, err error) {
	if req.Method != "POST" {
		err = fmt.Errorf("unvalid request method: %v", req.Method)
		return
	}

	err = req.ParseForm()
	if err != nil {
		return
	}
	params = req.Form
	log.Info("receive api params: %v", params)

	err = CheckRequiredParams(params, required)
	if err != nil {
		return
	}

	sign := CreateSign(params)
	if sign != params.Get("sign") {
		err = errors.New("sign mismatch")
		log.Debug("sign error, wanted: %v, provided: %v", sign, params.Get("sign"))
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

func CreateSign(params map[string][]string) (str string) {
	keys := SortStrKeys(params)

	for _, key := range keys {
		if key == "sign" {
			continue
		}
		str += key + "=" + params[key][0]
	}

	str += conf.Env.SecureKey
	str = fmt.Sprintf("%x", md5.Sum([]byte(str)))
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

func ResponseDone(w http.ResponseWriter, data string) {
	msg := map[string]interface{}{
		"code": SuccessCode,
	}

	if data != "" {
		msg["data"] = data
	}

	b, _ := json.Marshal(msg)
	w.Write(b)
}
