package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func TestMsAuthority(t *testing.T) {

	t.Run("GET /v1/MsAuthority", func(t *testing.T) {
		request := gorequest.New()
		resp, _, errs := request.Get("http://localhost:1323/v1/MsAuthority").
			Set("Content-Type", "application/json").
			End()

		if errs != nil {
			panic(errs)
		}

		assert.Equal(t, resp.StatusCode, 200, "htpp status should be equal")
	})

	t.Run("POST /v1/MsAuthority", func(t *testing.T) {

		m := map[string]interface{}{
			"userId": 0,
		}

		mJson, _ := json.Marshal(m)
		contentReader := bytes.NewReader(mJson)

		req, _ := http.NewRequest("POST", "http://localhost:1323/v1/MsAuthority", contentReader)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(req)

		assert.Equal(t, resp.StatusCode, 201, "htpp status should be equal")
	})

}
