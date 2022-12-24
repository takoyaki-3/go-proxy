package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"fmt"
	"net/url"
	"net/http"
	"net/http/httputil"
	"golang.org/x/crypto/acme/autocert"
	json "github.com/takoyaki-3/go-json"
)

type Domain struct {
	Domain	string	`domain`
	Host		string	`host`
	Scheme	string	`scheme`
}

var mapDomains map[string]Domain

func main() {
	// reverse-proxy
	rp := &httputil.ReverseProxy{Director: func(request *http.Request) {
		host := request.Host

		fmt.Println(host)

		url := *request.URL
		url.Scheme = mapDomains[host].Scheme
		url.Host = mapDomains[host].Host

		fmt.Println(host,"=>",url.Host)

		// Loggerを設定
		if host != "logger.api.takoyaki3.com"{
			fmt.Println(host)
			go Log("proxy",request.Host+request.RequestURI)
		} else {
			fmt.Println("skip log")
		}

		if request.Body != nil {
			buffer, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			req, err := http.NewRequest(request.Method, url.String(), bytes.NewBuffer(buffer))
			if err != nil {
				log.Fatal(err.Error())
			}
			req.Header = request.Header
			*request = *req	
		} else {
			req, err := http.NewRequest(request.Method, url.String(), nil)
			if err != nil {
				log.Fatal(err.Error())
			}
			req.Header = request.Header
			*request = *req
		}
	}}

	// initialize
	domains := []Domain{}
	if err:=json.LoadFromPath("./conf.json",&domains);err!=nil{
		log.Fatal(err)
	}
	mapDomains = map[string]Domain{}
	domainStrs := []string{}
	for _,domain := range domains{
		mapDomains[domain.Domain] = domain
		domains = append(domains,domain)
	}

	// start server
	log.Fatal(http.Serve(autocert.NewListener(domainStrs...), rp))
}

// ///
type LogStr struct {
	Contents []string `json:"contents"`
	Service  string   `json:"service"`
	Sign     string   `json:"sign"`
}

func Log(service, content string) error {
	var q LogStr
	q.Service = service
	q.Contents = append(q.Contents, content)
	str, _ := json.DumpToString(q)

	u := "https://logger.api.takoyaki3.com/add?json=" + url.QueryEscape(str)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	client := new(http.Client)
	resp, err := client.Do(req)

	defer resp.Body.Close()
	return err
}
