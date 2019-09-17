package client

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetResult(t *testing.T) {

	t.Log("Given a search engine server")
	{
		query := "memes"
		wantRes := []byte("Ok")

		server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/search" {
				t.Error("\t\tThe URL does not match")
				t.Errorf("\t\tWant: %s", "/serach")
				t.Errorf("\t\tGot: %s", req.URL.String())
			}

			if got := req.URL.Query().Get("q"); got != query {
				t.Error("\t\tThe URL query does not match")
				t.Errorf("\t\tWant: %s", query)
				t.Errorf("\t\tGot: %s", got)
			}

			if _, err := res.Write(wantRes); err != nil {
				t.Error("\t\tThe response can't be write")
				t.Errorf("\t\tDetails: %s", err)
			}
		}))
		defer server.Close()

		t.Log("And an http client")
		{
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error("\t\tThe server URL can't be parsed")
				t.Fatalf("\t\tDetails: %s", err)
			}
			client := &Client{
				u,
				&http.Client{},
			}
			gotRes, err := client.GetResults(query)
			t.Log("\tThen a error MUST NOT be returned")
			{
				if err != nil {
					t.Error("\t\tThe client call to GetResults returned an error")
					t.Errorf("\t\tDetails: %s", err)
				}
			}
			t.Log("\tAnd a response with the wanted results MUST be retrieved")
			{
				if string(gotRes) != string(wantRes) {
					t.Error("\t\tThe client call to GetResults returned an error")
					t.Errorf("\t\tWant: %s", wantRes)
					t.Errorf("\t\tGot: %s", gotRes)
				}
			}
		}
	}
}

func TestGetResultServerError(t *testing.T) {

	t.Log("Given a not working search engine server")
	{
		query := "memes"
		wantRes := []byte{}

		server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		t.Log("And an http client")
		{
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error("\t\tThe server URL can't be parsed")
				t.Fatalf("\t\tDetails: %s", err)
			}
			client := &Client{
				u,
				&http.Client{},
			}
			gotRes, err := client.GetResults(query)
			t.Log("\tThen a response MUST be empty")
			{
				if string(gotRes) != string(wantRes) {
					t.Error("\t\tThe client returned a non empty body")
					t.Errorf("\t\tWant: %s", wantRes)
					t.Errorf("\t\tGot: %s", gotRes)
				}
			}

			t.Log("\tAnd an error MUST ne returned")
			{
				if err == nil {
					t.Error("\t\tThe client call to GetResults did not return the expected")
					t.Errorf("\t\tWant: %s", "client: the server did not reply with a 200")
					t.Errorf("\t\tGot: %s", err)
				}
			}
		}
	}
}
