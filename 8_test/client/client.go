package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	uri    *url.URL
	client *http.Client
}

func New(base *url.URL) *Client {
	return &Client{
		base,
		&http.Client{},
	}
}

func (c *Client) GetResults(s string) ([]byte, error) {
	u := c.uri.ResolveReference(&url.URL{Path: "/search"})
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	q := req.URL.Query()
	q.Add("q", s)
	req.URL.RawQuery = q.Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []byte{}, errors.New("client: the server did not reply with a 200")
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
