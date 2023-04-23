package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BoardResp struct {
	Board []string
}

type Client struct {
	client http.Client
	host   string
	WpBot  bool
	token  string
}

func NewClient(wpBot bool, host string) Client {
	return Client{client: http.Client{}, WpBot: wpBot, host: host}
}

func (c *Client) StartGame() error {
	connectionString, err := url.JoinPath(c.host, "/api/game")
	if err != nil {
		return err
	}
	bodyJson, err := json.Marshal(StartingStruct{Wpbot: c.WpBot})
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(bodyJson)

	req, err := http.NewRequest(http.MethodPost, connectionString, body)
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

func (c *Client) GetBoard() (BoardResp, error) {
	connectionString, err := url.JoinPath(c.host, "/api/game/board")
	if err != nil {
		return BoardResp{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req, err := http.NewRequest(http.MethodGet, connectionString, http.NoBody)
	if err != nil {
		return BoardResp{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return BoardResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return BoardResp{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	newBoard := BoardResp{}
	err = json.NewDecoder(resp.Body).Decode(&newBoard)
	if err != nil {
		return BoardResp{}, err
	}
	return newBoard, nil
}

func (c *Client) GetStatus() (statusStruct, error) {
	connectionString, err := url.JoinPath(c.host, "/api/game/desc")
	if err != nil {
		return statusStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req, err := http.NewRequest(http.MethodGet, connectionString, http.NoBody)
	if err != nil {
		return statusStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return statusStruct{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return statusStruct{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	newStatus := statusStruct{}
	err = json.NewDecoder(resp.Body).Decode(&newStatus)
	if err != nil {
		return statusStruct{}, err
	}
	return newStatus, nil
}

func (c *Client) Fire(coordinates string) (fireStructResp, error) {
	connectionString, err := url.JoinPath(c.host, "/api/game/fire")
	if err != nil {
		return fireStructResp{}, fmt.Errorf("cannot create request: %w", nil)
	}

	bodyJson, err := json.Marshal(fireStruct{coordinates})
	if err != nil {
		return fireStructResp{}, err
	}

	body := bytes.NewBuffer(bodyJson)

	req, err := http.NewRequest(http.MethodPost, connectionString, body)
	if err != nil {
		return fireStructResp{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return fireStructResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fireStructResp{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	newFire := fireStructResp{}
	err = json.NewDecoder(resp.Body).Decode(&newFire)
	if err != nil {
		return fireStructResp{}, err
	}
	return newFire, nil
}
