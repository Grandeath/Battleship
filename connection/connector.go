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
	client         http.Client
	host           string
	token          string
	StartingHeader StartingStruct
}

func NewClient(host string) Client {
	return Client{client: http.Client{}, host: host}
}

func (c *Client) StartGame() error {
	connectionString, err := url.JoinPath(c.host, "/api/game")
	if err != nil {
		return err
	}
	bodyJson, err := json.Marshal(c.StartingHeader)
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

func (c *Client) GetLongDesc() (DescriptionStruct, error) {
	connectionString, err := url.JoinPath(c.host, "/api/game/desc")
	if err != nil {
		return DescriptionStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req, err := http.NewRequest(http.MethodGet, connectionString, http.NoBody)
	if err != nil {
		return DescriptionStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return DescriptionStruct{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return DescriptionStruct{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	newStatus := DescriptionStruct{}
	err = json.NewDecoder(resp.Body).Decode(&newStatus)
	if err != nil {
		return DescriptionStruct{}, err
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

func (c *Client) GetStatus() (StatusStruct, error) {
	connectionString, err := url.JoinPath(c.host, "/api/game")
	if err != nil {
		return StatusStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req, err := http.NewRequest(http.MethodGet, connectionString, http.NoBody)
	if err != nil {
		return StatusStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return StatusStruct{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return StatusStruct{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	newStatus := StatusStruct{}
	err = json.NewDecoder(resp.Body).Decode(&newStatus)
	if err != nil {
		return StatusStruct{}, err
	}
	return newStatus, nil
}

func (c *Client) GetPlayerList() (PlayerListStruct, error) {
	connectionString, err := url.JoinPath(c.host, "/api/game/list")
	if err != nil {
		return PlayerListStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}
	req, err := http.NewRequest(http.MethodGet, connectionString, http.NoBody)
	if err != nil {
		return PlayerListStruct{}, fmt.Errorf("cannot create request: %w", nil)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return PlayerListStruct{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PlayerListStruct{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	playerList := PlayerListStruct{}
	err = json.NewDecoder(resp.Body).Decode(&playerList)
	if err != nil {
		return PlayerListStruct{}, err
	}
	return playerList, nil
}
