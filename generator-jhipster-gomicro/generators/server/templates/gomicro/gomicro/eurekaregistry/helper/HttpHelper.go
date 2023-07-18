package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/micro/v3/service/logger"
	"net/http"
	"time"
)

type OperationBody interface {
}

// MakePostCall
func MakePostCall(urlToPost string, body OperationBody, headers map[string]string) (error, *http.Response) {
	msg := fmt.Sprintf("In MakePostCall to %s with body %s", urlToPost, body)
	logger.Infof(msg)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err := encoder.Encode(body)
	if err != nil {
		logger.Errorf("error while encoding " + err.Error())
	}
	msg = fmt.Sprintf("Request body  %+v", buffer.String())
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodPost, urlToPost, &buffer)
	if err != nil {
		logger.Errorf("Error while creating the http request  " + err.Error())
		return errors.New("Error while creating http request  " + err.Error()), nil
	}

	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Error while making post call " + err.Error())
		return errors.New("Error while making post call " + err.Error()), nil
	}
	logger.Infof("Successfull POST call with HTTP Status : " + resp.Status)
	return nil, resp
}

func MakePutCall(urlToPost string, body OperationBody, headers map[string]string) (error, *http.Response) {
	msg := fmt.Sprintf("In MakePostCall to %s with body %s", urlToPost, body)
	logger.Infof(msg)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}

	var buffer bytes.Buffer
	if body != nil {
		encoder := json.NewEncoder(&buffer)
		//encoder.SetIndent(" ", "\t")
		err := encoder.Encode(body)
		if err != nil {
			logger.Errorf("error while encoding " + err.Error())
		}
		msg = fmt.Sprintf("Prepared Request body  %+v", buffer.String())
		logger.Infof(msg)
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodPut, urlToPost, &buffer)
	if err != nil {
		logger.Errorf("Error while creating the http request  " + err.Error())
		return errors.New("Error while creating http request " + err.Error()), nil
	}

	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Error while making PUT call " + err.Error())
		return errors.New("Error while making PUT call " + err.Error()), nil
	}
	logger.Infof("PUT call Successfull with HTTP Status : " + resp.Status)
	return nil, resp
}
