package auth

import (
        "testing"
	"net/http"
)


func TestNoAuth(t *testing.T) {
    req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil) 
    result, err := GetAPIKey(req.Header) 
    if err != ErrNoAuthHeaderIncluded {
        t.Errorf("HTTP header with no authentication returned %v, %s", result, err)
    }
}

func TestAuth(t *testing.T) {
    req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil) 
    req.Header.Add("Authorization", "ApiKey abcde12345")
    result, err := GetAPIKey(req.Header) 
    if result != "abcde12345" || err != nil {
        t.Errorf("HTTP header with ApiKey=abcde12345 returned %v, %v", result, err)
    }
}

