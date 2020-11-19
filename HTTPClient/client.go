package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	maxTokenLive = 1800
	authPath     = "/00_AUTH"
	prjListPath  = "/11_PROJECT_LIST"
)

type Client struct {
	Address string
	Client  *http.Client
}

type Config struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// This type implements the http.RoundTripper interface
type AuthRoundTripper struct {
	Transport http.RoundTripper
	Token     string
	Username  string
	Password  string
	Address   string
	mx        sync.RWMutex
}

func (art *AuthRoundTripper) SetToken(token string) {
	art.mx.Lock()
	art.Token = token
	art.mx.Unlock()
}

func (art *AuthRoundTripper) GetToken() string {
	var t string
	art.mx.RLock()
	t = art.Token
	art.mx.RUnlock()
	return t
}

func (art *AuthRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	var err error
	clone := req.Clone(context.TODO()) // <- Не уверене что сюда можно засунуть req.Context()
	if req.Body != nil {
		clone.Body, err = req.GetBody() // Clone doesn't clone body
		if err != nil {
			return &http.Response{StatusCode: http.StatusBadRequest}, err
		}
	}

	if art.Token == "" {
		if err = art.auth(); err != nil { // if token empty then auth
			return &http.Response{StatusCode: http.StatusUnauthorized}, err
		}
	}

	req.Header.Add("Authorization", art.GetToken()) // add auth token to request
	res, err = art.Transport.RoundTrip(req)
	if err != nil {
		return res, err
	}

	if res.StatusCode == http.StatusUnauthorized { // token expires, auth agai(update token)
		if err = art.auth(); err != nil {
			return &http.Response{StatusCode: http.StatusUnauthorized}, err
		}
		clone.Header.Add("Authorization", art.GetToken())
		return art.Transport.RoundTrip(clone)
	}

	return res, nil
}

func (art *AuthRoundTripper) auth() error {
	data := url.Values{}
	data.Set("username", art.Username)
	data.Set("password", art.Password)
	data.Set("tokentime", strconv.Itoa(maxTokenLive))
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, art.Address+authPath, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body := struct {
		Result  string  `json:"result"`
		Token   *string `json:"token"`
		ReqID   *string `json:"reqId"`
		Message *string `json:"message"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return fmt.Errorf("decode body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("auth failed: %s", *body.Message)
	}

	art.SetToken(*body.Token)

	return nil
}

func NewClient(cfg Config) *Client {
	return &Client{
		Address: cfg.Address,
		Client: &http.Client{
			Timeout: time.Second * 30,
			Transport: &AuthRoundTripper{
				Transport: http.DefaultTransport,
				Address:   cfg.Address,
				Username:  cfg.Username,
				Password:  cfg.Password,
			},
		},
	}
}

type Project struct {
	ID                      string  `json:"ID"`
	PPID                    string  `json:"osOrgStructureProjectPpId"`
	Name                    string  `json:"osOrgStructureProjectName"`
	Base                    string  `json:"osOrgStructureProjectBase"`
	Status                  string  `json:"osOrgStructureProjectStatus"`
	LegalItem               string  `json:"osOrgStructureProjectLegalItem"`
	ProjectOwner            string  `json:"osOrgStructureProjectOwner"`
	OwneEmail               *string `json:"osOrgStructureProjectOwnerEmail"`
	ProjectResponsible      string  `json:"osOrgStructureProjectResponsible"`
	ProjectResponsibleEnail *string `json:"osOrgStructureProjectResponsibleEmail"`
}

func (c *Client) GetProjectsList() ([]Project, error) {
	// TODO
	// _ = prjListPath
	resp, err := c.Client.Get(c.Address + prjListPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	projects := []Project{}
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return projects, nil
}

func main() {
	c := NewClient(Config{
		Address:  "",
		Username: "",
		Password: "",
	})

	fmt.Println(c.GetProjectsList())
}
