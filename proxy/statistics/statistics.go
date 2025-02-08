package statistics

import (
	"fmt"
	"net/url"
	"time"
)

type ProxyStatistics struct {
	RequestCount map[url.URL]int
	ResponseTime map[url.URL]responseTime
}

func NewProxyStatistics() *ProxyStatistics {
	return &ProxyStatistics{
		RequestCount: make(map[url.URL]int),
		ResponseTime: make(map[url.URL]responseTime),
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

func (s *ProxyStatistics) Gather(url url.URL, callback func()) {
	fmt.Println("Gathering")
	measurement := newResponseTimeMeasurement(url)
	measurement.Begin()

	callback()

	measurement.End()
	s.AddCount(url)
	s.saveMeasurement(*measurement)
}

func (s *ProxyStatistics) saveMeasurement(measurement responseTimeMeasurement) {
	measurementVal, exists := s.ResponseTime[measurement.measuredUrl]

	if !exists {
		s.ResponseTime[measurement.measuredUrl] = *newResponseTime(measurement.measurement)
		return
	}

	measurementVal.times = append(measurementVal.times, measurement.measurement)
	measurementVal.mean += measurement.measurement / int64(len(measurementVal.times) + 1)

	s.ResponseTime[measurement.measuredUrl] = measurementVal
}

type responseTime struct {
	times []int64
	mean int64
}

func newResponseTime(time int64) *responseTime {
	return &responseTime{
		times: []int64{time},
		mean: time,
	}
}

type responseTimeMeasurement struct {
	measuredUrl url.URL
	beginTime time.Time
	measurement	int64
}

func newResponseTimeMeasurement(url url.URL) *responseTimeMeasurement {
	return &responseTimeMeasurement{
		measuredUrl: url,
	}
}

func (m *responseTimeMeasurement) Begin() {
	m.beginTime = time.Now()
}

func (m* responseTimeMeasurement) End() {
	m.measurement = time.Since(m.beginTime).Milliseconds()
}




