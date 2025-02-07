package statistics

import (
	"net/url"
	"testing"
)

func TestShouldAddRequestCount(t *testing.T) {
	stats := NewProxyStatistics()
	parsedUrl, _ := url.Parse("https://example.com")
	wantedCount := 2

	stats.AddCount(*parsedUrl)
	stats.AddCount(*parsedUrl)
	realCount, _ := stats.GetCount(*parsedUrl)

	if wantedCount != realCount {
		t.Errorf("Wrong count. Wanted count: %d Received count: %d", wantedCount, realCount)
	}
}

func TestShouldReturnErrorIfUrlNotCounted(t *testing.T) {
	stats := NewProxyStatistics()
	parsedUrl, _ := url.Parse("https://example.com")

	_, error := stats.GetCount(*parsedUrl)
	if error == nil {
		t.Error("Error should be returned for key that not exists")
	}
}
