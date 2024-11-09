package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/helloeave/json"
)

type NmosClient struct {
	BaseUrl         string
	Port            int
	RegistryVersion string
	HttpClient      *http.Client
	Transport       *http.Transport
}

type HttpResponse struct {
	Status        string
	StatusCode    int
	Header        http.Header
	ContentLength int64
	Body          []byte
}

type stop struct {
	error
}

func (c NmosClient) Post(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, c.BaseUrl+":"+strconv.Itoa(c.Port)+endpoint, bytes.NewBuffer([]byte{}))
}

func (c NmosClient) PostWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.MarshalSafeCollections(params)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(http.MethodPost, c.BaseUrl+":"+strconv.Itoa(c.Port)+endpoint, bytes.NewBuffer(json))
}

func (c NmosClient) Do(request *http.Request) (*HttpResponse, error) {
	var err error
	request.Header.Set("Content-Type", "application/json")

	response, err := c.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Status:        response.Status,
		StatusCode:    response.StatusCode,
		Header:        response.Header,
		ContentLength: response.ContentLength,
		Body:          body,
	}, nil
}

func (r HttpResponse) To(value *HttpResponse) {
	err := json.Unmarshal(r.Body, &value)
	if err != nil {
		value = nil
	}
}

func (c NmosClient) Keepalive(endpoint string, k chan string) {
	es := strings.Split(endpoint, "/")
	nodeId := es[len(es)-1]
	err := c.postKeepalive(endpoint, nodeId, k)
	log.Printf("Keepalive [%s]", nodeId)
	if err != nil {
		log.Printf("Keepalive failed for node [%s]. Retrying.", nodeId)
	}
	time.Sleep(5 * time.Second)
	c.Keepalive(endpoint, k)
}

func (c NmosClient) postKeepalive(endpoint string, nodeId string, k chan string) error {
	request, err := c.Post(endpoint)
	if err != nil {
		return err
	}
	response, err := c.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode == 200 {
		k <- nodeId
		return nil
	} else if response.StatusCode == 404 {
		log.Printf("Node [%s] returned 404 on keepalive.", nodeId)
		k <- "404" // Send signal to repost all resources
		return errors.New("404 Not Found")
	} else {
		return errors.New(string(response.Body))
	}
}
