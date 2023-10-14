//go:build httptest
// +build httptest

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

var (
	baseURL = "http://localhost:8080"
)

func Test_Get(t *testing.T) {
	t.Run("get: device", func(t *testing.T) {
		client := http.DefaultClient

		req, err := http.NewRequest("GET", baseURL+"/api/v1/devices/489", nil)

		if err != nil {
			t.Error("NewRequest failed", err)
		}

		req.Header.Set("Authorization", "Basic b3pvbjpyb3V0ZTI1Ng==")

		resp, err := client.Do(req)

		if err != nil {
			t.Error("get failed", err)
		}

		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		if resp.StatusCode != http.StatusOK {
			b, er := io.ReadAll(resp.Body)

			if er != nil {
				t.Error("ReadAll failed", er)
			}

			t.Errorf("Invalid status code! got: %d, want %d, response body: %s",
				resp.StatusCode,
				http.StatusOK,
				string(b))
		}

		type Device struct {
			ID        string    `json:"id"`
			Platform  string    `json:"platform"`
			UserID    string    `json:"userId"`
			EnteredAt time.Time `json:"enteredAt"`
		}

		type GetDevice struct {
			Device Device `json:"value"`
		}

		device := new(GetDevice)
		err = json.NewDecoder(resp.Body).Decode(device)

		if err != nil {
			t.Error("Unmarshal failed", err)
		}

		t.Log(device)
	})
}
