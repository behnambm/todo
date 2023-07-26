package todocommon

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func InterfaceToString(value interface{}) string {
	// Check if the value is already a string, and return it if it is.
	if str, ok := value.(string); ok {
		return str
	}

	// Check if the value is of numeric types and convert it to a string.
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	return ""
}

func GetEnvOrPanic(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Panicln(name)
	}
	return v
}

func PostJson(url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, data)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PostJsonWithAuth(url, token string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, data)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PutJsonWithAuth(url, token string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, data)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetWithAuth(url, token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DeleteWithAuth(url, token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
