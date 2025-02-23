package httpwrap

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"time"
)


func NewRequestGo(method, url string, bod any, response []any) error {
	if response == nil || reflect.TypeOf(response).Kind() != reflect.Ptr {
		return errors.New("response debe ser un puntero a una estructura")
	}

	var jsonData []byte
	var err error
	if bod != nil {
		jsonData, err = json.Marshal(bod)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url ,bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	if len(jsonData) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return nil
}