package mtd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var fileHelper = FileHelper{}

type HttpHelper struct{}

func (h *HttpHelper) RegParam(p string) bool {
	r, err := regexp.Compile(fileHelper.CheckConf().Reg)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

func (h *HttpHelper) Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// log.Println(request.Header["User-Agent"])
		handler(writer, request)
	}
}

func (h *HttpHelper) GetMap(r *http.Request) map[string][]string {
	return r.Form
}

func (h *HttpHelper) PostMap(r *http.Request) map[string][]string {
	return r.PostForm
}

func (h *HttpHelper) Get(r *http.Request, key string) string {
	param := strings.TrimSpace(r.URL.Query().Get(key))
	if h.RegParam(param) {
		return param
	} else {
		return ""
	}
}

func (h *HttpHelper) Post(r *http.Request, key string) string {
	param := strings.TrimSpace(r.PostFormValue(key))
	if h.RegParam(param) {
		return param
	} else {
		return ""
	}
}

func (h *HttpHelper) FormFile(w http.ResponseWriter, r *http.Request, key string) (bool, string) {
	f, fheader, err := r.FormFile("file")
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close file err: %v", err.Error())
		}
	}(f)
	if err != nil {
		return false, err.Error()
	}

	newf, err := os.OpenFile("../Temp/upload/"+fheader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, err.Error()
	}
	defer func(f io.Closer) {
		if err := f.Close(); err != nil {
			fmt.Printf("defer close file err: %v", err.Error())
		}
	}(newf)

	_, err = io.Copy(newf, f)
	if err != nil {
		return false, err.Error()
	}

	return true, newf.Name()
}

func (h *HttpHelper) HttpWrite(w http.ResponseWriter, Data interface{}) (int, error) {
	j, err := json.Marshal(Data)
	if err != nil {
		return 0, err
	}
	return w.Write(j)
}
