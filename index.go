package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type proxyHandler struct {
	url string
}

func (ph *proxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", ph.url, nil)
	response, _ := httpClient.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	_, _ = w.Write([]byte(body))
}

type homeHandler struct {
}

func (hh *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello Go!"))
}

func main() {
	url := "https://www.jianshu.com/"
	mux := http.NewServeMux()

	redirectHandler := http.RedirectHandler(url, 307)
	mux.Handle("/redirect", redirectHandler)

	proxyHandler := &proxyHandler{url}
	mux.Handle("/proxy", proxyHandler)

	homeHandler := &homeHandler{}
	mux.Handle("/", homeHandler)

	log.Println("Listening...")
	_ = http.ListenAndServe(":9090", mux)
}
