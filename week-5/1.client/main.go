package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

type App struct {
	client *http.Client
}

func main() {
	app := appWithHttpClient()
	result, _ := app.postNewItem()
	fmt.Println(string(result))
}

func appWithHttpClient() *App {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			MaxIdleConnsPerHost: 10,
		},
	}

	return &App{client: client}
}

func (a *App) getFromAPI() ([]byte, error) {
	resp, err := a.client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *App) getFromAPIWithClient() ([]byte, error) {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *App) getWithQueryString() ([]byte, error) {

	//id := "1"
	//params := "postId=" + url.QueryEscape(id)
	//path := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?%s", params)

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/comments", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("postId", "1")
	req.URL.RawQuery = q.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func (a *App) signInWithBearerToken() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/bearer", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"))

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *App) postNewItem() ([]byte, error) {
	values := map[string]interface{}{
		"name": "Ahmet",
		"age":  21,
	}

	bodyData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

//https://golang.cafe/blog/how-to-reuse-http-connections-in-go.html
//https://husni.dev/how-to-reuse-http-connection-in-go/
func (a *App) clientWithTrace() {
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("conn was reused: %t", info.Reused)
		},
	}

	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	req, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "http://localhost:8080/api/v1/integration-api/readyz", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	req, err = http.NewRequestWithContext(traceCtx, http.MethodGet, "http://localhost:8080/api/v1/integration-api/readyz", nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
