package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetRealTimeConfig_NilContext(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}

	expectedErr := "make request: net/http: nil Context"
	_, err = GetRealTimeConfig(nil, u)
	if err.Error() != expectedErr {
		t.Fatalf("expected: %v; got: %v", expectedErr, err.Error())
	}
}

func TestGetRealTimeConfig_ContextCanceled(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	expectedErr := context.Canceled
	_, err = GetRealTimeConfig(ctx, u)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expexted: %v, got: %v", expectedErr, err)
	}
}

func TestGetRealTimeConfig_BadResponseCode(t *testing.T) {
	code := http.StatusBadGateway
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusBadGateway)
		}))

	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}

	expectedErr := fmt.Sprintf("bad status code: %d", code)
	_, err = GetRealTimeConfig(context.Background(), u)
	if err == nil || err.Error() != expectedErr {
		t.Fatalf("expexted: %v, got: %v", expectedErr, err)
	}
}

func TestGetRealTimeConfig_BadResponseBody(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("{[}}"))
		}))

	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}

	expectedErr := "decode body"
	_, err = GetRealTimeConfig(context.Background(), u)
	if !strings.Contains(err.Error(), expectedErr) {
		t.Fatalf("expexted: %v, got: %v", expectedErr, err)
	}
}

func TestGetRealTimeConfig_Success(t *testing.T) {

	config := RealConfig{
		Version:   "1",
		ProjectID: "1",
		Values: map[string]string{
			"best_date_price_max_rps": "20",
			"boxpacker_endpoint":      "o3:///boxpacker.geo:grpc",
		},
	}
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(config)
		}))

	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}

	result, err := GetRealTimeConfig(context.Background(), u)
	if err != nil {
		t.Fatalf("unexpexted error: %v", err)
	}

	if !reflect.DeepEqual(*result, config) {
		t.Fatalf("expexted: %v, got: %v", *result, config)
	}
}
