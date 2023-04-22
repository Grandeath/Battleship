package connection

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Board struct {
	GotBoard []string
}

type Client struct {
	client      http.Client
	host        string
	WpBot       bool
	token       string
	BoardStruct Board
}

func (c *Client) StartGame() error {
	connectionString, err := url.JoinPath(c.host, "/api/game")
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, connectionString, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println(resp.Header.Get("X-Auth-Token"))
	c.token = resp.Header.Get("X-Auth-Token")

	return nil
}

func (c *Client) GetBoard() error {
	connectionString, err := url.JoinPath(c.host, "/api/game/board")
	req, err := http.NewRequest(http.MethodGet, connectionString, http.NoBody)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.token)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&c.BoardStruct)
	if err != nil {
		return err
	}

	return nil
}
