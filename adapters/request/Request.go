package request

import (
	"strings"

	"github.com/monaco-io/request"
)

type Response struct {
	Status int
	Data   interface{}
}

type sRequest struct {
	url           string
	method        string
	params        map[string]string
	headers       map[string]string
	body          map[string]interface{}
	multipartForm request.MultipartForm
}

type SRequest interface {
	send() (Response, error)
	SetMethod(string) *sRequest
	SetParams(map[string]string) *sRequest
	SetHeaders(map[string]string) *sRequest
	SetBody(map[string]interface{}) *sRequest
	SetMultipartForm(map[string]string, []string) *sRequest
	Get() (Response, error)
	Post() (Response, error)
	Put() (Response, error)
	Patch() (Response, error)
	Delete() (Response, error)
}

func Make(url string) SRequest {
	return &sRequest{
		url:     url,
		method:  "GET",
		headers: map[string]string{},
	}
}
func (sR *sRequest) send() (Response, error) {
	var result interface{}
	var rs Response
	c := request.Client{
		URL:    sR.url,
		Method: sR.method,
	}
	if len(sR.headers) == 0 {
		sR.headers = map[string]string{}
	}
	c.Header = sR.headers
	if len(sR.body) > 0 {
		c.JSON = sR.body
	}
	resp := c.Send().Scan(&result)
	rs.Status = resp.Response().StatusCode
	rs.Data = result
	if !resp.OK() {
		return rs, resp.Error()
	}
	return rs, nil
}

func (sR *sRequest) Get() (Response, error) {
	sR.method = "GET"
	queryStr := getUrlQueryString(sR.params)
	if queryStr != "" {
		if strings.Contains(sR.url, "?") {
			sR.url += "&" + queryStr
		} else {
			sR.url += "?" + queryStr
		}
	}
	return sR.send()
}

func getUrlQueryString(query map[string]string) string {
	str := ""
	for k, v := range query {
		str += k + "=" + v + "&"
	}
	str = strings.TrimSuffix(str, "&")
	return str
}

func (sR *sRequest) Post() (Response, error) {
	sR.method = "POST"
	return sR.send()
}
func (sR *sRequest) Put() (Response, error) {
	sR.method = "PUT"
	return sR.send()
}
func (sR *sRequest) Patch() (Response, error) {
	sR.method = "PATCH"
	return sR.send()
}
func (sR *sRequest) Delete() (Response, error) {
	sR.method = "DELETE"
	return sR.send()
}

func (sR *sRequest) SetMethod(m string) *sRequest {
	sR.method = m
	return sR
}
func (sR *sRequest) SetParams(p map[string]string) *sRequest {
	sR.params = p
	return sR
}
func (sR *sRequest) SetHeaders(h map[string]string) *sRequest {
	sR.headers = h
	return sR
}
func (sR *sRequest) SetBody(b map[string]interface{}) *sRequest {
	sR.body = b
	return sR
}
func (sR *sRequest) SetMultipartForm(m map[string]string, f []string) *sRequest {
	sR.multipartForm = request.MultipartForm{
		Files:  f,
		Fields: m,
	}
	return sR
}
