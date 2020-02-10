package http

import (
	"testing"
)

var (
	URL  string            = "http://abc.com"
	Head map[string]string = map[string]string{"code": "123"}
	Body []byte            = []byte("body")
)

func TestClient_SetURL(t *testing.T) {
	client := &Client{}
	client.SetURL(URL)
	t.Log(client.URL)
}

func TestClient_SetReqHeader(t *testing.T) {
	client := &Client{}
	client.SetURL(URL).SetReqHeader(Head)
	t.Log(client.ReqHeader)
}

func TestClient_SetReqBody(t *testing.T) {
	client := &Client{}
	client.SetURL(URL).SetReqHeader(Head).SetReqBody(Body)
	t.Log(client.ReqBody)
}

func TestClient_Do(t *testing.T) {
	client := &Client{}
	err := client.SetURL(URL).SetReqHeader(Head).SetReqBody(Body).Do()
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(client.ResHeader)
	t.Log(string(client.ResBody))

	client = &Client{}
	err = client.SetURL(URL).SetMethod("POST").SetReqHeader(Head).SetReqBody(Body).Do()
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(client.ResHeader)
	t.Log(string(client.ResBody))
}

func BenchmarkClient_SetReqHeader(b *testing.B) {
	client := &Client{}
	for i := 0; i < b.N; i++ {
		client.SetURL(URL).SetReqHeader(Head)
	}
}

func BenchmarkClient_SetReqBody(b *testing.B) {
	client := &Client{}
	for i := 0; i < b.N; i++ {
		client.SetURL(URL).SetReqBody(Body)
	}
}

func BenchmarkClient_Do(b *testing.B) {
	client := &Client{}
	var err error
	for i := 0; i < b.N; i++ {
		err = client.SetURL(URL).SetMethod("POST").SetReqBody(Body).Do()
		if err != nil {
			b.Log(err)
		}
	}
}
