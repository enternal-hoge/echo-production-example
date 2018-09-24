package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func TestMsUser(t *testing.T) {

	t.Run("GET /v1/MsUser", func(t *testing.T) {
		request := gorequest.New()
		resp, body, errs := request.Get("http://localhost:1323/v1/MsUser").
			Set("Content-Type", "application/json").
			End()

		if errs != nil {
			panic(errs)
		}

		fmt.Print(body)
		assert.Equal(t, resp.StatusCode, 200, "htpp status should be equal")
	})

	t.Run("POST /v1/MsUser", func(t *testing.T) {

		m := map[string]interface{}{
			"userId":      5,
			"userName":    "page",
			"password":    "pagepage@123",
			"phoneNumber": "090-0000-0000",
			"email":       "page@outlook.com",
			"address":     "unknown",
		}

		mJson, _ := json.Marshal(m)
		contentReader := bytes.NewReader(mJson)

		req, _ := http.NewRequest("POST", "http://localhost:1323/v1/MsUser", contentReader)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(req)

		assert.Equal(t, resp.StatusCode, 201, "htpp status should be equal")
	})

	t.Run("PUT /v1/MsUser", func(t *testing.T) {

		m := map[string]interface{}{
			"userId":      5,
			"userName":    "fuga",
			"password":    "fugafuga@123",
			"phoneNumber": "090-0000-0000",
			"email":       "fuga@outlook.com",
			"address":     "unknown",
		}

		mJson, _ := json.Marshal(m)
		contentReader := bytes.NewReader(mJson)

		req, _ := http.NewRequest("PUT", "http://localhost:1323/v1/MsUser", contentReader)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(req)

		assert.Equal(t, resp.StatusCode, 200, "htpp status should be equal")
	})

	t.Run("DELETE /v1/MsUser", func(t *testing.T) {

		m := map[string]interface{}{
			"userId": 5,
		}

		mJson, _ := json.Marshal(m)
		contentReader := bytes.NewReader(mJson)

		req, _ := http.NewRequest("DELETE", "http://localhost:1323/v1/MsUser", contentReader)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(req)

		assert.Equal(t, resp.StatusCode, 200, "htpp status should be equal")
	})

}
