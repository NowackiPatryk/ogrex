package proxy

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"net/http/httputil"
	"net/url"

	"example.com/ogrex/proxy/config"
)

type Proxy struct {
	config config.Config
}

func NewProxy(proxyConfig config.Config) *Proxy {
	return &Proxy{
		config: proxyConfig,
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
			targetUrl := service.Services[rand.IntN(len(service.Services))]

			targetUrlParsed, err := url.Parse(targetUrl)
			if err != nil {
				panic("Cant parse url")
			}

			proxy := httputil.NewSingleHostReverseProxy(targetUrlParsed)
			proxy.Director = func(r *http.Request) {
				r.URL.Scheme = targetUrlParsed.Scheme
				r.URL.Host = targetUrlParsed.Host
				r.URL.Path = targetUrlParsed.Path + r.URL.Path
				r.Host = targetUrlParsed.Host
			}

		http.Handle(service.Url, proxy)
	}
}


