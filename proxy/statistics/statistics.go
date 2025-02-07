package statistics

import (
	"fmt"
	"net/url"
)

type ProxyStatistics struct {
	RequestCount map[url.URL]int
}

func NewProxyStatistics() *ProxyStatistics {
	return &ProxyStatistics{
		RequestCount: make(map[url.URL]int),
	}
}

func (s *ProxyStatistics) AddCount(url url.URL) {
	if _, exists := s.RequestCount[url]; exists {
		s.RequestCount[url]++
		return
	}

	s.RequestCount[url] = 1
}

func (s *ProxyStatistics) GetCount(url url.URL) (int, error) {
	val, exists := s.RequestCount[url]
	if !exists {
		return 0, fmt.Errorf("no value in given key")
	}

	return val, nil
}
