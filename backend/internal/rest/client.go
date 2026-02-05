package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"
)

// Client handles authenticated requests to Yahoo Finance
type Client struct {
	httpClient *http.Client
	crumb      string
	crumbLock  sync.RWMutex
}

var (
	defaultClient *Client
	once          sync.Once
)

// GetClient returns the singleton instance of the Client, initializing it if necessary
func GetClient() *Client {
	once.Do(func() {
		jar, _ := cookiejar.New(nil)
		defaultClient = &Client{
			httpClient: &http.Client{
				Jar:     jar,
				Timeout: 10 * time.Second,
			},
		}
		// Try to initialize initially, but don't crash if it fails (it will retry on request)
		_ = defaultClient.RefreshSession()
	})
	return defaultClient
}

// RefreshSession fetches a new cookie and crumb
func (c *Client) RefreshSession() error {
	c.crumbLock.Lock()
	defer c.crumbLock.Unlock()

	// 1. Get Cookie from fc.yahoo.com
	// We use io.Discard because we only care about the Set-Cookie header which is handled by the Jar
	req, _ := http.NewRequest("GET", "https://fc.yahoo.com", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get cookie: %v", err)
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)

	// 2. Get Crumb
	req, _ = http.NewRequest("GET", "https://query1.finance.yahoo.com/v1/test/getcrumb", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get crumb: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code getting crumb: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read crumb body: %v", err)
	}

	c.crumb = string(body)
	fmt.Printf("Successfully refreshed session. Crumb: %s\n", c.crumb)
	return nil
}

// Get performs an authenticated GET request, appending the crumb parameter
func (c *Client) Get(url string) (*http.Response, error) {
	// Ensure we have a crumb
	if c.getCrumb() == "" {
		if err := c.RefreshSession(); err != nil {
			return nil, err
		}
	}

	// Append crumb to URL
	separator := "?"
	if contains(url, "?") {
		separator = "&"
	}
	authURL := fmt.Sprintf("%s%scrumb=%s", url, separator, c.getCrumb())

	req, _ := http.NewRequest("GET", authURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// If 401/429, try refreshing session once
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusTooManyRequests {
		resp.Body.Close()
		fmt.Println("Received 401/429, refreshing session...")
		if err := c.RefreshSession(); err != nil {
			return nil, err
		}

		// Update URL with new crumb
		authURL = fmt.Sprintf("%s%scrumb=%s", url, separator, c.getCrumb())
		req, _ = http.NewRequest("GET", authURL, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		return c.httpClient.Do(req)
	}

	return resp, nil
}

func (c *Client) getCrumb() string {
	c.crumbLock.RLock()
	defer c.crumbLock.RUnlock()
	return c.crumb
}

func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}
