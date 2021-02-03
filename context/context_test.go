package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestDoRequestWithContext_Timeout(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			time.Sleep(100 * time.Millisecond)
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	client := NewClientWithTimeout(50 * time.Millisecond)

	ch := make(chan Result, 1)
	client.DoRequestWithContext(context.Background(), server.URL, nil, ch)

	res := <-ch
	if res.Error == nil {
		t.Fatalf("expected error")
	}

	err, ok := res.Error.(*url.Error)

	if !ok {
		t.Fatalf("%v", reflect.TypeOf(res.Error))
	}

	if !err.Timeout() {
		t.Fatalf("expexted timeout error")
	}
}

func TestDoRequestWithContext_CtxCancel(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	client := NewClientWithTimeout(50 * time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	ch := make(chan Result, 1)
	client.DoRequestWithContext(ctx, server.URL, nil, ch)

	res := <-ch
	if res.Error == nil {
		t.Fatalf("expected error")
	}

	err, ok := res.Error.(*url.Error)
	if !ok {
		t.Fatalf("%v", reflect.TypeOf(res.Error))
	}

	expectedErr := context.Canceled
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expexted: %v, got: %v", err.Unwrap(), expectedErr)
	}
}

func TestDoRequestWithContext_BadStatus(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusBadRequest)
		}))

	defer server.Close()

	client := NewClientWithTimeout(50 * time.Millisecond)

	ch := make(chan Result, 1)
	client.DoRequestWithContext(context.Background(), server.URL, nil, ch)

	res := <-ch
	if res.Error == nil {
		t.Fatalf("expected error")
	}

	expectedErr := fmt.Errorf("bad status: %d", http.StatusBadRequest)

	if res.Error.Error() != expectedErr.Error() {
		t.Fatalf("expected: %v; got: %v", expectedErr, res.Error)
	}
}

func TestDoRequestWithContext_Success(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	client := NewClientWithTimeout(50 * time.Millisecond)

	ch := make(chan Result, 1)
	client.DoRequestWithContext(context.Background(), server.URL, nil, ch)

	res := <-ch
	if res.Error != nil {
		t.Fatalf("unexpected error: %v", res.Error)
	}
}

func TestDoRequestWithContext_NilContext(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	client := NewClientWithTimeout(50 * time.Millisecond)

	ch := make(chan Result, 1)
	client.DoRequestWithContext(nil, server.URL, nil, ch)

	res := <-ch
	if res.Error == nil {
		t.Fatalf("expected error")
	}

	expectedErr := errors.New("net/http: nil Context")

	if res.Error.Error() != expectedErr.Error() {
		t.Fatalf("expected: %v; got: %v", expectedErr, res.Error)
	}
}

func TestDoRequestWithContext_WaitGroup(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))

	defer server.Close()

	client := NewClientWithTimeout(50 * time.Millisecond)

	done := make(chan Result, 1)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ch := make(chan Result, 1)
		client.DoRequestWithContext(context.Background(), server.URL, wg, ch)
		wg.Wait()
		done <- <-ch
	}()

	select {
	case <-time.After(1 * time.Second):
		t.Fatalf("timeout")
	case res := <-done:
		if res.Error != nil {
			t.Fatalf("unexpected error: %v", res.Error)
		}
	}
}
