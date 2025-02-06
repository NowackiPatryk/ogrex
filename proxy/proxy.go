package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"example.com/ogrex/proxy/config"
	"example.com/ogrex/utils"
)

type Proxy struct {
	config         config.Config
	balancerQueues map[string]*utils.FifoQueue[url.URL]
}

func NewProxy(proxyConfig config.Config) *Proxy {
	return &Proxy{
		config:         proxyConfig,
		balancerQueues: make(map[string]*utils.FifoQueue[url.URL]),
	}
}

func (p *Proxy) Run() {
	p.setupServicesHandlers(p.config.Services)

	serverAddr := fmt.Sprintf(":%d", p.config.Server.Port)

	fmt.Println("Ogrex started at ", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func (p *Proxy) setupServicesHandlers(services map[string]config.ServiceConfig) {
	for _, service := range services {
		parsedUrls := p.parseServiceUrls(service.Services)
		p.setupBalancingForService(service.Url, parsedUrls)
		p.setupRedirectHandlersForService(service.Url)
	}
}

func (p *Proxy) setupBalancingForService(serviceUrl string, parsedUrls []url.URL) {
	balancerQueue := utils.NewFifoQueue(parsedUrls...)
	p.balancerQueues[serviceUrl] = balancerQueue
}

func (p *Proxy) setupRedirectHandlersForService(serviceUrl string) {
	http.HandleFunc(serviceUrl, func(w http.ResponseWriter, r *http.Request) {
		targetUrl := p.getTargetUrlForService(serviceUrl)

		proxy := httputil.NewSingleHostReverseProxy(&targetUrl)
		proxy.Director = func(r *http.Request) {
			r.URL.Scheme = targetUrl.Scheme
			r.URL.Host = targetUrl.Host
			// r.URL.Path = targetUrlParsed.Path + r.URL.Path
			r.Host = targetUrl.Host
		}

		proxy.ServeHTTP(w, r)
	})
}

func (p *Proxy) getTargetUrlForService(serviceUrl string) url.URL {
	queue := p.balancerQueues[serviceUrl]

	if queue.GetLength() == 1 {
		nextUrl, _ := queue.Peek()
		return nextUrl
	}

	nextUrl, _ := queue.TakeLast()
	queue.Insert(nextUrl)
	return nextUrl
}

func (p *Proxy) parseServiceUrls(serviceUrls []string) []url.URL {
	parsedUrls := []url.URL{}

	for _, stringUrl := range serviceUrls {
		targetUrlParsed, err := url.Parse(stringUrl)
		if err != nil {
			panic("Cant parse url")
		}
		parsedUrls = append(parsedUrls, *targetUrlParsed)
	}

	return parsedUrls
}
