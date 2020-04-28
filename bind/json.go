package bind

import (
	"bytes"
	"encoding/json"
)

func DecodeJSON(body []byte, obj interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(body))
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
