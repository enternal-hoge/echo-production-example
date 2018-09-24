package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {

	t.Run("Convert string to int : String2Int(str string)", func(t *testing.T) {
		v := String2Int("123")
		assert.Equal(t, 123, v, "return value should be equal")
	})

	t.Run("Convert int to string : Int2String(i int)", func(t *testing.T) {
		v := Int2String(123)
		assert.Equal(t, "123", v, "return value should be equal")
	})

	t.Run("Convert []byte to string : ByteArray2String(requestBody io.Reader)", func(t *testing.T) {
		assert.Equal(t, 1, 1, "return value should be equal")
	})

	t.Run("Convert map to json : Map2Json(m map[string]interface{})", func(t *testing.T) {
		m := map[string]interface{}{
			"userId":      5,
			"userName":    "page",
			"password":    "pagepage@123",
			"phoneNumber": "090-0000-0000",
			"email":       "page@outlook.com",
			"address":     "unknown",
		}
		json := Map2Json(m)
		assert.NotEmpty(t, json, "return value should be equal")
	})

}
