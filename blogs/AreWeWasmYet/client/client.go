package client

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

type CounterClient struct {
	httpClient *http.Client
}

func New() *CounterClient {
	return &CounterClient{
		httpClient: http.DefaultClient,
	}
}

func (c *CounterClient) IncermentCounter() error {
	req, err := http.NewRequest("put", "http://localhost:8080/add", nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}

func (c *CounterClient) GetCount() (int, error) {
	req, err := http.NewRequest("get", "http://localhost:8080/count", nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	v, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	count, _ := strconv.Atoi(string(v))
	return count, nil
}
