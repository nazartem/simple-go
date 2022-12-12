/* Пример отправки http-запроса на внешний сервис */

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// небольшой сервер, на который будут посылаться запросы
func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "getHandler: incoming request %#v\n", r)
		fmt.Fprintf(w, "getHandler: r.Url %#v\n", r.URL)
	})

	http.HandleFunc("/raw_body", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "postHandler: raw body %s\n", string(body))
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

// простой Get-запрос по url
func runGet() {
	url := "http://127.0.0.1:8080/?param=123&param2=test"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	// будем вычитывать и выводить только тело ответа
	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("http.Get body %#v\n\n\n", string(respBody))
}

// более сложный Get-запрос по url
func runGetFullReq() {
	// указывание метода и хедеров запроса
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"coursera/golang"},
		},
	}

	// пример парсинга url и выставки ещё одного значения в запрос
	req.URL, _ = url.Parse("http://127.0.0.1:8080/?id=42")
	req.URL.Query().Set("user", "rvasily")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("testGetFullReq resp %#v\n\n\n", string(respBody))
}

// пример отправки Post-запроса вместе с данными на сервер
func runTransportAndPost() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	data := `{"id": 42, "user": "rvasily"}`
	body := bytes.NewBufferString(data)

	// создаем новый запрос с методом Post и нужным url
	url := "http://127.0.0.1:8080/raw_body"
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	// передача сформированного запроса на сервер с помощью созданного клиента
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("runTransport %#v\n\n\n", string(respBody))
}

func mainReq() {
	go startServer()

	time.Sleep(100 * time.Millisecond)

	runGet()
	runGetFullReq()
	runTransportAndPost()
}
