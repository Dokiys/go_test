package std

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

// TestHttpServer 测试http服务端
func TestHttpServer(t *testing.T) {
	handler := func(respWriter http.ResponseWriter, req *http.Request) {
		// respWriter.Header().Set("Content-Type", "image/gif")
		respWriter.Header().Set("Content-Type", "text/plain")
		respWriter.Write([]byte("Hello Work!"))
	}
	mux := http.NewServeMux()
	mux.Handle("/lalala", http.HandlerFunc(handler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
	select {}
}

type Resp struct {}
type Req struct {}
// TestHttpSendGet 测试http发送Get请求
func TestHttpSendGet(t *testing.T) {
	req, err := http.NewRequest("GET", "http://www.4399.com", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	// 设置参数
	params := req.URL.Query()
	params.Add("id", strconv.FormatInt(1324, 10))
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	data := Resp{}
	assert.NoError(t, json.Unmarshal(body, &data))
}

// TestHttpSendPost 测试http发送Post请求
func TestHttpSendPost(t *testing.T) {
	param := Req{}
	bodyData, err := json.Marshal(param)
	assert.NoError(t, err)

	req, err := http.NewRequest( "POST", "http://www.4399.com", bytes.NewReader(bodyData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	data := Resp{}
	assert.NoError(t, json.Unmarshal(body, &data))
}
