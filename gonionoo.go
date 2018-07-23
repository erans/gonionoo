package gonionoo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// OOOURL is the OnionOO endpoint URL
	oooURL = "https://onionoo.torproject.org/"
	// OOOVersionMajor is the OnionOO major version
	oooVersionMajor = 6
	// OOOVersionMinor is the OnionOO minor version
	oooVersionMinor = 1
)

var validParameters = map[string]interface{}{
	"type":                nil,
	"running":             nil,
	"search":              nil,
	"lookup":              nil,
	"fingerprint":         nil,
	"country":             nil,
	"as":                  nil,
	"flag":                nil,
	"first_seen_days":     nil,
	"last_seen_days":      nil,
	"contact":             nil,
	"family":              nil,
	"fields":              nil,
	"order":               nil,
	"offset":              nil,
	"limit":               nil,
	"host_name":           nil,
	"recommended_version": nil,
}

var validMethods = map[string]interface{}{
	"summary":   nil,
	"details":   nil,
	"bandwidth": nil,
	"weights":   nil,
	"clients":   nil,
	"uptime":    nil,
}

func validateQueryParameters(query map[string]string) error {
	if query == nil {
		return nil
	}

	for k := range query {
		if _, ok := validParameters[k]; !ok {
			return fmt.Errorf("Invalid parameter '%s' in query", k)
		}
	}

	return nil
}

func validateMethod(method string) error {
	if method == "" {
		return fmt.Errorf("Method cannot be empty")
	}

	if _, ok := validMethods[method]; !ok {
		return fmt.Errorf("Invalid method '%s'", method)
	}

	return nil
}

func constructQueryParametersString(query map[string]string) string {
	if query == nil {
		return ""
	}

	var buffer = bytes.NewBufferString("")
	for key, val := range query {
		buffer.WriteString(fmt.Sprintf("%s=%s&", key, val))
	}

	return buffer.String()
}

func executeRequest(method string, query map[string]string, lastModifiedSince string, result interface{}) (string, error) {
	var err error
	if err = validateQueryParameters(query); err != nil {
		return "", err
	}

	if err = validateMethod(method); err != nil {
		return "", err
	}

	requestURL := fmt.Sprintf("%s%s?%s", oooURL, method, constructQueryParametersString(query))

	var request *http.Request
	if request, err = http.NewRequest("GET", requestURL, nil); err != nil {
		return "", err
	}

	request.Header.Add("Accept-Encoding", "gzip")

	if lastModifiedSince != "" {
		request.Header.Add("If-Modified-Since", lastModifiedSince)
	}

	client := new(http.Client)

	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return "", err
	}
	defer response.Body.Close()

	var reader io.ReadCloser
	// Check we actually go a gzipped response
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		if reader, err = gzip.NewReader(response.Body); err != nil {
			return "", err
		}

		defer reader.Close()
	default:
		reader = response.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return response.Header.Get("Last-Modified"), nil
}

// GetSummary returns a summary of the requested relays and bridges
func GetSummary(query map[string]string, lastModifiedSince string) (*Summary, string, error) {
	summary := new(Summary)

	if lastModified, err := executeRequest("summary", query, lastModifiedSince, &summary); err != nil {
		return nil, "", err
	} else {
		return summary, lastModified, nil
	}
}

// GetDetails returns detailed data of the requested relays and/or bridges
func GetDetails(query map[string]string, lastModifiedSince string) (*Details, string, error) {
	details := new(Details)

	if lastModified, err := executeRequest("details", query, lastModifiedSince, &details); err != nil {
		return nil, "", err
	} else {
		return details, lastModified, nil
	}
}

// GetBandwidth returns details bandwidth data for the requested relays/bridges
func GetBandwidth(query map[string]string, lastModifiedSince string) (*Bandwidth, string, error) {
	bandwidth := new(Bandwidth)

	if lastModified, err := executeRequest("bandwidth", query, lastModifiedSince, &bandwidth); err != nil {
		return nil, "", err
	} else {
		return bandwidth, lastModified, nil
	}
}

// GetWeights returns weights data for the requested relays/bridges
func GetWeights(query map[string]string, lastModifiedSince string) (*Weights, string, error) {
	weights := new(Weights)

	if lastModified, err := executeRequest("weights", query, lastModifiedSince, &weights); err != nil {
		return nil, "", err
	} else {
		return weights, lastModified, nil
	}
}

// GetClients returns clients data for the requested relays/bridges
func GetClients(query map[string]string, lastModifiedSince string) (*Clients, string, error) {
	clients := new(Clients)

	if lastModified, err := executeRequest("clients", query, lastModifiedSince, &clients); err != nil {
		return nil, "", err
	} else {
		return clients, lastModified, nil
	}
}

// GetUptime returns uptime data for the requested relays/bridges
func GetUptime(query map[string]string, lastModifiedSince string) (*Uptime, string, error) {
	uptime := new(Uptime)

	if lastModified, err := executeRequest("uptime", query, lastModifiedSince, &uptime); err != nil {
		return nil, "", err
	} else {
		return uptime, lastModified, nil
	}
}
