package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestRemoveDuplication(t *testing.T) {
	have := removeDuplication([]int{1, 2, 3, 3, 4, 4, 5, 1, 3, 5})
	want := []int{1, 2, 3, 4, 5}
	if !slices.Equal(have, want) {
		t.Errorf("Expected %v, but got %v", want, have)
	}
}

func TestPhoneNumberValidator(t *testing.T) {
	have := PhoneNumberValidator("920227833")
	want := true

	if have != want {
		t.Errorf("Expected %t", want)
	}
}

func TestXMLApiIntegration(t *testing.T) {
	reqBody := `<person><name>Meles</name><email>meles@example.com</email><phone>0911223344</phone></person>`
	req := httptest.NewRequest(http.MethodPost, "/xml", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/xml")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.StatusCode)
	}
	if !strings.Contains(string(body), "<name>Meles</name>") {
		t.Errorf("Expected response to contain <name>Meles</name>, got %s", body)
	}
}
